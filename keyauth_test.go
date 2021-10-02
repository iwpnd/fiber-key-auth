package keyauth

import (
	"net/http/httptest"
	"os"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
)

func TestRoute(t *testing.T) {
	os.Setenv("API_KEY_TEST", "valid")

	tests := []struct {
		description  string
		route        string
		expectedCode int
		apiKey       string
	}{
		{
			description:  "get HTTP status 403 when invalid key",
			route:        "/",
			expectedCode: 401,
			apiKey:       "invalid",
		},
		{
			description:  "get HTTP status 403 when no key",
			route:        "/",
			expectedCode: 401,
		},
		{
			description:  "get HTTP status 200 when authorized",
			route:        "/",
			expectedCode: 200,
			apiKey:       "valid",
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

		assert.Equalf(t, test.expectedCode, resp.StatusCode, test.description)
	}
}
