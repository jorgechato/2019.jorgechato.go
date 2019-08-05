package project

type (
	Repositories struct {
		Total int `json:"total_count"`
		Items []struct {
			Name        string `json:"name"`
			FullName    string `json:"full_name"`
			Homepage    string `json:"homepage,omitempty"`
			Description string `json:"description"`
			CreatedAt   string `json:"created_at"`
			UpdatedAt   string `json:"updated_at"`
			Url         string `json:"html_url"`
			Language    string `json:"language"`
		} `json:"items"`
	}
)
