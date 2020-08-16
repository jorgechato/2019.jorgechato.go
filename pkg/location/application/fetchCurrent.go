package application

import (
	"time"

	"jorgechato.com/pkg/location/domain"
)

func (l *LocationUseCase) FetchCurrent(input *domain.Location) *domain.Metadata {
	now := time.Now()

	if len(input.AllTrips) == 0 {
		return &input.Base
	}

	for _, trip := range input.AllTrips {
		timeZone, _ := time.LoadLocation(trip.TimezoneID)

		if int64(trip.CheckIn) < now.In(timeZone).Unix() && (&trip.CheckOut != nil || int64(trip.CheckOut) > now.In(timeZone).Unix()) {
			// current trip
			input.Current = trip.Steps[len(trip.Steps)-1].Metadata
		} else {
			input.Current = input.Base
		}
	}

	return &input.Current
}
