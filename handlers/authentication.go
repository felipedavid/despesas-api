package handlers

import (
	"context"
	"fmt"
	"net/http"

	"github.com/markbates/goth/gothic"
)

func oauthAuthentication(w http.ResponseWriter, r *http.Request) error {
	provider := r.PathValue("provider")
	r = r.WithContext(context.WithValue(context.Background(), "provider", provider))

	user, err := gothic.CompleteUserAuth(w, r)
	if err != nil {
		gothic.BeginAuthHandler(w, r)
		return nil
	}

	fmt.Printf("User: %+v", user)
	return nil
}

func oauthCallback(w http.ResponseWriter, r *http.Request) error {
	provider := r.PathValue("provider")
	r = r.WithContext(context.WithValue(context.Background(), "provider", provider))

	user, err := gothic.CompleteUserAuth(w, r)
	if err != nil {
		return err
	}

	fmt.Printf("User: %+v", user)

	return nil
}
