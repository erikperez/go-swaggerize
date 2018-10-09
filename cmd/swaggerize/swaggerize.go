package main

import (
	"github.com/erikperez/go-swaggerize/internal/app"
	"github.com/erikperez/go-swaggerize/pkg/swagger"
	"github.com/erikperez/go-swaggerize/pkg/swaggerizer"
)

func main() {

	swag := swagger.NewSwagger("sms.admin.prisguiden.no", "/")
	swag.SetInfo(&swagger.SwaggerInfo{
		Title: "tester",
		// License: &SwaggerLicense{},
		// Contact: &SwaggerContact{},
	})

	//Example usage:
	routes := []swaggerizer.SwaggerizeRoute{}
	routes = append(routes, swaggerizer.SwaggerizeRoute{
		Group: "send",
		Route: "/send/message",
		Verb:  "post",
		Model: message{},
	})
	routes = append(routes, swaggerizer.SwaggerizeRoute{
		Group: "status",
		Route: "/status",
		Verb:  "get",
		Model: getStatus{},
	})
	routes = append(routes, swaggerizer.SwaggerizeRoute{
		Group: "user",
		Route: "/user/{username}",
		Verb:  "get",
		Model: getUser{},
	})
	routes = append(routes, swaggerizer.SwaggerizeRoute{
		Group: "user",
		Route: "/user/{username}",
		Verb:  "post",
		Model: putUser{},
	})
	routes = append(routes, swaggerizer.SwaggerizeRoute{
		Group: "user",
		Route: "/user/{username}",
		Verb:  "put",
		Model: putUser{},
		Responses: []swaggerizer.SwaggerizeResponse{
			swaggerizer.SwaggerizeResponse{
				Name:        "default",
				Description: "The default response when everything is ok",
			},
			swaggerizer.SwaggerizeResponse{
				Name:        "200",
				Model:       putUserResponse{},
				Description: "The user has been put",
			},
		},
	})

	app.Swaggerize(swag, routes)

}

type message struct {
	MessageID   string `json:"messageid" swagger:"required:true;in:query;"`
	ServiceID   int    `json:"serviceid"`
	ServiceName string `json:"servicename"`
	Sno         string `json:"sno"`
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

type putUserResponse struct {
	Success bool
	Error   string
}
