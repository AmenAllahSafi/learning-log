package main
import "fmt"
func main() {
	fmt.Println("Welcome")
	var name string
	var age int
	fmt.Print("What's your name? : ")
	fmt.Scanln(&name)
	fmt.Print("What's your age? : ")
	fmt.Scanln(&age)
	fmt.Printf("Hi %s, you're %d years old.\n", name, age)
	if age < 18 {
		fmt.Println("minor")
		}else if age > 18 && age < 65 {
			fmt.Println("adult")
			}else{
				fmt.Println("senior")
				}
	var num int 
	fmt.Print("A number please: ")
	fmt.Scanln(&num)
	if num % 2 == 0 {
		fmt.Println("Your number is even")
		}else{
			fmt.Println("Your number is odd")
			}
	switch {
		case num >= 0:
			fmt.Println("small number")
		case num >= 10:
			fmt.Println("slightly better")
			fallthrough
		case num >= 20:
			fmt.Println("big number")
		}
	for i := 1 ; i <= num; i++ {
		fmt.Println(i)
		if i % 3 == 0{
			continue
			}
		if i % 13 == 0 {
			break
			}
		for j := 1 ; j <= i ; j++ {
			fmt.Printf("%d * %d ", j, i)
			fmt.Printf("%d * %d \n", j, i)
			}		
		}		
	fmt.Printf("age : %T num : %T\n", age, num)		
	}
