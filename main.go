package main

import "gitlab.intelligentb.com/cafu/supply/pilot-management/endpoint"

//------------------------------------------------------------
// This is the main file that starts the application.
//-------------------------------------------------------------

//------------------------------------------------------------
// This is the entry/starting point of our application.
//-------------------------------------------------------------
func main() {
	endpoint.StartApp(":8080")
}
