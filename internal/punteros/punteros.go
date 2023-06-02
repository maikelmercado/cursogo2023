package punteros

func DemoPuntero() {

	num := 1

	println(num)

	aumentar(&num)

	println(num)

}

func aumentar(num *int) {
	*num = *num + 5
}
