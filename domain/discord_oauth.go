package domain

import (
	"fmt"
	"net/url"
)

type DiscordOauth struct {
	ClientID     string
	RedirectURI  string
	Scope        string
	ResponseType string
}

func (d DiscordOauth) GetDiscordAuthUrl(state string) string {
	params := url.Values{}
	params.Add("response_type", d.ResponseType)
	params.Add("redirect_uri", d.RedirectURI)
	params.Add("client_id", d.ClientID)
	params.Add("scope", d.Scope)
	params.Add("state", state)
	params.Add("prompt", "consent")

	return fmt.Sprintf("https://discord.com/api/oauth2/authorize?%s", params.Encode())
}
