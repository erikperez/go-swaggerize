package swaggerizer

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
