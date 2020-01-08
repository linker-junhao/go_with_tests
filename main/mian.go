package main

import (
	"github.com/linkerGitHub/go_with_tests/mocking"
	"os"
)

/*
func MyGreeterHandler(w http.ResponseWriter, r *http.Request){
	dependencyInjection.Greet(w, "world")
}

func main()  {
	http.ListenAndServe(":5000", http.HandlerFunc(MyGreeterHandler))
}*/

func main() {
	mocking.Countdown(os.Stdout)
}
