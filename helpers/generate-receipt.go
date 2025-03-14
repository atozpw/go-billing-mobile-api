package helpers

import (
	"strconv"
	"time"

	"github.com/atozpw/go-billing-mobile-api/configs"
	"github.com/go-pdf/fpdf"
)

func GenerateReceiptTirtaOgan(id string) {

	var receipt struct {
		ByrNo       string
		PelNo       string
		PelNama     string
		PelAlamat   string
		RekGol      string
		DkdKd       string
		RekPeriode  string
		TglBaca     string
		RekStanlalu string
		RekStankini string
		RekPakai    string
		RekUangair  string
		RekBeban    string
		RekAngsuran string
		RekDenda    string
		RekMaterai  string
		ByrTotal    string
		LokIp       string
		KarId       string
	}

	configs.DB.Raw("SELECT b.byr_no, a.pel_no, a.pel_nama, a.pel_alamat, a.rek_gol, a.dkd_kd, CONCAT(MONTHNAME_ID(a.rek_bln), ' ', a.rek_thn) AS rek_periode, a.tgl_baca, a.rek_stanlalu, a.rek_stankini, (a.rek_stankini - a.rek_stanlalu) AS rek_pakai, a.rek_uangair, (a.rek_adm + a.rek_meter) AS rek_beban, a.rek_angsuran, a.rek_denda, a.rek_materai, b.byr_total, b.lok_ip, b.kar_id FROM tm_rekening a JOIN tm_pembayaran b ON b.rek_nomor = a.rek_nomor AND b.byr_sts > 0 WHERE a.rek_nomor = ?", id).Scan(&receipt)

	strToIntUangAir, _ := strconv.Atoi(receipt.RekUangair)
	strToIntBeban, _ := strconv.Atoi(receipt.RekBeban)
	strToIntAngsuran, _ := strconv.Atoi(receipt.RekAngsuran)
	strToIntDenda, _ := strconv.Atoi(receipt.RekDenda)
	strToIntMaterai, _ := strconv.Atoi(receipt.RekMaterai)
	strToIntTotal, _ := strconv.Atoi(receipt.ByrTotal)

	pdf := fpdf.New("L", "mm", "A5", "")
	pdf.AddPage()
	pdf.Ln(10)
	pdf.Image("images/logo-ogan-ilir.jpg", 11, 13, 15, 0, false, "", 0, "")
	pdf.SetFont("Arial", "B", 16)
	pdf.CellFormat(60, 0, "", "", 0, "C", false, 0, "")
	pdf.CellFormat(130, 0, "BUKTI PEMBAYARAN", "", 1, "C", false, 0, "")
	pdf.Ln(10)
	pdf.SetFont("Arial", "B", 8)
	pdf.CellFormat(60, 5, "Yth. Bapak/Ibu", "", 0, "L", false, 0, "")
	pdf.CellFormat(26, 5, "No. Pelanggan", "TB", 0, "L", false, 0, "")
	pdf.CellFormat(21, 5, "Golongan", "TB", 0, "L", false, 0, "")
	pdf.CellFormat(19, 5, "Rayon", "TB", 0, "L", false, 0, "")
	pdf.CellFormat(26, 5, "Bulan", "TB", 0, "L", false, 0, "")
	pdf.SetFont("Arial", "B", 9)
	pdf.CellFormat(38, 5, "#"+receipt.ByrNo, "", 1, "R", false, 0, "")
	pdf.SetFont("Arial", "B", 8)
	pdf.CellFormat(60, 5, receipt.PelNama, "", 0, "L", false, 0, "")
	pdf.SetFontStyle("")
	pdf.CellFormat(26, 5, receipt.PelNo, "", 0, "L", false, 0, "")
	pdf.CellFormat(21, 5, receipt.RekGol, "", 0, "L", false, 0, "")
	pdf.CellFormat(19, 5, receipt.DkdKd, "", 0, "L", false, 0, "")
	pdf.CellFormat(26, 5, receipt.RekPeriode, "", 1, "L", false, 0, "")
	pdf.CellFormat(60, 5, receipt.PelAlamat, "", 2, "L", false, 0, "")
	pdf.SetFont("Arial", "B", 10)
	pdf.CellFormat(60, 5, "", "", 0, "L", false, 0, "")
	pdf.CellFormat(130, 5, "RINCIAN PERHITUNGAN BIAYA AIR", "", 1, "C", false, 0, "")
	pdf.Ln(2)
	pdf.SetFont("Arial", "", 8)
	pdf.CellFormat(60, 5, "Bayarlah rekening secara tepat waktu", "", 0, "L", false, 0, "")
	pdf.SetFont("Arial", "B", 8)
	pdf.CellFormat(26, 5, "Tanggal Baca", "TB", 0, "L", false, 0, "")
	pdf.CellFormat(21, 5, "Stand Lalu", "TB", 0, "L", false, 0, "")
	pdf.CellFormat(19, 5, "Stand Kini", "TB", 0, "L", false, 0, "")
	pdf.CellFormat(26, 5, "Pemakaian", "TB", 0, "L", false, 0, "")
	pdf.CellFormat(38, 5, "Biaya Air", "TB", 1, "L", false, 0, "")
	pdf.SetFont("Arial", "", 8)
	pdf.CellFormat(60, 5, "untuk menghindari denda dan penutupan", "", 0, "L", false, 0, "")
	pdf.CellFormat(26, 5, receipt.TglBaca, "", 0, "L", false, 0, "")
	pdf.CellFormat(21, 5, receipt.RekStanlalu, "", 0, "L", false, 0, "")
	pdf.CellFormat(19, 5, receipt.RekStankini, "", 0, "L", false, 0, "")
	pdf.CellFormat(26, 5, receipt.RekPakai, "", 0, "L", false, 0, "")
	pdf.CellFormat(38, 5, CurrencyFormat(strToIntUangAir), "", 1, "L", false, 0, "")
	pdf.CellFormat(60, 5, "sambungan instalasi air minum.", "", 1, "L", false, 0, "")
	pdf.CellFormat(60, 5, "", "", 0, "L", false, 0, "")
	pdf.SetFont("Arial", "B", 10)
	pdf.CellFormat(130, 5, "RINGKASAN BIAYA", "", 1, "C", false, 0, "")
	pdf.Ln(2)
	pdf.SetFont("Arial", "", 8)
	pdf.CellFormat(60, 5, "Email : pdam.t@yahoo.com", "", 0, "L", false, 0, "")
	pdf.CellFormat(45, 5, "Pemakaian Air :", "", 0, "R", false, 0, "")
	pdf.CellFormat(50, 5, CurrencyFormat(strToIntUangAir), "", 1, "R", false, 0, "")
	pdf.CellFormat(60, 5, "Facebook : pdamtirtaoganilir", "", 0, "L", false, 0, "")
	pdf.CellFormat(45, 5, "Beban Tetap :", "", 0, "R", false, 0, "")
	pdf.CellFormat(50, 5, CurrencyFormat(strToIntBeban), "", 1, "R", false, 0, "")
	pdf.CellFormat(60, 5, "Web : tirtaogan.com", "", 0, "L", false, 0, "")
	pdf.CellFormat(45, 5, "Angsuran :", "", 0, "R", false, 0, "")
	pdf.CellFormat(50, 5, CurrencyFormat(strToIntAngsuran), "", 1, "R", false, 0, "")
	pdf.CellFormat(60, 5, "", "", 0, "L", false, 0, "")
	pdf.CellFormat(45, 5, "Denda :", "", 0, "R", false, 0, "")
	pdf.CellFormat(50, 5, CurrencyFormat(strToIntDenda), "", 1, "R", false, 0, "")
	pdf.CellFormat(60, 5, "Kantor Pusat :", "", 0, "L", false, 0, "")
	pdf.CellFormat(45, 5, "Materai :", "", 0, "R", false, 0, "")
	pdf.CellFormat(50, 5, CurrencyFormat(strToIntMaterai), "", 1, "R", false, 0, "")
	pdf.CellFormat(60, 5, "Telp : 0711-7584400", "", 0, "L", false, 0, "")
	pdf.CellFormat(45, 5, "Total Bulan Ini :", "", 0, "R", false, 0, "")
	pdf.CellFormat(50, 5, CurrencyFormat(strToIntTotal), "", 1, "R", false, 0, "")
	pdf.Ln(2)
	pdf.SetFontStyle("B")
	pdf.CellFormat(60, 5, "NPWP : 02.543.074.5.312.000", "", 0, "L", false, 0, "")
	pdf.CellFormat(130, 5, "Terbilang : "+NumberToWords(strToIntTotal)+" Rupiah", "TB", 1, "C", false, 0, "")
	pdf.Ln(4)
	pdf.SetFontStyle("I")
	pdf.SetFontSize(7)
	pdf.CellFormat(190, 4, "["+time.Now().Format("02/01/2006 15:04")+"|"+receipt.LokIp+"|"+receipt.KarId+"]", "", 1, "R", false, 0, "")
	pdf.CellFormat(190, 4, "Rekening ini dibuat oleh komputer, tanda tangan pejabat PDAM tidak diperlukan dan sebagai bukti pembayaran yang sah.", "", 1, "R", false, 0, "")
	pdf.OutputFileAndClose("storages/public/INV-" + id + ".pdf")

}

func GenerateReceiptTirtaDeli(id string) {

	var receipt struct {
		ByrNo       string
		PelNo       string
		PelNama     string
		PelAlamat   string
		RekGol      string
		DkdKd       string
		RekPeriode  string
		TglBaca     string
		RekStanlalu string
		RekStankini string
		RekPakai    string
		RekUangair  string
		RekBeban    string
		RekAngsuran string
		RekDenda    string
		RekMaterai  string
		ByrTotal    string
		LokIp       string
		KarId       string
	}

	configs.DB.Raw("SELECT b.byr_no, a.pel_no, a.pel_nama, a.pel_alamat, a.rek_gol, a.dkd_kd, CONCAT(MONTHNAME_ID(a.rek_bln), ' ', a.rek_thn) AS rek_periode, a.tgl_baca, a.rek_stanlalu, a.rek_stankini, (a.rek_stankini - a.rek_stanlalu) AS rek_pakai, a.rek_uangair, (a.rek_adm + a.rek_meter) AS rek_beban, a.rek_angsuran, a.rek_denda, a.rek_materai, b.byr_total, b.lok_ip, b.kar_id FROM tm_rekening a JOIN tm_pembayaran b ON b.rek_nomor = a.rek_nomor AND b.byr_sts > 0 WHERE a.rek_nomor = ?", id).Scan(&receipt)

	strToIntUangAir, _ := strconv.Atoi(receipt.RekUangair)
	strToIntBeban, _ := strconv.Atoi(receipt.RekBeban)
	strToIntAngsuran, _ := strconv.Atoi(receipt.RekAngsuran)
	strToIntDenda, _ := strconv.Atoi(receipt.RekDenda)
	strToIntMaterai, _ := strconv.Atoi(receipt.RekMaterai)
	strToIntTotal, _ := strconv.Atoi(receipt.ByrTotal)

	pdf := fpdf.New("L", "mm", "A5", "")
	pdf.AddPage()
	pdf.Ln(10)
	pdf.Image("images/logo-deli-serdang.png", 11, 13, 15, 0, false, "", 0, "")
	pdf.SetFont("Arial", "B", 16)
	pdf.CellFormat(60, 0, "", "", 0, "C", false, 0, "")
	pdf.CellFormat(130, 0, "BUKTI PEMBAYARAN", "", 1, "C", false, 0, "")
	pdf.Ln(10)
	pdf.SetFont("Arial", "B", 8)
	pdf.CellFormat(60, 5, "Yth. Bapak/Ibu", "", 0, "L", false, 0, "")
	pdf.CellFormat(26, 5, "No. Pelanggan", "TB", 0, "L", false, 0, "")
	pdf.CellFormat(21, 5, "Golongan", "TB", 0, "L", false, 0, "")
	pdf.CellFormat(19, 5, "Rayon", "TB", 0, "L", false, 0, "")
	pdf.CellFormat(26, 5, "Bulan", "TB", 0, "L", false, 0, "")
	pdf.SetFont("Arial", "B", 9)
	pdf.CellFormat(38, 5, "#"+receipt.ByrNo, "", 1, "R", false, 0, "")
	pdf.SetFont("Arial", "B", 8)
	pdf.CellFormat(60, 5, receipt.PelNama, "", 0, "L", false, 0, "")
	pdf.SetFontStyle("")
	pdf.CellFormat(26, 5, receipt.PelNo, "", 0, "L", false, 0, "")
	pdf.CellFormat(21, 5, receipt.RekGol, "", 0, "L", false, 0, "")
	pdf.CellFormat(19, 5, receipt.DkdKd, "", 0, "L", false, 0, "")
	pdf.CellFormat(26, 5, receipt.RekPeriode, "", 1, "L", false, 0, "")
	pdf.CellFormat(60, 5, receipt.PelAlamat, "", 2, "L", false, 0, "")
	pdf.SetFont("Arial", "B", 10)
	pdf.CellFormat(60, 5, "", "", 0, "L", false, 0, "")
	pdf.CellFormat(130, 5, "RINCIAN PERHITUNGAN BIAYA AIR", "", 1, "C", false, 0, "")
	pdf.Ln(2)
	pdf.SetFont("Arial", "", 8)
	pdf.CellFormat(60, 5, "Bayarlah rekening secara tepat waktu", "", 0, "L", false, 0, "")
	pdf.SetFont("Arial", "B", 8)
	pdf.CellFormat(26, 5, "Tanggal Baca", "TB", 0, "L", false, 0, "")
	pdf.CellFormat(21, 5, "Stand Lalu", "TB", 0, "L", false, 0, "")
	pdf.CellFormat(19, 5, "Stand Kini", "TB", 0, "L", false, 0, "")
	pdf.CellFormat(26, 5, "Pemakaian", "TB", 0, "L", false, 0, "")
	pdf.CellFormat(38, 5, "Biaya Air", "TB", 1, "L", false, 0, "")
	pdf.SetFont("Arial", "", 8)
	pdf.CellFormat(60, 5, "untuk menghindari denda dan penutupan", "", 0, "L", false, 0, "")
	pdf.CellFormat(26, 5, receipt.TglBaca, "", 0, "L", false, 0, "")
	pdf.CellFormat(21, 5, receipt.RekStanlalu, "", 0, "L", false, 0, "")
	pdf.CellFormat(19, 5, receipt.RekStankini, "", 0, "L", false, 0, "")
	pdf.CellFormat(26, 5, receipt.RekPakai, "", 0, "L", false, 0, "")
	pdf.CellFormat(38, 5, CurrencyFormat(strToIntUangAir), "", 1, "L", false, 0, "")
	pdf.CellFormat(60, 5, "sambungan instalasi air minum.", "", 1, "L", false, 0, "")
	pdf.CellFormat(60, 5, "", "", 0, "L", false, 0, "")
	pdf.SetFont("Arial", "B", 10)
	pdf.CellFormat(130, 5, "RINGKASAN BIAYA", "", 1, "C", false, 0, "")
	pdf.Ln(2)
	pdf.SetFont("Arial", "", 8)
	pdf.CellFormat(60, 5, "Email : perumdam@tirtadeli.co.id", "", 0, "L", false, 0, "")
	pdf.CellFormat(45, 5, "Pemakaian Air :", "", 0, "R", false, 0, "")
	pdf.CellFormat(50, 5, CurrencyFormat(strToIntUangAir), "", 1, "R", false, 0, "")
	pdf.CellFormat(60, 5, "Web : http://www.tirtadeli.co.id", "", 0, "L", false, 0, "")
	pdf.CellFormat(45, 5, "Beban Tetap :", "", 0, "R", false, 0, "")
	pdf.CellFormat(50, 5, CurrencyFormat(strToIntBeban), "", 1, "R", false, 0, "")
	pdf.CellFormat(60, 5, "Kantor Pusat Lubuk Pakam", "", 0, "L", false, 0, "")
	pdf.CellFormat(45, 5, "Angsuran :", "", 0, "R", false, 0, "")
	pdf.CellFormat(50, 5, CurrencyFormat(strToIntAngsuran), "", 1, "R", false, 0, "")
	pdf.CellFormat(60, 5, "No. Telp : 061 795 2911", "", 0, "L", false, 0, "")
	pdf.CellFormat(45, 5, "Denda :", "", 0, "R", false, 0, "")
	pdf.CellFormat(50, 5, CurrencyFormat(strToIntDenda), "", 1, "R", false, 0, "")
	pdf.CellFormat(60, 5, "Pengaduan Pelanggan", "", 0, "L", false, 0, "")
	pdf.CellFormat(45, 5, "Materai :", "", 0, "R", false, 0, "")
	pdf.CellFormat(50, 5, CurrencyFormat(strToIntMaterai), "", 1, "R", false, 0, "")
	pdf.CellFormat(60, 5, "No. HP : 0812 6974 6240", "", 0, "L", false, 0, "")
	pdf.CellFormat(45, 5, "Total Bulan Ini :", "", 0, "R", false, 0, "")
	pdf.CellFormat(50, 5, CurrencyFormat(strToIntTotal), "", 1, "R", false, 0, "")
	pdf.Ln(2)
	pdf.SetFontStyle("B")
	pdf.CellFormat(60, 5, "NPWP : 01.457.276.2-125.000", "", 0, "L", false, 0, "")
	pdf.CellFormat(130, 5, "Terbilang : "+NumberToWords(strToIntTotal)+" Rupiah", "TB", 1, "C", false, 0, "")
	pdf.Ln(4)
	pdf.SetFontStyle("I")
	pdf.SetFontSize(7)
	pdf.CellFormat(190, 4, "["+time.Now().Format("02/01/2006 15:04")+"|"+receipt.LokIp+"|"+receipt.KarId+"]", "", 1, "R", false, 0, "")
	pdf.CellFormat(190, 4, "Rekening ini dibuat oleh komputer, tanda tangan pejabat PDAM tidak diperlukan dan sebagai bukti pembayaran yang sah.", "", 1, "R", false, 0, "")
	pdf.OutputFileAndClose("storages/public/INV-" + id + ".pdf")

}
