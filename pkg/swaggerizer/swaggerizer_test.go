package swaggerizer

import (
	"testing"

	"github.com/erikperez/go-swaggerize/pkg/swagger"
)

func TestSwaggerize(t *testing.T) {
	swag := swagger.NewSwagger("myapi.example.com", "/")
	swag.SetInfo(&swagger.SwaggerInfo{
		Title: "tester",
		License: &swagger.SwaggerLicense{
			Name: "Choose a license",
			URL:  "https://choosealicense.com/",
		},
		Contact: &swagger.SwaggerContact{
			Name: "erikperez",
			URL:  "github.com/erikperez",
		},
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
	routes = append(routes, SwaggerizeRoute{
		Group: "user",
		Route: "/user/{username}",
		Verb:  "put",
		Model: putUser{},
		Responses: []SwaggerizeResponse{
			SwaggerizeResponse{
				Name:        "default",
				Description: "The default response when everything is ok",
			},
			SwaggerizeResponse{
				Name:        "200",
				Model:       putUserResponse{},
				Description: "The user has been put",
			},
		},
	})
	o, err := Swaggerize(swag, routes)
	if err != nil {
		t.Logf(err.Error())
		t.Fail()
		return
	}
	t.Logf(o)
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
