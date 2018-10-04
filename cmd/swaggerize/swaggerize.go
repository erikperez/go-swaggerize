package main

import (
	"encoding/json"
	"fmt"
	"reflect"
	"strconv"
	"strings"
)

type message struct {
	MessageID   string `json:"messageid" swagger:"required:true;in:query;"`
	ServiceID   int    `json:"serviceid"`
	ServiceName string `json:"servicename"`
	Sno         string `json:"sno"`
}

type SwaggerizeOptions struct {
	Name             string
	Required         bool
	In               string
	CollectionFormat string
	Enum             []string
}

func main() {

	swag := NewSwagger("sms.admin.prisguiden.no", "/")
	swag.setInfo(&SwaggerInfo{
		Title: "tester",
		// License: &SwaggerLicense{},
		// Contact: &SwaggerContact{},
	})

	//Example usage:
	routes := []SwaggerizeRoute{}
	routes = append(routes, SwaggerizeRoute{
		Group: "send",
		Route: "/send/message",
		Verb:  "post",
		Model: message{},
	})
	routes = append(routes, SwaggerizeRoute{
		Group: "status",
		Route: "/status",
		Verb:  "get",
		Model: getStatus{},
	})
	routes = append(routes, SwaggerizeRoute{
		Group: "user",
		Route: "/user/{username}",
		Verb:  "get",
		Model: getUser{},
	})
	routes = append(routes, SwaggerizeRoute{
		Group: "user",
		Route: "/user/{username}",
		Verb:  "post",
		Model: putUser{},
	})

	swaggerize(swag, routes)

}

type SwaggerizeRoute struct {
	Group string
	Route string
	Verb  string
	Model interface{}
}

type SwaggerRouteDefinition struct {
	ModelName  *string
	Definition *SwaggerDefinition
	Params     []SwaggerPathItemParameter
}

func swaggerize(swag *SwaggerModel, routes []SwaggerizeRoute) {
	for _, route := range routes {

		routeVerb := strings.ToLower(route.Verb)
		if route.Group != "" {
			swag.addTag(SwaggerTag{Name: route.Group})
		}
		routeDefinition := parseStructToDefinition(route.Model)
		hasModel := routeDefinition.ModelName != nil && routeDefinition.Definition != nil
		if hasModel {
			swag.addDefinition(*routeDefinition.ModelName, *routeDefinition.Definition)
		}

		var postMethod *SwaggerPathItem
		var getMethod *SwaggerPathItem
		var putMethod *SwaggerPathItem
		var deleteMethod *SwaggerPathItem

		var genericMethod = &SwaggerPathItem{Tags: []string{route.Group},
			Consumes:   []string{},
			Produces:   []string{},
			Responses:  getDefaultResponse(),
			Parameters: []SwaggerPathItemParameter{},
		}

		if hasModel && routeVerb != "get" {
			genericMethod.Produces = append(genericMethod.Produces, "application/json")
			genericMethod.Consumes = append(genericMethod.Consumes, "application/json")
			genericMethod.addParameter(SwaggerPathItemParameter{
				In:       "body",
				Name:     "body",
				Required: true,
				Schema:   &SwaggerSchema{Ref: "#/definitions/" + *routeDefinition.ModelName},
			})
		}
		hasParams := len(routeDefinition.Params) > 0
		if hasParams {
			for i := 0; i < len(routeDefinition.Params); i++ {
				genericMethod.addParameter(SwaggerPathItemParameter{
					In:       routeDefinition.Params[i].In,
					Name:     routeDefinition.Params[i].Name,
					Required: routeDefinition.Params[i].Required,
					Type:     routeDefinition.Params[i].Type,
					Enum:     routeDefinition.Params[i].Enum,
				})
			}
		}

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

		// if route.Verb == "post" {
		// 	postMethod = &SwaggerPathItem{Tags: []string{route.Group},
		// 		Consumes:   []string{"application/json"},
		// 		Produces:   []string{"application/json"},
		// 		Responses:  getDefaultResponse(),
		// 		Parameters: []SwaggerPathItemParameter{},
		// 	}

		// 	if hasModel {
		// 		postMethod.addParameter(SwaggerPathItemParameter{
		// 			In:       "body",
		// 			Name:     "body",
		// 			Required: true,
		// 			Schema:   &SwaggerSchema{Ref: "#/definitions/" + *routeDefinition.ModelName},
		// 		})
		// 	}
		// 	hasParams := len(routeDefinition.Params) > 0
		// 	if hasParams {
		// 		for i := 0; i < len(routeDefinition.Params); i++ {
		// 			postMethod.addParameter(SwaggerPathItemParameter{
		// 				In:       routeDefinition.Params[i].In,
		// 				Name:     routeDefinition.Params[i].Name,
		// 				Required: routeDefinition.Params[i].Required,
		// 				Type:     routeDefinition.Params[i].Type,
		// 				Enum:     routeDefinition.Params[i].Enum,
		// 			})
		// 		}
		// 	}
		// } else if route.Verb == "get" {
		// 	getMethod = &SwaggerPathItem{Tags: []string{route.Group},
		// 		Consumes:   []string{"application/json"},
		// 		Produces:   []string{"application/json"},
		// 		Responses:  getDefaultResponse(),
		// 		Parameters: []SwaggerPathItemParameter{},
		// 	}

		// 	hasParams := len(routeDefinition.Params) > 0
		// 	if hasParams {
		// 		for i := 0; i < len(routeDefinition.Params); i++ {
		// 			getMethod.addParameter(SwaggerPathItemParameter{
		// 				In:       routeDefinition.Params[i].In,
		// 				Name:     routeDefinition.Params[i].Name,
		// 				Required: routeDefinition.Params[i].Required,
		// 				Type:     routeDefinition.Params[i].Type,
		// 				Enum:     routeDefinition.Params[i].Enum,
		// 			})
		// 		}
		// 	}
		// }

		swaggerPathMethods := SwaggerPathMethods{
			Post:   postMethod,
			Put:    putMethod,
			Delete: deleteMethod,
			Get:    getMethod,
		}
		swag.addPath(route.Route, swaggerPathMethods)
	}
	out, _ := json.Marshal(swag)
	fmt.Println(string(out))
}

func getDefaultResponse() map[string]SwaggerPathResponse {
	defaultResponse := make(map[string]SwaggerPathResponse)
	defaultResponse["default"] = SwaggerPathResponse{Description: "Default response"}
	return defaultResponse
}

func parseStructToDefinition(v interface{}) SwaggerRouteDefinition {
	fields := reflect.TypeOf(v)
	values := reflect.ValueOf(v)
	name := values.Type().Name()
	defType := "object"

	routeParams := []SwaggerPathItemParameter{}
	definition := &SwaggerDefinition{Type: defType}

	for i := 0; i < values.NumField(); i++ {
		field := fields.Field(i)
		value := values.Field(i)

		var reflectedType string
		switch value.Kind() {
		case reflect.Bool:
			reflectedType = "boolean"
		case reflect.String:
			reflectedType = "string"
		case reflect.Int:
			reflectedType = "integer"
		case reflect.Int32:
			reflectedType = "integer"
		case reflect.Int64:
			reflectedType = "integer"
		default:
			reflectedType = "string"
		}

		prop := SwaggerDefinitionProperty{
			Type:   reflectedType,
			Format: field.Type.String(),
		}

		tag := field.Tag.Get("swagger")
		paramOptions := parseParamsOptions(tag)
		if paramOptions != nil {
			paramName := field.Name
			if paramOptions.Name != "" {
				paramName = paramOptions.Name
			}
			routeParams = append(routeParams, SwaggerPathItemParameter{
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

		definition.addProperty(field.Name, prop)
	}

	return SwaggerRouteDefinition{ModelName: &name, Definition: definition, Params: routeParams}
}

func parseParamsOptions(tag string) *SwaggerizeOptions {
	if tag == "" {
		return nil
	}
	ret := &SwaggerizeOptions{}
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
				ret.Enum = strings.Split(p, ",")
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

type getStatus struct {
	Status []string `json:"status" swagger:"required:true;in:query;multiple:true;enum:['available','pending','sold']"`
}

type getUser struct {
	Username string `json:"username" swagger:"required:true;in:path;name:username"`
}

type putUser struct {
	Username string `json:"username" swagger:"required:true;in:path;name:username"`
	Email    string
}
