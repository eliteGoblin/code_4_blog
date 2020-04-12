package ginko_cart

// Cart represents the state of a buyer's shopping cart
type Cart struct {
	items []Item
	totalItems int
	totalUniqueItems int
	totalPrice float64
}

func NewCart() *Cart {
	return &Cart{
		items: make([]Item, 0),
	}
}

// Item represents any item available for sale
type Item struct {
	Name  string
	Price float64
	Qty   int
}

// AddItem adds an item to the cart
func (c *Cart) AddItem(item Item) error {
	c.items = append(c.items, item)
	c.totalItems += item.Qty
	c.totalUniqueItems += 1
	c.totalPrice += float64(item.Qty)* item.Price
	return nil
}

// TotalUnits returns the total number of units across all items in the cart
func (c *Cart) TotalItems() int {
	return c.totalItems // return a random number as a placeholder
}

// TotalUniqueItems returns the number of unique items in the cart
func (c *Cart) TotalUniqueItems() int {
	return c.totalUniqueItems
}

// TotalPrice returns the total price of items in the cart
func (c *Cart) TotalPrice() float64 {
	return c.totalPrice
}