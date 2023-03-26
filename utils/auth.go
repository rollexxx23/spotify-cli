package utils

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/brianstrauch/spotify"
	"github.com/pkg/browser"
	"github.com/rollexxx23/spotify-cli/constants"

	"github.com/spf13/viper"
)

func Login() (*spotify.Token, error) {
	verifier, challenge, err := spotify.CreatePKCEVerifierAndChallenge()
	if err != nil {
		return nil, err
	}

	state, err := spotify.GenerateRandomState()
	if err != nil {
		return nil, err
	}

	scopes := []string{spotify.ScopePlaylistReadPrivate, spotify.ScopeUserLibraryModify, spotify.ScopeUserModifyPlaybackState, spotify.ScopeUserReadPlaybackState}
	uri := spotify.BuildPKCEAuthURI(constants.ClientID, constants.RedirectURI, challenge, state, scopes...)

	if err := browser.OpenURL(uri); err != nil {
		panic(err)
	}

	code, err := ListenForCode(state)
	if err != nil {
		return nil, err
	}

	token, err := spotify.RequestPKCEToken(constants.ClientID, code, constants.RedirectURI, verifier)
	if err != nil {
		return nil, err
	}

	return token, nil
}

func ListenForCode(state string) (code string, err error) {
	server := &http.Server{Addr: ":1024"}

	http.HandleFunc("/callback", func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Query().Get("state") != state || r.URL.Query().Get("error") != "" {
			err = errors.New("authorization failed")
			fmt.Fprintln(w, "Failure. Please Try Again.")
		} else {
			code = r.URL.Query().Get("code")
			fmt.Fprintln(w, "Success!, Now You Close The Browser Window.")
		}

		go func() {
			server.Shutdown(context.Background())
		}()
	})

	if err := server.ListenAndServe(); err != http.ErrServerClosed {
		return "", err
	}

	return
}

// save token to configuration file

func SaveTokens(token *spotify.Token) error {

	expiresIn := time.Now().Unix() + int64(token.ExpiresIn)
	viper.Set("expires_in", expiresIn)
	encryptedAccessToken := encrypt([]byte("asuperstrong32bitpasswordgohere!"), token.AccessToken)
	viper.Set("access_token", encryptedAccessToken)

	encryptedRefreshToken := encrypt([]byte("asuperstrong32bitpasswordgohere!"), token.RefreshToken)
	viper.Set("refresh_token", encryptedRefreshToken)
	err := viper.WriteConfig()
	if err != nil {
		return err
	}
	return nil
}

// get spotify API

func GetSpotifyAPI() (*spotify.API, error) {
	if time.Now().Unix() > viper.GetInt64("expiration") {
		if err := refresh(); err != nil {
			return nil, err
		}
	}

	encryptedToken := viper.GetString("access_token")
	token := decrypt([]byte("asuperstrong32bitpasswordgohere!"), encryptedToken)
	if token == "" {
		return nil, errors.New("Not Logged In. You Need To Login First")
	}

	return spotify.NewAPI(token), nil
}

// refresh

func refresh() error {
	encryptedRefreshToken := viper.GetString("refresh_token")
	refresh := decrypt([]byte("asuperstrong32bitpasswordgohere!"), encryptedRefreshToken)
	token, err := spotify.RefreshPKCEToken(refresh, constants.ClientID)
	if err != nil {
		return err
	}

	return SaveTokens(token)
}
