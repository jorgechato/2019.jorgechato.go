package utils

var (
	VERSION string = "1.0"
	AUTHOR  string = ""
	TAG     string = ""
	PORT    string = "5000"
	HOST    string = "0.0.0.0"

	LOGPATH string = "out/api.jorgechato.com.log"

	APIBASE = "/v1"
)

var (
	// KEY Env variables
	GITHUB_TOKEN = "GITHUB_TOKEN"
	GITHUB_USER  = "GITHUB_USER"
	GITHUB_TOPIC = "GITHUB_TOPIC"
	GIST_LIST    = "GIST_LIST"

	MAILCHIMP_KEY  = "MAILCHIMP_KEY"
	MAILCHIMP_DS   = "MAILCHIMP_DS"
	MAILCHIMP_LIST = "MAILCHIMP_LIST"

	LOCATION_USER = "LOCATION_USER"
)

const (
	CACHE_EXPIRED      = 60 // minutes
	CACHE_CLEAN        = 7  // seconds
	CACHE_KEY_LOCATION = "location"

	LOCATION_API  = "https://www.polarsteps.com/api/users/byusername/%v"
	MAILCHIMP_API = "https://%v.api.mailchimp.com/3.0/lists/%v/members/%v"
)
