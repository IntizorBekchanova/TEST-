package main

import "fmt"

func main() {

	a := 100 
	b := 10

	c := a + b // с = 110


	c1 := a * b  // с = 1000


	c2 := a - b  // с = 90

	
	c3 := a / b  // с = 10

    c4 := a % b  // с = 0
	fmt.Println(c, c1, c2, c3, c4)

}	
//----------------------------------------------------------------

package main

import "fmt"

func main() {
    var t int = 93
    a := 100
    i := a - t
    fmt.Println(i)
}
//----------------------------------------------------------------

package main

import "fmt"

func main(){
    var a int
    fmt.Scan(&a)
    
    c :=(a*2)+100
    fmt.Println(c)
}