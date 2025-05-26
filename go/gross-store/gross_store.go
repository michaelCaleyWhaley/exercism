package gross

// Units stores the Gross Store unit measurements.
func Units() map[string]int {
	return map[string]int{
		"quarter_of_a_dozen": 3,
		"half_of_a_dozen":    6,
		"dozen":              12,
		"small_gross":        120,
		"gross":              144,
		"great_gross":        1728,
	}
}

// NewBill creates a new bill.
func NewBill() map[string]int {
	return map[string]int{}
}

// AddItem adds an item to customer bill.
func AddItem(bill, units map[string]int, item, unit string) bool {
	unitPrice, isExistingUnit := units[unit]
	if !isExistingUnit {
		return false
	}

	_, isOnBillAlready := bill[item]
	if isOnBillAlready {
		bill[item] += unitPrice
	} else {
		bill[item] = unitPrice
	}

	return true
}

// RemoveItem removes an item from customer bill.
func RemoveItem(bill, units map[string]int, item, unit string) bool {
	_, isOnBill := bill[item]
	quantity, isExistingUnit := units[unit]

	if !isOnBill || !isExistingUnit {
		return false
	}

	newQuantity := bill[item] - quantity
	isLessThanZero := newQuantity < 0
	if isLessThanZero {
		return false
	}

	if newQuantity == 0 {
		delete(bill, item)
		return true
	}

	bill[item] -= quantity
	return true
}

// GetItem returns the quantity of an item that the customer has in his/her bill.
func GetItem(bill map[string]int, item string) (int, bool) {
	quantity, isInBill := bill[item]

	if !isInBill {
		return 0, false
	}

	return quantity, true
}
