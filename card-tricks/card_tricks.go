package cards

// FavoriteCards returns a slice with the cards 2, 6 and 9 in that order.
func FavoriteCards() []int {
	return []int{2, 6, 9}
}

// GetItem retrieves an item from a slice at given position.
// If the index is out of range, we want it to return -1.
func GetItem(slice []int, index int) int {
	if len(slice) > index && !(index < 0) {
		return slice[index]
	} else {
		return -1
	}
}

// SetItem writes an item to a slice at given position overwriting an existing value.
// If the index is out of range the value needs to be appended.
func SetItem(slice []int, index, value int) []int {
	if len(slice) > index && !(index < 0) {
		slice[index] = value
		return slice
	} else {
		slice = append(slice, value)
		return slice
	}
}

// PrependItems adds an arbitrary number of values at the front of a slice.
func PrependItems(slice []int, values ...int) []int {

	var preSlice []int
	preSlice = append(preSlice, values...)

	slice = append(preSlice, slice...)
	return slice
}

// RemoveItem removes an item from a slice by modifying the existing slice.
func RemoveItem(slice []int, index int) []int {
	var preSlice []int
	var postSlice []int
	if len(slice) > index && index >= 0 {
		preSlice = slice[:index]
		postSlice = slice[index+1:]
	} else {
		return slice
	}

	preSlice = append(preSlice, postSlice...)
	return preSlice

}
