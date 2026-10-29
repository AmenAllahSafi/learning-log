package main
import "fmt"
func main() {
	var age int = 100000
	name := "sonic"
	fmt.Println(age)
	fmt.Println(name)
	power := 99999999
	height := 0.95
	isSuper := false
	companion := "Tails"
	fmt.Println(power)
	fmt.Println(height)
	fmt.Println(isSuper)
	fmt.Println(companion)
	var sonicLovers int //it will display zero
	var sonicWeakness float64 //it will diplay a zero with a decimal
	var isSonicSmart bool //false
	var sonicFavWord string // nothing will be displayed
	fmt.Println(sonicLovers)
	fmt.Println(sonicWeakness)
	fmt.Println(isSonicSmart)
	fmt.Println(sonicFavWord)
	const pi = 3.14159
	a := 17
	b := 5
	fmt.Println("sum",a + b,"diff", a - b,"product", a * b,"remainder", a % b,"division", a / b ) //it will remain as int 
	score := 10
	fmt.Println(score) 
	score += 5
	fmt.Println(score)
	score *= 2
	fmt.Println(score)
	score -= 3
	fmt.Println(score)
	score /= 3
	fmt.Println(score)
	counter := 0
	counter++
	counter++
	counter++
	counter--
	fmt.Println(counter)
	/*
	 * user_name = bob it works but not preferd 
	 * UserName = bob it's intended for exported variables
	 * userName = bob this is the better choice everyone does it and it looks cool.
	 * */
	 
	}
