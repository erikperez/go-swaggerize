package app

import "github.com/erikperez/go-swaggerize/pkg/swagger"

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
