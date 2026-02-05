// cloth.go
package entities

type Cloth struct {
	IDCloth    int32   `json:"id_clothes" gorm:"column:id_clothes;primaryKey;autoIncrement"`
	Name       string  `json:"name" gorm:"column:name;not null"`
	Description *string `json:"description,omitempty" gorm:"column:description"`
	Size       *string `json:"size,omitempty" gorm:"column:size"`
	Price      *float64 `json:"price,omitempty" gorm:"column:price;type:decimal(10,2)"`
	Stock      *int32  `json:"stock,omitempty" gorm:"column:stock;default:0"`
	ImageURL   *string `json:"image_url,omitempty" gorm:"column:imagen_url"`
}

// Setters
func (c *Cloth) SetIDCloth(id int32) {
	c.IDCloth = id
}

func (c *Cloth) SetName(name string) {
	c.Name = name
}

func (c *Cloth) SetDescription(description string) {
	c.Description = &description
}

func (c *Cloth) SetSize(size string) {
	c.Size = &size
}

func (c *Cloth) SetPrice(price float64) {
	c.Price = &price
}

func (c *Cloth) SetStock(stock int32) {
	c.Stock = &stock
}

func (c *Cloth) SetImageURL(imageURL string) {
	c.ImageURL = &imageURL
}

// Getters
func (c *Cloth) GetIDCloth() int32 {
	return c.IDCloth
}

func (c *Cloth) GetName() string {
	return c.Name
}

func (c *Cloth) GetDescription() string {
	if c.Description == nil {
		return ""
	}
	return *c.Description
}

func (c *Cloth) GetSize() string {
	if c.Size == nil {
		return ""
	}
	return *c.Size
}

func (c *Cloth) GetPrice() float64 {
	if c.Price == nil {
		return 0.0
	}
	return *c.Price
}

func (c *Cloth) GetStock() int32 {
	if c.Stock == nil {
		return 0
	}
	return *c.Stock
}

func (c *Cloth) GetImageURL() string {
	if c.ImageURL == nil {
		return ""
	}
	return *c.ImageURL
}

// Constructor básico con campos requeridos
func NewCloth(name string) *Cloth {
	return &Cloth{
		Name: name,
	}
}

// Constructor completo
func NewClothFull(
	name string,
	description *string,
	size *string,
	price *float64,
	stock *int32,
	imageURL *string,
) *Cloth {
	return &Cloth{
		Name:        name,
		Description: description,
		Size:        size,
		Price:       price,
		Stock:       stock,
		ImageURL:    imageURL,
	}
}