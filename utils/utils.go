package utils

// EqualIntSlices .. compares two integer slices
// returns true, if same else false
func EqualIntSlices(x, y []int) bool {

	if len(x) != len(y) {
		return false
	}

	for i, val := range x {
		if y[i] != val {
			return false
		}
	}

	return true
}

// EqualIntSlices2D .. compares two integer slices of two dimensions
// returns true, if same else false
func EqualIntSlices2D(x, y [][]int) bool {
	var arrx, arry []int

	if len(x) != len(y) {
		return false
	}

	for _, val := range x {
		for _, str := range val {
			arrx = append(arrx, str)
		}
	}

	for _, val := range y {
		for _, str := range val {
			arry = append(arry, str)
		}
	}

	for i := range x {
		if arrx[i] != arry[i] {
			return false
		}
	}

	return true
}

// IsPalindromeString ... Determine whether a string is a palindrome
func IsPalindromeString(input string) bool {
	len := len(input)
	for i := 0; i < len/2; i++ {
		if input[i] != input[len-i-1] {
			return false
		}
	}

	return true
}

// ReverseString ... Reverse a string
func ReverseString(s string) string {
	runes := []rune(s)
	len := len(s)

	for i, j := 0, len-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}

	return string(runes)
}

// MinTwoIntegers .. finds minimum of two intergers and returns minimum
func MinTwoIntegers(x, y int) int {
	if x < y {
		return x
	}

	return y
}

// MaxTwoIntegers .. finds maximum of two intergers and returns maximum
func MaxTwoIntegers(x, y int) int {
	if x < y {
		return x
	}

	return y
}
