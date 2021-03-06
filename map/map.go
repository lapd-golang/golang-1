package main
import (
	"fmt"
)

type People struct{
	Name string
	House map[string] string
}

type Child struct{
	Name string
	House map[string] string
}

func main() {
	set := map[int]string {}
	if set != nil {
		fmt.Println(set, len(set))
		set[1] = "hello"
		fmt.Println(set, len(set))
	} else {
		set = make(map[int]string)
		set[0]="hello"
		fmt.Println(set)
	}
	fmt.Println("--------------------")

	var set1  map[int] interface {}
	if set1 != nil {
		fmt.Println(set1)
	} else {
		set1 = make(map[int]interface{})
		set1[0]="hello"
		set1[2]=100
		fmt.Println(set1)
	}
	fmt.Println("--------------------")

	var set2  map[string]*People 
//	set2 := map[string]*People{}
	if set2 != nil {
		fmt.Println(set2, len(set2))
	} else {
		if k, ok := set2["1"]; ok == true {
			fmt.Println("nil, ", k, len(set2))

		}
		set2 = make(map[string]*People)
		set2["2"] = &People{"wang", map[string]string{"Beijing" : "Haidian"}}
		fmt.Println(set2, len(set2))
	}
	fmt.Println("--------------------")

	ch := new(Child)
	ch.Name = "Qing"
	if ch.House == nil {
		ch.House = make(map[string] string)
		ch.House["Beijing"]="Haidian"
	} else {
	//	fmt.Println(ch)
	}
	ch.House["TJ"] = "Hepin"
	fmt.Println(ch)
	fmt.Println("--------------------")

	mm := ch.House
	fmt.Println(mm)
}
