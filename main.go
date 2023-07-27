package main

import (
	"fmt"
	"log"

	"github.com/xuri/excelize/v2"
)

func main() {
	filePath := "200.xlsx"
	xlsxFile, err := excelize.OpenFile(filePath)
	if err != nil {
		log.Fatalf("Error opening XLSX file: %v", err)
	}
	sheet := xlsxFile.GetSheetList()
	fmt.Println("sheet name ", sheet)
	// Set the row number for the row to be formatted (10th row in this case)
	rowNumbers := []int{5, 8, 10}
	columns, err := xlsxFile.GetCols("Sheet1")
	if err != nil {
		log.Fatalf("Error getting column count: %v", err)
	}
	totalCols := len(columns)
	fill := excelize.Fill{Type: "pattern", Pattern: 1, Color: []string{"#FF0000"}}
	styleID, err := xlsxFile.NewStyle(&excelize.Style{Fill: fill})
	if err != nil {
		log.Fatalf("Error creating style: %v", err)
	}

	// Apply the style to each cell in the 10th row
	for _, rowNumber := range rowNumbers {
		for col := 1; col <= totalCols; col++ {
			cellName, err := excelize.CoordinatesToCellName(col, rowNumber)
			if err != nil {
				log.Fatalf("Error converting coordinates to cell name: %v", err)
			}
			xlsxFile.SetCellStyle("Sheet1", cellName, cellName, styleID)
		}
	}

	newFilePath := "output.xlsx"
	if err := xlsxFile.SaveAs(newFilePath); err != nil {
		log.Fatalf("Error saving XLSX file: %v", err)
	}
	fmt.Printf("Successfully added a red background color  '%s'.\n", newFilePath)

}
