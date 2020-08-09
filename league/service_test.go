package league_test

import (
	"context"
	"github.com/imjoseangel/cigame/github"
	"github.com/imjoseangel/cigame/league"
	"io/ioutil"
	"testing"
)

type stubAliasService string

func (s stubAliasService) GetAlias(string) string {
	return string(s)
}

func TestGithubIntegrationsService_GetIntegrations(t *testing.T) {
	commitService := github.NewService(github.NewClient("", ioutil.Discard))
	service := league.NewService(commitService, stubAliasService("Bob"))
	_, err := service.GetStats(context.Background(), "quii", []string{"cigame"})

	if err != nil {
		t.Fatalf("Failed to get integrations %s", err)
	}
}
