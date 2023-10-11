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
		message, bmi := bmi.Bmi(weigh, tall)
		messageBmi := fmt.Sprintf("Ahora mismo tu IMC es %f", bmi)
		fmt.Println(messageBmi)
		fmt.Println(message)
	case "mario":
		for {
			var pyramidHeight int
			fmt.Print("Altura de la piramide:")
			fmt.Scan(&pyramidHeight)
			pyramid, err := mario.Mario(pyramidHeight)

			if err == nil {
				fmt.Println(pyramid)
				break
			}
		}
	}
}
