package bigint

import (
	"errors"
	"strconv"
	"strings"
)

type Bigint struct {
	Value string
}

var ErrorBadInput = errors.New("bad input, please input only number")

// deleteZeros...
func deleteZeros(input string) string {
	var result = ""
	check := false
	if strings.HasPrefix(input, "-") {
		check = true
		input = input[1:]
	}
	input = strings.TrimPrefix(input, "+")
	for _, item := range input {
		if len(result) == 0 && item == '0' {
			continue
		}
		result += string(item)
	}
	if check && result != "0" {
		result = "-" + result
	}
	return result
}

func validateInput(input string) (bool, string) {
	allowed := "1234567890"
	var err bool
	start := 0
	if strings.HasPrefix(input, "+") || strings.HasPrefix(input, "-") {
		start = 1
	}

	arr := strings.Split(input[start:], "")
	for _, v := range arr {
		if !strings.Contains(allowed, v) {
			err = true
		}
	}

	return err, input
}

func NewInt(num string) (Bigint, error) {
	err, num := validateInput(num)
	if err {
		return Bigint{Value: num}, ErrorBadInput
	}
	num = deleteZeros(num)
	return Bigint{Value: num}, nil
}

func (z *Bigint) Set(num string) error {
	err, num := validateInput(num)
	if err {
		return ErrorBadInput
	}
	num = deleteZeros(num)
	z.Value = num
	return nil
}

//Add ...
func Add(first, second Bigint) (r Bigint) {
	var err error
	first, err = NewInt(first.Value)
	if err != nil {
		panic("Error input1!")
	}
	second, err = NewInt(second.Value)
	if err != nil {
		panic("Error input2!")
	}
	num1 := first.Value
	num2 := second.Value
	var result string
	if num1 == "" {
		r.Value = num2
		return
	}
	if num2 == "" {
		r.Value = num1
		return
	}
	if num1[0] != '-' && num2[0] != '-' {
		result = add(num1, num2)
		r.Value = result
		return
	}
	if num1[0] == '-' && num2[0] == '-' {
		arr1 := strings.Split(num1, "-")
		str1 := strings.Join(arr1, "")
		arr2 := strings.Split(num2, "-")
		str2 := strings.Join(arr2, "")
		result = add(str1, str2)
		result = "-" + result
		r.Value = result
		return
	} //-123 13
	if num1[0] == '-' || num2[0] == '-' {
		pre1 := num1
		pre2 := num2
		if num1[0] == '-' {
			num1 = num1[1:]
		}
		if num2[0] == '-' {
			num2 = num2[1:]
		}
		result = sub(num1, num2)
		if (len(num1) > len(num2) ||
			len(num1) == len(num2) &&
				int(num1[0]) > int(num2[0])) &&
			pre1[0] == '-' {
			result = "-" + result
		}
		if (len(num2) > len(num1) ||
			len(num2) == len(num1) &&
				int(num2[0]) > int(num1[0])) &&
			pre2[0] == '-' {
			result = "-" + result
		}
	}
	r.Value = result
	return
}

//Subtraction ...
func Sub(a, b Bigint) (c Bigint) {
	son1 := a.Value
	son2 := b.Value
	var natija string
	if son1[0] != '-' && son2[0] != '-' {
		natija = sub(son1, son2)
		if len(son2) > len(son1) || (len(son2) == len(son1) && int(son2[0]) > int(son1[0])) {
			natija = "-" + natija
		}
		c.Value = natija
		return
	}
	if son1[0] == '-' && son2[0] == '-' {
		pre1 := son1
		arr1 := strings.Split(son1, "-")
		str1 := strings.Join(arr1, "")
		arr2 := strings.Split(son2, "-")
		str2 := strings.Join(arr2, "")
		natija = sub(str1, str2)
		if (len(son1) > len(son2) || (len(son1) == len(son2) && int(son2[0]) > int(son1[0]))) && pre1[0] == '-' {
			natija = "-" + natija
		}
		c.Value = natija
		return
	}
	if son1[0] == '-' || son2[0] == '-' {
		pre1 := son1
		if son1[0] == '-' {
			arr1 := strings.Split(son1, "-")
			str1 := strings.Join(arr1, "")
			son1 = str1
		}
		if son2[0] == '-' {
			arr2 := strings.Split(son2, "-")
			str2 := strings.Join(arr2, "")
			son2 = str2
		}
		natija = add(son1, son2)
		if pre1[0] == '-' {
			natija = "-" + natija
		}
	}
	c.Value = natija
	return
}

func add(a, b string) (result string) {
	arr1 := strings.Split(a, "")
	arr2 := strings.Split(b, "")
	var res []string
	if len(arr1) > len(arr2) {
		res = arr1
	} else if len(arr1) == len(arr2) {
		res = arr1
	} else {
		res = arr2
	}
	var sl1 = make([]int64, len(res))
	var sl2 = make([]int64, len(res))
	for i := 0; i < len(arr1); i++ {
		n1, _ := strconv.ParseInt(arr1[len(arr1)-1-i], 0, 8)
		sl1[len(sl1)-1-i] = n1
	}
	for i := 0; i < len(arr2); i++ {
		n2, _ := strconv.ParseInt(arr2[len(arr2)-1-i], 0, 8)
		sl2[len(sl2)-1-i] = n2
	}
	var r0, r1, r2 int64
	for j := (len(sl1) - 1); j > -1; j-- {
		son1 := sl1[j] + sl2[j] + r0
		r1 = son1 / 10
		r2 = son1 % 10
		r0 = r1
		s := strconv.FormatInt(r2, 10)
		result += s
		if j == 0 && r1 != 0 {
			result += strconv.FormatInt(r1, 10)
		}
	}
	rune_arr := []rune(result)
	var rev []rune
	for i := len(rune_arr) - 1; i >= 0; i-- {
		rev = append(rev, rune_arr[i])
	}
	result = string(rev)
	return

}

func sub(str1, str2 string) (natija string) {
	if len(str2) > len(str1) {
		str1, str2 = str2, str1
	}
	if len(str2) == len(str1) {
		for i := 0; i < len(str1); i++ {
			if int(str2[i]) > int(str1[i]) {
				str1, str2 = str2, str1
				break
			}
		}
	}
	arr1 := strings.Split(str1, "")
	arr2 := strings.Split(str2, "")
	var res []string
	if len(arr1) > len(arr2) {
		res = arr1
	} else if len(arr1) == len(arr2) {
		res = arr1
	} else {
		res = arr2
	}
	var sl1 = make([]int64, len(res))
	var sl2 = make([]int64, len(res))
	for i := 0; i < len(arr1); i++ {
		n1, _ := strconv.ParseInt(arr1[len(arr1)-1-i], 0, 8)
		sl1[len(sl1)-1-i] = n1
	}
	for i := 0; i < len(arr2); i++ {
		n2, _ := strconv.ParseInt(arr2[len(arr2)-1-i], 0, 8)
		sl2[len(sl2)-1-i] = n2
	}
	var s string
	var ayirma int64
	for j := (len(sl1) - 1); j > -1; j-- {
		if sl1[j] >= sl2[j] {
			ayirma = sl1[j] - sl2[j]
			s = strconv.FormatInt(ayirma, 10)
			natija += s
		} else if sl1[j] < sl2[j] {
			a := sl1[j] + int64(10)
			if j != 0 {
				sl1[j-1] = sl1[j-1] - 1
			}
			ayirma = a - sl2[j]
			s = strconv.FormatInt(ayirma, 10)
			natija += s
		}

	}
	rune_arr := []rune(natija)
	if string(rune_arr[len(rune_arr)-1]) == "0" {
		rune_arr = rune_arr[:len(rune_arr)-1]
	}
	var rev []rune
	for i := len(rune_arr) - 1; i >= 0; i-- {
		rev = append(rev, rune_arr[i])
	}
	natija = string(rev)
	var j int
	for k := len(natija) - 1; k >= 0; k-- {
		if string(natija[k]) == "0" {
			j++
		}
	}
	if j == len(natija) {
		natija = "0"
	}
	return
}

//Multiply
func Multiply(s1, s2 Bigint) (multy Bigint) {
	if s1.Value[0] == '-' && s2.Value[0] == '-' {
		s1 = Bigint{Value: s1.Value[1:]}
		s2 = Bigint{Value: s2.Value[1:]}
		multy = mul(s1, s2)
		return
	}
	if s1.Value[0] == '-' {
		s1 = Bigint{Value: s1.Value[1:]}
		multy = mul(s1, s2)
		multy.Value = "-" + multy.Value
		return
	}
	if s2.Value[0] == '-' {
		s2 = Bigint{Value: s2.Value[1:]}
		multy = mul(s1, s2)
		multy.Value = "-" + multy.Value
		return
	}
	multy = mul(s1, s2)
	return
}

func mul(n1, n2 Bigint) Bigint {
	num1 := n1.Value
	num2 := n2.Value
	if num1 == "0" || num2 == "0" {
		return Bigint{Value: "0"}
	}
	res := ""
	result := Bigint{Value: "0"}
	var slc []string
	var remain int
	if len(num1) < len(num2) {
		num1, num2 = num2, num1
	}
	for i := len(num2) - 1; i >= 0; i-- {
		res = ""
		remain = 0
		for j := len(num1) - 1; j >= 0; j-- {
			k1, _ := strconv.Atoi(string(num1[j]))
			k2, _ := strconv.Atoi(string(num2[i]))
			r := k1*k2 + remain
			remain = r / 10
			r = r % 10
			res = strconv.Itoa(r) + res
		}
		if remain != 0 {
			res = strconv.Itoa(remain) + res
		}
		for k := len(num2) - 1 - i; k > 0; k-- {
			res += "0"
		}
		slc = append(slc, res)
	}
	var value Bigint
	for _, v := range slc {
		value = Bigint{Value: v}
		result = Add(result, value)
	}
	return result
}

//Mod ...
func Mod(a, b Bigint) Bigint {
	err := errors.New("ZeroDivisionError: integer division or modulo by zero")
	string1 := a.Value
	string2 := b.Value
	result := ""
	if len(string1) < len(string2) {
		return Bigint{Value: string1}
	}
	if string2 == "0" {
		panic(err)
	}
	if string(string1[0]) == "-" && string(string2[0]) != "-" {
		result = mod(string1[1:], string2)
		if result == "0" {
			return Bigint{Value: result}
		}
		result = "-" + result
		return Bigint{Value: result}
	}
	if string(string1[0]) != "-" && string(string2[0]) == "-" {
		result = mod(string1, string2[1:])
		if result == "0" {
			return Bigint{Value: result}
		}
		result = "-"
		return Bigint{Value: result}
	}
	if string(string1[0]) == "-" && string(string2[0]) == "-" {
		result = "-" + mod(string1[1:], string2[1:])
		return Bigint{Value: result}
	}
	result = mod(string1, string2)
	return Bigint{Value: result}
}

//mod ...
func mod(n1, n2 string) (resolve string) {
	if len(n1) < len(n2) {
		resolve = n1
		return
	}
	var len = len(n1)
	var num int
	var rem = 0
	for i := 0; i < len; i++ {
		num = rem*10 + int(n1[i]-'0')
		d, _ := strconv.Atoi(n2)
		rem = num % d
	}
	resolve = strconv.Itoa(rem)
	return
}

func (x Bigint) Abs() Bigint {
	if x.Value[0] == '-' {
		return Bigint{
			Value: x.Value[1:],
		}
	}
	return Bigint{
		Value: x.Value,
	}
}
