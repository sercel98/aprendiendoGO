package main

import (
	"fmt"
	"github.com/360EntSecGroup-Skylar/excelize"
	"strconv"
)

func main() {

	xlsx, err := excelize.OpenFile("./src/readExcel/Worksheet_Cervantes.xlsx")

	if err != nil {
		fmt.Println(err)
		return
	}

	asignaturasGrupoIDs := []string{}
	periodos := [5]string{"960","961","962","963","988"}
	// Obtengo la filas de la hoja "asignacion"
	rows, _ := xlsx.GetRows("asignacion_academina")

	for i, row := range rows {

		if i < 1 {
			continue
		}
		asignaturasGrupoIDs = append(asignaturasGrupoIDs, row[6])


	}

	rows, _ = xlsx.GetRows("periodos_asignaturas")
	contador := 0
	for i:=0; i < len(asignaturasGrupoIDs); i++ {

		text := asignaturasGrupoIDs[i]

		for j:= 0 ; j < 5 ; j++ {
			index := strconv.Itoa(contador + j + 2)
			xlsx.SetCellValue("periodos_asignaturas", "B"+index, text)
			textPeriodo := periodos[j]
			xlsx.SetCellValue("periodos_asignaturas", "C"+index, textPeriodo)
			if j < 3 {
				xlsx.SetCellValue("periodos_asignaturas", "D"+index, "10")
			} else if j == 3{
				xlsx.SetCellValue("periodos_asignaturas", "D"+index, "70")
			} else {
				xlsx.SetCellValue("periodos_asignaturas", "D"+index, "100")
			}

		}
		contador+=5
	}

	if err := xlsx.SaveAs("./src/readExcel/Worksheet_Cervantes.xlsx"); err != nil {
		fmt.Println(err)
	}

	}




