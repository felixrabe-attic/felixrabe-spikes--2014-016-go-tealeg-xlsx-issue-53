package main

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/tealeg/xlsx"
	"gopkg.in/felixrabe-go/misc.v0"
)

func main() {
	test("01_empty_file", emptyFile)
	test("02_empty_sheet", emptySheet)
	test("03_empty_row", emptyRow)
	test("04_empty_cell", emptyCell)
	test("05_one_value", oneValue)
	test("06_two_by_three", twoByThree)
}

// Tests

func emptyFile(f *xlsx.File) {}

func emptySheet(f *xlsx.File) {
	f.AddSheet("Test 1")
}

func emptyRow(f *xlsx.File) {
	f.AddSheet("Test 1").AddRow()
}

func emptyCell(f *xlsx.File) {
	f.AddSheet("Test 1").AddRow().AddCell()
}

func oneValue(f *xlsx.File) {
	f.AddSheet("Test 1").AddRow().AddCell().Value = "Test Value"
}

func twoByThree(f *xlsx.File) {
	sheet := f.AddSheet("Test 1")
	for iRow := 1; iRow < 3; iRow++ {
		row := sheet.AddRow()
		for iCol := 1; iCol < 4; iCol++ {
			cell := row.AddCell()
			cell.Value = fmt.Sprintf("Row %d Col %d", iRow, iCol)
		}
	}
}

// Infrastructure

var outputDir string

func init() {
	outputDir = misc.ThisDirJoin("output")
	if err := os.MkdirAll(outputDir, 0777); err != nil {
		misc.Fatal(err)
	}
}

func test(name string, fn func(*xlsx.File)) {
	file := xlsx.NewFile()
	fn(file)
	if err := file.Save(filepath.Join(outputDir, name+".xlsx")); err != nil {
		misc.Fatal(err)
	}
}
