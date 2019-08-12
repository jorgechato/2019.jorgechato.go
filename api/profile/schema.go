package profile

type (
	Profile struct {
		Name    string `json:"name"`
		Picture string `json:"avatar_url"`
		Company string `json:"company"`
		Bio     string `json:"bio"`
		Url     string `json:"html_url"`
		Email   string `json:"email,omitempty"`
	}
)
