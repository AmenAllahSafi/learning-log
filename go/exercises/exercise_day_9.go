package main
import "fmt"
func main() {
	a := 10
	b := 20
	fmt.Println(a == b, a != b)
	// output will be false then true
	fmt.Println(a < b, a > b, a <= b, a >= b)// true false true false
	s1 := "go"
	s2 := "Go"
	fmt.Println(s1 == s2)//false
	fmt.Println(5 == 5.0)// i know go is strict with types i never experminted this but i guess it's an error UPDATE it ran fine.
}
