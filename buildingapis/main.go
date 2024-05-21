package main

import (
	// "crypto/rand"
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"strconv"
	"time"
	"log"
	"github.com/gorilla/mux"
)

//Model for course -file

type Course struct{
	CourseId string `json:"courseid"`
	CourseName string `json:"coursename"`
	CoursePrice int    `json:"price"`
	Author *Author `json:"author"`
}

type Author struct{
	Fullname string `json:"fullname"`
	Website string `json:"website"`
}

//fake DB

var courses []Course 

//middleware, helper-file
// func IsEmpty(c *Course) bool{
// //    return c.CourseId == "" &&  c.CourseName ==""
//     return c.CourseName == "" 
// }

func (c *Course) IsEmpty() bool {
    return c.CourseName == "" 
}


func main() {
	fmt.Println("apis")
	r:=mux.NewRouter()

	//seeding 
	courses=append(courses, Course{CourseId: "2",CourseName: "ReactJs",CoursePrice:299,
	Author:&Author{Fullname: "Kumar",Website:"lco.dev" } ,
 })

 courses = append(courses, Course{
	CourseId:    "3",
	CourseName:  "Python Basics",
	CoursePrice: 199,
	Author:      &Author{Fullname: "John Doe", Website: "example.com"},
 })
	//listen to a port

     r.HandleFunc("/",serveHome).Methods("GET")
     r.HandleFunc("/courses",getAllCourses).Methods("GET")
     r.HandleFunc("/course/{id}",getOneCourse).Methods("GET")
     r.HandleFunc("/course",createOneCourse).Methods("POST")
     r.HandleFunc("/course/{id}",updateCourse).Methods("PUT")
     r.HandleFunc("/course/{id}",deleteOneCourse).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":5500",r))




	r.HandleFunc("/",serveHome)
}

//controllers -file 

//serve home route


func serveHome(w http.ResponseWriter , r *http.Request){
	w.Write([]byte("<h1>Welcome to Api by LearncoderOnline</h1>"))
}

func getAllCourses(w http.ResponseWriter, r *http.Request){
	fmt.Println("Get All courses")
	w.Header().Set("Content-Type","application/json")
	json.NewEncoder(w).Encode(courses)
}

func getOneCourse(w http.ResponseWriter, r*http.Request){
	fmt.Println("Get one course")
	w.Header().Set("Content-Type","application/json")

	//grab id from request 
	
	params:=mux.Vars(r)
	fmt.Println(params)

	//loop through courses,find matching id and return the response 
	for _,course:=range courses{
		if course.CourseId == params["id"]{
			json.NewEncoder(w).Encode(course)
			return 
		}
		json.NewEncoder(w).Encode("No course id is found")
		return 
	}
}

func createOneCourse(w http.ResponseWriter,r *http.Request){
	fmt.Println("Create one course")
	fmt.Println("Request Method:", r.Method)

	w.Header().Set("Content-Type","application/json")

   //what if entire body is empty

   if r.Body == nil{
	 json.NewEncoder(w).Encode("Please send some data")
   }
  
   //what about the data is sent 
   var course Course 
  _= json.NewDecoder(r.Body).Decode(&course)

  if  course.IsEmpty(){
	json.NewEncoder(w).Encode("No data inside JSON")
	return 
  }
  //Todo check only if title is duplicate 
  //loop ,title matches with course.coursename,json
   
//   for index,course := range courses{
// 	if course.courseName == 
//   }

   //generate unique id,string 
   //append course into courses
 rand.Seed(time.Now().UnixNano())
 course.CourseId=strconv.Itoa(rand.Intn(100))
 courses=append(courses, course)
 json.NewEncoder(w).Encode(course)
 return
}





func updateCourse(w http.ResponseWriter,r *http.Request){
 fmt.Println("Update a course")
 w.Header().Set("Content-Type","application/json")
 // first - grab id from req 
 params:= mux.Vars(r)
 //loop, id ,remove , add with my ID 

 for index,course:=range courses{
	if course.CourseId == params["id"]{
		courses=append(courses[:index],courses[index+1:]... )
		var course Course
		_=json.NewDecoder(r.Body).Decode(&course)
		course.CourseId=params["id"]
		courses=append(courses, course)
		json.NewEncoder(w).Encode(course)
		return
	}


 }
 //Todo send a response when id is not found
}

func deleteOneCourse(w http.ResponseWriter,r *http.Request){
	fmt.Println("Delete one course")
	w.Header().Set("Content-Type","application/json")
	params:=mux.Vars(r)

	//loop ,id perform remove operation 
	for index,course:=range courses{
		if course.CourseId == params["id"]{
			courses=append(courses[:index],courses[index+1:]...)
			break
		}
	}
}


