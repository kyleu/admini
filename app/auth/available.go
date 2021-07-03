package auth

import (
	"fmt"
	"sort"

	"github.com/kyleu/admini/app/util"
)

var (
	AvailableProviderNames map[string]string
	AvailableProviderKeys  []string
)

func initAvailable() {
	if AvailableProviderNames == nil {
		AvailableProviderNames = map[string]string{
			"amazon": "Amazon", "apple": "Apple", "auth0": "Auth0", "azuread": "Azure AD",
			"battlenet": "Battlenet", "bitbucket": "Bitbucket", "box": "Box",
			"dailymotion": "Dailymotion", "deezer": "Deezer", "digitalocean": "Digital Ocean", "discord": "Discord", "dropbox": "Dropbox",
			"eveonline": "Eve Online",
			"facebook":  "Facebook", "fitbit": "Fitbit",
			"gitea": "Gitea", "github": "Github", "gitlab": "Gitlab", "google": "Google", "gplus": "Google Plus",
			"heroku":    "Heroku",
			"instagram": "Instagram", "intercom": "Intercom",
			"kakao":  "Kakao",
			"lastfm": "Last FM", "line": "LINE", "linkedin": "Linkedin",
			"mastodon": "Mastodon", "meetup": "Meetup.com", "microsoft": "Microsoft", "microsoftonline": "Microsoft Online",
			"naver": "Naver", "nextcloud": "NextCloud",
			"okta": "Okta", "onedrive": "Onedrive", "openid-connect": "OpenID Connect",
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
	} else {
		keys := []string{"\"" + id + "_key\"", "\"" + id + "_secret\""}
		switch id {
		case "auth0":
			keys = append(keys, "\"auth0_domain\"")
		case "microsoft":
			keys = append(keys, "\"microsoft_tenant\"")
		case "nextcloud":
			keys = append(keys, "\"nextcloud_url\"")
		}
		return fmt.Sprintf("To enable %s, set %s as environment variables", n, util.OxfordComma(keys))
	}
}