package swagger

// Model is the struct of the swagger spec
type Model struct {
	Swagger             string                        `json:"swagger"`
	Info                *Info                         `json:"info"`
	Host                string                        `json:"host"`
	BasePath            string                        `json:"basePath"`
	Tags                []Tag                         `json:"tags"`
	Schemes             []string                      `json:"schemes"`
	Paths               map[string]PathMethods        `json:"paths"`
	Definitions         map[string]Definition         `json:"definitions,omitempty"`
	SecurityDefinitions map[string]SecurityDefinition `json:"securityDefinitions,omitempty"`
	ExternalDocs        *ExternalDocs                 `json:"externalDocs,omitempty"`
}

// NewSwagger creates an instance of Model
func NewSwagger(host string, basePath string) *Model {
	return &Model{
		Swagger:             "2.0",
		Host:                host,
		BasePath:            basePath,
		Tags:                []Tag{},
		Paths:               make(map[string]PathMethods),
		Definitions:         make(map[string]Definition),
		SecurityDefinitions: make(map[string]SecurityDefinition),
		Schemes:             []string{},
	}
}

// SetInfo sets the info on the swagger model.
func (s *Model) SetInfo(info *Info) *Model {
	s.Info = info
	return s
}
func contains(s []Tag, e Tag) bool {
	for _, a := range s {
		if a.Name == e.Name {
			return true
		}
	}
	return false
}

// AddTag adds a tag if it does not exist
func (s *Model) AddTag(tag Tag) *Model {
	if !contains(s.Tags, tag) {
		s.Tags = append(s.Tags, tag)
	}
	return s
}

// AddDefinition set a definition on Model.Definitions map. Name is used as key.
func (s *Model) AddDefinition(name string, definition Definition) *Model {
	s.Definitions[name] = definition
	return s
}

// SetDefinitions sets the Model.Definitions map.
func (s *Model) SetDefinitions(definitions map[string]Definition) *Model {
	s.Definitions = definitions
	return s
}

// AddPath adds PathMethods on a path's name. Supports: Get, Post, Put, Delete
func (s *Model) AddPath(name string, definition PathMethods) *Model {
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

func (s *Model) SetPaths(paths map[string]PathMethods) *Model {
	s.Paths = paths
	return s
}

// Info is a holder object used to define the swagger spec and serialize to JSON
type Info struct {
	Description    string   `json:"description"`
	Version        string   `json:"version"`
	Title          string   `json:"title"`
	TermsOfService string   `json:"termsOfService,omitempty"`
	Contact        *Contact `json:"contact,omitempty"`
	License        *License `json:"license,omitempty"`
}

// Contact is a holder object used to define the swagger spec and serialize to JSON
type Contact struct {
	Name  string `json:"name,omitempty"`
	Email string `json:"email,omitempty"`
	URL   string `json:"url,omitempty"`
}

// License is a holder object used to define the swagger spec and serialize to JSON
type License struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

// Tag is a holder object used to define the swagger spec and serialize to JSON
type Tag struct {
	Name         string        `json:"name,omitempty"`
	Description  string        `json:"description,omitempty"`
	ExternalDocs *ExternalDocs `json:"externalDocs,omitempty"`
}

// ExternalDocs is a holder object used to define the swagger spec and serialize to JSON
type ExternalDocs struct {
	Description string `json:"description,omitempty"`
	URL         string `json:"url,omitempty"`
}

// Path is a holder object used to define the swagger spec and serialize to JSON
type Path struct {
	Verbs map[string]PathItem
}

func (path *Path) AddVerb(name string, pathItem PathItem) *Path {
	if path.Verbs == nil {
		path.Verbs = make(map[string]PathItem)
	}
	path.Verbs[name] = pathItem
	return path
}

// PathItem is a holder object used to define the swagger spec and serialize to JSON
type PathItem struct {
	Tags        []string                      `json:"tags,omitempty"`
	Summary     string                        `json:"summary,omitempty"`
	Description string                        `json:"description,omitempty"`
	OperationID string                        `json:"operationId,omitempty"`
	Consumes    []string                      `json:"consumes,omitempty"`
	Produces    []string                      `json:"produces,omitempty"`
	Parameters  []PathItemParameter           `json:"parameters,omitempty"`
	Responses   map[string]PathResponse       `json:"responses,omitempty"`
	Security    map[string]SecurityDefinition `json:"security,omitempty"`
}

// AddParameter adds a PathItemParameter on a PathItems' parameter
func (pathItem *PathItem) AddParameter(parameter PathItemParameter) *PathItem {
	pathItem.Parameters = append(pathItem.Parameters, parameter)
	return pathItem
}

// PathMethods is a holder object used to define the swagger spec and serialize to JSON
type PathMethods struct {
	Post   *PathItem `json:"post,omitempty"`
	Get    *PathItem `json:"get,omitempty"`
	Put    *PathItem `json:"put,omitempty"`
	Delete *PathItem `json:"delete,omitempty"`
}

// PathItemParameter is a holder object used to define the swagger spec and serialize to JSON
type PathItemParameter struct {
	Ref              string   `json:"$ref,omitempty"`
	In               string   `json:"in,omitempty"` //query, header, path, formdata, or body
	Name             string   `json:"name,omitempty"`
	Description      string   `json:"description,omitempty"`
	Required         bool     `json:"required,omitempty"`
	Enum             []string `json:"enum,omitempty"`
	Type             string   `json:"type,omitempty"`
	Format           string   `json:"format,omitempty"`
	Schema           *Schema  `json:"schema,omitempty"`
	CollectionFormat string   `json:"collectionFormat,omitempty"`
}

// Schema is a holder object used to define the swagger spec and serialize to JSON
type Schema struct {
	Ref              string  `json:"$ref,omitempty"`
	Format           string  `json:"format,omitempty"`
	Title            string  `json:"title,omitempty"`
	Description      string  `json:"description,omitempty"`
	Default          string  `json:"default,omitempty"`
	Maximum          float64 `json:"maximum,omitempty"`
	Minimum          float64 `json:"minimum,omitempty"`
	ExclusiveMaximum bool    `json:"exclusiveMaximum,omitempty"`
}

// PathResponse is a holder object used to define the swagger spec and serialize to JSON
type PathResponse struct {
	Ref         string            `json:"$ref,omitempty"`
	Description string            `json:"description,omitempty"`
	Headers     map[string]Header `json:"headers,omitempty"`
	Schema      Schema            `json:"schema,omitempty"`
	Examples    []string          `json:"examples,omitempty"`
}

// Header is a holder object used to define the swagger spec and serialize to JSON
type Header struct {
	Ref         string `json:"$ref,omitempty"`
	Type        string `json:"type,omitempty"`
	Format      string `json:"format,omitempty"`
	Description string `json:"description,omitempty"`
}

// SecurityDefinition is a holder object used to define the swagger spec and serialize to JSON
type SecurityDefinition struct {
	Type             string            `json:"type,omitempty"`
	Name             string            `json:"name,omitempty"`
	In               string            `json:"in,omitempty"` //query or header
	AuthorizationURL string            `json:"authorizationUrl,omitempty"`
	TokenURL         string            `json:"tokenUrl,omitempty"`
	Flow             string            `json:"flow,omitempty"`
	Scopes           map[string]string `json:"scopes,omitempty"`
}

// Definition is a holder object used to define the swagger spec and serialize to JSON
type Definition struct {
	Type       string                        `json:"type,omitempty"`
	Properties map[string]DefinitionProperty `json:"properties,omitempty"`
	XML        DefinitionXML                 `json:"xml,omitempty"`
}

// AddProperty is used to add a property to a swagger route definition
func (definition *Definition) AddProperty(name string, prop DefinitionProperty) *Definition {
	if definition.Properties == nil {
		definition.Properties = make(map[string]DefinitionProperty)
	}
	definition.Properties[name] = prop
	return definition
}

// DefinitionProperty is a holder object used to define the swagger spec and serialize to JSON
type DefinitionProperty struct {
	Ref              string               `json:"$ref,omitempty"`
	Type             string               `json:"type,omitempty"`
	Format           string               `json:"format,omitempty"`
	Items            []DefinitionProperty `json:"items,omitempty"`
	Default          string               `json:"default,omitempty"`
	Maximum          float32              `json:"maximum,omitempty"`
	ExclusiveMaximum bool                 `json:"exclusiveMaximum,omitempty"`
	Minimum          float32              `json:"minimum,omitempty"`
	MaxLength        int                  `json:"maxLength,omitempty"`
	MinLength        int                  `json:"minLength,omitempty"`
	Pattern          string               `json:"pattern,omitempty"`
	MaxItems         string               `json:"maxItems,omitempty"`
	MinItems         string               `json:"minItems,omitempty"`
	UniqueItems      bool                 `json:"uniqueItems,omitempty"`
	MultipleOf       float32              `json:"multipleOf,omitempty"`
	Enum             []string             `json:"enum,omitempty"`
}

// DefinitionXML is a holder object used to define the swagger spec and serialize to JSON
type DefinitionXML struct {
	Name      string `json:"name,omitempty"`
	Namespace string `json:"namespace,omitempty"`
	Prefix    string `json:"prefix,omitempty"`
	Attribute bool   `json:"attribute,omitempty"`
	Wrapped   bool   `json:"wrapped,omitempty"`
}
