package location

type Polarsteps struct {
	Alltrips              []Alltrips     `json:"alltrips"`
	CountryCount          int            `json:"country_count"`
	CountryCountFollowers int            `json:"country_count_followers"`
	CountryCountPublic    int            `json:"country_count_public"`
	CreationDate          string         `json:"creation_date"`
	Currency              string         `json:"currency"`
	Description           string         `json:"description"`
	Email                 string         `json:"email"`
	FbID                  interface{}    `json:"fb_id"`
	FbToken               interface{}    `json:"fb_token"`
	FirstName             string         `json:"first_name"`
	FollowRequests        []interface{}  `json:"follow_requests"`
	Followees             []interface{}  `json:"followees"`
	Followers             []interface{}  `json:"followers"`
	ID                    int            `json:"id"`
	LastModified          interface{}    `json:"last_modified"`
	LastName              string         `json:"last_name"`
	LivingLocation        LivingLocation `json:"living_location"`
	LivingLocationID      int            `json:"living_location_id"`
	LivingLocationName    string         `json:"living_location_name"`
	Locale                string         `json:"locale"`
	ProfileImagePath      string         `json:"profile_image_path"`
	ProfileImageThumbPath string         `json:"profile_image_thumb_path"`
	SavedSpots            []interface{}  `json:"saved_spots"`
	SentFollowRequests    []interface{}  `json:"sent_follow_requests"`
	Stats                 struct {
		Continents                    []string `json:"continents"`
		CountryCodes                  []string `json:"country_codes"`
		FurthestPlaceFromHomeCountry  string   `json:"furthest_place_from_home_country"`
		FurthestPlaceFromHomeKm       int      `json:"furthest_place_from_home_km"`
		FurthestPlaceFromHomeLocation string   `json:"furthest_place_from_home_location"`
		KmCount                       float64  `json:"km_count"`
		LastTripEndDate               float64  `json:"last_trip_end_date"`
		LikeCount                     float64  `json:"like_count"`
		StepCount                     float64  `json:"step_count"`
		TimeTraveledInSeconds         float64  `json:"time_traveled_in_seconds"`
		TripCount                     float64  `json:"trip_count"`
		WorldPercentage               int      `json:"world_percentage"`
	} `json:"stats"`
	Synchronized         bool        `json:"synchronized"`
	TemperatureIsCelsius bool        `json:"temperature_is_celsius"`
	TravellerType        interface{} `json:"traveller_type"`
	UnitIsKm             bool        `json:"unit_is_km"`
	Username             string      `json:"username"`
	UUID                 string      `json:"uuid"`
	Visibility           int         `json:"visibility"`
}

type Alltrips struct {
	AllSteps []struct {
		Location  LivingLocation `json:"location"`
		EndDate   float64        `json:"end_date"`
		StartDate float64        `json:"start_date"`
	} `json:"all_steps"`
	CoverPhoto                  interface{} `json:"cover_photo"`
	CoverPhotoPath              interface{} `json:"cover_photo_path"`
	CoverPhotoThumbPath         interface{} `json:"cover_photo_thumb_path"`
	EndDate                     float64     `json:"end_date"`
	FbPublishStatus             interface{} `json:"fb_publish_status"`
	FeatureDate                 interface{} `json:"feature_date"`
	FeatureText                 interface{} `json:"feature_text"`
	Featured                    interface{} `json:"featured"`
	FeaturedPriorityForNewUsers interface{} `json:"featured_priority_for_new_users"`
	ID                          int         `json:"id"`
	IsDeleted                   bool        `json:"is_deleted"`
	Language                    interface{} `json:"language"`
	Likes                       int         `json:"likes"`
	Name                        string      `json:"name"`
	OpenGraphID                 interface{} `json:"open_graph_id"`
	Slug                        string      `json:"slug"`
	StartDate                   float64     `json:"start_date"`
	StepCount                   int         `json:"step_count"`
	Summary                     string      `json:"summary"`
	TimezoneID                  string      `json:"timezone_id"`
	TotalKm                     float64     `json:"total_km"`
	Type                        interface{} `json:"type"`
	UserID                      int         `json:"user_id"`
	UUID                        string      `json:"uuid"`
	Views                       int         `json:"views"`
	Visibility                  int         `json:"visibility"`
}

type LivingLocation struct {
	Accuracy     interface{} `json:"accuracy"`
	CountryCode  string      `json:"country_code"`
	Detail       string      `json:"detail"`
	Elevation    interface{} `json:"elevation"`
	FullDetail   string      `json:"full_detail"`
	ID           int         `json:"id"`
	IsDeleted    bool        `json:"is_deleted"`
	LastModified interface{} `json:"last_modified"`
	Lat          float64     `json:"lat"`
	Lon          float64     `json:"lon"`
	Name         string      `json:"name"`
	Precision    interface{} `json:"precision"`
	Synchronized bool        `json:"synchronized"`
	UUID         string      `json:"uuid"`
	Venue        string      `json:"venue"`
}
