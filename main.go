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

type Drink struct {
	Name string
}

func (d *Drink) DoStuff() string {
	return d.Name
}

type Americano struct {
	Drink
}

type CaramelLatte struct {
	Drink
}

func CreateProduct(name string) Product {
	switch name {
	case "Nasi Goreng":
		return &NasiGoreng{
			Food: Food{Name: "Nasi Goreng"},
		}
	case "Ayam Goreng":
		return &AyamGoreng{
			Food: Food{Name: "Ayam Goreng"},
		}
	case "Americano":
		return &Americano{
			Drink: Drink{Name: "Americano"},
		}
	case "CaramelLatte":
		return &CaramelLatte{
			Drink: Drink{Name: "CaramelLatte"},
		}
	default:
		return nil
	}
}
func main() {
	nasiGoreng := CreateProduct("Nasi Goreng")
	fmt.Println(nasiGoreng.DoStuff())

	ayamGoreng := CreateProduct("Ayam Goreng")
	fmt.Println(ayamGoreng.DoStuff())

	americano := CreateProduct("Americano")
	fmt.Println(americano.DoStuff())

	caramelLatte := CreateProduct("CaramelLatte")
	fmt.Println(caramelLatte.DoStuff())
}
