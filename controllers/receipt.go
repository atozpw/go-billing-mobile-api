package controllers

import (
	"bytes"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"strconv"

	"github.com/atozpw/go-billing-mobile-api/configs"
	"github.com/atozpw/go-billing-mobile-api/helpers"
	"github.com/atozpw/go-billing-mobile-api/models"
	"github.com/gin-gonic/gin"
)

func ReceiptToWhatsapp(c *gin.Context) {

	var body struct {
		Number string
		TrxId  string
	}

	if c.Bind(&body) != nil {
		c.JSON(http.StatusBadRequest, models.ResponseOnlyMessage{
			Code:    400,
			Message: "Gagal memuat Request Body",
		})
		return
	}

	var customer struct {
		PelNo    string
		PelNama  string
		ByrTgl   string
		ByrTotal string
	}

	var bills []struct {
		RekNomor  string
		RekPeriod string
		ByrTotal  string
	}

	configs.DB.Raw("SELECT a.pel_no, a.pel_nama, CONCAT(DAY(b.byr_tgl), ' ', MONTHNAME_ID(MONTH(b.byr_tgl)), ' ', YEAR(b.byr_tgl)) AS byr_tgl, SUM(b.byr_total) + 1500 AS byr_total FROM tm_rekening a JOIN tm_pembayaran b ON b.rek_nomor = a.rek_nomor WHERE b.byr_no = ? AND b.byr_sts > 0 GROUP BY b.byr_no", body.TrxId).Scan(&customer)

	configs.DB.Raw("SELECT a.rek_nomor, CONCAT(MONTHNAME_ID(a.rek_bln), ' ', a.rek_thn) AS rek_period, b.byr_total FROM tm_rekening a JOIN tm_pembayaran b ON b.rek_nomor = a.rek_nomor WHERE b.byr_no = ? AND b.byr_sts > 0", body.TrxId).Scan(&bills)

	billPeriod := ""

	for i := 0; i < len(bills); i++ {
		strToIntTotal, _ := strconv.Atoi(bills[i].ByrTotal)
		billPeriod += bills[i].RekPeriod + ": Rp" + helpers.CurrencyFormat(strToIntTotal) + "\n"
	}

	strToIntTotal, _ := strconv.Atoi(customer.ByrTotal)

	// WhatsappSendText(body.Number, customer.PelNo, customer.PelNama, billPeriod, customer.ByrTgl, helpers.CurrencyFormat(strToIntTotal))

	// time.After(time.Second * 4)

	// for i := 0; i < len(bills); i++ {
	// 	if i > 0 {
	// 		time.After(time.Second * 4)
	// 	}
	// 	WhatsappSendFile(body.Number, os.Getenv("STORAGE_URL_PATH")+"/INV-"+bills[i].RekNomor+".pdf", "INV-"+bills[i].RekNomor+".pdf")
	// }

	// WhatsappSendFile(body.Number, os.Getenv("STORAGE_URL_PATH")+"/INV-"+body.TrxId+".pdf", "INV-"+body.TrxId+".pdf")

	WahaBroadcast(body.Number, customer.PelNo, customer.PelNama, billPeriod, customer.ByrTgl, helpers.CurrencyFormat(strToIntTotal), os.Getenv("STORAGE_URL_PATH")+"/INV-"+body.TrxId+".pdf", "INV-"+body.TrxId+".pdf")

	c.JSON(http.StatusOK, models.ResponseOnlyMessage{
		Code:    200,
		Message: "Bukti pembayaran berhasil terkirim",
	})

}

func WhatsappSendText(number string, customerNo string, customerName string, billPeriod string, trxDate string, trxAmount string) {

	url := os.Getenv("WHATSAPP_ENDPOINT") + "/send-text"
	method := "POST"

	message := "Pelanggan Yth.\n"
	message += "Terima kasih telah melakukan pembayaran rekening air.\n"
	message += "Pada tanggal " + trxDate + ".\n"
	message += "Untuk nomor pelanggan " + customerNo + ", atas nama " + customerName + ".\n\n"
	message += "Dengan rincian:\n"
	message += billPeriod
	message += "Biaya Layanan: Rp1,500\n"
	message += "Total Pembayaran: Rp" + trxAmount

	payload := &bytes.Buffer{}
	writer := multipart.NewWriter(payload)
	_ = writer.WriteField("number", number)
	_ = writer.WriteField("message", message)
	err := writer.Close()

	if err != nil {
		fmt.Println(err)
		return
	}

	client := &http.Client{}
	req, err := http.NewRequest(method, url, payload)

	if err != nil {
		fmt.Println(err)
		return
	}

	req.Header.Set("Content-Type", writer.FormDataContentType())
	res, err := client.Do(req)

	if err != nil {
		fmt.Println(err)
		return
	}

	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(string(body))

}

func WhatsappSendFile(number string, file string, filename string) {

	url := os.Getenv("WHATSAPP_ENDPOINT") + "/send-file-from-url"
	method := "POST"

	payload := &bytes.Buffer{}
	writer := multipart.NewWriter(payload)
	_ = writer.WriteField("number", number)
	_ = writer.WriteField("file", file)
	_ = writer.WriteField("filename", filename)
	err := writer.Close()

	if err != nil {
		fmt.Println(err)
		return
	}

	client := &http.Client{}
	req, err := http.NewRequest(method, url, payload)

	if err != nil {
		fmt.Println(err)
		return
	}

	req.Header.Set("Content-Type", writer.FormDataContentType())
	res, err := client.Do(req)

	if err != nil {
		fmt.Println(err)
		return
	}

	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(string(body))

}

func WahaBroadcast(number string, customerNo string, customerName string, billPeriod string, trxDate string, trxAmount string, file string, filename string) {

	message := "Pelanggan Yth.\n"
	message += "Terima kasih telah melakukan pembayaran rekening air.\n"
	message += "Pada tanggal " + trxDate + ".\n"
	message += "Untuk nomor pelanggan " + customerNo + ", atas nama " + customerName + ".\n\n"
	message += "Dengan rincian:\n"
	message += billPeriod
	message += "Biaya Layanan: Rp1,500\n"
	message += "Total Pembayaran: Rp" + trxAmount

	configs.DB.Exec("INSERT INTO wa_broadcasts (customer_no, number, message, file, filename, created_at, updated_at) VALUES (?, ?, ?, ?, ?, NOW(), NOW())", customerNo, number, message, file, filename)

}
