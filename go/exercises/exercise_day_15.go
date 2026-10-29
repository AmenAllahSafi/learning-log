package main
import "fmt"
func main(){
	//1.
	day := 10
	switch day {
		case 1:
		    fmt.Println("Monday")
		case 2:
			fmt.Println("Tuesday")
		case 3:
			fmt.Println("Wendesday")
		case 4:
			fmt.Println("Thursday")
		case 5:
			fmt.Println("Friday")
		case 6:
			fmt.Println("Saturday")
		case 7:
			fmt.Println("Sunday")
		default:
			fmt.Println("Invalid day")
		}
		//2.
		grade := "N"
		switch grade {
			case "A":
				fmt.Println("Execellent!")
			case "B":
				fmt.Println("Good job!")
			case "C":
				fmt.Println("needs improvment")
			default: 
				fmt.Println("Unkown garde")
			}
		//3.
		score := 90
		switch {
			case score == 90:
				fmt.Println("A")
			case score == 80:
				fmt.Println("B")
			case score == 70:
				fmt.Println("C")
			default:
				fmt.Println("F")
			}
		//4.
		num := 2
		switch num {
			case 1:
				fmt.Println("one")
			case 2:
				fmt.Println("two")
				fallthrough
			case 3:
				fmt.Println("three")
			case 4:
				fmt.Println("four")
				//it will print two three
			}
		//5.
		temp := 75 
		switch {
			case temp <= 32:
				fmt.Println("Freezing")
			case temp <= 60:
				fmt.Println("Cold")
			case temp <= 85:
				fmt.Println("Warm")
				fallthrough // by defualt fall though makes the next case execute
			default:
				fmt.Println("Hot")
			}	
	}
