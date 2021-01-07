package main

import "fmt"

// Algoritmo para escoger el mayor y el menor número de un conjunto de l números

/*
Sea un conjunto de L números  (l > 0 ), los cuales son leídos, diseñe un algoritmo que calcule
el número mayor y el menor del conjunto dado. Asuma que el usuario ingresa un entero positivo l

Análisis: Si l = 10, o el conjunto de números cuyos valores son 10, 29, 87, 25, 67, 77, 87, 98, 99, 23.
El algoritmo debe sacar el mayor de los números, es decir, el valor 99, y el menor de los números, es decir el valor 10.
*/

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