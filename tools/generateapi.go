package tools

import (
	"github.com/Mrzrb/astra"
	"github.com/Mrzrb/astra/astTraversal"
	"github.com/Mrzrb/astra/inputs"
	"github.com/Mrzrb/astra/outputs"
	"github.com/gin-gonic/gin"
)

func GenerateDoc(engine *gin.Engine) error {
	gen := getGen(engine)
	config := astra.Config{
		Title:   "Example API",
		Version: "1.0.0",
		Host:    "localhost",
		Port:    8080,
	}
	gen.SetConfig(&config)
	err := gen.Parse()
	if err != nil {
		panic(err)
	}

	return err
}

func getGen(engine *gin.Engine) *astra.Service {
	gen := astra.New(
		inputs.WithGinInput(engine),
		outputs.WithOpenAPIOutput("openapi.generated.yaml"),
		astra.WithCustomFunc(func(contextVarName string, contextFuncBuilder *astra.ContextFuncBuilder) (*astra.Route, error) {
			funcType, err := contextFuncBuilder.Traverser.Type()
			if err != nil {
				return nil, err
			}

			// handleError(c, statusCode, err)
			if funcType.Name() == "RenderJsonFail" {
				// Because we know explicitly that the first argument is the context, we can ignore it
				// We only need to concern ourselves with the status code, which is a unique case
				// We can also ignore the error, as it will be parsed as a string
				return contextFuncBuilder.Ignored().ExpressionResult().Build(func(route *astra.Route, params []any) (*astra.Route, error) {
					// Params is a list of the arguments returned from the function
					// [0] is ignored
					// [1] is the status code (int)
					// [2] is the error (type)
					// We only need to concern ourselves with the status code

					returnType := astra.ReturnType{
						ContentType: "text/plain",
						StatusCode:  400,
						Field: astra.Field{
							Type: "string",
						},
					}

					// Create the return type for this explicit error code

					// A custom utility function that prevents duplicate return types
					route.ReturnTypes = astra.AddReturnType(route.ReturnTypes, returnType)
					return route, nil
				})
			}

			// handleSuccess(c, statusCode, data)
			if funcType.Name() == "RenderJsonSucc" {
				// Because we know explicitly that the first argument is the context, we can ignore it
				// We only need to concern ourselves with the status code, which is a unique case
				// We can also ignore the error, as it will be parsed as a string
				return contextFuncBuilder.Ignored().ExpressionResult().Build(func(route *astra.Route, params []any) (*astra.Route, error) {
					// Params is a list of the arguments returned from the function
					// [0] is ignored
					// [1] is the status code (int)
					// [2] is the data (type), this comes as a result type from the astTraversal package
					dataType := params[1].(astTraversal.Result)

					returnType := astra.ReturnType{
						ContentType: "application/json",
						StatusCode:  200,
						Field:       astra.ParseResultToField(dataType),
					}

					route.ReturnTypes = astra.AddReturnType(route.ReturnTypes, returnType)

					return route, nil
				})
			}

			// When the function is not one of our custom functions, we return nil, nil
			// This will tell Astra to continue with the default behaviour for this function
			return nil, nil
		}),
	)
	return gen
}
