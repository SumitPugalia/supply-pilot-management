package main

//------------------------------------------------------------
// This is the main file that starts the application.
//-------------------------------------------------------------
import (
	"pilot-management/router"
)

//------------------------------------------------------------
// This is the entry/starting point of our application.
//-------------------------------------------------------------
func main() {
	router.StartApp(":8080")
}
