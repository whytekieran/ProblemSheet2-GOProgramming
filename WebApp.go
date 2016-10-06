package main

//Using the gorilla framework
import (
	"github.com/gorilla/mux"
	"net/http"
	"fmt"
	"os"
	"io/ioutil"
	"bytes"
)

const (
	PORT = ":8080"
)

func helloWorld(w http.ResponseWriter, r *http.Request){
	response := "Hello World"
	fmt.Fprintln(w,response)
}

func serverHelloPage(w http.ResponseWriter, r *http.Request){
	fileName := "templates/Home.html"

	_, err := os.Stat(fileName)	//Check the file exists
	if err != nil { //If it doesnt exist
		fileName = "templates/PageNotFound.html" //The file we want to display back to the user is set to the following
	}

	http.ServeFile(w, r, fileName)		//Serve the file back to the client
}

//reverse word method for the GET
func reverse(w http.ResponseWriter, r *http.Request){
	vars := mux.Vars(r)//Get any variables passed in the request
	text := vars["text"] //get the variable called text from the URL
	reverseText := reverseWord(text);
	w.Write([]byte(reverseText))
}

//reverse word method for the POST
func reversePost(w http.ResponseWriter, r *http.Request){

	var bodyBytes []byte
	if r.Body != nil {
		bodyBytes, _ = ioutil.ReadAll(r.Body)
	}
	r.Body = ioutil.NopCloser(bytes.NewBuffer(bodyBytes))
	bodyString := string(bodyBytes)
	reverseText := reverseWord(bodyString);
	w.Write([]byte(reverseText))
}

//Function to reverse a word
func reverseWord(word string) string {

	// Convert string to rune slice.
	data := []rune(word)
	result := []rune{}

	// Add runes in reverse order.
	for i := len(data) - 1; i >= 0; i-- {
		result = append(result, data[i])
	}

	// Return new string.
	return string(result)
}

func main(){

	rtr := mux.NewRouter()
	//Make sure URL casing matches whats specified here
	rtr.HandleFunc("/Reverse/{text}", reverse)
	rtr.HandleFunc("/Reverse", reversePost).Methods("POST")
	rtr.HandleFunc("/HelloWorld", helloWorld)
	rtr.HandleFunc("/", serverHelloPage)

	http.Handle("/",rtr)
	http.ListenAndServe(PORT,nil)
}