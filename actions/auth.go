package actions

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
	"github.com/tacnoman/mustard-api/core"
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
func AuthHandler(c echo.Context) error {
	url := googleOauthConfig.AuthCodeURL("randomstate")

	return c.Redirect(302, url)
}

// AuthCallbackHandler is
func AuthCallbackHandler(c echo.Context) error {
	accessToken, err := googleOauthConfig.Exchange(oauth2.NoContext, c.QueryParam("code"))
	if err != nil {
		fmt.Println("Caanot get access token")
		fmt.Println(err.Error())
		return c.JSON(500, err)
	}

	fmt.Println(accessToken)
	resp, err := http.Get(fmt.Sprintf("https://www.googleapis.com/oauth2/v1/userinfo?access_token=%s", accessToken.AccessToken))
	if err != nil {
		fmt.Println(err.Error())
		return c.JSON(500, err)
	}

	defer resp.Body.Close()
	content, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		return c.JSON(500, err)
	}

	authDTO := dtos.GoogleAuthDTO{}
	json.Unmarshal(content, &authDTO)

	user := models.User{}
	user.InsertOrUpdate(&authDTO)

	mySigningKey := core.GetEnv("JWT_TOKEN", "secret")
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, user)
	tokenString, _ := token.SignedString([]byte(mySigningKey))

	return c.JSON(200, map[string]string{"token": tokenString})
}
