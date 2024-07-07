package controllers

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/go-pdf/fpdf"
)

func TestReceipt(c *gin.Context) {
	pdf := fpdf.New("L", "mm", "A5", "")
	pdf.AddPage()
	pdf.Ln(8)
	pdf.SetFont("Arial", "B", 16)
	pdf.CellFormat(55, 0, "", "", 0, "C", false, 0, "")
	pdf.CellFormat(134, 0, "BUKTI PEMBAYARAN", "", 1, "C", false, 0, "")
	pdf.Ln(10)
	pdf.SetFont("Arial", "B", 8)
	pdf.CellFormat(55, 6, "Yth. Bapak/Ibu", "", 0, "L", false, 0, "")
	pdf.CellFormat(27, 6, "No. Pelanggan", "TB", 0, "L", false, 0, "")
	pdf.CellFormat(22, 6, "Golongan", "TB", 0, "L", false, 0, "")
	pdf.CellFormat(20, 6, "Rayon", "TB", 0, "L", false, 0, "")
	pdf.CellFormat(27, 6, "Bulan", "TB", 0, "L", false, 0, "")
	pdf.SetFont("Arial", "B", 9)
	pdf.CellFormat(38, 6, "#1234567890123456", "", 1, "R", false, 0, "")
	pdf.SetFont("Arial", "B", 8)
	pdf.CellFormat(55, 6, "Asep Suprapto, S.Kom", "", 0, "L", false, 0, "")
	pdf.SetFontStyle("")
	pdf.CellFormat(27, 6, "00000000", "", 0, "L", false, 0, "")
	pdf.CellFormat(22, 6, "R2222", "", 0, "L", false, 0, "")
	pdf.CellFormat(20, 6, "BC0101", "", 0, "L", false, 0, "")
	pdf.CellFormat(27, 6, "September 2024", "", 1, "L", false, 0, "")
	pdf.CellFormat(55, 6, "Jl. Karang Luas I Gg. Keluarga", "", 2, "L", false, 0, "")
	pdf.SetFont("Arial", "B", 10)
	pdf.CellFormat(55, 6, "", "", 0, "L", false, 0, "")
	pdf.CellFormat(134, 10, "RINCIAN PERHITUNGAN BIAYA AIR", "", 1, "C", false, 0, "")
	pdf.SetFont("Arial", "", 8)
	pdf.CellFormat(55, 6, "Web : http://www.tirtadeli.co.id", "", 0, "L", false, 0, "")
	pdf.SetFont("Arial", "B", 8)
	pdf.CellFormat(27, 6, "Tgl. Baca", "TB", 0, "L", false, 0, "")
	pdf.CellFormat(22, 6, "Lalu", "TB", 0, "L", false, 0, "")
	pdf.CellFormat(20, 6, "Kini", "TB", 0, "L", false, 0, "")
	pdf.CellFormat(27, 6, "Pakai", "TB", 0, "L", false, 0, "")
	pdf.CellFormat(38, 6, "Total", "TB", 1, "L", false, 0, "")
	pdf.SetFont("Arial", "", 8)
	pdf.CellFormat(55, 3, "e-Mail : pdamtirtadeli@yahoo.co.id", "", 0, "L", false, 0, "")
	pdf.CellFormat(27, 6, "1", "", 0, "L", false, 0, "")
	pdf.CellFormat(22, 6, "10", "", 0, "L", false, 0, "")
	pdf.CellFormat(20, 6, "30", "", 0, "L", false, 0, "")
	pdf.CellFormat(27, 6, "20", "", 0, "L", false, 0, "")
	pdf.CellFormat(38, 6, "100.000", "", 1, "L", false, 0, "")
	pdf.SetFont("Arial", "B", 10)
	pdf.CellFormat(55, 6, "", "", 0, "L", false, 0, "")
	pdf.CellFormat(134, 10, "RINGKASAN BIAYA", "", 1, "C", false, 0, "")
	pdf.SetFont("Arial", "", 8)
	pdf.CellFormat(55, 4, "No. Telp : 061 - 795 2911", "", 0, "L", false, 0, "")
	pdf.CellFormat(45, 4, "Pemakaian Air :", "", 0, "R", false, 0, "")
	pdf.CellFormat(50, 4, "80.000", "", 1, "R", false, 0, "")
	pdf.CellFormat(55, 4, "(Kantor Pusat Lubuk Pakam)", "", 0, "L", false, 0, "")
	pdf.CellFormat(45, 4, "Beban Tetap :", "", 0, "R", false, 0, "")
	pdf.CellFormat(50, 4, "20.000", "", 1, "R", false, 0, "")
	pdf.CellFormat(55, 4, "", "", 0, "L", false, 0, "")
	pdf.CellFormat(45, 4, "Angsuran :", "", 0, "R", false, 0, "")
	pdf.CellFormat(50, 4, "0", "", 1, "R", false, 0, "")
	pdf.CellFormat(55, 4, "No. HP : 0812 6974 6240", "", 0, "L", false, 0, "")
	pdf.CellFormat(45, 4, "Denda :", "", 0, "R", false, 0, "")
	pdf.CellFormat(50, 4, "0", "", 1, "R", false, 0, "")
	pdf.CellFormat(55, 4, "(Pengaduan Pelanggan)", "", 0, "L", false, 0, "")
	pdf.CellFormat(45, 4, "Materai :", "", 0, "R", false, 0, "")
	pdf.CellFormat(50, 4, "0", "", 1, "R", false, 0, "")
	pdf.CellFormat(55, 4, "", "", 0, "L", false, 0, "")
	pdf.CellFormat(45, 4, "Total Bulan Ini :", "", 0, "R", false, 0, "")
	pdf.CellFormat(50, 4, "100.000", "", 1, "R", false, 0, "")
	pdf.Ln(2)
	pdf.CellFormat(55, 5, "", "", 0, "L", false, 0, "")
	pdf.SetFontStyle("B")
	pdf.CellFormat(27, 5, "Terbilang :", "TB", 0, "R", false, 0, "")
	pdf.CellFormat(108, 5, "SERATUS SEMBILAN PULUH SEMBILAN RIBU LIMA RATUS RUPIAH", "TB", 1, "C", false, 0, "")
	pdf.Ln(8)
	pdf.SetFontStyle("I")
	pdf.CellFormat(190, 5, "[16/04/2024 10:00|192.168.135.33|1602019bc]", "", 1, "R", false, 0, "")
	pdf.CellFormat(190, 5, "Rekening ini dibuat oleh komputer, tanda tangan pejabat PDAM tidak diperlukan dan sebagai bukti pembayaran yang sah.", "", 1, "R", false, 0, "")
	pdf.OutputFileAndClose("hello2.pdf")
}

func strDelimit(str string, sepstr string, sepcount int) string {
	pos := len(str) - sepcount
	for pos > 0 {
		str = str[:pos] + sepstr + str[pos:]
		pos = pos - sepcount
	}
	return str
}

func TestPdf(c *gin.Context) {
	pdf := fpdf.New("P", "mm", "A4", "")
	type countryType struct {
		nameStr, capitalStr, areaStr, popStr string
	}
	countryList := make([]countryType, 0, 8)
	header := []string{"Country", "Capital", "Area (sq km)", "Pop. (thousands)"}
	loadData := func(fileStr string) {
		fl, err := os.Open(fileStr)
		if err == nil {
			scanner := bufio.NewScanner(fl)
			var c countryType
			for scanner.Scan() {
				// Austria;Vienna;83859;8075
				lineStr := scanner.Text()
				list := strings.Split(lineStr, ";")
				if len(list) == 4 {
					c.nameStr = list[0]
					c.capitalStr = list[1]
					c.areaStr = list[2]
					c.popStr = list[3]
					countryList = append(countryList, c)
				} else {
					err = fmt.Errorf("error tokenizing %s", lineStr)
				}
			}
			fl.Close()
			if len(countryList) == 0 {
				err = fmt.Errorf("error loading data from %s", fileStr)
			}
		}
		if err != nil {
			pdf.SetError(err)
		}
	}
	// Simple table
	basicTable := func() {
		left := (210.0 - 4*40) / 2
		pdf.SetX(left)
		for _, str := range header {
			pdf.CellFormat(40, 7, str, "1", 0, "", false, 0, "")
		}
		pdf.Ln(-1)
		for _, c := range countryList {
			pdf.SetX(left)
			pdf.CellFormat(40, 6, c.nameStr, "1", 0, "", false, 0, "")
			pdf.CellFormat(40, 6, c.capitalStr, "1", 0, "", false, 0, "")
			pdf.CellFormat(40, 6, c.areaStr, "1", 0, "", false, 0, "")
			pdf.CellFormat(40, 6, c.popStr, "1", 0, "", false, 0, "")
			pdf.Ln(-1)
		}
	}
	// Better table
	improvedTable := func() {
		// Column widths
		w := []float64{40.0, 35.0, 40.0, 45.0}
		wSum := 0.0
		for _, v := range w {
			wSum += v
		}
		left := (210 - wSum) / 2
		// 	Header
		pdf.SetX(left)
		for j, str := range header {
			pdf.CellFormat(w[j], 7, str, "1", 0, "C", false, 0, "")
		}
		pdf.Ln(-1)
		// Data
		for _, c := range countryList {
			pdf.SetX(left)
			pdf.CellFormat(w[0], 6, c.nameStr, "LR", 0, "", false, 0, "")
			pdf.CellFormat(w[1], 6, c.capitalStr, "LR", 0, "", false, 0, "")
			pdf.CellFormat(w[2], 6, strDelimit(c.areaStr, ",", 3),
				"LR", 0, "R", false, 0, "")
			pdf.CellFormat(w[3], 6, strDelimit(c.popStr, ",", 3),
				"LR", 0, "R", false, 0, "")
			pdf.Ln(-1)
		}
		pdf.SetX(left)
		pdf.CellFormat(wSum, 0, "", "T", 0, "", false, 0, "")
	}
	// Colored table
	fancyTable := func() {
		// Colors, line width and bold font
		pdf.SetFillColor(255, 0, 0)
		pdf.SetTextColor(255, 255, 255)
		pdf.SetDrawColor(128, 0, 0)
		pdf.SetLineWidth(.3)
		pdf.SetFont("", "B", 0)
		// 	Header
		w := []float64{40, 35, 40, 45}
		wSum := 0.0
		for _, v := range w {
			wSum += v
		}
		left := (210 - wSum) / 2
		pdf.SetX(left)
		for j, str := range header {
			pdf.CellFormat(w[j], 7, str, "1", 0, "C", true, 0, "")
		}
		pdf.Ln(-1)
		// Color and font restoration
		pdf.SetFillColor(224, 235, 255)
		pdf.SetTextColor(0, 0, 0)
		pdf.SetFont("", "", 0)
		// 	Data
		fill := false
		for _, c := range countryList {
			pdf.SetX(left)
			pdf.CellFormat(w[0], 6, c.nameStr, "LR", 0, "", fill, 0, "")
			pdf.CellFormat(w[1], 6, c.capitalStr, "LR", 0, "", fill, 0, "")
			pdf.CellFormat(w[2], 6, strDelimit(c.areaStr, ",", 3),
				"LR", 0, "R", fill, 0, "")
			pdf.CellFormat(w[3], 6, strDelimit(c.popStr, ",", 3),
				"LR", 0, "R", fill, 0, "")
			pdf.Ln(-1)
			fill = !fill
		}
		pdf.SetX(left)
		pdf.CellFormat(wSum, 0, "", "T", 0, "", false, 0, "")
	}
	loadData("countries.txt")
	pdf.SetFont("Arial", "", 14)
	pdf.AddPage()
	basicTable()
	pdf.AddPage()
	improvedTable()
	pdf.AddPage()
	fancyTable()
	pdf.OutputFileAndClose("hello.pdf")
}
