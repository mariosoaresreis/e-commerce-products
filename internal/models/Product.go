package models

import "time"

type Product struct {
	Id                  int64
	Name                string
	PhoneticName        string
	Description         string
	PhoneticDescription string
	SKU                 int64
	Departments         []int64
	Tags                []string
	Dimension           ProductDimension
	TechnicalDetails    TechnicalDetail
	CreatedAt           time.Time  `json:"createdAt"`
	UpdatedAt           *time.Time `json:"updatedAt,omitempty"`
	ArchivedAt          *time.Time `json:"archivedAt,omitempty"`
}

type ProductDimension struct {
	Height float32
	Width  float32
	Depth  float32
	Weight float32
}

type TechnicalDetail struct {
	Color               string
	Material            string
	ItemPackageQuantity int
}
