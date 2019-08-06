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
)

const (
	CACHE_EXPIRED      = 60 // minutes
	CACHE_CLEAN        = 7  // seconds
	CACHE_KEY_LOCATION = "location"

	LOCATION_API = "https://www.polarsteps.com/api/users/byusername/JorgeChato"

	GITHUB_USER  = "jorgechato"
	GITHUB_TOPIC = "public"
)
