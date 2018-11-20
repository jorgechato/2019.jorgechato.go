package types

type User struct {
	Model
	Name           string           `json:"name"`
	Thumbmail      string           `json:"thumbnail"`
	Nick           string           `json:"nick"`
	Password       string           `json:"-"`
	Email          string           `json:"email"`
	Role           string           `json:"role"`
	SocialNetworks []*SocialNetwork `json:"social_networks"`
	Articles       []*Article       `json:"articles"`
}

type SocialNetwork struct {
	Model
	Name string `json:"name"`
	Url  string `json:"url"`
}
