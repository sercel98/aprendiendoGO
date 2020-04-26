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

	personas, error := cargarPersonas(xlsx)

	fmt.Println(len(personas))
	if error != nil {
		fmt.Println(err)
		return
	}

	// Obtengo la filas de la hoja "asignacion"
	rows, _ := xlsx.GetRows("Personas")

	for i, row := range rows {

		if i < 2 {
			continue
		}

		personaTipoDoc := row[5]

		var tipoDoc int
		switch personaTipoDoc {
		case "TI:TARJETA DE IDENTIDAD":
			tipoDoc = 2
		case "CC":
			tipoDoc = 1
		case "CC:CÉDULA DE CIUDADANÍA":
			tipoDoc = 1
		case "CE:CÉDULA DE EXTRANJERÍA":
			tipoDoc = 3
		case "NES:NÚMERO ESTABLECIDO POR LA SECRETARÍA":
			tipoDoc = 8
		case "NIP:NÚMERO DE IDENTIFICACIÓN PERSONAL":
			tipoDoc = 6
		case "NUIP:NÚMERO UNICO DE IDENTIFICACIÓN PERSONAL":
			tipoDoc = 7
		case "RC:REGISTRO CIVIL DE NACIMIENTO":
			tipoDoc = 5
		case "REGITRO CIVIL":
			tipoDoc = 5
		case "VENEZOLANO":
			tipoDoc = 3
		}
		personaNumDoc := row[6]

		personaId := personas[personaNumDoc][tipoDoc]

		index := strconv.Itoa(i + 1)

		xlsx.SetCellValue("Personas", "A"+index, personaId)
	}

	// Guardar Cambios
	if err := xlsx.SaveAs("./src/readExcel/Worksheet_Cervantes.xlsx"); err != nil {
		fmt.Println(err)
	}

}

func cargarPersonas(file *excelize.File) (map[string]map[int]int, error) {

	rows, error := file.GetRows("Personas_db")
	personas := make(map[string]map[int]int)

	if error != nil {
		return nil, errors.New("error al cargar la hoja Personas_db")
	}

	for i, row := range rows {

		if i == 0 {
			continue
		}
		if row[0] == "" {
			break
		}

		tipoDoc, err := strconv.Atoi(row[2])
		if err != nil {
			return nil, errors.New("error al convertir el id del tipo de documento " + strconv.Itoa(i))
		}

		numDocumento := row[3]
		persona_id, err := strconv.Atoi(row[0])
		if err != nil {
			return nil, errors.New("error al convertir el id de la persona" + strconv.Itoa(i))
		}
		documentMap := make(map[int]int)
		documentMap[tipoDoc] = persona_id
		personas[numDocumento] = documentMap
	}
	return personas, nil
}
