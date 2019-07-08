package main

func main() {
	a := App{}
	a.Initialize("root", "root", "ORDERS")
	a.Run(":8080")
}
