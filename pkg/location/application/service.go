package application

import "jorgechato.com/pkg/location/domain"

type LocationUseCase struct {
	locationUseCase domain.LocationUseCase
}

func NewLocationService(service domain.LocationUseCase) LocationUseCase {
	return LocationUseCase{
		service,
	}
}
