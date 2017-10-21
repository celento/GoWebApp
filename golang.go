// by <@celento>. Feel free to use it for anything cool !

package main

import ("fmt"
"net/http")



//  w parameter is of the http.ResponseWriter type.

func index(w http.ResponseWriter, r *http.Request) {
	// Fprintf formats by what is specified, and writes to the first parameter, w.
	fmt.Fprintf(w, "Hi there! Welcome to the Web App.")
}

func login(w http.ResponseWriter, r *http.Request) {
	// Fprintf formats by what is specified, and writes to the first parameter, w.	
	fmt.Fprintf(w, "Welcome to the Login Page.")
}

func signup(w http.ResponseWriter, r *http.Request) {
	// Fprintf formats by what is specified, and writes to the first parameter, w.	
	fmt.Fprintf(w, "Welcome to the Online SignUp Page.")
}

func store(w http.ResponseWriter, r *http.Request) {
	// Fprintf formats by what is specified, and writes to the first parameter, w.	
	fmt.Fprintf(w, "Welcome to the Online Store Page.")
}


func about(w http.ResponseWriter, r *http.Request) {
	// Fprintf formats by what is specified, and writes to the first parameter, w	
	fmt.Fprintf(w, "Welcome to the about page.")
}

func main() {

	// HandleFunc is used to link functions to the corresponding paths.
	// "/" is the index or root file of the web app. 
 	http.HandleFunc("/", index)

	//  Linking other pages to the corresponding functions.
	http.HandleFunc("/store/", store) 
	http.HandleFunc("/login/", login)
	http.HandleFunc("/signup/", login)	
	http.HandleFunc("/about/", about)

	// Defining a port for the app to run. This example will run on port 8000
	// The second argument here is server stuff. We will use nil for this example.
	http.ListenAndServe(":8000", nil) 



}

// *********************************************************************************//