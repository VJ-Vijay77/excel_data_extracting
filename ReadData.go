package main

import (
	"fmt"

	"github.com/xuri/excelize/v2"
)

func main(){
	f,err := excelize.OpenFile("test.xlsx")
	if err != nil {
		panic(err)
	}

	c1,_ := f.GetCellValue("Sheet1","A1")
	fmt.Println("Value in A1 : ",c1)


	c2,_ := f.GetCellValue("Sheet1","B2")
	fmt.Println("Value in B2 : ",c2)


	c3,_ := f.GetCellValue("Sheet1","A3")
	fmt.Println("Value in A3 : ",c3)

}