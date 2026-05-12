package gross

// Units stores the Gross Store unit measurements.
func Units() map[string]int {
	m := map[string]int{
		"quarter_of_a_dozen": 3,
		"half_of_a_dozen":    6,
		"dozen":              12,
		"small_gross":        120,
		"gross":              144,
		"great_gross":        1728,
	}

	return m

}

// NewBill creates a new bill.
func NewBill() map[string]int {
	m := make(map[string]int)
	return m
}

// AddItem adds an item to customer bill.
func AddItem(bill, units map[string]int, item, unit string) bool {

	_, exists := units[unit]

	if !exists {
		return false
	} else {
		_, exists := bill[item]

		if exists {
			bill[item] += units[unit]
			return true

		} else {
			bill[item] = units[unit]
			_, exists := bill[item]
			if exists {
				return true
			} else {
				return false
			}
		}
	}

}

// RemoveItem removes an item from customer bill.
func RemoveItem(bill, units map[string]int, item, unit string) bool {

	_, exists := units[unit]

	if !exists {
		return false
	} else {
		value, exists := bill[item]

		if !exists {
			return false
		} else {

			if (value - units[unit]) < 0 {
				return false
			}

			if (value - units[unit]) == 0 {
				delete(bill, item)
				return true
			} else {
				bill[item] -= units[unit]
				return true
			}
		}
	}

}

// GetItem returns the quantity of an item that the customer has in his/her bill.
func GetItem(bill map[string]int, item string) (int, bool) {
	value, exists := bill[item]

	if !exists {
		return 0, false
	} else {
		return value, true
	}
}
