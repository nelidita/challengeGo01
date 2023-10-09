package main

import (
	"fmt"
	"main/bmi"
	"main/evenodd"
	"main/loops"
	"main/mario"
	"main/physics"
	"os"
)

func main() {
	args := os.Args
	functionName := args[1]

	switch functionName {
	case "evenodd":
		var num int
		fmt.Println("Que numero quieres evaluar")
		fmt.Scan(&num)
		fmt.Println(evenodd.EvenOdd(num))
	case "ohm":
		var v float32
		var r float32
		var i float32
		fmt.Println("Dame los valores de v,r e i")
		fmt.Scan(&v, &r, &i)
		fmt.Println(physics.Ohm(v, r, i))
	case "foobar":
		var limit int
		fmt.Println("Cual sera el limite de FooBar")
		fmt.Scan(&limit)
		fmt.Println(loops.FooBar(limit))
	case "bmi":
		var weigh float32
		var tall float32
		fmt.Println("¿Cuánto pesas? (no mientas)")
		fmt.Scan(&weigh)
		fmt.Println("¿Qué altura tienes? (descalzo) ")
		fmt.Scan(&tall)
		fmt.Println(bmi.Bmi(weigh, tall))
	case "mario":
		for {
			var pyramidHeight int
			fmt.Print("Altura de la piramide:")
			fmt.Scan(&pyramidHeight)
			if pyramidHeight < 1 || pyramidHeight > 8 {
				continue
			} else {
				fmt.Println(mario.Mario(pyramidHeight))
				break
			}
		}
	}
}
