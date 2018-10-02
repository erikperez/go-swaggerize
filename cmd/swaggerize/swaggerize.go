package main

import (
	"encoding/json"
	"fmt"
	"reflect"
	"strconv"
)

type model struct {
	MessageID   string `json:"messageid"`
	ServiceID   int    `json:"serviceid"`
	ServiceName string `json:"servicename"`
	Sno         string `json:"sno"`
}

func main() {

	// path := "/send"
	m := model{}

	swag := NewSwagger("sms.admin.prisguiden.no", "/")
	swag.setInfo(&SwaggerInfo{
		Title: "tester",
		// License: &SwaggerLicense{},
		Contact: &SwaggerContact{},
	})
	swag.addTag(SwaggerTag{Name: "send"})
	swag.addDefinition(parseStructToDefinition(m))

	swaggerpath := SwaggerPath{}
	swaggerpath.addVerb("post", SwaggerPathItem{Tags: []string{"send"}, Consumes: []string{"application/json"}})

	// swag.addPath(path, swaggerpath)

	out, _ := json.Marshal(swag)
	fmt.Println(string(out))

	// val := reflect.Indirect(reflect.ValueOf(m))
	// fmt.Println(val.Type().Name())
	// for i := 0; i < val.NumField(); i++ {
	// 	field := val.Type().Field(i)
	// 	fmt.Println(field.Name, field.Type)
	// }
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
