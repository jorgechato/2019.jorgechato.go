package domain

type (
	GoodreadsResponse struct {
		Books []Book `xml:"reviews>review"`
	}

	Book struct {
		Score    string `json:"score" xml:"rating"`
		Metadata struct {
			Title     string   `json:"title" xml:"title_without_series"`
			Cover     string   `json:"cover" xml:"image_url"`
			Goodreads string   `json:"goodreads_url" xml:"link"`
			Authors   []Author `json:"author_list" xml:"authors>author"`
			ISBN      string   `json:"isbn" xml:"isbn"`
		} `json:"book" xml:"book"`
	}

	Author struct {
		Name string `json:"name" xml:"name"`
	}
)
