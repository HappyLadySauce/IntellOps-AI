package gstr

import "strings"

// Str returns part of `haystack` string starting from and including
// the first occurrence of `needle` to the end of `haystack`.
// See http://php.net/manual/en/function.strstr.php.
func Str(haystack string, needle string) string {
	if needle == "" {
		return ""
	}
	pos := strings.Index(haystack, needle)
	if pos == NotFoundIndex {
		return ""
	}
	return haystack[pos+len([]byte(needle))-1:]
}

// StrEx returns part of `haystack` string starting from and excluding
// the first occurrence of `needle` to the end of `haystack`.
func StrEx(haystack string, needle string) string {
	if s := Str(haystack, needle); s != "" {
		return s[1:]
	}
	return ""
}

// StrTill returns part of `haystack` string ending to and including
// the first occurrence of `needle` from the start of `haystack`.
func StrTill(haystack string, needle string) string {
	pos := strings.Index(haystack, needle)
	if pos == NotFoundIndex || pos == 0 {
		return ""
	}
	return haystack[:pos+1]
}

// StrTillEx returns part of `haystack` string ending to and excluding
// the first occurrence of `needle` from the start of `haystack`.
func StrTillEx(haystack string, needle string) string {
	pos := strings.Index(haystack, needle)
	if pos == NotFoundIndex || pos == 0 {
		return ""
	}
	return haystack[:pos]
}

// SubStr returns a portion of string `str` specified by the `start` and `length` parameters.
// The parameter `length` is optional, it uses the length of `str` in default.
func SubStr(str string, start int, length ...int) (substr string) {
	strLength := len(str)
	if start < 0 {
		if -start > strLength {
			start = 0
		} else {
			start = strLength + start
		}
	} else if start > strLength {
		return ""
	}
	realLength := 0
	if len(length) > 0 {
		realLength = length[0]
		if realLength < 0 {
			if -realLength > strLength-start {
				realLength = 0
			} else {
				realLength = strLength - start + realLength
			}
		} else if realLength > strLength-start {
			realLength = strLength - start
		}
	} else {
		realLength = strLength - start
	}

	if realLength == strLength {
		return str
	} else {
		end := start + realLength
		return str[start:end]
	}
}

// SubStrRune returns a portion of string `str` specified by the `start` and `length` parameters.
// SubStrRune considers parameter `str` as unicode string.
// The parameter `length` is optional, it uses the length of `str` in default.
func SubStrRune(str string, start int, length ...int) (substr string) {
	// Converting to []rune to support unicode.
	var (
		runes       = []rune(str)
		runesLength = len(runes)
	)

	strLength := runesLength
	if start < 0 {
		if -start > strLength {
			start = 0
		} else {
			start = strLength + start
		}
	} else if start > strLength {
		return ""
	}
	realLength := 0
	if len(length) > 0 {
		realLength = length[0]
		if realLength < 0 {
			if -realLength > strLength-start {
				realLength = 0
			} else {
				realLength = strLength - start + realLength
			}
		} else if realLength > strLength-start {
			realLength = strLength - start
		}
	} else {
		realLength = strLength - start
	}
	end := start + realLength
	if end > runesLength {
		end = runesLength
	}
	return string(runes[start:end])
}

// StrLimit returns a portion of string `str` specified by `length` parameters, if the length
// of `str` is greater than `length`, then the `suffix` will be appended to the result string.
func StrLimit(str string, length int, suffix ...string) string {
	if len(str) < length {
		return str
	}
	suffixStr := defaultSuffixForStrLimit
	if len(suffix) > 0 {
		suffixStr = suffix[0]
	}
	return str[0:length] + suffixStr
}

// StrLimitRune returns a portion of string `str` specified by `length` parameters, if the length
// of `str` is greater than `length`, then the `suffix` will be appended to the result string.
// StrLimitRune considers parameter `str` as unicode string.
func StrLimitRune(str string, length int, suffix ...string) string {
	runes := []rune(str)
	if len(runes) < length {
		return str
	}
	suffixStr := defaultSuffixForStrLimit
	if len(suffix) > 0 {
		suffixStr = suffix[0]
	}
	return string(runes[0:length]) + suffixStr
}

// SubStrFrom returns a portion of string `str` starting from first occurrence of and including `need`
// to the end of `str`.
func SubStrFrom(str string, need string) (substr string) {
	pos := Pos(str, need)
	if pos < 0 {
		return ""
	}
	return str[pos:]
}

// SubStrFromEx returns a portion of string `str` starting from first occurrence of and excluding `need`
// to the end of `str`.
func SubStrFromEx(str string, need string) (substr string) {
	pos := Pos(str, need)
	if pos < 0 {
		return ""
	}
	return str[pos+len(need):]
}

// SubStrFromR returns a portion of string `str` starting from last occurrence of and including `need`
// to the end of `str`.
func SubStrFromR(str string, need string) (substr string) {
	pos := PosR(str, need)
	if pos < 0 {
		return ""
	}
	return str[pos:]
}

// SubStrFromREx returns a portion of string `str` starting from last occurrence of and excluding `need`
// to the end of `str`.
func SubStrFromREx(str string, need string) (substr string) {
	pos := PosR(str, need)
	if pos < 0 {
		return ""
	}
	return str[pos+len(need):]
}
