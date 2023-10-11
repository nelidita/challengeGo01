package bmi

func Bmi(weigh float32, tall float32) (string, float32) {
	// Fórmula: peso (kg) / [estatura (m)]2
	// Si el IMC es inferior a 18,5 -> 'Tienes bajo peso, añade más patatas al caldo'
	// Si el IMC es mayor o igual a 18,5 pero menor a 25 -> 'Tienes un peso normal, te tengo una sana envidia'
	// Si el IMC es mayor o igual a 25 -> 'Tienes sobrepeso, lo sé, la pandemia nos ha afectado a todos'
	//Ahora mismo tu IMC es 23.37

	bmi := weigh / (tall * tall)
	var evaluation string

	if bmi < 18.5 {
		evaluation = "Tienes bajo peso, añade más papas al caldo"
	} else if bmi >= 18.5 && bmi < 25 {
		evaluation = "Tienes un peso normal, te tengo una sana envidia"
	} else if bmi >= 25 {
		evaluation = "Tienes sobrepeso, lo se, la pandemia nos ha afectado a todos"
	}

	return evaluation, bmi

}
