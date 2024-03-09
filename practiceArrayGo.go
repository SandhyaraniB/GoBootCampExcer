package main

import "fmt"

func main(){

	var a []int

	fmt.Println("Array is size  : ", len(a))

        fmt.Println("Array   : ", a)

	var b [3]string

	b[2]="Sandhya"

	fmt.Println("Array 4th value : ",b)
        fmt.Println("Array  : ", b[2])

	c:= [5]int{1,2,3,4,5}
	fmt.Println("Array values : ", c)

	var twoD [2][3]int
	for i:=0; i<2;i++{
		for j:=0; j<3;j++{
		twoD[i][j]=j

		}

	}
	    fmt.Println("2d: ", twoD)


 nums := []int{2, 3, 4}
    sum := 0
    for _, num := range nums {
        sum += num
    }
    fmt.Println("sum:", sum)


}
