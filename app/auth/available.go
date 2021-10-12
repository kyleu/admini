package auth

import (
	"fmt"
	"os"
	"sort"

	"github.com/kyleu/admini/app/util"
)

var (
	AvailableProviderNames map[string]string
	AvailableProviderKeys  []string
)

const (
	auth0Key         = "auth0"
	microsoftKey     = "microsoft"
	nextcloudKey     = "nextcloud"
	OpenIDConnectKey = "openid_connect"
)

func initAvailable() {
	if AvailableProviderNames == nil {
		openIDConnectName := os.Getenv(OpenIDConnectKey + "_name")
		if openIDConnectName == "" {
			openIDConnectName = "OpenID Connect"
		}
		AvailableProviderNames = map[string]string{
			"amazon": "Amazon", "apple": "Apple", auth0Key: "Auth0", "azuread": "Azure AD",
			"battlenet": "Battlenet", "bitbucket": "Bitbucket", "box": "Box",
			"dailymotion": "Dailymotion", "deezer": "Deezer", "digitalocean": "Digital Ocean", "discord": "Discord", "dropbox": "Dropbox",
			"eveonline": "Eve Online",
			"facebook":  "Facebook", "fitbit": "Fitbit",
			"gitea": "Gitea", "github": "Github", "gitlab": "Gitlab", "google": "Google", "gplus": "Google Plus",
			"heroku":    "Heroku",
			"instagram": "Instagram", "intercom": "Intercom",
			"kakao":  "Kakao",
			"lastfm": "Last FM", "line": "LINE", "linkedin": "Linkedin",
			"mastodon": "Mastodon", "meetup": "Meetup.com", microsoftKey: "Microsoft", "microsoftonline": "Microsoft Online",
			"naver": "Naver", nextcloudKey: "NextCloud",
			"okta": "Okta", "onedrive": "Onedrive", OpenIDConnectKey: openIDConnectName,
			"paypal":     "Paypal",
			"salesforce": "Salesforce", "seatalk": "SeaTalk", "shopify": "Shopify", "slack": "Slack",
			"soundcloud": "SoundCloud", "spotify": "Spotify", "steam": "Steam", "strava": "Strava", "stripe": "Stripe",
			"twitch": "Twitch", "twitter": "Twitter", "typetalk": "Typetalk",
			"uber":  "Uber",
			"vk":    "VK",
			"wepay": "Wepay",
			"xero":  "Xero",
			"yahoo": "Yahoo", "yammer": "Yammer", "yandex": "Yandex",
		}
		AvailableProviderKeys = nil
		for k := range AvailableProviderNames {
			AvailableProviderKeys = append(AvailableProviderKeys, k)
		}
		sort.Strings(AvailableProviderKeys)
	}
}

func ProviderUsage(id string, enabled bool) string {
	n, ok := AvailableProviderNames[id]
	if !ok {
		return "INVALID PROVIDER [" + id + "]"
	}
	if enabled {
		return n + " is already configured"
	}
	keys := []string{"\"" + id + "_key\"", "\"" + id + "_secret\""}
	switch id {
	case auth0Key:
		keys = append(keys, "\"auth0_domain\"")
	case microsoftKey:
		keys = append(keys, "\"microsoft_tenant\"")
	case nextcloudKey:
		keys = append(keys, "\"nextcloud_url\"")
	case OpenIDConnectKey:
		keys = append(keys, "\"openid_connect_url\"", "\"openid_connect_name\"")
	}
	return fmt.Sprintf("To enable %s, set %s as environment variables", n, util.OxfordComma(keys, "and"))
}
