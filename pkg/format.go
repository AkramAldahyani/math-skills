package mathskills

import "strconv"

// Changing the data set format from []byte to []float64 to make dealing with them easier
func Format(arr []byte) []float64 {
	var a []string
	var word string
	for i := 0; i < len(arr); i++ {
		if string(arr[i]) == "\n" || string(arr[i]) == " " {
			a = append(a, (word))
			word = ""
		}
		if (string(arr[i]) >= "0" && string(arr[i]) <= "9") || string(arr[i]) == "." {
			word += string(arr[i])
		}
	}
	if word != "" {
		a = append(a, (word))
	}
	var res []float64
	for i := 0; i < len(a); i++ {
		value, err := strconv.ParseFloat(a[i], 64)
		if CheckError(err) {
			value = 0
		}
		res = append(res, value)
	}
	return res
}
