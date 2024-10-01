package pdfGenerator

import (
	"fmt"
)

func WriteHTMLToPDF(path string, templatePath string, data interface{}) error {

	r := NewRequestPdf("")

	//path for download pdf
	outputPath := path

	if err := r.ParseTemplate(templatePath, data); err == nil {

		// Generate PDF with custom arguments
		args := []string{"no-pdf-compression"}

		// Generate PDF
		ok, err := r.GeneratePDF(outputPath, args)

		if err != nil {
			return err
		}
		fmt.Println(ok, "pdf generated successfully")
	} else {
		return err
	}

	return nil
}
