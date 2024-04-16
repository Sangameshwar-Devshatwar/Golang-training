package main

import "fmt"

func main() {
	rect := Rectangle{}

	rect.SetLength(1)

	rect.SetWidth(2)
	rr := rect.Area()
	fmt.Println(rr)
	rect.SetLength(23)

	rect.SetWidth(34)
	res := rect.Perimiter()
	fmt.Println(res)
	len := rect.GetLength()
	fmt.Println(len)
	wid := rect.GetWidth()
	fmt.Println(wid)
	re := rect.Area()
	fmt.Println(re)

}

type Rectangle struct {
	Length float64
	Width  float64
}

func (r *Rectangle) GetLength() float64 {
	return r.Length
}
func (r *Rectangle) SetLength(val float64) {
	r.Length = val

}
func (r *Rectangle) GetWidth() float64 {
	return r.Width
}
func (r *Rectangle) SetWidth(val float64) {
	r.Width = val
}
func (r *Rectangle) Area() float64 {
	return r.Length * r.Width
}
func (r *Rectangle) Perimiter() float64 {
	return 2 * (r.Length + r.Width)
}
