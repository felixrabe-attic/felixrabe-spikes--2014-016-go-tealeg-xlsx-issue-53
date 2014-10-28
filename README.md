This is a test for https://github.com/tealeg/xlsx/issues/53.

The files have been output into `output/`, then opened with Excel 2010 (32-bit
German on Win 7) and immediately saved into `output-excel/`. (File
`01_empty_file.xlsx` could not be opened at all by Excel.)

The files have been produced using:

    go get github.com/tealeg/xlsx gopkg.in/felixrabe-go/misc.v0
    go run xlsx-test.go

The second repair dialog always mentioned this error message with the same line / column numbers (sorry for my German):

    Excel hat die Überprüfung und Reparatur auf Dateiebene abgeschlossen. Einige Teile dieser Arbeitsmappe wurden repariert oder verworfen.
    Reparierter Teil: /xl/worksheets/sheet1.xml-Part mit XML-Fehler.  Schwerwiegender Fehler Zeile 4, Spalte 17.

(/xl/worksheets/sheet1.xml-Part with XML error. Fatal error line 4, column 17.)
