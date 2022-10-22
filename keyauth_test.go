package keyauth

import (
	"io"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
)

func TestRoute(t *testing.T) {
	os.Setenv("API_KEY_TEST", "valid")

	tests := []struct {
		description     string
		route           string
		expectedCode    int
		expectedMessage string
		apiKey          string
	}{
		{
			description:     "get HTTP status 403 when invalid key",
			route:           "/",
			expectedCode:    401,
			apiKey:          "invalid",
			expectedMessage: "invalid api key",
		},
		{
			description:     "get HTTP status 403 when no key",
			route:           "/",
			expectedCode:    401,
			expectedMessage: "no api key",
		},
		{
			description:     "get HTTP status 200 when authorized",
			route:           "/",
			expectedCode:    200,
			expectedMessage: "Ahoi!",
			apiKey:          "valid",
		},
	}

	app := fiber.New()
	app.Use(New())
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Ahoi!")
	})

	for _, test := range tests {
		req := httptest.NewRequest("GET", test.route, nil)

		if test.apiKey != "" {
			req.Header.Set("x-api-key", test.apiKey)
		}

		resp, _ := app.Test(req, 1)
		defer resp.Body.Close()

		message, _ := io.ReadAll(resp.Body)
		assert.Equal(t, test.expectedMessage, string(message))
		assert.Equalf(t, test.expectedCode, resp.StatusCode, test.description)
	}
}

func TestRouteWithStructuredLogs(t *testing.T) {
	os.Setenv("API_KEY_TEST", "valid")

	tests := []struct {
		description     string
		route           string
		expectedMessage string
		expectedCode    int
		apiKey          string
	}{
		{
			description:     "get HTTP status 403 when invalid key",
			route:           "/",
			expectedCode:    401,
			expectedMessage: `{"message":"invalid api key"}`,
			apiKey:          "invalid",
		},
		{
			description:     "get HTTP status 403 when no key",
			route:           "/",
			expectedCode:    401,
			expectedMessage: `{"message":"no api key"}`,
		},
		{
			description:     "get HTTP status 200 when authorized",
			route:           "/",
			expectedCode:    200,
			expectedMessage: "Ahoi!",
			apiKey:          "valid",
		},
	}

	app := fiber.New()
	app.Use(New(WithStructuredLog()))
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Ahoi!")
	})

	for _, test := range tests {
		req := httptest.NewRequest("GET", test.route, nil)

		if test.apiKey != "" {
			req.Header.Set("x-api-key", test.apiKey)
		}

		resp, _ := app.Test(req, 1)
		defer resp.Body.Close()
		message, _ := io.ReadAll(resp.Body)

		assert.Equal(t, test.expectedMessage, string(message))
		assert.Equalf(t, test.expectedCode, resp.StatusCode, test.description)
	}
}
