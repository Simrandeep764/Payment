package Models

type Checkout struct {
	CheckoutId uint `gorm:"primary_key;auto_increment:true" json:"pid"`
	CartId     uint `json:"cid"`
	CustomerId uint `json:"custid"`
}

type Cart struct {
	CartId     uint `gorm:"primary_key;auto_increment:false" json:"cid"`
	ProductId  int  `gorm:"primary_key;auto_increment:false;type:varchar(100)" json:"prodid"`
	CustomerId uint `json:"custid"`
	Qty        int  `json:"prodqty" gorm:"type:int"`
}

type Inventory struct {
	ProductId    uint   `gorm:"primary_key;auto_increment" json:"prodid"`
	ProductName  string `json:"prodname" gorm:"type:varchar(100)"`
	ProductQty   int    `json:"prodqty" gorm:"type:int"`
	ProductPrice int    `json:"prodprice" gorm:"type:int"`
}

func (c *Inventory) TableName() string {
	return "inventory"
}
func (c *Cart) TableName() string {
	return "cart"
}
func (p *Checkout) TableName() string {
	return "Checkout"
}
