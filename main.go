package main

import "pilot-management/router"

//------------------------------------------------------------
// This is the main file that starts the application.
//-------------------------------------------------------------

//------------------------------------------------------------
// This is the entry/starting point of our application.
//-------------------------------------------------------------
func main() {
	router.StartApp(":8080")
}
