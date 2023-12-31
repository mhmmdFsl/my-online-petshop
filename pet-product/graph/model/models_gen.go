// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

type NewProduct struct {
	Name     string `json:"name"`
	Price    int    `json:"price"`
	ImageURL string `json:"imageUrl"`
}

type Product struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	Price     int    `json:"price"`
	ImageURL  string `json:"imageUrl"`
	CreatedAt string `json:"createdAt"`
	UpdatedAt string `json:"updatedAt"`
}

type UpdateProduct struct {
	ID       string  `json:"id"`
	Name     string  `json:"name"`
	Price    *int    `json:"price,omitempty"`
	ImageURL *string `json:"imageUrl,omitempty"`
}
