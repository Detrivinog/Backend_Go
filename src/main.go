package main

import (
	"fmt"
	"strconv"
	"time"
)

func geometriaCuadrado(lado int) (perimetro, area int) {
	return lado * 4, lado * lado
}

func main() {
	// Declaración de constantes
	const pi float64 = 3.14
	const pi2 = 3.1415

	fmt.Println("pi:", pi)
	fmt.Println("pi2:", pi2)

	// Declaracion de variables enteras
	base := 12
	var altura int = 14
	var area int

	fmt.Println(base, altura, area)

	// Zero values
	var a int
	var b float64
	var c string
	var d bool

	fmt.Println("Zero values: ", a, b, c, d)

	// Paquete fmt

	myName := "David"
	myLastName := "Triviño"
	myAge := 26
	myHeight := 1.70

	fmt.Printf("Mi nombre es %s %s, tengo %d y mi estatura es %f\n", myName, myLastName, myAge, myHeight)

	description := fmt.Sprintf("Mi nombre es %s %s, tengo %d y mi estatura es %f", myName, myLastName, myAge, myHeight)
	fmt.Println(description)

	// Tipos de dato
	fmt.Printf("myName: %T\n", myName)

	// Funciones
	perimetro, area := geometriaCuadrado(4)
	fmt.Printf("Perimetro: %d\nArea: %d\n", perimetro, area)

	perimetro, _ = geometriaCuadrado(5)
	fmt.Printf("Perimetro: %d\n", perimetro)

	// Cliclos: for condicional
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}

	// For While
	fmt.Println("For while")

	counter := 0
	for counter < 10 {
		fmt.Println(counter)
		counter++
	}

	// Capturando valor y error
	myValue, err := strconv.ParseInt("NaN", 0, 64)

	// Validando errores.
	if err != nil {
		fmt.Println("%v\n", err)
	} else {
		fmt.Println(myValue)
	}

	// Mapa clave valor.
	m := make(map[string]int)
	m["key"] = 6
	fmt.Println(m["key"])

	// Slice de enteros.
	s := []int{1, 2, 3}
	for index, value := range s {
		fmt.Println(index)
		fmt.Println(value)
	}
	s = append(s, 16)
	for index, value := range s {
		fmt.Println(index)
		fmt.Println(value)
	}

	//Apuntadores
	g := 25
	fmt.Println(g)
	h := &g
	fmt.Println(h)
	i := *h
	fmt.Println(i)

	// Canales

	chanel := make(chan int)
	go doSomething(chanel)
	<-chanel

}

func doSomething(c chan int) {
	time.Sleep(3 * time.Second)
	fmt.Println("done")
	c <- 1
}
