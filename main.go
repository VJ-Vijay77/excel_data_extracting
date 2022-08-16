package main

import (
	"log"

	"github.com/xuri/excelize/v2"
)

func main(){
	f := excelize.NewFile()
	f.SetCellValue("Sheet1","A1",100)
	f.SetCellValue("Sheet1","B2",50)
	f.SetCellValue("Sheet1","A2","testingdone")
	f.SetCellValue("Sheet1","A3","vijaydinesh77vj@gmail.com")
	if err := f.SaveAs("test.xlsx"); err != nil {
		log.Fatalln(err)
	}

}