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
	swag.SetInfo(&swagger.Info{
		Title: "My API Example",
		License: &swagger.License{
			Name: "Choose a license",
			URL:  "https://choosealicense.com/",
		},
		Contact: &swagger.Contact{
			Name: "erikperez",
			URL:  "github.com/erikperez",
		},
	})
	routes := []swaggerizer.Route{}
	routes = append(routes, swaggerizer.Route{
		Group: "status",
		Route: "/status",
		Verb:  "get",
		Model: getStatus{},
	})

	o, _ := swaggerizer.Swaggerize(swag, routes)
	t.Logf(o)
}
