/*
Package main circle test  API

This application provides a RESTful API for the Dcircle test.

Security:
	- api_key:

SecurityDefinitions:
    api_key:
        type: apiKey
        name: Authorization
        in: header
	Version:   1.0.0
	Host:      ERTAPI
	License: ERT Proprietary API

Consumes:
	- application/json

Produces:
	- application/json

swagger:meta
*/
package main

import "github.com/ert4web/swagger-test/app"

//go:generate swagger generate spec -i ../tags.json -o ../swagger.json
func main() {
	app.GetUserHandlerFunc()
}
