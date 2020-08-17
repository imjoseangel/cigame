// main declares the CLI that spins up the server of
// our API.
// It takes some arguments, validates if they're valid
// and match the expected type and then intiialize the
// server.
package main

import (
	"log"
	"os"

	"github.com/alexflint/go-arg"
	"github.com/cirocosta/hello-swagger/swagger/models"
	"github.com/cirocosta/hello-swagger/swagger/restapi"
	"github.com/cirocosta/hello-swagger/swagger/restapi/operations"
	"github.com/go-openapi/loads"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/swag"
)

// cliArgs defines the configuration that the CLI
// expects. By using a struct we can very easily
// aggregate them into an object and check what are
// the expected types.
// If we need to mock this later it's just a matter
// of reusing the struct.
type cliArgs struct {
	Port int `arg:"-p, help:port to listen to"`
}

var (
	// args is a reference to an instantiation of
	// the configuration that the CLI expects but
	// with some values set.
	// By setting some values in advance we provide
	// default values that the user might provide
	// or not.
	args = &cliArgs{
		Port: 8080,
	}
)

// main parses the arguments from the CLI as specified
// by our configuration described in `cliArgs` and then
// populates the `args` reference we defined in the `vars`
// section above.
func main() {
	arg.MustParse(args)

	swaggerSpec, err := loads.Analyzed(restapi.SwaggerJSON, "")
	if err != nil {
		log.Fatalln(err)
	}

	api := operations.NewHelloAPI(swaggerSpec)
	server := restapi.NewServer(api)
	defer server.Shutdown()

	server.Port = args.Port

	// Implement the handler functionality.
	// As all we need to do is give an implementation to the interface
	// we can just override the `api` method giving it a method with a valid
	// signature (we didn't need to have this implementation here, it could
	// even come from a different package).
	api.GetHostnameHandler = operations.GetHostnameHandlerFunc(
		func(params operations.GetHostnameParams) middleware.Responder {
			response, err := os.Hostname()
			if err != nil {
				return operations.NewGetHostnameDefault(500).WithPayload(&models.Error{
					Code:    500,
					Message: swag.String("failed to retrieve hostname"),
				})
			}

			return operations.NewGetHostnameOK().WithPayload(response)
		})

	// Start listening using having the handlers and port
	// already set up.
	if err := server.Serve(); err != nil {
		log.Fatalln(err)
	}
}
