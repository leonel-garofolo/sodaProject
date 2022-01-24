package app

import (
	"fmt"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert" // add Testify package
)

func TestAPIHttpRequest(t *testing.T) {
	api := Server{}
	context := api.Start()

	testRequests := []struct {
		description  string // description of the test case
		route        string // route path to test
		expectedCode int    // expected HTTP status code
	}{
		// First test case
		{
			description:  "get HTTP status 200",
			route:        "/api",
			expectedCode: 200,
		},
		// Second test case
		{
			description:  "get HTTP status 404, when route is not exists",
			route:        "/not-found",
			expectedCode: 404,
		},
	}

	fmt.Println(len(testRequests))

	// Iterate through test single test cases
	for _, test := range testRequests {
		fmt.Println("TEST: " + test.route)
		// Create a new http request with the route from the test case
		req := httptest.NewRequest("GET", "http://localhost:3000/api", nil)

		// Perform the request plain with the app,
		// the second argument is a request latency
		// (set to -1 for no latency)
		resp, _ := context.App.Test(req, 1)

		// Verify, if the status code is as expected
		assert.Equalf(t, test.expectedCode, resp.StatusCode, test.description)
	}

}
