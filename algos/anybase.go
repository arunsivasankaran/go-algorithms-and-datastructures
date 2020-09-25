package algos

const charIntMap = "0123456789abcdef"

// NumberToAnyBase returns the string representation of a given decimal number in the base provided
func NumberToAnyBase(num int, base int) string {

	if num < base {
		return string(charIntMap[num])
	}

	return NumberToAnyBase(num/base, base) + string(charIntMap[num%base])
}
