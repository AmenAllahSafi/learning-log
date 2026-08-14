package main
import "fmt"
func main() {
	//1 basic not
	isRaining := true
	fmt.Println(!isRaining) //false
	//2.double negation
	fmt.Println(!!isRaining)
	//3.not with comparisons
	score := 45
	passing := score >= 50
	fmt.Println(passing)
	notPassing := !passing
	fmt.Println(notPassing)
	//4.precedence puzzle
	a := true
	b := false
	c := true
	fmt.Println(!a && b || c)//true
	fmt.Println(!(a && b) || c)//true
	fmt.Println(!(a && (b || c)))//false 
	//5.basic verbs 
	name := "Sonic"
	age := 30
	height := 1.75
	isFast := true
	fmt.Printf("name: %s, age: %d, height: %.1f, isfast: %t\n", name,age,height,isFast)
	//.6 decimal precision
	pi := 3.14159265
	fmt.Printf("pi: %f, pi: %.2f pi: %.4f\n", pi,pi,pi)
	//.7 The universel verb %v
	fmt.Printf("name: %v\n" ,name)
	fmt.Printf("age: %v\n" ,age)
	fmt.Printf("height: %v\n" ,height)
	fmt.Printf("fast: %v\n" ,isFast)
	//8. Combine everything
	fmt.Printf("Score: %d, Passing: %t\n",score, passing)
	
	}
