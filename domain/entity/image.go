package entity

type Image struct {
	ID        int
	URL       string
	CreatedBy UserID
}

type ImageList []Image

func (il ImageList) GetIDs() []int {
	ids := make([]int, 0, len(il))
	for i := range il {
		ids = append(ids, il[i].ID)
	}
	return ids
}
