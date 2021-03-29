// download the package using "go get github.com/jung-kurt/gofpdf"

package main

import (
	"fmt"
	"io/ioutil"
	"log"

	"github.com/jung-kurt/gofpdf"
)

func main() {

	file := "radey.txt"

	content, err := ioutil.ReadFile(file)

	if err != nil {
		log.Fatalf("%s file not found", file)
	}

	pdf := gofpdf.New("p", "mm", "A4", "")

	pdf.AddPage()
	pdf.SetFont("Arial", "B", 14)
	pdf.MultiCell(198, 5, string(content), "0", "0", false)

	_ = pdf.OutputFileAndClose("radey.pdf")

	fmt.Println("PDF created")

}
