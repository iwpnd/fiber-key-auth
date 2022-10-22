package keyauth

import (
	"os"
	"strings"

	"github.com/gofiber/fiber/v2"
)

type handler struct {
	json bool
}

// Option type
type Option func(*handler)

// WithStructuredLog
func WithStructuredLog() Option {
	return func(h *handler) {
		h.json = true
	}
}

func getKeysInEnv() []string {
	var keys []string

	for _, e := range os.Environ() {
		pair := strings.SplitN(e, "=", 2)
		key := pair[0]
		value := pair[1]

		if strings.HasPrefix(key, "API_KEY_") {
			keys = append(keys, value)
		}
	}

	return keys
}

func keyInKeys(key string, keys []string) bool {
	for _, k := range keys {
		if key == k {
			return true
		}
	}

	return false
}

func (h *handler) keyAuth(c *fiber.Ctx) error {
	key := c.Get("x-api-key")
	if key == "" {
		if h.json {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"message": "no api key",
			})
		}
		return fiber.NewError(fiber.StatusUnauthorized, "no api key")
	}

	keys := getKeysInEnv()

	if keyInKeys(key, keys) {
		return c.Next()
	}

	if h.json {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "invalid api key",
		})
	}

	return fiber.NewError(fiber.StatusUnauthorized, "invalid api key")

}

// New exports a keyauth middleware handler
func New(options ...Option) fiber.Handler {
	h := &handler{
		json: false,
	}

	for _, option := range options {
		option(h)
	}

	return h.keyAuth
}
