package lib

import (
	"context"
	"fmt"
	"net/http"

	"golang.org/x/oauth2"
)

// Handle callback
func Callback(w http.ResponseWriter, r *http.Request) {

	var oauth2Config oauth2.Config

	// Extract the code and state from the query parameters
	code := r.URL.Query().Get("code")
	if code == "" {
		http.Error(w, "Missing code", http.StatusBadRequest)
		return
	}

	// Exchange the authorization code for tokens
	token, err := oauth2Config.Exchange(context.Background(), code)
	if err != nil {
		http.Error(w, "Failed to exchange token: "+err.Error(), http.StatusInternalServerError)
		return
	}

	// Token received: Access and ID token can be used to authenticate
	fmt.Fprintf(w, "Access Token: %s\n", token.AccessToken)
	fmt.Fprintf(w, "ID Token: %s\n", token.Extra("id_token"))
}
