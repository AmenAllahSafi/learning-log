package main
import "fmt"
func main() {
	//logical AND
	isAdult := true
	hasTicket := false
	fmt.Println(isAdult && hasTicket) //false
	//Logical OR
	fmt.Println(isAdult || hasTicket) //true
	//combine with comparisons
	age := 25
	hasID := true
	fmt.Println(age >= 18 && hasID == true)
	//OR with comparission
	temp := 15
	fmt.Println(temp < 10 || temp > 30)
	/*if the temp was 5 it will output True
	 * if the temp was 35 it will output true aswell  
	 */
	 //turth table
	 /*
	  * false and true = false
	  * true and true = true
	  * false and false = false
	  * 
	  * false or true = true
	  * true or true = true
	  * false or false = false
	  */
	  //combine AND and OR
	  hasMembership := true
	  hasCoupon := false
	  purchaseAmount := 50
	  getsDiscount := (hasMembership || hasCoupon) && purchaseAmount > 50
	  fmt.Println(getsDiscount)
	
	}
