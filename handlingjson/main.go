package main

import (
	"encoding/json"
	"fmt"
)

type course struct{
 Name string `json:"coursename"`
 Price int 
 Platform string `json:"website"`
 Password string  `json:"-"`
 Tags []string `json:"tags,omitempty"`

}


func main() {
	fmt.Println("Json data")
	// EncodeJson()
	DecodeJson()
}

func EncodeJson(){
	lcoCourses:=[]course{
		{"ReactJs Bootcamp",299,"LearncodeOnline.in","abcd123",[]string{"web-dev","js"}},
		{"JavaScript Bootcamp",199,"LearncodeOnline.in","bcda123",[]string{"full-stack","js"}},
		{"Angular Bootcamp",399,"LearncodeOnline.in","cdab123",nil},
	} 

	//package this data as Json data 

	finalJson,err:=json.MarshalIndent(lcoCourses, "lco","\t")
	if err != nil{
		panic(err)
	}
    fmt.Printf("%s\n",finalJson)
    // fmt.Println(finalJson)


}

func DecodeJson(){
	jsonDataFromWeb:=[]byte(`
	
	{
        "coursename": "ReactJs Bootcamp",
        "Price": 299,
        "website": "LearncodeOnline.in",
        "Password": "abcd123",
        "Tags": [ "web-dev", "js"]
    }
   `)
   var lcoCourse course
   checkValid:=json.Valid(jsonDataFromWeb)
   if checkValid{
	fmt.Println("JSON WAS VALID")
	json.Unmarshal(jsonDataFromWeb,&lcoCourse)
	fmt.Printf("%#v\n",lcoCourse)
   }else{
	fmt.Println("Json was not valid")
   }

   //some cases wherev you just want to add data to key value

   var myOnlineData map[string] interface{}
	json.Unmarshal(jsonDataFromWeb,&myOnlineData)

	for k,v:=range myOnlineData{
		fmt.Printf("key is %v and value is %v and type is :%T\n",k,v,v)
		
	}
   



}