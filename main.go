// main.go

package main

func main() {
	a := App{}
	// You need to set your Username and Password here
	a.Initialize("root", "root", "rest_api_example")

	a.Run(":8083")
}
