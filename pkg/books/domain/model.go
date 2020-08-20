package domain

type (
	GoodreadsResponse struct {
		Books []Book `xml:"reviews>review"`
	}

	Book struct {
		Score    string `json:"score" xml:"rating"`
		Metadata struct {
			Title     string  `json:"title" xml:"title_without_series"`
			Cover     string  `json:"image_url" xml:"image_url"`
			Goodreads string  `json:"goodreads_url" xml:"link"`
			Autors    []Autor `json:"author_list" xml:"authors>author"`
		} `json:"book" xml:"book"`
	}

	Autor struct {
		Name string `json:"name" xml:"name"`
	}
)
