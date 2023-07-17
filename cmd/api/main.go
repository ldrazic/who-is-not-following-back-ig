package main

import (
	"github.com/ldrazic/who-is-not-following-back-ig/pkg/instagram"
	"github.com/ldrazic/who-is-not-following-back-ig/pkg/shared"
)

func main() {
	httpClient := shared.NewHTTPClient()
	igRepo := instagram.NewInstagramRepository(httpClient)
	igService := instagram.NewInstagramService(*igRepo)
	igService.CompareFollowingThatDoesntFollowBack()
}
