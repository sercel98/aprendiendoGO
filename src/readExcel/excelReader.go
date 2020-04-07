package main

import (
	"fmt"
	"github.com/360EntSecGroup-Skylar/excelize"
	"strconv"
)

func main() {

	xlsx, err := excelize.OpenFile("./src/readExcel/docentes.xlsx")
	if err != nil {
		fmt.Println(err)
		return
	}

	// Obtengo la filas de la hoja "asignacion"
	rows, _ := xlsx.GetRows("asignacion")

	for i, row := range rows {

		if i == 0 {
			continue
		}

		colaboradorName := row[17]
		colaboradorId := returnColaboradorId(colaboradorName, xlsx)

		grupoName := row[15]
		grupoId := returnGrupoId(grupoName, xlsx)

		index := strconv.Itoa(i)

		if colaboradorId > 0 {
			xlsx.SetCellValue("asignacion", "E"+index, colaboradorId)
		}
		if grupoId > 0 {
			xlsx.SetCellValue("asignacion", "C"+index, grupoId)
		}

		fmt.Print(colaboradorId, "\t")
		fmt.Print(colaboradorName, "\t")
		fmt.Println()
	}

	// Save xlsx file by the given path.
	if err := xlsx.SaveAs("./src/readExcel/docentes.xlsx"); err != nil {
		fmt.Println(err)
	}

}

func returnColaboradorId(name string, file *excelize.File) int {

	rows, _ := file.GetRows("docentes")
	for _, row := range rows {
		if row[1] == name {
			colaboradorId, error := strconv.Atoi(row[0])
			if error != nil {
				fmt.Println(error)
				return -1
			}
			return colaboradorId
		}
	}
	return -1
}

func returnGrupoId(name string, file *excelize.File) int {

	rows, _ := file.GetRows("grupos")
	for _, row := range rows {
		if row[1] == name {
			groupId, error := strconv.Atoi(row[0])
			if error != nil {
				fmt.Println(error)
				return -1
			}
			return groupId
		}
	}

	return -1
}
