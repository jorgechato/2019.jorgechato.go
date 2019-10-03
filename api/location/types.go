package location

import (
	"sync"
	"time"
)

type (
	Where struct {
		Name     string     `json:"name"`
		Stats    Stats      `json:"stats"`
		Location Location   `json:"location"`
		Origin   Location   `json:"origin"`
		Next     []Location `json:"next_locations,omitempty"`
	}

	Location struct {
		City        string  `json:"city,omitempty"`
		Country     string  `json:"country"`
		CountryCode string  `json:"country_code,omitempty"`
		Lat         float64 `json:"lat,omitempty"`
		Lon         float64 `json:"lon,omitempty"`
		Url         string  `json:"url,omitempty"`
		Thumbnail   string  `json:"thumbnail,omitempty"`
		Next        string  `json:"next,omitempty"`
		In          int64   `json:"check_in,omitempty"`
		Out         int64   `json:"check_out,omitempty"`
		TimezoneID  string  `json:"timezone_id"`
	}

	Stats struct {
		Continents []string `json:"continents"`
		Countries  []string `json:"countries"`
	}
)

func (w *Where) build(p Polarsteps) {
	wg := sync.WaitGroup{}
	now := time.Now()

	w.Name = p.FirstName

	w.Stats = Stats{
		Continents: p.Stats.Continents,
		Countries:  p.Stats.CountryCodes,
	}

	w.Location = Location{
		City:    p.LivingLocation.Name,
		Country: p.LivingLocation.Detail,
		Lat:     p.LivingLocation.Lat,
		Lon:     p.LivingLocation.Lon,
	}

	w.Origin = Location{
		Lat: p.LivingLocation.Lat,
		Lon: p.LivingLocation.Lon,
	}

	for _, location := range p.Alltrips {
		timeZone, _ := time.LoadLocation(location.TimezoneID)

		if int64(location.StartDate) > now.In(timeZone).Unix() {

			wg.Add(1)

			go func(location Alltrips) {
				defer wg.Done()

				next := Location{
					Country:    location.Name,
					In:         int64(location.StartDate),
					Out:        int64(location.EndDate),
					TimezoneID: location.TimezoneID,
				}

				w.Next = append(
					w.Next,
					next,
				)
			}(location)

		} else if int64(location.EndDate) > now.In(timeZone).Unix() && len(location.AllSteps) > 0 {

			step := location.AllSteps[len(location.AllSteps)-1]

			w.Location = Location{
				Country:     step.Location.Detail,
				City:        step.Location.Name,
				Lat:         step.Location.Lat,
				Lon:         step.Location.Lon,
				In:          int64(step.StartDate),
				CountryCode: step.Location.CountryCode,
				TimezoneID:  location.TimezoneID,
			}
		}
	}

	wg.Wait()
}
