package mario

func Mario(num int) string {
	pyramid := ""

	for i := 1; i <= num; i++ {
		hash := ""
		spaces := ""
		line := ""
		// Imprime espacios en blanco a la izquierda
		for j := 1; j <= num-i; j++ {
			spaces += " "
		}

		// Imprime "#" en la pirámide
		for k := 1; k <= i; k++ {
			hash += "#"
		}

		// Imprime cada linea
		line = spaces + hash + "\n"

		//Concatena cada linea y crea la piramide
		pyramid += line
	}
	return pyramid
}
