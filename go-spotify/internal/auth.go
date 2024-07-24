package internal

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"net/url"
	"os"
	"strings"

	"github.com/zeindevs/gospotify/config"
	"github.com/zeindevs/gospotify/util"
)

type AuthService struct {
	http *Http
	cfg  *config.Config
}

func NewAuthService(cfg *config.Config) *AuthService {
	return &AuthService{
		http: NewHttp(),
		cfg:  cfg,
	}
}

func (as *AuthService) Login(clientID string) (string, error) {
	state, err := util.GenerateRandomID(16)
	if err != nil {
		return "", err
	}

	data := url.Values{}
	data.Set("response_type", "code")
	data.Set("client_id", clientID)
	data.Set("scope", "user-read-private user-read-email user-read-playback-state user-modify-playback-state user-read-currently-playing")
	data.Set("redirect_uri", "http://localhost:9001/callback")
	data.Set("state", state)

	url := fmt.Sprintf("https://accounts.spotify.com/authorize?%s", data.Encode())

	return url, nil
}

func (as *AuthService) ClientLogin() (any, error) {
	data := url.Values{}
	data.Set("grant_type", "client_credentials")
	data.Set("client_id", as.cfg.CLIENT_ID)
	data.Set("client_secret", as.cfg.CLIENT_SECRET)

	as.http.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	res, err := as.http.Post("https://accounts.spotify.com/api/token", strings.NewReader(data.Encode()))

	var val any
	json.Unmarshal(res, &val)

	if err := os.WriteFile("client.json", res, 06444); err != nil {
		return nil, err
	}

	return val, err
}

func (as *AuthService) Callback(code string, state string) (any, error) {
	data := url.Values{}
	data.Set("grant_type", "authorization_code")
	data.Set("code", code)
	data.Set("redirect_uri", "http://localhost:9001/callback")

	token := base64.StdEncoding.EncodeToString([]byte(as.cfg.CLIENT_ID + ":" + as.cfg.CLIENT_SECRET))

	as.http.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	as.http.Header.Add("Authorization", "Basic "+token)
	res, err := as.http.Post("https://accounts.spotify.com/api/token", strings.NewReader(data.Encode()))

	if err := os.WriteFile("secret.json", res, 06444); err != nil {
		return nil, err
	}

	var val any
	json.Unmarshal(res, &val)

	return val, err
}
