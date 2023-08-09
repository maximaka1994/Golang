package main

import "fmt"

func main() {

	var num int
	fmt.Println("Введите номер числа из ряда Фибоначи:")
	fmt.Scanf("%d\n", &num)

	nm := 1
	buf := 0
	var buf1 int
	for i := 1; i < num; i++ {
		buf1 = nm
		nm += buf
		buf = buf1
	}

	fmt.Println("Значение числа:", nm)
}
