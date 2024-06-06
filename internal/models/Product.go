package models

type Product struct {
	Id                  int64
	Name                string
	PhoneticName        string
	Description         string
	PhoneticDescription string
	SKU                 int64
	Tags                []string
	Dimension           ProductDimension
	TechnicalDetails    TechnicalDetail
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
