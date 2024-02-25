package gross

// Units stores the Gross Store unit measurements.
func Units() map[string]int {
	gross := make(map[string]int)
	gross["quarter_of_a_dozen"] = 3
	gross["half_of_a_dozen"] = 6
	gross["dozen"] = 12
	gross["small_gross"] = 120
	gross["gross"] = 144
	gross["great_gross"] = 1728

	return gross
}

// NewBill creates a new bill.
func NewBill() map[string]int {
	bill := make(map[string]int)
	return bill
}

// AddItem adds an item to customer bill.
func AddItem(bill, units map[string]int, item, unit string) bool {
	grossValue, exists := units[unit]
	if exists {
		billValue, billExists := bill[item]
		if billExists {
			bill[item] = grossValue + billValue
		} else {
			bill[item] = grossValue
		}
		return true
	} else {
		return false
	}
}

// RemoveItem removes an item from customer bill.
func RemoveItem(bill, units map[string]int, item, unit string) bool {
	unitVal, ok := units[unit]
	if !ok {
		return false
	}
	current, ok := bill[item]
	if !ok {
		return false
	}
	newVal := current - unitVal
	if newVal < 0 {
		return false
	} else if newVal == 0 {
		delete(bill, item)
	} else {
		bill[item] = newVal
	}
	return true
}

// GetItem returns the quantity of an item that the customer has in his/her bill.
func GetItem(bill map[string]int, item string) (int, bool) {
	b, ok := bill[item]
	if !ok {
		b = 0
	}
	return b, ok
}
