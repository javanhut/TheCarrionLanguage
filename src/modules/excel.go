package modules

import (
	"sync"

	"github.com/javanhut/TheCarrionLanguage/src/object"
	"github.com/xuri/excelize/v2"
)

// Global Excel file handle registry
var (
	excelHandles     = make(map[int64]*excelize.File)
	nextExcelHandle  int64 = 1
	excelHandleMutex sync.RWMutex
)

// Helper function to extract string value from object
func extractStringExcel(obj object.Object) (string, bool) {
	switch v := obj.(type) {
	case *object.String:
		return v.Value, true
	case *object.Instance:
		if value, exists := v.Env.Get("value"); exists {
			if strVal, ok := value.(*object.String); ok {
				return strVal.Value, true
			}
		}
		if v.Grimoire != nil && v.Grimoire.Name == "String" {
			if value, exists := v.Env.Get("value"); exists {
				if strVal, ok := value.(*object.String); ok {
					return strVal.Value, true
				}
			}
		}
		return v.Inspect(), true
	default:
		return obj.Inspect(), true
	}
}

// Helper function to extract integer value from object
func extractIntExcel(obj object.Object) (int64, bool) {
	switch v := obj.(type) {
	case *object.Integer:
		return v.Value, true
	default:
		return 0, false
	}
}

// Get Excel handle from registry
func getExcelHandle(handleID int64) (*excelize.File, bool) {
	excelHandleMutex.RLock()
	defer excelHandleMutex.RUnlock()
	file, exists := excelHandles[handleID]
	return file, exists
}

// Store Excel handle in registry
func storeExcelHandle(file *excelize.File) int64 {
	excelHandleMutex.Lock()
	defer excelHandleMutex.Unlock()
	handleID := nextExcelHandle
	nextExcelHandle++
	excelHandles[handleID] = file
	return handleID
}

// Remove Excel handle from registry
func removeExcelHandle(handleID int64) {
	excelHandleMutex.Lock()
	defer excelHandleMutex.Unlock()
	delete(excelHandles, handleID)
}

var ExcelBuiltins = map[string]*object.Builtin{
	"excelOpen": {
		Fn: func(args ...object.Object) object.Object {
			if len(args) != 1 {
				return &object.Error{Message: "excelOpen requires 1 argument: path"}
			}

			pathStr, ok := extractStringExcel(args[0])
			if !ok {
				return &object.Error{Message: "excelOpen: path must be a string"}
			}

			file, err := excelize.OpenFile(pathStr)
			if err != nil {
				return &object.Error{Message: "excelOpen: failed to open file '" + pathStr + "': " + err.Error()}
			}

			handleID := storeExcelHandle(file)
			return &object.Integer{Value: handleID}
		},
	},

	"excelCreate": {
		Fn: func(args ...object.Object) object.Object {
			file := excelize.NewFile()
			handleID := storeExcelHandle(file)
			return &object.Integer{Value: handleID}
		},
	},

	"excelClose": {
		Fn: func(args ...object.Object) object.Object {
			if len(args) != 1 {
				return &object.Error{Message: "excelClose requires 1 argument: handleID"}
			}

			handleID, ok := extractIntExcel(args[0])
			if !ok {
				return &object.Error{Message: "excelClose: handleID must be an integer"}
			}

			file, exists := getExcelHandle(handleID)
			if !exists {
				return &object.Error{Message: "excelClose: invalid excel handle"}
			}

			err := file.Close()
			removeExcelHandle(handleID)

			if err != nil {
				return &object.Error{Message: "excelClose: failed to close file: " + err.Error()}
			}

			return &object.None{}
		},
	},

	"excelGetSheets": {
		Fn: func(args ...object.Object) object.Object {
			if len(args) != 1 {
				return &object.Error{Message: "excelGetSheets requires 1 argument: handleID"}
			}

			handleID, ok := extractIntExcel(args[0])
			if !ok {
				return &object.Error{Message: "excelGetSheets: handleID must be an integer"}
			}

			file, exists := getExcelHandle(handleID)
			if !exists {
				return &object.Error{Message: "excelGetSheets: invalid excel handle"}
			}

			sheetList := file.GetSheetList()
			elements := make([]object.Object, len(sheetList))
			for i, name := range sheetList {
				elements[i] = &object.String{Value: name}
			}

			return &object.Array{Elements: elements}
		},
	},

	"excelReadSheet": {
		Fn: func(args ...object.Object) object.Object {
			if len(args) != 2 {
				return &object.Error{Message: "excelReadSheet requires 2 arguments: handleID, sheetName"}
			}

			handleID, ok := extractIntExcel(args[0])
			if !ok {
				return &object.Error{Message: "excelReadSheet: handleID must be an integer"}
			}

			sheetName, ok := extractStringExcel(args[1])
			if !ok {
				return &object.Error{Message: "excelReadSheet: sheetName must be a string"}
			}

			file, exists := getExcelHandle(handleID)
			if !exists {
				return &object.Error{Message: "excelReadSheet: invalid excel handle"}
			}

			rows, err := file.GetRows(sheetName)
			if err != nil {
				return &object.Error{Message: "excelReadSheet: failed to read sheet '" + sheetName + "': " + err.Error()}
			}

			// Convert to 2D array
			rowElements := make([]object.Object, len(rows))
			for i, row := range rows {
				cellElements := make([]object.Object, len(row))
				for j, cell := range row {
					cellElements[j] = &object.String{Value: cell}
				}
				rowElements[i] = &object.Array{Elements: cellElements}
			}

			return &object.Array{Elements: rowElements}
		},
	},

	"excelReadRow": {
		Fn: func(args ...object.Object) object.Object {
			if len(args) != 3 {
				return &object.Error{Message: "excelReadRow requires 3 arguments: handleID, sheetName, rowNum"}
			}

			handleID, ok := extractIntExcel(args[0])
			if !ok {
				return &object.Error{Message: "excelReadRow: handleID must be an integer"}
			}

			sheetName, ok := extractStringExcel(args[1])
			if !ok {
				return &object.Error{Message: "excelReadRow: sheetName must be a string"}
			}

			rowNum, ok := extractIntExcel(args[2])
			if !ok {
				return &object.Error{Message: "excelReadRow: rowNum must be an integer"}
			}

			file, exists := getExcelHandle(handleID)
			if !exists {
				return &object.Error{Message: "excelReadRow: invalid excel handle"}
			}

			rows, err := file.GetRows(sheetName)
			if err != nil {
				return &object.Error{Message: "excelReadRow: failed to read sheet '" + sheetName + "': " + err.Error()}
			}

			// Row numbers are 1-based
			if rowNum < 1 || int(rowNum) > len(rows) {
				return &object.Error{Message: "excelReadRow: row number out of range"}
			}

			row := rows[rowNum-1]
			cellElements := make([]object.Object, len(row))
			for i, cell := range row {
				cellElements[i] = &object.String{Value: cell}
			}

			return &object.Array{Elements: cellElements}
		},
	},

	"excelReadCell": {
		Fn: func(args ...object.Object) object.Object {
			if len(args) != 3 {
				return &object.Error{Message: "excelReadCell requires 3 arguments: handleID, sheetName, cellRef"}
			}

			handleID, ok := extractIntExcel(args[0])
			if !ok {
				return &object.Error{Message: "excelReadCell: handleID must be an integer"}
			}

			sheetName, ok := extractStringExcel(args[1])
			if !ok {
				return &object.Error{Message: "excelReadCell: sheetName must be a string"}
			}

			cellRef, ok := extractStringExcel(args[2])
			if !ok {
				return &object.Error{Message: "excelReadCell: cellRef must be a string (e.g., 'A1')"}
			}

			file, exists := getExcelHandle(handleID)
			if !exists {
				return &object.Error{Message: "excelReadCell: invalid excel handle"}
			}

			value, err := file.GetCellValue(sheetName, cellRef)
			if err != nil {
				return &object.Error{Message: "excelReadCell: failed to read cell: " + err.Error()}
			}

			return &object.String{Value: value}
		},
	},

	"excelWriteCell": {
		Fn: func(args ...object.Object) object.Object {
			if len(args) != 4 {
				return &object.Error{Message: "excelWriteCell requires 4 arguments: handleID, sheetName, cellRef, value"}
			}

			handleID, ok := extractIntExcel(args[0])
			if !ok {
				return &object.Error{Message: "excelWriteCell: handleID must be an integer"}
			}

			sheetName, ok := extractStringExcel(args[1])
			if !ok {
				return &object.Error{Message: "excelWriteCell: sheetName must be a string"}
			}

			cellRef, ok := extractStringExcel(args[2])
			if !ok {
				return &object.Error{Message: "excelWriteCell: cellRef must be a string (e.g., 'A1')"}
			}

			valueStr, ok := extractStringExcel(args[3])
			if !ok {
				return &object.Error{Message: "excelWriteCell: value must be a string"}
			}

			file, exists := getExcelHandle(handleID)
			if !exists {
				return &object.Error{Message: "excelWriteCell: invalid excel handle"}
			}

			err := file.SetCellValue(sheetName, cellRef, valueStr)
			if err != nil {
				return &object.Error{Message: "excelWriteCell: failed to write cell: " + err.Error()}
			}

			return &object.None{}
		},
	},

	"excelSave": {
		Fn: func(args ...object.Object) object.Object {
			if len(args) < 1 || len(args) > 2 {
				return &object.Error{Message: "excelSave requires 1-2 arguments: handleID, [path]"}
			}

			handleID, ok := extractIntExcel(args[0])
			if !ok {
				return &object.Error{Message: "excelSave: handleID must be an integer"}
			}

			file, exists := getExcelHandle(handleID)
			if !exists {
				return &object.Error{Message: "excelSave: invalid excel handle"}
			}

			var err error
			if len(args) == 2 {
				pathStr, ok := extractStringExcel(args[1])
				if !ok {
					return &object.Error{Message: "excelSave: path must be a string"}
				}
				err = file.SaveAs(pathStr)
			} else {
				err = file.Save()
			}

			if err != nil {
				return &object.Error{Message: "excelSave: failed to save file: " + err.Error()}
			}

			return &object.None{}
		},
	},

	"excelNewSheet": {
		Fn: func(args ...object.Object) object.Object {
			if len(args) != 2 {
				return &object.Error{Message: "excelNewSheet requires 2 arguments: handleID, sheetName"}
			}

			handleID, ok := extractIntExcel(args[0])
			if !ok {
				return &object.Error{Message: "excelNewSheet: handleID must be an integer"}
			}

			sheetName, ok := extractStringExcel(args[1])
			if !ok {
				return &object.Error{Message: "excelNewSheet: sheetName must be a string"}
			}

			file, exists := getExcelHandle(handleID)
			if !exists {
				return &object.Error{Message: "excelNewSheet: invalid excel handle"}
			}

			index, err := file.NewSheet(sheetName)
			if err != nil {
				return &object.Error{Message: "excelNewSheet: failed to create sheet: " + err.Error()}
			}

			return &object.Integer{Value: int64(index)}
		},
	},

	"excelDeleteSheet": {
		Fn: func(args ...object.Object) object.Object {
			if len(args) != 2 {
				return &object.Error{Message: "excelDeleteSheet requires 2 arguments: handleID, sheetName"}
			}

			handleID, ok := extractIntExcel(args[0])
			if !ok {
				return &object.Error{Message: "excelDeleteSheet: handleID must be an integer"}
			}

			sheetName, ok := extractStringExcel(args[1])
			if !ok {
				return &object.Error{Message: "excelDeleteSheet: sheetName must be a string"}
			}

			file, exists := getExcelHandle(handleID)
			if !exists {
				return &object.Error{Message: "excelDeleteSheet: invalid excel handle"}
			}

			err := file.DeleteSheet(sheetName)
			if err != nil {
				return &object.Error{Message: "excelDeleteSheet: failed to delete sheet: " + err.Error()}
			}

			return &object.None{}
		},
	},
}
