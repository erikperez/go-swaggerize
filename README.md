# go-swaggerize
Give your models som swag.

[![CircleCI](https://circleci.com/gh/erikperez/go-swaggerize/tree/master.svg?style=svg)](https://circleci.com/gh/erikperez/go-swaggerize/tree/master)

### Example usage:
```
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
```
