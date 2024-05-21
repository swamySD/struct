package main

import (
	// "bufio"
	"fmt"
	// "math/rand"
	// "time"
	// "sort"
	// "os"
	// "strconv"
	// "strings"
	// "time"
)

// "bufio"
// "fmt"
// "os"

//  type is a blue print
// no inheritance in golang ;no super or parent claases instead we usecustom  type
 type User struct{
	Name string 
	Email string
	Status bool 
	Age int
	
 }

func main() {
	// welcome :="Welcome to userInput"
	// fmt.Println(welcome)
	// reader:=bufio.NewReader(os.Stdin)
	// fmt.Println("Enter the rating for our Pizza")
	// //comma ok//err ok
	// input,_:=reader.ReadString('\n')
	// fmt.Println("Thanks for rating, ",input)
	// fmt.Printf("type of this rating is %T",input)
	// fmt.Println("Welcome to our pizza app")
	// fmt.Println("Please rate our pizza between 1 and 5")

	// reader:=bufio.NewReader(os.Stdin)
	// input,_:=reader.ReadString('\n')
	// fmt.Println("Thanks for rating," ,input)
	// numRating,err := strconv.ParseFloat(strings.TrimSpace(input),64)
	// if err != nil{
	// 	fmt.Println(err)
	// 	return
	// }
	// fmt.Println("Added 1 to your rating",numRating+1)

	//package time
	//for present time
	// presentTime:=time.Now()
	// fmt.Println(presentTime)

	// fmt.Println(presentTime.Format("01-02-2006 15:04:05 Monday"))

	// //for creating time
	// createdDate:=time.Date(2020,time.August,12,23,23,0,0,time.UTC)
	// fmt.Println(createdDate)
	// fmt.Println(createdDate.Format("01-02-2006 Monday"))

	//pointers

	// var one *int
	// fmt.Println(one)
	//Here reference meanse & percent
	// myNumber := 23
	// var ptr = &myNumber
	// fmt.Println("Value of actual pointer is", ptr) //this gives the memory address of the myNumber
	// fmt.Println("value of a pointer is", *ptr)     //this gives the value actual value

	// fruits := []string{"apple", "bananana", "orange"}
	// fmt.Println(fruits)

	// fruits = append(fruits, "Mango", "grapes")
	// fmt.Println(fruits)

	// highScore := make([]int, 4)
	// highScore[0] = 234
	// highScore[1] = 334
	// highScore[2] = 434
	// highScore[3] = 534
	// highScore = append(highScore, 34, 34, 22, 556, 5676)

	// fmt.Println(highScore)
	// sort.Ints(highScore)

	//remove a value from slice based on index

	// courses := []string{"html", "css", "javascript", "ReactJs", "NextJs"}
	// fmt.Println(courses)
	// var index int=2
	// courses=append(courses[:index],courses[index+1:]...)
	// fmt.Println(courses) 
	
	//maps

		// languages:=make(map[string]string)
		// languages["RB"]="Ruby"
		// languages["PY"]="python"
		// languages["JS"]="Javascript"
		// languages["RS"]="React"
	     
		// fmt.Println("List of all languages: ",languages)
		// fmt.Println("Js shorts for:", languages["JS"])

		// delete(languages,"RB")
		// fmt.Println(languages)

		// //loops 

		// for key,value:=range languages{
		// 	fmt.Printf("For key %v value is %v\n",key,value)
		// }



		//structs are object like in javascript

		// hitesh:=User{"Hitesh","hitesh@go.dev",true,16}
		// fmt.Printf("hitesh details are %+v\n",hitesh)

	//if esle			
	// loginCount:=10
	// var result string 
	// if loginCount <10{
	// 	result="Regular user"
	// }else if loginCount > 10{
	// 	result ="Watch out"
	// }else{
	// 	result ="Exactly 10 login count"
	// }
	// fmt.Println(result)
	
	// if num:=3; num <10{
	// 	fmt.Println("Num is less than 10")
	// }else{
	// 	fmt.Println("Number is not less than 10")
	// }
   
	//switch case 
// 	rand.Seed(time.Now().UnixNano())
// 		diceNumber:=rand.Intn(6)+1 
// 		fmt.Println("Value of dice is",diceNumber)
// 		switch diceNumber {
// 		case 1:
// 			fmt.Println("Dice value is 1 you can open")
// 		case 2:
// 			fmt.Println("You can move 2 spot")
// 		case 3:
// 			fmt.Println("You can move 3 spot")	
// 		case 4:
// 			fmt.Println("You can move 4 spot")
// 		case 5:
// 			fmt.Println("You can move 5 spot")
// 		case 6:
// 			fmt.Println("You can move 6 spot")	
				

// 		}
	// days:=[]string{"Sunday","Tueesday","Wednesday","Friday","Saturday"}
    // fmt.Println(days)
	// for d:=0;d<len(days);d++{
	// 	fmt.Println(days[d])
	// }

	// for i:=range days{
	// 	fmt.Println(days[i])
	// }

	// for i,index:=range days{
	// 	fmt.Println(days[i],index)
		
	// }
	// result:=adder(2, 4)
	// fmt.Println("Result is: ",result)
	// proResult,message:=proAdder(1,3,3,4,5,4,3,3,2,2,4)
	
	// fmt.Println("Pro Result is",proResult)
	// fmt.Println("Pro messagee is",message)

	//Methods in go lang 

	hitesh:=User{"Hitesh","hitesh@go.dev",true,25}
	 hitesh.GetStatus()
	hitesh.NewMail()
 }

 func adder(a int, b int)int{
	return a+b
 }

 func proAdder(values ...int,)(int,string){
	total:=0 
	for _,value :=range values{
		total+=value
	}
	return total,"Hi pro result is found"
 }

func (U User)GetStatus(){
fmt.Println("Is User active",U.Status)
}
func (U User) NewMail(){
	U.Email="test@go.dev"
	 fmt.Println("Hello")
	fmt.Println("Email of this user is:",U.Email)

}

 