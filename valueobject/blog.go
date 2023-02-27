package valueobject

type Blog struct {
	Slug      string `json:"slug,omitempty"`
	CreatedAt string `json:"createdAt,omitempty"`
	ReadCount int    `json:"readCount,omitempty"`
}
