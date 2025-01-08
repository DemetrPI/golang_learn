package cards

import "slices"

// FavoriteCards returns a slice with the cards 2, 6 and 9 in that order.
func FavoriteCards() []int {
	return []int{2, 6, 9}
}

// GetItem retrieves an item from a slice at given position.
// If the index is out of range, we want it to return -1.
func GetItem(slice []int, index int) int {
	var card int
	if index > len(slice) || index < 0 {
		card = -1
	} else {
		card = slice[index]
	}
	return card
}

// SetItem writes an item to a slice at given position overwriting an existing value.
// If the index is out of range the value needs to be appended.
func SetItem(slice []int, index, value int) []int {

	if index < 0 || index >= len(slice) {
		slice = append(slice, value)
	} else {
		slice[index] = value
	}

	return slice
}

// PrependItems adds an arbitrary number of values at the front of a slice.
func PrependItems(slice []int, values ...int) []int {
	return append(values, slice...)

}

// RemoveItem removes an item from a slice by modifying the existing slice.
func RemoveItem(slice []int, index int) []int {
	if index >= 0 && index < len(slice) {
		slice = slices.Delete(slice, index, index+1)
	}
	return slice
}
