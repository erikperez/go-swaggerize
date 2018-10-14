package swaggerizer

import "github.com/erikperez/go-swaggerize/pkg/swagger"

// SwaggerizeRoute is a holder object used to define a Swagger Route Definition
type SwaggerizeRoute struct {
	Group     string
	Route     string
	Verb      string
	Model     interface{}
	Responses []SwaggerizeResponse
	Produces  []string
	Consumes  []string
}

// SwaggerizeResponse is a holder object used to define a response object for a request and/or route
type SwaggerizeResponse struct {
	Name        string
	Description string
	Model       interface{}
	Headers     []SwaggerizeResponseHeader
}

// SwaggerizeResponseHeader is used by SwaggerizeResponse to define a response's headers
type SwaggerizeResponseHeader struct {
	Name        string
	Type        string
	Format      string
	Description string
}

// swaggerizeOptions internal struct used to parse struct tags on the models.
type swaggerizeOptions struct {
	Name             string
	Required         bool
	In               string
	CollectionFormat string
	Enum             []string
}

// swaggerizeRouteDefinition is an internal struct used to parse a route definition
type swaggerizeRouteDefinition struct {
	ModelName  *string
	Definition *swagger.SwaggerDefinition
	Params     []swagger.SwaggerPathItemParameter
}

// swaggerizeResponseDefinition is an internal struct used to parse a route definition's response.
type swaggerizeResponseDefinition struct {
	Definition *swagger.SwaggerDefinition
	ModelName  *string
}
