package main

import "fmt"

type Product interface {
	DoStuff() string
}

type Food struct {
	Name string
}

func (f *Food) DoStuff() string {
	return f.Name
}

type NasiGoreng struct {
	Food
}

type AyamGoreng struct {
	Food
}

func CreateFood(name string) Product {
	switch name {
	case "Nasi Goreng":
		return &NasiGoreng{
			Food: Food{Name: "Nasi Goreng"},
		}
	case "Ayam Goreng":
		return &AyamGoreng{
			Food: Food{Name: "Ayam Goreng"},
		}
	default:
		return nil
	}
}
func main() {
	nasiGoreng := CreateFood("Nasi Goreng")
	fmt.Println(nasiGoreng.DoStuff())

	ayamGoreng := CreateFood("Ayam Goreng")
	fmt.Println(ayamGoreng.DoStuff())
}
