package maps

// Units store the Dozen Store unit measurement
func Units() map[string]int {
	panic("Please implement the Units() function")
}

// NewBill create a new bill
func NewBill() map[string]int {
	panic("Please implement the NewBill() function")
}

// AddItem add item to customer bill
func AddItem(bill map[string]int, item string, unit string) bool {
	panic("Please implement the AddItem() function")
}

// RemoveItem remove item from customer bill
func RemoveItem(bill map[string]int, item string, unit string) bool {
	panic("Please implement the RemoveItem() function")
}

// GetItem return the quantity of item that the customer has in his/her bill
func GetItem(bill map[string]int, item string) (int, bool) {
	panic("Please implement the GetItem() function")
}
