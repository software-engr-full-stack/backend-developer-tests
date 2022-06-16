package main

import (
    "net/http"

    "github.com/aws/aws-lambda-go/lambda"
    "github.com/awslabs/aws-lambda-go-api-proxy/httpadapter"

    "github.com/software-engr-full-stack/backend-developer-tests/rest-service/pkg/handlers"
)

func main() {
    http.HandleFunc("/people/", handlers.People)

    lambda.Start(httpadapter.New(http.DefaultServeMux).ProxyWithContext)
}
