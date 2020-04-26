package main

import (
	"errors"
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

	docentes, error := cargarDocentes(xlsx)
	grupos, error := cargarGrupos(xlsx)
	asignaturas, error := cargarAsignaturas(xlsx)

	if error != nil {
		fmt.Println(err)
		return
	}
	// Obtengo la filas de la hoja "asignacion"
	rows, _ := xlsx.GetRows("test")

	for i, row := range rows {

		if i == 0 {
			continue
		}

		colaboradorName := row[2]
		colaboradorId := docentes[colaboradorName]

		grupoName := row[0]
		grupoId := grupos[grupoName]

		asignaturaName := row[1]
		asignaturaId := asignaturas[asignaturaName]

		index := strconv.Itoa(i + 1)

		xlsx.SetCellValue("test", "E"+index, colaboradorId)
		xlsx.SetCellValue("test", "F"+index, grupoId)
		xlsx.SetCellValue("test", "G"+index, asignaturaId)
	}

	// Guardar Cambios
	if err := xlsx.SaveAs("./src/readExcel/Worksheet_Cervantes.xlsx"); err != nil {
		fmt.Println(err)
	}

}

func cargarGrupos(file *excelize.File) (map[string]int, error) {

	rows, error := file.GetRows("grupos")
	grupos := make(map[string]int)

	if error != nil {
		fmt.Println("error al cargar la hoja grupos")
		return nil, errors.New("error al cargar la hoja grupos")
	}

	for i, row := range rows {

		if i == 0 {
			continue
		}

		if i > 28 {
			break
		}
		idGrupo, err := strconv.Atoi(row[3])
		if err != nil {
			return nil, errors.New("error al convertir el id del grupo" + strconv.Itoa(i))
		}

		nombreGrupo := row[4]
		grupos[nombreGrupo] = idGrupo
	}
	return grupos, nil
}

func cargarDocentes(file *excelize.File) (map[string]int, error) {

	rows, error := file.GetRows("Docentes")
	docentes := make(map[string]int)

	if error != nil {
		fmt.Println("error al cargar la hoja docentes")
		return nil, errors.New("error al cargar la hoja docentes")
	}

	for i, row := range rows {
		if i == 0 {
			continue
		}

		if i > 35 {
			break
		}

		idDocente, err := strconv.Atoi(row[9])
		if err != nil {
			return nil, errors.New("error al convertir el id del docente en" + strconv.Itoa(i))
		}

		nombreDocente := row[6]
		docentes[nombreDocente] = idDocente

	}
	return docentes, nil
}

func cargarAsignaturas(file *excelize.File) (map[string]int, error) {

	rows, error := file.GetRows("Asignaturas")
	asignaturas := make(map[string]int)

	if error != nil {
		fmt.Println("error al cargar la hoja asignaturas")
		return nil, errors.New("error al cargar la hoja asignaturas")
	}

	for i, row := range rows {
		if i == 0 {
			continue
		}

		if i > 27 {
			break
		}

		idAsignatura, err := strconv.Atoi(row[4])
		if err != nil {
			return nil, errors.New("error al convertir el id de la asignatura" + strconv.Itoa(i))
		}

		nombreAsignatura := row[0]
		asignaturas[nombreAsignatura] = idAsignatura

	}
	return asignaturas, nil
}
