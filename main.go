package main
import "fmt"
 import "net/http"

// func index (w http.ResponseWriter, r *http.Request){
  /*
w.Write([]byte("Wisdom Welcome to Golang"))
  }
*/
func main()  {
  
  // http.HandleFunc("/index", index)
 
  // http.ListenAndServe(":8000",nil)
 
  var name string 
  var age int
  
for{
  fmt.Printf("Enter name : \n\n")
  fmt.Scan(&name)
  fmt.Printf("Enter age : \n\n")
  fmt.Scan(&age)
  
  fmt.Printf(" %#v became a Golang developer at age %v \n",name,age)
  
  var hobbies [3]string
   hobbies[0] = "Math"
   hobbies[1] = "Music"
   hobbies[2] = "Tech"
  
  fmt.Printf("The hobbies is of type %T and all %v hobby are %v %v %v \n",hobbies,name,hobbies[0],hobbies[1],hobbies[2])
   } 
}
