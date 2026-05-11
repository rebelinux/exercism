package cards

// FavoriteCards returns a slice with the cards 2, 6 and 9 in that order.
func FavoriteCards() []int {
	c := []int{2, 6, 9}
	return c
}

// GetItem retrieves an item from a slice at given position.
// If the index is out of range, we want it to return -1.
func GetItem(slice []int, index int) int {
	var c int

	if index >= len(slice) || index == -1 {
		c = -1
	} else {
		c = slice[index]
	}

	return c

}

// SetItem writes an item to a slice at given position overwriting an existing value.
// If the index is out of range the value needs to be appended.
func SetItem(slice []int, index, value int) []int {
	c := GetItem(slice, index)
	var n []int

	if c == -1 {
		n = append(slice, value)
	} else {
		slice[index] = value
		n = slice
	}

	return n

}

// PrependItems adds an arbitrary number of values at the front of a slice.
func PrependItems(slice []int, values ...int) []int {
	var s []int
	var v []int = values
	if len(values) == 0 {
		s = slice
	} else {
		s = append(v, slice...)
	}

	return s
}

// RemoveItem removes an item from a slice by modifying the existing slice.
func RemoveItem(slice []int, index int) []int {
	var s []int
	if index <= -1 || index >= len(slice) {
		s = slice
	} else {
		s = append(slice[:index], slice[index+1:]...)
	}

	return s

}
