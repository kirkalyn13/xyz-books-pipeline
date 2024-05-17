package model

type (

	// Book response struct
	Book struct {
		ID              uint      `json:"id,omitempty"`
		Title           string    `json:"title,omitempty"`
		ISBN13          string    `json:"isbn13,omitempty"`
		ISBN10          string    `json:"isbn10,omitempty"`
		ListPrice       float64   `json:"list_price,omitempty"`
		PublicationYear int       `json:"publication_year,omitempty"`
		ImageURL        string    `json:"image_url,omitempty"`
		Edition         string    `json:"edition,omitempty"`
		PublisherID     *uint     `json:"publisher_id,omitempty"`
		Publisher       Publisher `json:"publisher,omitempty"`
		Authors         []*Author `json:"authors,omitempty"`
	}

	// Author response struct
	Author struct {
		ID         uint   `json:"id,omitempty"`
		FirstName  string `json:"first_name,omitempty"`
		LastName   string `json:"last_name,omitempty"`
		MiddleName string `json:"middle_name"`
	}

	// Publisher response struct
	Publisher struct {
		ID   uint   `json:"id,omitempty"`
		Name string `json:"name,omitempty"`
	}

	// BooksResponse response struct
	BooksResponse struct {
		Books []Book `json:"books"`
	}
)
