package main

import (
	d "Disks/disk"
	ds "Disks/disks"
	"fmt"
)

func main() {
	var min, max, a int
	min = 0
	max = 0
	var dd d.Disk
	var dodo ds.Disks
	fmt.Println("Введите букву для окончания ввода")
	for i := 0; i > -1; i++ {
		n, err := fmt.Scan(&a)
		if err != nil {
			break
		}
		if n == -1 {
			break
		}
		if i == 0 {
			min = a
			max = a
		} else {
			if a < min {
				min = a
			}
			if a > max {
				max = a
			}
		}
		dd.Len = a
		dodo.Que = append(dodo.Que, dd)
	}
	//dd.Len = 10
	//dodo.Que = append(dodo.Que, dd)
	//fmt.Println(dd.Len)
	fmt.Println("Самый маленький:", min, "Самый большой:", max)
	fmt.Println(dodo)
}
