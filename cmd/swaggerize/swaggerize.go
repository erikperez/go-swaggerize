package main

import (
	"encoding/json"
	"fmt"
	"reflect"
	"strconv"
)

type message struct {
	MessageID   string `json:"messageid"`
	ServiceID   int    `json:"serviceid"`
	ServiceName string `json:"servicename"`
	Sno         string `json:"sno"`
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
		swag.addTag(SwaggerTag{Name: route.Group})
		str, def := parseStructToDefinition(route.Model)
		swag.addDefinition(str, def)

		defaultResponse := getDefaultResponse()
		swaggerPathMethods := SwaggerPathMethods{
			Post: &SwaggerPathItem{Tags: []string{route.Group},
				Consumes:  []string{"application/json"},
				Produces:  []string{"application/json"},
				Responses: defaultResponse,
				Parameters: []SwaggerPathItemParameter{
					SwaggerPathItemParameter{
						In:       "body",
						Name:     "body",
						Required: true,
						Schema:   SwaggerSchema{Ref: "#/definitions/" + str},
					},
				},
			}}
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
