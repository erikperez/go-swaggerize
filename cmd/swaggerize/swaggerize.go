package main

import (
	"encoding/json"
	"fmt"
	"reflect"
	"strconv"
)

type message struct {
	MessageID   string `json:"messageid" swagger:"required:true;in:query;"`
	ServiceID   int    `json:"serviceid"`
	ServiceName string `json:"servicename"`
	Sno         string `json:"sno"`
}

type getStatus struct {
	Status []string `json:"status" swagger:"required:true;in:query;multiple:true;enum:['available','pending','sold']"`
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

	swaggerize(swag, routes)

}

type SwaggerizeRoute struct {
	Group string
	Route string
	Verb  string
	Model interface{}
}

func swaggerize(swag *SwaggerModel, routes []SwaggerizeRoute) {
	for _, route := range routes {

		if route.Group != "" {
			swag.addTag(SwaggerTag{Name: route.Group})
		}
		var modelName string
		// if route.Model != nil {
		modelName, def := parseStructToDefinition(route.Model)
		swag.addDefinition(modelName, def)
		// }

		var postMethod *SwaggerPathItem
		var getMethod *SwaggerPathItem
		var putMethod *SwaggerPathItem
		var deleteMethod *SwaggerPathItem

		if route.Verb == "post" {
			postMethod = &SwaggerPathItem{Tags: []string{route.Group},
				Consumes:   []string{"application/json"},
				Produces:   []string{"application/json"},
				Responses:  getDefaultResponse(),
				Parameters: []SwaggerPathItemParameter{},
			}

			hasModel := len(swag.Definitions) > 0
			if hasModel {
				postMethod.addParameter(SwaggerPathItemParameter{
					In:       "body",
					Name:     "body",
					Required: true,
					Schema:   &SwaggerSchema{Ref: "#/definitions/" + modelName},
				})
			}
		} else if route.Verb == "get" {
			getMethod = &SwaggerPathItem{Tags: []string{route.Group},
				Consumes:   []string{"application/json"},
				Produces:   []string{"application/json"},
				Responses:  getDefaultResponse(),
				Parameters: []SwaggerPathItemParameter{},
			}

			hasModel := len(swag.Definitions) > 0
			if hasModel {
				getMethod.addParameter(SwaggerPathItemParameter{
					In:       "query",
					Name:     "status",
					Required: true,
					Type:     "string",
				})
			}
		}

		swaggerPathMethods := SwaggerPathMethods{
			Post:   postMethod,
			Put:    putMethod,
			Delete: deleteMethod,
			Get:    getMethod,
		}
		swag.addPath(route.Route, swaggerPathMethods)

		out, _ := json.Marshal(swag)
		fmt.Println(string(out))
	}

}

func getDefaultResponse() map[string]SwaggerPathResponse {
	defaultResponse := make(map[string]SwaggerPathResponse)
	defaultResponse["default"] = SwaggerPathResponse{Description: "Default response"}
	return defaultResponse
}

func parseStructToDefinition(v interface{}) (string, SwaggerDefinition) {
	fields := reflect.TypeOf(v)
	values := reflect.ValueOf(v)
	name := values.Type().Name()
	defType := "object"
	def := SwaggerDefinition{Type: defType}

	for i := 0; i < values.NumField(); i++ {
		field := fields.Field(i)
		value := values.Field(i)
		var reflectedType string
		switch value.Kind() {
		case reflect.Bool:
			reflectedType = "boolean"
		case reflect.String:
			v := value.String()
			fmt.Print(v, "\n")
			reflectedType = "string"
		case reflect.Int:
			v := strconv.FormatInt(value.Int(), 10)
			fmt.Print(v, "\n")
			reflectedType = "integer"
		case reflect.Int32:
			v := strconv.FormatInt(value.Int(), 10)
			fmt.Print(v, "\n")
			reflectedType = "integer"
		case reflect.Int64:
			v := strconv.FormatInt(value.Int(), 10)
			fmt.Print(v, "\n")
			reflectedType = "integer"
		default:
			reflectedType = "string"
		}

		def.addProperty(field.Name, SwaggerDefinitionProperty{
			Type:   reflectedType,
			Format: field.Type.String(),
		})
	}

	return name, def
}
