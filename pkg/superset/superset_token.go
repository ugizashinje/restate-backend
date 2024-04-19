package superset

import (
	b64 "encoding/base64"
	"encoding/json"
	"fmt"
	"os"
	"strings"
	"time"
	"warrant-api/pkg/config"
	"warrant-api/pkg/messages"

	"github.com/go-resty/resty/v2"
	"gopkg.in/guregu/null.v4"
)

var Token string

const Period time.Duration = 60 * time.Second

// claimns {"fresh":true,"iat":1712656179,"jti":"846d0b3a-508b-4e45-85b1-8aea4b2bca7a","type":"access","sub":1,"nbf":1712656179,"exp":171265707

type SupersetClaims struct {
	Fresh bool   `json:"fresh"`
	Iat   int64  `json:"iat"`
	Jti   string `json:"jti"`
	Type  string `json:"type"`
	Sub   int    `json"sub"`
	Nbf   int64  `json"nbf"`
	Exp   int64  `json"exp"`
}

type SupersetAuth struct {
	Password string `json:"password"`
	Provider string `json:"provider"`
	Refresh  bool   `json:"refresh"`
	Username string `json:"username"`
}

type AuthSuccess struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}
type RefreshSuccess struct {
	AccessToken string `json:"access_token"`
}

type AuthError struct {
	Message null.String `json:"message"`
	Msg     null.String `json:"msg"`
}

func MaintainToken(c chan int) {
	client := resty.New()
	loginBody := SupersetAuth{
		Username: config.Superset.Username,
		Password: config.Superset.Password,
		Provider: "db",
		Refresh:  true,
	}
	loginUrl := config.Superset.Url + "/security/login"
	loginSuccess := AuthSuccess{}
	loginError := &AuthError{}

	refreshUrl := config.Superset.Url + "/security/refresh"
	refreshSuccess := RefreshSuccess{}
	for {

		_, err := client.R().
			EnableTrace().
			SetBody(loginBody).
			SetResult(&loginSuccess).
			SetError(&loginError).
			Post(loginUrl)
		if loginSuccess.AccessToken == "" {
			fmt.Println("Invalid superset credential")
			os.Exit(3)
		}
		Token = loginSuccess.AccessToken
		fmt.Println("login -=> ", Token)
		c <- 1

		if err = tokenValid(Token); err != nil {
			continue
		}
		for {
			if tokenValid(Token) == nil {
				fmt.Println("Token valid")
				time.Sleep(Period)
				continue
			}

			if _, err := client.R().
				EnableTrace().
				SetBody(nil).
				SetResult(&refreshSuccess).
				SetError(loginError).SetAuthToken(loginSuccess.RefreshToken).
				Post(refreshUrl); err != nil {
				break
			}
			Token = refreshSuccess.AccessToken
			fmt.Println("refreshed -=> ", Token)

		}
	}
}

func tokenValid(t string) error {
	splited := strings.Split(t, ".")
	if len(splited) != 3 {
		fmt.Println("Invalid token ", t)
		os.Exit(3)
	}
	b64claims := splited[1]
	claimsJson, err := b64.RawStdEncoding.DecodeString(b64claims)
	if err != nil {
		fmt.Print("Superset jwt error ", err.Error())
	}
	claims := SupersetClaims{}
	if err = json.Unmarshal(claimsJson, &claims); err != nil {
		fmt.Println("Error ", err.Error())
		os.Exit(5)
	}
	fmt.Println("expiru : ", time.Unix(claims.Exp, 0))
	if time.Now().After(time.Unix(claims.Exp, 0).Add(Period)) {
		return messages.Errorf(400, "Superset Token expired")
	}
	return nil
}
