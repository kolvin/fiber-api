package main_test

import (
	"net/http/httptest"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert" // add Testify package
)

func TestExample(t *testing.T) {
	tests := []struct {
		description  string // description of the test case
		route        string // route path to test
		expectedCode int    // expected HTTP status code
	}{
		{
			description:  "get HTTP status 200",
			route:        "/",
			expectedCode: 200,
		},
		{
			description:  "get HTTP status 404, when route is not exists",
			route:        "/not-found",
			expectedCode: 404,
		},
	}

	// Define Fiber app.
	app := fiber.New()

	// Create route with GET method for test
	app.Get("/", func(c *fiber.Ctx) error {
		// Return simple string as response
		return c.SendString("Hello shit world")
	})

	for _, test := range tests {
		// Create a new http request with the route from the test case
		req := httptest.NewRequest("GET", test.route, nil)

		// Perform the request plain with the app,
		// the second argument is a request latency
		// (set to -1 for no latency)
		resp, _ := app.Test(req, 1)

		// Verify, if the status code is as expected
		assert.Equalf(t, test.expectedCode, resp.StatusCode, test.description)
	}
}
