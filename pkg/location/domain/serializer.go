package domain

type LocationSerializer interface {
	Decode(input []byte) (*Location, error)
	Encode(input *Location) ([]byte, error)
}

type TripSerializer interface {
	Decode(input []byte) (*Trip, error)
	Encode(input *Trip) ([]byte, error)
}
