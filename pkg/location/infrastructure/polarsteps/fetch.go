package polarsteps

import "jorgechato.com/pkg/location/application/json"

func Fetch() ([]byte, error) {
	rawLocation, err := GetProfile()

	jsonLocation := &json.Location{}
	location, _ := jsonLocation.Decode(rawLocation)

	for key, val := range location.AllTrips {
		rawTrip, err := GetTripByID(val.ID)

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
