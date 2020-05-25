package application

import "jorgechato.com/pkg/location/domain"

func (l *LocationUseCase) FetchNext(input *domain.Location) *[]domain.Trip {
	if len(input.Next) == 0 {
		input.Next = []domain.Trip{}
	}

	for _, trip := range input.AllTrips {
		if len(trip.PlannedSteps) > 0 {
			trip.Steps = trip.PlannedSteps
			input.Next = append(input.Next, trip)
		}
	}

	return &input.Next
}
