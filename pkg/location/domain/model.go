package domain

type (
	Location struct {
		Base     Metadata `json:"base_location" polarsteps:"living_location"`
		Current  Metadata `json:"current_location,omitempty"`
		Next     []Trip   `json:"next_locations,omitempty"`
		AllTrips []Trip   `json:"-" polarsteps:"alltrips,omitempty"`

		Stats struct {
			Continents []string `json:"continents" polarsteps:"continents"`
			Countries  []string `json:"countries" polarsteps:"country_codes"`
		} `json:"stats" polarsteps:"stats"`
	}

	Trip struct {
		Name         string  `json:"name" polarsteps:"name"`
		Thumbnail    string  `json:"thumbnail,omitempty" polarsteps:"cover_photo_path" polarsteps_planned:"cover_photo_path"`
		TimezoneID   string  `json:"timezone_id" polarsteps:"timezone_id" polarsteps_planned:"timezone_id"`
		ID           int     `json:"-" polarsteps:"id"`
		CheckIn      float64 `json:"check_in" polarsteps:"start_date" polarsteps_planned:"start_date"`
		CheckOut     float64 `json:"check_out,omitempty" polarsteps:"end_date,omitempty"`
		Steps        []Steps `json:"steps" polarsteps:"all_steps"`
		PlannedSteps []Steps `json:"-" polarsteps_planned:"planned_steps"`
	}

	Steps struct {
		TimezoneID string   `json:"timezone_id,omitempty" polarsteps:"timezone_id" polarsteps_planned:"timezone_id"`
		Metadata   Metadata `json:"metadata" polarsteps:"location" polarsteps_planned:"location"`
		CheckIn    float64  `json:"check_in" polarsteps:"start_date" polarsteps_planned:"start_time"`
		CheckOut   float64  `json:"check_out,omitempty" polarsteps:"end_date,omitempty" polarsteps_planned:"-"`
	}

	Metadata struct {
		City        string  `json:"city" polarsteps:"name" polarsteps_planned:"name"`
		Country     string  `json:"country" polarsteps:"detail" polarsteps_planned:"detail"`
		CountryCode string  `json:"country_code" polarsteps:"country_code" polarsteps_planned:"country_code"`
		Lat         float64 `json:"lat" polarsteps:"lat" polarsteps_planned:"lat"`
		Lon         float64 `json:"lon" polarsteps:"lon" polarsteps_planned:"lon"`
		Thumbnail   string  `json:"thumbnail,omitempty"`
	}
)
