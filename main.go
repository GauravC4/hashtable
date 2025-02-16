package main

func main() {
	mm := MyMapConstructor()
	mm.Set(23, 203)
	mm.Set(25, 205)
	mm.Set(23, 207)
	mm.Set(23, 209)
	mm.Set(27, 207)

	mm.Remove(25)
	mm.Remove(23)
	mm.Remove(25)
	mm.Remove(27)

}
