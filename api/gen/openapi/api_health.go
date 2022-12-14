/*
 * igusaya_blog
 *
 * 個人用内製Blog
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

import (
	"net/http"
	"strings"
)

// HealthApiController binds http requests to an api service and writes the service results to the http response
type HealthApiController struct {
	service      HealthApiServicer
	errorHandler ErrorHandler
}

// HealthApiOption for how the controller is set up.
type HealthApiOption func(*HealthApiController)

// WithHealthApiErrorHandler inject ErrorHandler into controller
func WithHealthApiErrorHandler(h ErrorHandler) HealthApiOption {
	return func(c *HealthApiController) {
		c.errorHandler = h
	}
}

// NewHealthApiController creates a default api controller
func NewHealthApiController(s HealthApiServicer, opts ...HealthApiOption) Router {
	controller := &HealthApiController{
		service:      s,
		errorHandler: DefaultErrorHandler,
	}

	for _, opt := range opts {
		opt(controller)
	}

	return controller
}

// Routes returns all the api routes for the HealthApiController
func (c *HealthApiController) Routes() Routes {
	return Routes{
		{
			"HealthGet",
			strings.ToUpper("Get"),
			"/health",
			c.HealthGet,
		},
	}
}

// HealthGet -
func (c *HealthApiController) HealthGet(w http.ResponseWriter, r *http.Request) {
	result, err := c.service.HealthGet(r.Context())
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, w)

}
