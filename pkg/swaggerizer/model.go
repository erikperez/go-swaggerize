package swaggerizer

import "github.com/erikperez/go-swaggerize/pkg/swagger"

// Route is a holder object used to define a Swagger Route Definition
type Route struct {
	Group     string
	Route     string
	Verb      string
	Model     interface{}
	Responses []Response
	Produces  []string
	Consumes  []string
}

// Response is a holder object used to define a response object for a request and/or route
type Response struct {
	Name        string
	Description string
	Model       interface{}
	Headers     []ResponseHeader
}

// ResponseHeader is used by Response to define a response's headers
type ResponseHeader struct {
	Name        string
	Type        string
	Format      string
	Description string
}

// Options internal struct used to parse struct tags on the models.
type options struct {
	Name             string
	Required         bool
	In               string
	CollectionFormat string
	Enum             []string
}

// RouteDefinition is an internal struct used to parse a route definition
type routeDefinition struct {
	ModelName  *string
	Definition *swagger.Definition
	Params     []swagger.PathItemParameter
}

// ResponseDefinition is an internal struct used to parse a route definition's response.
type responseDefinition struct {
	Definition *swagger.Definition
	ModelName  *string
}
