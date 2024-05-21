package main 
import (
	"fmt"
)

type bot interface{
	getGreeting() string
}

type englishBot struct{}
type spanishBot struct{}

const LoginToken string="dfdddd"
func main(){
	var isLoggedIn bool =true
	fmt.Println(isLoggedIn)
	fmt.Printf("variable is of type: %T \n ", isLoggedIn) 

	var smallVal uint8=255
	fmt.Println(smallVal)
	fmt.Printf("variable is of type: %T \n ", smallVal)

	var smallFloat float64=255.4543278233342348
	fmt.Println(smallFloat)
	fmt.Printf("variable is of type: %T \n ", smallFloat)

	var anotherVariable string 
	fmt.Println(anotherVariable)
     
	var website="learncodeonline.in"
	
	fmt.Println(website)
	fmt.Println(LoginToken)
    
	

	// var fruits [4]string
	// fruits[0]="apple"
	// fruits[1]="Tomato"
	// fruits[3]="Peach"
	// fmt.Println(fruits)
	// var vegList=[5] string{"Potato","beans","mushroom"}
	// fmt.Println("Vegy list is:",vegList)

	eb:=englishBot{}
	sb:=spanishBot{}
	printGreeting(eb)
	printGreeting(sb)
}

func printGreeting(b bot){
	fmt.Println(b.getGreeting())
}


 func (englishBot)   getGreeting() string {
	//very custom logic for generating an english 
	return "Hi there"
 }

 func (spanishBot)   getGreeting() string{
	//very custom logic for generating an english 
	return "Hola!"
 }
	



