package main

import "fmt"

func main()  {
	var l, n, mayor, menor int
	fmt.Println("Leer el valor de número de números a ser leidos")
	fmt.Scanf("%d", &l)

	fmt.Println(("Se lee el primer número"))
	fmt.Scanf("%d", &n)


	mayor = n
	menor = n


	for i := 1; i < l; i++ {
		fmt.Println("See lee el nuevo número")
		fmt.Scanf("%d", &n)

		if n >= mayor {
			mayor = n
		} else if n <= menor {
			menor = n
		}
	}

	fmt.Println("El mayor número leido es ", mayor)
	fmt.Println("El menor número leido es ", menor)
}