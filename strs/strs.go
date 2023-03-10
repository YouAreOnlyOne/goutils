package strs

import (
	"net"
	"reflect"
	"regexp"
	"strconv"
	"unicode"
	"unsafe"
)

func StrLength(str string) int {
	count := 0
	for _, c := range str {
		if unicode.Is(unicode.Han, c) {
			count++
		} else {
			count++
		}
	}
	return count
}

func ChEnLength(str string) int {
	count := 0
	for _, c := range str {
		if unicode.Is(unicode.Han, c) {
			count += 2
		} else {
			count++
		}
	}
	return count
}

func CheckPasswordRule(password string) bool {
	if len(password) < 11 {
		return false
	}
	var digit, upper, lower, special bool
	for _, c := range password {
		if unicode.IsDigit(c) {
			digit = true
		} else if unicode.IsUpper(c) {
			upper = true
		} else if unicode.IsLower(c) {
			lower = true
		} else {
			special = true
		}
	}
	if !digit || !upper || !lower || !special {
		return false
	}
	return true
}

func SubnetMatch(subnet string) (string, string, error) {
	ip, ipv4Net, err := net.ParseCIDR(subnet)
	if err != nil {
		return "", "", err
	}
	return ipv4Net.String(), ip.String(), err
}

func CheckIP(ip string) string {
	result := net.ParseIP(ip)
	if result != nil {
		return result.String()
	}
	return ""
}

func Byte2Any(b []byte, t reflect.Type) interface{} {
	data := Byte2Str(b)
	switch t.Kind() {
	case reflect.Int:
		res, _ := strconv.Atoi(data)
		return res
	case reflect.Int64:
		res, _ := strconv.ParseUint(data, 10, 64)
		return res
	case reflect.Uint:
		res, _ := strconv.ParseUint(data, 10, 64)
		return res
	case reflect.Uint64:
		res, _ := strconv.ParseUint(data, 10, 64)
		return res
	case reflect.Float64:
		res, _ := strconv.ParseFloat(data, 64)
		return res
	case reflect.Bool:
		res, _ := strconv.ParseBool(data)
		return res
	default:
		return data
	}
}

func Byte2Str(b []byte) string {
	return *(*string)(unsafe.Pointer(&b))
}

func Str2Uint64(s string) uint64 {
	res, _ := strconv.ParseUint(s, 10, 64)
	return res
}

func Str2Float(s string) float64 {
	res, err := strconv.ParseFloat(s, 64)
	if err != nil {
		return 0
	}
	//return math.Trunc(res) * 1e-2
	res, _ = NewFromFloat(res).Round(2).Float64()
	return res
}

func CheckMobile(phone string) bool {
	// ????????????
	// ^1???????????????
	// [345789]{1} ????????????345789 ?????????
	// \\d \d????????? ???????????? {9} ???9???
	// $ ?????????
	regRuler := "^1[345789]{1}\\d{9}$"

	// ??????????????????
	reg := regexp.MustCompile(regRuler)

	// ?????? MatchString ????????????
	return reg.MatchString(phone)
}
