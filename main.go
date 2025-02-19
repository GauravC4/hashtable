package main

// go to main_test.go for actual code

func main() {
	mm := MyMapOAConstructor()
	mm.Set(23, 203)
	mm.Set(25, 205)
	mm.Set(35, 207)
	mm.Set(45, 204)
	mm.Set(55, 201)
	mm.Set(23, 207)
	mm.Set(23, 209)

	mm.Remove(25)
	mm.Remove(23)
	mm.Remove(25)
	mm.Remove(27)

}
