package entity

type Image struct {
	ID        int
	URL       string
	CreatedBy User
}

type ImageList []Image
