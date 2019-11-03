package actions

import (
	"io/ioutil"
	"log"

	"github.com/gobuffalo/buffalo"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

// AuthHandler is a default handler to serve up
// a home page.
func AuthHandler(c buffalo.Context) error {
	data, err := ioutil.ReadFile("./config/google.json")
	if err != nil {
		log.Fatal(err)
	}

	conf, err := google.JWTConfigFromJSON(data, "https://www.googleapis.com/auth/bigquery")
	if err != nil {
		log.Fatal(err)
	}

	client := conf.Client(oauth2.NoContext)
	client.Get("...")
	return c.Render(200, r.JSON(map[string]string{"message": "Welcome to Buffalo!"}))
}

// AuthCallbackHandler is a default handler to serve up
// a home page.
func AuthCallbackHandler(c buffalo.Context) error {
	return c.Render(200, r.JSON(map[string]string{"message": "Welcome to Buffalo!"}))
}
