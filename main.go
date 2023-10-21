package main

func main() {
	println("Hello world")
	a := 5
	b := &a // Указывает на память, где находится переменная a
	println(b)
	println(a + *b) // *b = a
	println(a)
	println(a + a)
}
