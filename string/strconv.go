package main
import (
	"fmt"
	"strconv"
)

func main(){
	str := strconv.Itoa(strconv.IntSize)
	str1 := strconv.Itoa(int(1))
	fmt.Println("int to string", str, str1)
	fmt.Println("Variables:", strconv.ErrRange, "\n", strconv.ErrSyntax)

	str2 := strconv.FormatInt(int64(1), 10)
	fmt.Println("int64 to string", str2)
	var i8 int8 = 127
//	str22:= strconv.FormatInt(int64(i8), 10)
	str22:= strconv.FormatInt((i8), 10)
	fmt.Println("int64 to string", str22)

	str3 := strconv.FormatFloat(float64(0.008), 'f', -1, 32)
	fmt.Println("float64 to string", str3)

	i, _ := strconv.Atoi("10")
	fmt.Println("sting to int", i)
	
	i64, _ := strconv.ParseInt("1234", 10, 64)
	fmt.Println("sting to int64", i64)

	f32, _ := strconv.ParseFloat("3.1415926535", 32)
	fmt.Println("sting to float32", f32)

	fmt.Println("----------------------------------------------")
	ss := make([]byte, 0, 100)
	ss = strconv.AppendInt(ss, 1234, 10)
	ss = strconv.AppendBool(ss, false)
	ss = strconv.AppendQuote(ss, "abcdefg")
	ss = strconv.AppendQuoteRune(ss, '单')
	fmt.Println(string(ss))
	fmt.Println("----------------------------------------------")
}
