package config

import (
	"gmauth/domain"
	"os"
)

var discordOAuth *domain.DiscordOauth

func SetupDiscordOauth() {
	clientID := os.Getenv("CLIENTID")
	redirectURI := os.Getenv("REDIRECTURI")
	responseType := "code"
	scope := "identify email connections guilds"

	setDiscordOauth(&domain.DiscordOauth{ClientID: clientID, ResponseType: responseType, Scope: scope, RedirectURI: redirectURI})
}

func setDiscordOauth(DiscordOAuth *domain.DiscordOauth) {
	discordOAuth = DiscordOAuth
}

func GetDiscordOAuth() domain.DiscordOauth {
	return *discordOAuth
}
