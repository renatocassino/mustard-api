package actions

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/gobuffalo/buffalo"
	"github.com/gobuffalo/pop"
	"github.com/tacnoman/mustard-api/dtos"
	"github.com/tacnoman/mustard-api/models"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

var (
	googleOauthConfig = &oauth2.Config{
		RedirectURL:  "http://localhost:8000/auth/callback",
		ClientID:     "831276509280-m7ckgviuile76hutdgqibe7g9ll2hhh6.apps.googleusercontent.com",
		ClientSecret: "OT4qa5gyAHDPX-nXF3IQ307A",
		Scopes: []string{
			"https://www.googleapis.com/auth/plus.me",
			"https://www.googleapis.com/auth/userinfo.email",
			"https://www.googleapis.com/auth/userinfo.profile",
		},
		Endpoint: google.Endpoint,
	}
)

// AuthHandler is a default handler to serve up
// a home page.
func AuthHandler(c buffalo.Context) error {
	url := googleOauthConfig.AuthCodeURL("randomstate")

	return c.Redirect(302, url)
}

// AuthCallbackHandler is a default handler to serve up
// a home page.
func AuthCallbackHandler(c buffalo.Context) error {
	// Get the DB connection from the context
	tx, ok := c.Value("tx").(*pop.Connection)
	if !ok {
		return fmt.Errorf("no transaction found")
	}

	accessToken, err := googleOauthConfig.Exchange(oauth2.NoContext, c.Param("code"))
	if err != nil {
		return c.Render(500, r.JSON(err))
	}

	fmt.Println(accessToken)
	resp, err := http.Get(fmt.Sprintf("https://www.googleapis.com/oauth2/v1/userinfo?access_token=%s", accessToken.AccessToken))
	if err != nil {
		fmt.Println(err.Error())
		return c.Render(500, r.JSON(err))
	}

	defer resp.Body.Close()
	content, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		return c.Render(500, r.Plain(err.Error()))
	}

	authDTO := dtos.GoogleAuthDTO{}
	json.Unmarshal(content, &authDTO)

	user := models.User{}
	user.InsertOrUpdate(&authDTO, tx)

	return c.Render(200, r.JSON(user))
}
