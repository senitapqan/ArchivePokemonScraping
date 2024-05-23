package models

type Pockemon struct {
	Name     string `json:"name"`
	Price    string `json:"price"`
	ImageUrl string `json:"image_url"`
}

type PockemonBuilder struct {
	name     string `json:"name"`
	price    string `json:"price"`
	imageUrl string `json:"image_url"`
}

func NewPockemonBuilder() *PockemonBuilder {
	return &PockemonBuilder{}
}

func (b *PockemonBuilder) SetName(name string) *PockemonBuilder {
	b.name = name
	return b
}

func (b *PockemonBuilder) SetImageUrl(imageUrl string) *PockemonBuilder {
	b.imageUrl = imageUrl
	return b
}

func (b *PockemonBuilder) SetPrice(priceAmount, priceSymbol string) *PockemonBuilder {
	b.price = priceAmount + priceSymbol
	return b
}

func (b PockemonBuilder) Build() *Pockemon {
	return &Pockemon{
		Name: b.name,
		Price: b.price,
		ImageUrl: b.imageUrl,
	}
}
