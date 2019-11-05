package types

// APIRoute is an abstract data type for defining a route
type APIRoute struct {
	Method   string
	Path     string
	Callback *func(...interface{})
}

// APIConfig contains configuration data for the API controller
type APIConfig struct {
	BasePath string

	DBService
	AuthorRepository
	ClippingRepository
	TitleRepository
	ParserService
	UploadService
}

// APIController is the main controller for the application,
// which manages a REST server
type APIController interface {
	Setup(config *APIConfig, routes []APIRoute) error
}
