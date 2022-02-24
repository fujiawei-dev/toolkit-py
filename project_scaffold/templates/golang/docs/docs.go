{{GOLANG_HEADER}}

package {{GOLANG_PACKAGE}}

import "github.com/swaggo/swag"

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0.0",
	Host:             "localhost:8080",
	BasePath:         "/api/v1",
	Schemes:          []string{"http"},
	Title:            "Swagger Example",
	Description:      "Automatically generate RESTful API documentation with Swagger 2.0 for Go.",
	InfoInstanceName: "swagger",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
