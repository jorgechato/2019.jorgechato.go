package polarsteps

import (
	"fmt"
	"net/url"

	"jorgechato.com/pkg/location/application/json"
	"jorgechato.com/pkg/shared/infrastructure/api"
)

func FetchTripByID(id int) ([]byte, error) {
	u := url.URL{
		Scheme: api.HTTPS,
		Host:   BASE_API,
		Path:   fmt.Sprintf(TRIPS, id),
	}

	return api.Get(u.String())
}

func FetchTrips() ([]byte, error) {
	rawLocation, err := FetchProfile()

	jsonLocation := &json.Location{}
	location, _ := jsonLocation.Decode(rawLocation)

	for key, val := range location.AllTrips {
		rawTrip, err := FetchTripByID(val.ID)

		if err != nil {
			byteBody, _ := jsonLocation.Encode(location)
			return byteBody, err
		}

		jsonTrip := &json.Trip{Tag: TAG}
		filledTrip, _ := jsonTrip.Decode(rawTrip)
		location.AllTrips[key] = *filledTrip
	}

	byteBody, _ := jsonLocation.Encode(location)
	return byteBody, err
}
