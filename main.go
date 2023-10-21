package main

func main() {
	println("Hello world")
	a := 5
	b := &a
	println(b)
	println(a + *b)
}
