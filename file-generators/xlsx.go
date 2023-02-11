package filegenerators

import (
	"fmt"
	"reflect"
	"report-generator/models"

	excelize "github.com/xuri/excelize/v2"
)

// type Row interface {
// 	models.VisitVrach
// }

func ServeExcel(visits []models.VisitVrach) (*excelize.File, func()) {

	f := excelize.NewFile()
	closeFn := func() {
		if err := f.Close(); err != nil {
			fmt.Println(err)
		}
	}

	sheetName := f.GetSheetName(f.GetActiveSheetIndex())
	streamWriter, err := f.NewStreamWriter(sheetName)
	if err != nil {
		fmt.Println(err)
		return nil, nil
	}

	if err != nil {
		fmt.Println(err)
		return nil, nil
	}

	header := make([]interface{}, 50)
	model_type := reflect.ValueOf(visits[0])
	clm_cnt := model_type.NumField()
	for colID := 0; colID < clm_cnt; colID++ {
		header[colID] = model_type.Type().Field(colID).Name
	}
	if err := streamWriter.SetRow("A1", header, excelize.RowOpts{Height: 45, Hidden: false}); err != nil {
		fmt.Println(err)
		return nil, nil
	}

	if err != nil {
		fmt.Println(err)
		return nil, nil
	}

	for rowID, visit := range visits {
		row := make([]interface{}, 50)
		model_type := reflect.ValueOf(visit)
		for colID := 0; colID < clm_cnt; colID++ {
			row[colID] = model_type.Field(colID).Interface()
		}

		cell, _ := excelize.CoordinatesToCellName(1, rowID+2)
		if err := streamWriter.SetRow(cell, row); err != nil {
			fmt.Println(err)
			return nil, nil
		}
	}

	if err := streamWriter.Flush(); err != nil {
		fmt.Println(err)
		return nil, nil
	}

	return f, closeFn
}
