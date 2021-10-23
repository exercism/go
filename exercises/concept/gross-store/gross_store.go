package gross

// Units stores the Gross Store unit measurements.
func Units() map[string]int {
	panic("Please implement the Units() function")
}

// NewBill creates a new bill.
func NewBill() map[string]int {
	panic("Please implement the NewBill() function")
}

// AddItem adds an item to customer bill.
func AddItem(bill, units map[string]int, item, unit string) bool {
	panic("Please implement the AddItem() function")
}

// RemoveItem removes an item from customer bill.
func RemoveItem(bill, units map[string]int, item, unit string) bool {
	panic("Please implement the RemoveItem() function")
}

// GetItem returns the quantity of an item that the customer has in his/her bill.
func GetItem(bill map[string]int, item string) (int, bool) {
	panic("Please implement the GetItem() function")
}
