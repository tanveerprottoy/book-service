package main

import "github.com/tanveerprottoy/book-service/internal/app/bookservice"

// starting point for the app
func main() {
	a := bookservice.NewApp()
	a.Run()
}
