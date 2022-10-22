# About

On deployment inject API keys authorized to use your service. Every call to a private
endpoint of your service has to include a `header['x-api-key']` attribute that is
validated against the API keys (starting with: `API_KEY_`) in your environment.
If it is present, a request is authorized. If it is not fiber returns `401 Unauthorized`.
Use this either as a middleware the usage.

## Built With

- [fiber](https://github.com/gofiber/fiber/v2)

<!-- GETTING STARTED -->

## Getting Started

### Installation

```sh
go get github.com/iwpnd/fiber-key-auth
```

## Usage

As Middleware:

```go
package main

import (
  "os"

  "github.com/iwpnd/fiber-key-auth"
  "github.com/gofiber/fiber/v2"
)

os.Setenv("API_KEY_TEST", "valid")

func main() {
    app := fiber.New()

    app.Use(keyauth.New())

    // or optionally with structured error message
    // app.Use(keyauth.New(WithStructuredErrorMsg()))

    app.Get("/", func(c *fiber.Ctx) error {
        return c.SendString("Hello, World 👋!")
    })

    app.Listen(":3000")
}
```

Now try to access your `/` route.

```bash
curl localhost:3000

>> "no api key" # or {"message": "no api key"}

curl localhost:3000 -H "x-api-key: invalid"

>> "invalid api key" # or {"message": "invalid api key"}

curl localhost:3000 -H "x-api-key: valid"

>> "Hello, World 👋!"
```

## License

Distributed under the MIT License. See `LICENSE` for more information.

<!-- CONTACT -->

## Contact

Benjamin Ramser - [@imwithpanda](https://twitter.com/imwithpanda) - ahoi@iwpnd.pw  
Project Link: [https://github.com/iwpnd/fiber-key-auth](https://github.com/iwpnd/fiber-key-auth)
