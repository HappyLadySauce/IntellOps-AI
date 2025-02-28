package gtag

const (
	Default           = "default"      // Default value tag of struct field for receiving parameters from HTTP request.
	DefaultShort      = "d"            // Short name of Default.
	Param             = "param"        // Parameter name for converting certain parameter to specified struct field.
	ParamShort        = "p"            // Short name of Param.
	Valid             = "valid"        // Validation rule tag for struct of field.
	ValidShort        = "v"            // Short name of Valid.
	NoValidation      = "nv"           // No validation for specified struct/field.
	ORM               = "orm"          // ORM tag for ORM feature, which performs different features according scenarios.
	Arg               = "arg"          // Arg tag for struct, usually for command argument option.
	Brief             = "brief"        // Brief tag for struct, usually be considered as summary.
	Root              = "root"         // Root tag for struct, usually for nested commands management.
	Additional        = "additional"   // Additional tag for struct, usually for additional description of command.
	AdditionalShort   = "ad"           // Short name of Additional.
	Path              = `path`         // Route path for HTTP request.
	Method            = `method`       // Route method for HTTP request.
	Domain            = `domain`       // Route domain for HTTP request.
	Mime              = `mime`         // MIME type for HTTP request/response.
	Consumes          = `consumes`     // MIME type for HTTP request.
	Summary           = `summary`      // Summary for struct, usually for OpenAPI in request struct.
	SummaryShort      = `sm`           // Short name of Summary.
	SummaryShort2     = `sum`          // Short name of Summary.
	Description       = `description`  // Description for struct, usually for OpenAPI in request struct.
	DescriptionShort  = `dc`           // Short name of Description.
	DescriptionShort2 = `des`          // Short name of Description.
	Example           = `example`      // Example for struct, usually for OpenAPI in request struct.
	ExampleShort      = `eg`           // Short name of Example.
	Examples          = `examples`     // Examples for struct, usually for OpenAPI in request struct.
	ExamplesShort     = `egs`          // Short name of Examples.
	ExternalDocs      = `externalDocs` // External docs for struct, always for OpenAPI in request struct.
	ExternalDocsShort = `ed`           // Short name of ExternalDocs.
	GConv             = "gconv"        // GConv defines the converting target name for specified struct field.
	GConvShort        = "c"            // GConv defines the converting target name for specified struct field.
	Json              = "json"         // Json tag is supported by stdlib.
	Security          = "security"     // Security defines scheme for authentication. Detail to see https://swagger.io/docs/specification/authentication/
	In                = "in"           // Swagger distinguishes between the following parameter types based on the parameter location. Detail to see https://swagger.io/docs/specification/describing-parameters/
)

// StructTagPriority defines the default priority tags for Map*/Struct* functions.
// Note that, the `gconv/param` tags are used by old version of package.
// It is strongly recommended using short tag `c/p` instead in the future.
var StructTagPriority = []string{
	GConv, Param, GConvShort, ParamShort, Json,
}
