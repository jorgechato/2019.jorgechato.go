package domain

type LocationUseCase interface {
	FetchCurrent(*Location) *Metadata
	FetchNext(*Location) *[]Trip
}
