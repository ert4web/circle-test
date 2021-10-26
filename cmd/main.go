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

import "fmt"

func main() {

}

// GetUserSectionRequest structure for user request
// swagger:parameters GetUserDetails
type GetUserSectionRequest struct {
	// The resource xid for which to retrieve product data
	// in: path
	// required: true
	UserXID string `json:"userXID"`
}

// API Response Response for user deatils
// swagger:response userDetails
type UserSectionWithENT struct {
	// in: body
	Body struct {
		Name   string `json:"name,omitempty"`
		Age    int    `json:"age,omitempty"`
		UserID string `json:"user_id,omitempty"`
	}
}

// APIErrorResponse Response for errors
// swagger:response errorsResponse
type ErrorResponse struct {
	// in: body
	Body struct {
		// in: body
		// required: true
		Err string `json:"error"`
	}
}

//swagger:response GetUserResponse
type GetUserResponse struct {
}

// GetUserSection fetches a section for a user
// swagger:route GET /v1/user/sections/{userXID} UserOperations GetUserDetails
//
// Gets a section for a user from roster API
//
// ---
// Consumes:
// - application/json
// Produces:
// - application/json
// Responses:
// 	 200: userDetails
// 	 400: errorsResponse
// 	 401: errorsResponse
// 	 500: errorsResponse

func GetUserHandlerFunc() {
	fmt.Println("test handler for swagger")
}
