package docs

import (
	"testing"

	"github.com/erikperez/go-swaggerize/pkg/swagger"
	"github.com/erikperez/go-swaggerize/pkg/swaggerizer"
)

type getStatus struct {
	Status []string `swagger:"required:true;in:query;multiple:true;enum:['available','pending','sold']"`
}

func TestExample(t *testing.T) {
	swag := swagger.NewSwagger("myapi.example.com", "/")
	swag.SetInfo(&swagger.SwaggerInfo{
		Title: "My API Example",
		License: &swagger.SwaggerLicense{
			Name: "Choose a license",
			URL:  "https://choosealicense.com/",
		},
		Contact: &swagger.SwaggerContact{
			Name: "erikperez",
			URL:  "github.com/erikperez",
		},
	})
	routes := []swaggerizer.SwaggerizeRoute{}
	routes = append(routes, swaggerizer.SwaggerizeRoute{
		Group: "status",
		Route: "/status",
		Verb:  "get",
		Model: getStatus{},
	})

	o, _ := swaggerizer.Swaggerize(swag, routes)
	t.Logf(o)
}
