package main

import (
	"github.com/aquasecurity/lmdrouter"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"net/http"
	"time"

	"context"
	"log"
)

var router *lmdrouter.Router

func init() {
	//router = lmdrouter.NewRouter("/api", loggerMiddleware, authMiddleware)
	router = lmdrouter.NewRouter("/api", loggerMiddleware)
	router.Route("GET", "/", listSomethings)
	/*router.Route("POST", "/", postSomething, someOtherMiddleware)
	router.Route("GET", "/:id", getSomething)
	router.Route("PUT", "/:id", updateSomething)
	router.Route("DELETE", "/:id", deleteSomething)*/
}

func main() {
	lambda.Start(router.Handler)
}

// the rest of the code is a redacted example, it will probably reside in a
// separate package inside your project

type listSomethingsInput struct {
	ID            string `lambda:"path.id"`              // a path parameter declared as :id
	ShowSomething bool   `lambda:"query.show_something"` // a query parameter named "show_something"
}

type postSomethingInput struct {
	Title string    `json:"title"`
	Date  time.Time `json:"date"`
}

func listSomethings(ctx context.Context, req events.APIGatewayProxyRequest) (
	res events.APIGatewayProxyResponse,
	err error,
) {
	// parse input from request and path parameters
	var input listSomethingsInput
	err = lmdrouter.UnmarshalRequest(req, false, &input)
	if err != nil {
		return lmdrouter.HandleError(err)
	}

	// call some business logic that generates an output struct
	// ...

	return lmdrouter.MarshalResponse(http.StatusOK, nil, output)
}

func postSomethings(ctx context.Context, req events.APIGatewayProxyRequest) (
	res events.APIGatewayProxyResponse,
	err error,
) {
	// parse input from request body
	var input postSomethingInput
	err = lmdrouter.UnmarshalRequest(req, true, &input)
	if err != nil {
		return lmdrouter.HandleError(err)
	}

	// call some business logic that generates an output struct
	// ...

	return lmdrouter.MarshalResponse(http.StatusCreated, nil, output)
}

func loggerMiddleware(next lmdrouter.Handler) lmdrouter.Handler {
	return func(ctx context.Context, req events.APIGatewayProxyRequest) (
		res events.APIGatewayProxyResponse,
		err error,
	) {
		// [LEVEL] [METHOD PATH] [CODE] EXTRA
		format := "[%s] [%s %s] [%d] %s"
		level := "INF"
		var code int
		var extra string

		res, err = next(ctx, req)
		if err != nil {
			level = "ERR"
			code = http.StatusInternalServerError
			extra = " " + err.Error()
		} else {
			code = res.StatusCode
			if code >= 400 {
				level = "ERR"
			}
		}

		log.Printf(format, level, req.HTTPMethod, req.Path, code, extra)

		return res, err
	}
}
