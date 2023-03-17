package a

import "fmt"

func main() {
	mapValue := map[string]int{"a": 1, "b": 2}
	a := mapValue["a"] // want `map key exists check is not performed`
	fmt.Println(a)

	b, ok := mapValue["b"]
	if !ok {
		fmt.Println("b is not exists")
	} else {
		fmt.Println(b)
	}
}
