package swagger

// SwaggerModel is the struct of the swagger spec
type SwaggerModel struct {
	Swagger             string                               `json:"swagger"`
	Info                *SwaggerInfo                         `json:"info"`
	Host                string                               `json:"host"`
	BasePath            string                               `json:"basePath"`
	Tags                []SwaggerTag                         `json:"tags"`
	Schemes             []string                             `json:"schemes"`
	Paths               map[string]SwaggerPathMethods        `json:"paths"`
	Definitions         map[string]SwaggerDefinition         `json:"definitions,omitempty"`
	SecurityDefinitions map[string]SwaggerSecurityDefinition `json:"securityDefinitions,omitempty"`
	ExternalDocs        *SwaggerExternalDocs                 `json:"externalDocs,omitempty"`
}

// NewSwagger creates an instance of SwaggerModel
func NewSwagger(host string, basePath string) *SwaggerModel {
	return &SwaggerModel{
		Swagger:             "2.0",
		Host:                host,
		BasePath:            basePath,
		Tags:                []SwaggerTag{},
		Paths:               make(map[string]SwaggerPathMethods),
		Definitions:         make(map[string]SwaggerDefinition),
		SecurityDefinitions: make(map[string]SwaggerSecurityDefinition),
		Schemes:             []string{},
	}
}

// SetInfo sets the info on the swagger model.
func (s *SwaggerModel) SetInfo(info *SwaggerInfo) *SwaggerModel {
	s.Info = info
	return s
}
func contains(s []SwaggerTag, e SwaggerTag) bool {
	for _, a := range s {
		if a.Name == e.Name {
			return true
		}
	}
	return false
}

// AddTag adds a tag if it does not exist
func (s *SwaggerModel) AddTag(tag SwaggerTag) *SwaggerModel {
	if !contains(s.Tags, tag) {
		s.Tags = append(s.Tags, tag)
	}
	return s
}

// AddDefinition set a definition on SwaggerModel.Definitions map. Name is used as key.
func (s *SwaggerModel) AddDefinition(name string, definition SwaggerDefinition) *SwaggerModel {
	s.Definitions[name] = definition
	return s
}

// SetDefinitions sets the SwaggerModel.Definitions map.
func (s *SwaggerModel) SetDefinitions(definitions map[string]SwaggerDefinition) *SwaggerModel {
	s.Definitions = definitions
	return s
}

// AddPath adds SwaggerPathMethods on a path's name. Supports: Get, Post, Put, Delete
func (s *SwaggerModel) AddPath(name string, definition SwaggerPathMethods) *SwaggerModel {
	if val, ok := s.Paths[name]; ok {
		if definition.Post != nil {
			val.Post = definition.Post
		} else if definition.Get != nil {
			val.Get = definition.Get
		} else if definition.Put != nil {
			val.Put = definition.Put
		} else if definition.Delete != nil {
			val.Delete = definition.Delete
		}
		s.Paths[name] = val
	} else {
		s.Paths[name] = definition
	}
	return s
}

func (s *SwaggerModel) SetPaths(paths map[string]SwaggerPathMethods) *SwaggerModel {
	s.Paths = paths
	return s
}

// SwaggerInfo is a holder object used to define the swagger spec and serialize to JSON
type SwaggerInfo struct {
	Description    string          `json:"description"`
	Version        string          `json:"version"`
	Title          string          `json:"title"`
	TermsOfService string          `json:"termsOfService,omitempty"`
	Contact        *SwaggerContact `json:"contact,omitempty"`
	License        *SwaggerLicense `json:"license,omitempty"`
}

// SwaggerContact is a holder object used to define the swagger spec and serialize to JSON
type SwaggerContact struct {
	Name  string `json:"name,omitempty"`
	Email string `json:"email,omitempty"`
	URL   string `json:"url,omitempty"`
}

// SwaggerLicense is a holder object used to define the swagger spec and serialize to JSON
type SwaggerLicense struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

// SwaggerTag is a holder object used to define the swagger spec and serialize to JSON
type SwaggerTag struct {
	Name         string               `json:"name,omitempty"`
	Description  string               `json:"description,omitempty"`
	ExternalDocs *SwaggerExternalDocs `json:"externalDocs,omitempty"`
}

// SwaggerExternalDocs is a holder object used to define the swagger spec and serialize to JSON
type SwaggerExternalDocs struct {
	Description string `json:"description,omitempty"`
	URL         string `json:"url,omitempty"`
}

// SwaggerPath is a holder object used to define the swagger spec and serialize to JSON
type SwaggerPath struct {
	Verbs map[string]SwaggerPathItem
}

func (path *SwaggerPath) AddVerb(name string, pathItem SwaggerPathItem) *SwaggerPath {
	if path.Verbs == nil {
		path.Verbs = make(map[string]SwaggerPathItem)
	}
	path.Verbs[name] = pathItem
	return path
}

// SwaggerPathItem is a holder object used to define the swagger spec and serialize to JSON
type SwaggerPathItem struct {
	Tags        []string                             `json:"tags,omitempty"`
	Summary     string                               `json:"summary,omitempty"`
	Description string                               `json:"description,omitempty"`
	OperationID string                               `json:"operationId,omitempty"`
	Consumes    []string                             `json:"consumes,omitempty"`
	Produces    []string                             `json:"produces,omitempty"`
	Parameters  []SwaggerPathItemParameter           `json:"parameters,omitempty"`
	Responses   map[string]SwaggerPathResponse       `json:"responses,omitempty"`
	Security    map[string]SwaggerSecurityDefinition `json:"security,omitempty"`
}

// AddParameter adds a SwaggerPathItemParameter on a SwaggerPathItems' parameter
func (pathItem *SwaggerPathItem) AddParameter(parameter SwaggerPathItemParameter) *SwaggerPathItem {
	pathItem.Parameters = append(pathItem.Parameters, parameter)
	return pathItem
}

// SwaggerPathMethods is a holder object used to define the swagger spec and serialize to JSON
type SwaggerPathMethods struct {
	Post   *SwaggerPathItem `json:"post,omitempty"`
	Get    *SwaggerPathItem `json:"get,omitempty"`
	Put    *SwaggerPathItem `json:"put,omitempty"`
	Delete *SwaggerPathItem `json:"delete,omitempty"`
}

// SwaggerPathItemParameter is a holder object used to define the swagger spec and serialize to JSON
type SwaggerPathItemParameter struct {
	Ref              string         `json:"$ref,omitempty"`
	In               string         `json:"in,omitempty"` //query, header, path, formdata, or body
	Name             string         `json:"name,omitempty"`
	Description      string         `json:"description,omitempty"`
	Required         bool           `json:"required,omitempty"`
	Enum             []string       `json:"enum,omitempty"`
	Type             string         `json:"type,omitempty"`
	Format           string         `json:"format,omitempty"`
	Schema           *SwaggerSchema `json:"schema,omitempty"`
	CollectionFormat string         `json:"collectionFormat,omitempty"`
}

// SwaggerSchema is a holder object used to define the swagger spec and serialize to JSON
type SwaggerSchema struct {
	Ref              string  `json:"$ref,omitempty"`
	Format           string  `json:"format,omitempty"`
	Title            string  `json:"title,omitempty"`
	Description      string  `json:"description,omitempty"`
	Default          string  `json:"default,omitempty"`
	Maximum          float64 `json:"maximum,omitempty"`
	Minimum          float64 `json:"minimum,omitempty"`
	ExclusiveMaximum bool    `json:"exclusiveMaximum,omitempty"`
}

// SwaggerPathResponse is a holder object used to define the swagger spec and serialize to JSON
type SwaggerPathResponse struct {
	Ref         string                   `json:"$ref,omitempty"`
	Description string                   `json:"description,omitempty"`
	Headers     map[string]SwaggerHeader `json:"headers,omitempty"`
	Schema      SwaggerSchema            `json:"schema,omitempty"`
	Examples    []string                 `json:"examples,omitempty"`
}

// SwaggerHeader is a holder object used to define the swagger spec and serialize to JSON
type SwaggerHeader struct {
	Ref         string `json:"$ref,omitempty"`
	Type        string `json:"type,omitempty"`
	Format      string `json:"format,omitempty"`
	Description string `json:"description,omitempty"`
}

// SwaggerSecurityDefinition is a holder object used to define the swagger spec and serialize to JSON
type SwaggerSecurityDefinition struct {
	Type             string            `json:"type,omitempty"`
	Name             string            `json:"name,omitempty"`
	In               string            `json:"in,omitempty"` //query or header
	AuthorizationURL string            `json:"authorizationUrl,omitempty"`
	TokenURL         string            `json:"tokenUrl,omitempty"`
	Flow             string            `json:"flow,omitempty"`
	Scopes           map[string]string `json:"scopes,omitempty"`
}

// SwaggerDefinition is a holder object used to define the swagger spec and serialize to JSON
type SwaggerDefinition struct {
	Type       string                               `json:"type,omitempty"`
	Properties map[string]SwaggerDefinitionProperty `json:"properties,omitempty"`
	XML        SwaggerDefinitionXML                 `json:"xml,omitempty"`
}

// AddProperty is used to add a property to a swagger route definition
func (definition *SwaggerDefinition) AddProperty(name string, prop SwaggerDefinitionProperty) *SwaggerDefinition {
	if definition.Properties == nil {
		definition.Properties = make(map[string]SwaggerDefinitionProperty)
	}
	definition.Properties[name] = prop
	return definition
}

// SwaggerDefinitionProperty is a holder object used to define the swagger spec and serialize to JSON
type SwaggerDefinitionProperty struct {
	Ref              string                      `json:"$ref,omitempty"`
	Type             string                      `json:"type,omitempty"`
	Format           string                      `json:"format,omitempty"`
	Items            []SwaggerDefinitionProperty `json:"items,omitempty"`
	Default          string                      `json:"default,omitempty"`
	Maximum          float32                     `json:"maximum,omitempty"`
	ExclusiveMaximum bool                        `json:"exclusiveMaximum,omitempty"`
	Minimum          float32                     `json:"minimum,omitempty"`
	MaxLength        int                         `json:"maxLength,omitempty"`
	MinLength        int                         `json:"minLength,omitempty"`
	Pattern          string                      `json:"pattern,omitempty"`
	MaxItems         string                      `json:"maxItems,omitempty"`
	MinItems         string                      `json:"minItems,omitempty"`
	UniqueItems      bool                        `json:"uniqueItems,omitempty"`
	MultipleOf       float32                     `json:"multipleOf,omitempty"`
	Enum             []string                    `json:"enum,omitempty"`
}

// SwaggerDefinitionXML is a holder object used to define the swagger spec and serialize to JSON
type SwaggerDefinitionXML struct {
	Name      string `json:"name,omitempty"`
	Namespace string `json:"namespace,omitempty"`
	Prefix    string `json:"prefix,omitempty"`
	Attribute bool   `json:"attribute,omitempty"`
	Wrapped   bool   `json:"wrapped,omitempty"`
}
