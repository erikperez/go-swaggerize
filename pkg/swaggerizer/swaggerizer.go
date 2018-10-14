package swaggerizer

import (
	"encoding/json"
	"reflect"
	"strconv"
	"strings"

	"github.com/erikperez/go-swaggerize/pkg/swagger"
)

// Swaggerize converts an array of Routes into a Swagger 2.0 model (swagger.Model)
func Swaggerize(swag *swagger.Model, routes []Route) (string, error) {
	for _, route := range routes {
		routeVerb := strings.ToLower(route.Verb)
		routeDefinition := parseStructToDefinition(route.Model)
		hasParams := len(routeDefinition.Params) > 0
		hasModel := routeDefinition.ModelName != nil && routeDefinition.Definition != nil

		if route.Group != "" {
			swag.AddTag(swagger.Tag{Name: route.Group})
		}

		if len(route.Produces) == 0 {
			route.Produces = append(route.Produces, "application/json")
		}

		if len(route.Consumes) == 0 {
			route.Consumes = append(route.Consumes, "application/json")
		}

		var genericMethod = &swagger.PathItem{
			Tags:       []string{route.Group},
			Consumes:   []string{},
			Produces:   []string{},
			Parameters: []swagger.PathItemParameter{},
		}

		responses, responseDefinitions := parseResponses(route.Responses)
		genericMethod.Responses = responses
		for _, responseDefinition := range responseDefinitions {
			swag.AddDefinition(*responseDefinition.ModelName, *responseDefinition.Definition)
		}

		if hasModel {
			swag.AddDefinition(*routeDefinition.ModelName, *routeDefinition.Definition)

			if routeVerb != "get" {
				genericMethod.Produces = route.Produces
				genericMethod.Consumes = route.Consumes

				genericMethod.AddParameter(swagger.PathItemParameter{
					In:       "body",
					Name:     "body",
					Required: true,
					Schema:   &swagger.Schema{Ref: "#/definitions/" + *routeDefinition.ModelName},
				})
			}
		}

		if hasParams {
			for i := 0; i < len(routeDefinition.Params); i++ {
				genericMethod.AddParameter(swagger.PathItemParameter{
					In:       routeDefinition.Params[i].In,
					Name:     routeDefinition.Params[i].Name,
					Required: routeDefinition.Params[i].Required,
					Type:     routeDefinition.Params[i].Type,
					Enum:     routeDefinition.Params[i].Enum,
				})
			}
		}

		var postMethod *swagger.PathItem
		var getMethod *swagger.PathItem
		var putMethod *swagger.PathItem
		var deleteMethod *swagger.PathItem

		switch routeVerb {
		case "get":
			getMethod = genericMethod
			break
		case "post":
			postMethod = genericMethod
			break
		case "delete":
			deleteMethod = genericMethod
			break
		case "put":
			putMethod = genericMethod
			break
		}

		swaggerPathMethods := swagger.PathMethods{
			Post:   postMethod,
			Put:    putMethod,
			Delete: deleteMethod,
			Get:    getMethod,
		}
		swag.AddPath(route.Route, swaggerPathMethods)
	}
	out, err := json.Marshal(swag)
	if err != nil {
		return "", err
	}

	return string(out), nil
}

func parseResponses(responses []Response) (map[string]swagger.PathResponse, []responseDefinition) {
	ret := make(map[string]swagger.PathResponse)
	definitions := []responseDefinition{}
	if len(responses) == 0 {
		ret["default"] = swagger.PathResponse{Description: "Default response"}
	} else {
		for i := 0; i < len(responses); i++ {
			response := responses[i]
			resp := swagger.PathResponse{Description: response.Description}
			if response.Model != nil {
				m := parseStructToDefinition(response.Model)
				if m.Definition != nil {
					definitions = append(definitions, responseDefinition{
						Definition: m.Definition,
						ModelName:  m.ModelName,
					})
					resp.Schema = swagger.Schema{Ref: "#/definitions/" + *m.ModelName}
				}
			}
			ret[response.Name] = resp
		}
	}
	return ret, definitions
}

func parseStructToDefinition(v interface{}) routeDefinition {
	fields := reflect.TypeOf(v)
	values := reflect.ValueOf(v)
	structName := values.Type().Name()
	defType := "object"

	routeParams := []swagger.PathItemParameter{}
	definition := &swagger.Definition{Type: defType}

	for i := 0; i < values.NumField(); i++ {
		field := fields.Field(i)
		value := values.Field(i)

		var reflectedType string
		var reflectedFormat string
		switch value.Kind() {
		case reflect.Bool:
			reflectedType = "boolean"
		case reflect.String:
			reflectedType = "string"
		case reflect.Int:
			reflectedType = "integer"
			reflectedFormat = "int32"
		case reflect.Int32:
			reflectedType = "integer"
			reflectedFormat = "int32"
		case reflect.Int64:
			reflectedType = "long"
			reflectedFormat = "int64"
		case reflect.Float32:
		case reflect.Float64:
			reflectedType = "float"
			reflectedFormat = "float"
		default:
			reflectedType = "string"
		}

		if reflectedFormat == "" {
			reflectedFormat = strings.ToLower(field.Type.String())
		}

		prop := swagger.DefinitionProperty{
			Type:   reflectedType,
			Format: reflectedFormat,
		}

		tag := field.Tag.Get("swagger")
		paramName := field.Name
		paramOptions := parseParamsOptions(tag)
		if paramOptions != nil {

			if paramOptions.Name != "" {
				paramName = paramOptions.Name
			}

			routeParams = append(routeParams, swagger.PathItemParameter{
				Required:         paramOptions.Required,
				In:               paramOptions.In,
				CollectionFormat: paramOptions.CollectionFormat,
				Enum:             paramOptions.Enum,
				Name:             paramName,
				Type:             reflectedType,
				Format:           field.Type.String(),
			})
			prop.Enum = paramOptions.Enum
		}

		definition.AddProperty(paramName, prop)
	}

	return routeDefinition{ModelName: &structName, Definition: definition, Params: routeParams}
}

func parseParamsOptions(tag string) *options {
	if tag == "" {
		return nil
	}
	ret := &options{}
	splitted := strings.Split(tag, ";")
	for i := 0; i < len(splitted); i++ {
		splitVar := strings.Split(splitted[i], ":")

		switch splitVar[0] {
		case "required":
			{
				p, err := strconv.ParseBool(splitVar[1])
				if err == nil {
					ret.Required = p
				}
				break
			}
		case "in":
			{
				p := splitVar[1]
				ret.In = p
				break
			}
		case "multiple":
			{
				p, err := strconv.ParseBool(splitVar[1])
				if err == nil && p {
					ret.CollectionFormat = "multi"
				}
				break
			}
		case "enum":
			{
				p := splitVar[1]
				p = p[1 : len(p)-1]
				val := []string{}
				for _, v := range strings.Split(p, ",") {
					val = append(val, v[1:len(v)-1])
				}
				ret.Enum = val //strings.Split(p, ",")
				break
			}
		case "name":
			{
				p := splitVar[1]
				ret.Name = p
				break
			}
		default:
			break

		}
	}
	return ret
}
