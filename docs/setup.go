package docs

import "os"

func SetupSwaggerInfo() {
	// programmatically set swagger info
	SwaggerInfo.Title = "Selling Management API"
	SwaggerInfo.Description = "Selling Management API"
	SwaggerInfo.Version = "1.0"
	SwaggerInfo.Host = os.Getenv("SwaggerHost") + ":" + os.Getenv("AppPort")
	SwaggerInfo.BasePath = "/api"
	SwaggerInfo.Schemes = []string{"http", "https"}
}
