package swaggerizer

import "github.com/erikperez/go-swaggerize/pkg/swagger"

type SwaggerizeRoute struct {
	Group     string
	Route     string
	Verb      string
	Model     interface{}
	Responses []SwaggerizeResponse
	Produces  []string
	Consumes  []string
}

type SwaggerizeResponse struct {
	Name        string
	Description string
	Model       interface{}
	Headers     []SwaggerizeResponseHeader
}

type SwaggerizeResponseHeader struct {
	Name        string
	Type        string
	Format      string
	Description string
}
type swaggerizeOptions struct {
	Name             string
	Required         bool
	In               string
	CollectionFormat string
	Enum             []string
}

type swaggerizeRouteDefinition struct {
	ModelName  *string
	Definition *swagger.SwaggerDefinition
	Params     []swagger.SwaggerPathItemParameter
}

type swaggerizeResponseDefinition struct {
	Definition *swagger.SwaggerDefinition
	ModelName  *string
}
