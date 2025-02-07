package main

import (
	"log"
	"net/http"
)

//Define a home handler function which writes a byte slice containing 
// "Hello from SnippetBox" as the response body
func home(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello from SnippetBox"))
}

func snippetView(w http.ResponseWriter, r *http.Request){
	w.Write([]byte("Display a specific snippet..."))
}

func snippetCreate(w http.ResponseWriter,r *http.Request){
	w.Write([]byte("Display a form for creating a new snippet..."))
}

func main() {
	//Used the http.NewServeMux() function to initialize a new servermus (router btw), then

	// register the home functon as the handler for the  "/" URL pattern.
	mux := http.NewServeMux()
	mux.HandleFunc("/{$}", home)
	mux.HandleFunc("/snippet/view",snippetView)
	mux.HandleFunc("/snippet/create", snippetCreate)

	// Prints a log message to say that the server is starting
	log.Print("Server startig on : 4000")
	// Use the http.ListenAndServe() function to start a new web server. We pass in
	// two parameters: the TCP network address to listen on (in this case ":4000")
	// and the servemux we just created. If http.ListenAndServe() returns an error
	// we use the log.Fatal() function to log the error message and exit. Note
	// that any error returned by http.ListenAndServe() is always non-nil.
	err := http.ListenAndServe(":4000", mux) // Should always be in format "host:port"
	log.Fatal(err)
}
