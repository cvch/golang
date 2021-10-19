package main

import "fmt"

var tabla [10]int
var matriz [5][7]int

func main() {
	tabla[0] = 1
	tabla[5] = 15

	fmt.Println(tabla)

	tabla2 := [10]int{5, 6, 98, 1, 4, 0, 3, 54, 99, 11}
	fmt.Println(tabla2)

	for i := 0; i < len(tabla2); i++ {
		fmt.Println(tabla2[i])
	}

	matriz[3][5] = 1
	fmt.Println("Matriz")
	fmt.Println(matriz)

	matrizSlice := []int{2, 5, 4}
	fmt.Println(matrizSlice)

	variante2()
	variante3()
	variante4()
}

func variante2() {
	elementos := [5]int{1, 2, 3, 4, 5}
	// desde la posicion 3 hasta la ultima [3:]
	// desde la 1 posicion hasta la 4 [:4]
	// desde la posicion 3 hasta la 4 [3:4] 4 excluyente

	porcion := elementos[3:]
	fmt.Println(porcion)
}

func variante3() {
	elementos := make([]int, 5, 20)
	fmt.Printf("\nLargo %d, capacidad %d", len(elementos), cap(elementos))

}

func variante4() {
	// Si hacemos append sobre un slice
	nums := make([]int, 0, 0)
	for i := 0; i < 9; i++ {
		nums = append(nums, i)
	}
	fmt.Printf("\nLargo %d, Capacidad %d\n", len(nums), cap(nums))

}
