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
	pdf.CellFormat(60, 5, "Email : perumdam@tirtadeli.co.id", "", 0, "L", false, 0, "")
	pdf.SetFont("Arial", "B", 8)
	pdf.CellFormat(26, 5, "Tanggal Baca", "TB", 0, "L", false, 0, "")
	pdf.CellFormat(21, 5, "Stand Lalu", "TB", 0, "L", false, 0, "")
	pdf.CellFormat(19, 5, "Stand Kini", "TB", 0, "L", false, 0, "")
	pdf.CellFormat(26, 5, "Pemakaian", "TB", 0, "L", false, 0, "")
	pdf.CellFormat(38, 5, "Total", "TB", 1, "L", false, 0, "")
	pdf.SetFont("Arial", "", 8)
	pdf.CellFormat(60, 5, "Web : http://www.tirtadeli.co.id", "", 0, "L", false, 0, "")
	pdf.CellFormat(26, 5, receipt.TglBaca, "", 0, "L", false, 0, "")
	pdf.CellFormat(21, 5, receipt.RekStanlalu, "", 0, "L", false, 0, "")
	pdf.CellFormat(19, 5, receipt.RekStankini, "", 0, "L", false, 0, "")
	pdf.CellFormat(26, 5, receipt.RekPakai, "", 0, "L", false, 0, "")
	pdf.CellFormat(38, 5, CurrencyFormat(strToIntUangAir), "", 1, "L", false, 0, "")
	pdf.CellFormat(60, 5, "", "", 1, "L", false, 0, "")
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

func GenerateReceiptTirtaDeliSmall(id string) {

	var payment struct {
		ByrNo     string
		ByrTgl    string
		KarNama   string
		PelNo     string
		PelNama   string
		PelAlamat string
		RekGol    string
	}

	configs.DB.Raw("SELECT a.byr_no, DATE_FORMAT(a.byr_tgl, '%Y-%m-%d %H:%i:%s') AS byr_tgl, b.kar_nama, c.pel_no, c.pel_nama, c.pel_alamat, c.rek_gol FROM tm_pembayaran a JOIN tm_karyawan b ON b.kar_id = a.kar_id JOIN tm_rekening c ON c.rek_nomor = a.rek_nomor WHERE a.byr_no = ? AND a.byr_sts > 0 ORDER BY a.rek_nomor DESC LIMIT 1", id).Scan(&payment)

	var bills []struct {
		RekThn      string
		RekBln      string
		RekStanlalu string
		RekStankini string
		RekPakai    string
		RekUangair  string
		RekBeban    string
		RekDenda    string
		RekMaterai  string
		RekTotal    string
	}

	configs.DB.Raw("SELECT b.rek_thn, MONTHNAME_ID(b.rek_bln) AS rek_bln, b.rek_stanlalu, b.rek_stankini, (b.rek_stankini - b.rek_stanlalu) AS rek_pakai, b.rek_uangair, (b.rek_adm + b.rek_meter) AS rek_beban, b.rek_denda, b.rek_materai, (b.rek_denda + b.rek_materai + b.rek_total) AS rek_total FROM tm_pembayaran a JOIN tm_rekening b ON b.rek_nomor = a.rek_nomor WHERE a.byr_no = ? AND a.byr_sts > 0", id).Scan(&bills)

	pdf := fpdf.New("P", "mm", "A5", "")
	pdf.AddPage()
	pdf.Ln(10)
	pdf.Image("images/logo-deli-serdang.png", 11, 16, 15, 0, false, "", 0, "")
	pdf.SetFont("Courier", "", 8)
	pdf.SetFontStyle("B")
	pdf.CellFormat(17, 0, "", "", 0, "L", false, 0, "")
	pdf.CellFormat(100, 0, "PERUMDA AIR MINUM TIRTA DELI", "", 1, "L", false, 0, "")
	pdf.SetFontStyle("")
	pdf.CellFormat(17, 8, "", "", 0, "L", false, 0, "")
	pdf.CellFormat(100, 8, "NPWP : 01.457.276.2-125.000", "", 1, "L", false, 0, "")
	pdf.Ln(4)
	pdf.SetFontSize(9)
	pdf.SetFontStyle("B")
	pdf.CellFormat(130, 4, "BUKTI PEMBAYARAN", "", 1, "C", false, 0, "")
	pdf.SetFontSize(8)
	pdf.SetFontStyle("")
	pdf.CellFormat(130, 4, "---------------------------------------------------------------------------", "", 1, "L", false, 0, "")
	pdf.CellFormat(130, 4, "No. Transaksi   : "+payment.ByrNo, "", 1, "L", false, 0, "")
	pdf.CellFormat(130, 4, "Waktu           : "+payment.ByrTgl, "", 1, "L", false, 0, "")
	pdf.CellFormat(130, 4, "Kasir           : "+payment.KarNama, "", 1, "L", false, 0, "")
	pdf.Ln(2)
	pdf.CellFormat(130, 4, "No. Pelanggan   : "+payment.PelNo, "", 1, "L", false, 0, "")
	pdf.CellFormat(130, 4, "Nama            : "+payment.PelNama, "", 1, "L", false, 0, "")
	pdf.CellFormat(130, 4, "Alamat          : "+payment.PelAlamat, "", 1, "L", false, 0, "")
	pdf.CellFormat(130, 4, "Golongan        : "+payment.RekGol, "", 1, "L", false, 0, "")
	pdf.CellFormat(130, 4, "---------------------------------------------------------------------------", "", 1, "L", false, 0, "")

	totalRekening := 0

	for i := 0; i < len(bills); i++ {
		strToIntPakai, _ := strconv.Atoi(bills[i].RekPakai)
		strToIntUangAir, _ := strconv.Atoi(bills[i].RekUangair)
		strToIntBeban, _ := strconv.Atoi(bills[i].RekBeban)
		strToIntDenda, _ := strconv.Atoi(bills[i].RekDenda)
		strToIntMaterai, _ := strconv.Atoi(bills[i].RekMaterai)
		strToIntTotal, _ := strconv.Atoi(bills[i].RekTotal)

		totalRekening += strToIntTotal

		if i > 0 {
			pdf.Ln(2)
		}

		pdf.SetFontStyle("B")
		pdf.CellFormat(60, 4, "Rekening "+bills[i].RekBln+" "+bills[i].RekThn, "", 0, "L", false, 0, "")
		pdf.SetFontStyle("")
		pdf.CellFormat(70, 4, "Stand Meter     : "+bills[i].RekStanlalu+" - "+bills[i].RekStankini, "", 1, "L", false, 0, "")
		pdf.CellFormat(60, 4, "", "", 0, "L", false, 0, "")
		pdf.CellFormat(70, 4, "Pemakaian       : "+CurrencyFormat(strToIntPakai)+" m3", "", 1, "L", false, 0, "")
		pdf.CellFormat(60, 4, "", "", 0, "L", false, 0, "")
		pdf.CellFormat(70, 4, "Uang Air        : "+CurrencyFormat(strToIntUangAir), "", 1, "L", false, 0, "")
		pdf.CellFormat(60, 4, "", "", 0, "L", false, 0, "")
		pdf.CellFormat(70, 4, "Beban Tetap     : "+CurrencyFormat(strToIntBeban), "", 1, "L", false, 0, "")
		pdf.CellFormat(60, 4, "", "", 0, "L", false, 0, "")
		pdf.CellFormat(70, 4, "Denda           : "+CurrencyFormat(strToIntDenda), "", 1, "L", false, 0, "")
		pdf.CellFormat(60, 4, "", "", 0, "L", false, 0, "")
		pdf.CellFormat(70, 4, "Materai         : "+CurrencyFormat(strToIntMaterai), "", 1, "L", false, 0, "")
		pdf.CellFormat(60, 4, "", "", 0, "L", false, 0, "")
		pdf.CellFormat(70, 4, "Total           : "+CurrencyFormat(strToIntTotal), "", 1, "L", false, 0, "")
	}

	pdf.CellFormat(130, 4, "---------------------------------------------------------------------------", "", 1, "L", false, 0, "")
	pdf.CellFormat(60, 4, "Total Rekening", "", 0, "L", false, 0, "")
	pdf.CellFormat(70, 4, "                : "+CurrencyFormat(totalRekening), "", 1, "L", false, 0, "")
	pdf.CellFormat(60, 4, "Biaya Layanan", "", 0, "L", false, 0, "")
	pdf.CellFormat(70, 4, "                : 1,500", "", 1, "L", false, 0, "")
	pdf.SetFontStyle("B")
	pdf.CellFormat(60, 4, "Grand Total", "", 0, "L", false, 0, "")
	pdf.CellFormat(70, 4, "                : "+CurrencyFormat(totalRekening+1500), "", 1, "L", false, 0, "")
	pdf.Ln(5)
	pdf.SetFontStyle("I")
	pdf.CellFormat(130, 3, "Rekening ini dibuat oleh komputer, tanda tangan pejabat PDAM tidak", "", 1, "L", false, 0, "")
	pdf.CellFormat(130, 3, "diperlukan dan sebagai bukti pembayaran yang sah.", "", 1, "L", false, 0, "")
	pdf.OutputFileAndClose("storages/public/INV-" + id + ".pdf")

}
