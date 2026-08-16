/*Part one:
 * 1. the diference between declaring and calling is that in calling you use an already decalred function. while declaring you make the function
 * 2. we use the blank identifier to skip that return value
 * 3. a variable function scope means that the variable can't be used outside the function.
 * 4. it makes the code look cleaner
 */
package main
import "fmt"
func greet(name string) string {
	return "Hello, " + name + " !"
	}
func square(i int) int {
	return i * i
	} 
func addThree(i,y,x int) (sum int) {
	sum = i + y + x
	return //rewritten 
	}
func divide(a,b int) (int,int) {
	quotation := a / b 
	remainder := a % b
	return quotation , remainder
	}		
func minMax(a,b,c int) (bigger, smaller int) {
	if a >= b && a >= c {
		bigger = a		
		} else if b >= a && b >= c {
			bigger = b
			} else if c >= a && c >= b {
				bigger = c
				}
	if a <= b && a <= c {
		smaller = a
		} else if b <= a && b <= c {
		smaller = b	
			} else if c <= a && c <= b {
				smaller = c
				}
					
		return bigger , smaller
	}
func main() {
	fmt.Println(greet("sonic"))
	fmt.Println(square(99))
	fmt.Println(addThree(6,5,4))
	quotation , remainder := divide(5,9)
	fmt.Println(quotation,remainder)
	_ , b := divide(3,6)
	fmt.Println(b)
	fmt.Println(minMax(5,5,2))
	}
//ex g; it's missing a return keyword
