package controllers

import (
	"net/http"
	"time"

	"github.com/atozpw/go-billing-mobile-api/configs"
	"github.com/atozpw/go-billing-mobile-api/helpers"
	"github.com/atozpw/go-billing-mobile-api/models"
	"github.com/gin-gonic/gin"
)

func PaymentIndex(c *gin.Context) {

	authId := helpers.AuthSession(c.GetHeader("Authorization"))
	date := c.Query("date")

	var payments []struct {
		ByrNo      string `json:"id"`
		PelNo      string `json:"customerNo"`
		PelNama    string `json:"customerName"`
		ByrTgl     string `json:"trxDate"`
		ByrJam     string `json:"trxTime"`
		RekThn     string `json:"billYear"`
		RekBln     string `json:"billMonth"`
		RekPakai   string `json:"waterUsage"`
		RekUangair string `json:"amount"`
		RekAdm     string `json:"adminFee"`
		RekMeter   string `json:"meterCost"`
		RekDenda   string `json:"additionalAmount"`
		ByrTotal   string `json:"trxTotal"`
	}

	result := configs.DB.Raw("SELECT b.byr_no, a.pel_no, a.pel_nama, DATE_FORMAT(b.byr_tgl, '%Y-%m-%d') AS byr_tgl, DATE_FORMAT(b.byr_tgl, '%H:%i:%s') AS byr_jam, a.rek_thn, MONTHNAME_ID(a.rek_bln) as rek_bln, (a.rek_stankini - a.rek_stanlalu) AS rek_pakai, a.rek_uangair, a.rek_adm, a.rek_meter, a.rek_denda, b.byr_total FROM tm_rekening a JOIN tm_pembayaran b ON b.rek_nomor = a.rek_nomor WHERE DATE_FORMAT(b.byr_tgl, '%Y-%m-%d') = ? AND b.kar_id = ? AND b.byr_sts > 0 ORDER BY b.byr_tgl", date, authId).Scan(&payments)

	if result.RowsAffected == 0 {
		c.JSON(http.StatusNotFound, models.ResponseWithData{
			Code:    404,
			Message: "Data tidak ditemukan",
			Data:    []int{},
		})
	} else {
		c.JSON(http.StatusOK, models.ResponseWithData{
			Code:    200,
			Message: "Data Pembayaran",
			Data:    payments,
		})
	}

}

func PaymentFind(c *gin.Context) {

	authId := helpers.AuthSession(c.GetHeader("Authorization"))

	var bills []struct {
		RekNomor   string `json:"id"`
		RekThn     string `json:"year"`
		RekBln     string `json:"month"`
		RekUangair string `json:"amount"`
		RekAdm     string `json:"adminFee"`
		RekMeter   string `json:"meterCost"`
		RekDenda   string `json:"additionalAmount"`
		RekLayanan string `json:"serviceFee"`
		RekTotal   string `json:"total"`
	}

	configs.DB.Raw("SELECT a.rek_nomor, b.rek_thn, MONTHNAME(CONCAT(b.rek_thn, '-', b.rek_bln, '-1')) AS rek_bln, b.rek_uangair, b.rek_adm, b.rek_meter, b.rek_denda, 0 AS rek_layanan, b.rek_total FROM tm_pembayaran a JOIN tm_rekening b ON b.rek_nomor = a.rek_nomor WHERE a.byr_no = ? AND a.kar_id = ? AND a.byr_sts > 0", c.Param("id"), authId).Scan(&bills)

	var payment struct {
		ByrNo     string      `json:"id"`
		ByrTgl    string      `json:"trxDate"`
		KarNama   string      `json:"cashier"`
		PelNo     string      `json:"customerNo"`
		PelNama   string      `json:"customerName"`
		PelAlamat string      `json:"customerAddress"`
		RekGol    string      `json:"customerGroup"`
		Detail    interface{} `json:"billDetails"`
	}

	payment.Detail = bills

	paymentResult := configs.DB.Raw("SELECT a.byr_no, DATE_FORMAT(a.byr_tgl, '%Y-%m-%d %H:%i:%s') AS byr_tgl, b.kar_nama, c.pel_no, c.pel_nama, c.pel_alamat, c.rek_gol FROM tm_pembayaran a JOIN tm_karyawan b ON b.kar_id = a.kar_id JOIN tm_rekening c ON c.rek_nomor = a.rek_nomor WHERE a.byr_no = ? AND a.kar_id = ? AND a.byr_sts > 0 ORDER BY a.rek_nomor DESC LIMIT 1", c.Param("id"), authId).Scan(&payment)

	if paymentResult.RowsAffected == 0 {
		c.JSON(http.StatusNotFound, models.ResponseWithData{
			Code:    404,
			Message: "Data tidak ditemukan",
			Data:    []int{},
		})
	} else {
		c.JSON(http.StatusOK, models.ResponseWithData{
			Code:    200,
			Message: "Data Pembayaran",
			Data:    payment,
		})
	}

}

func PaymentStore(c *gin.Context) {

	type Bill struct {
		Id     string
		Amount string
	}

	var body struct {
		Id     string
		Amount string
		UserId string
		Bills  []Bill
	}

	if c.Bind(&body) != nil {
		c.JSON(http.StatusBadRequest, models.ResponseOnlyMessage{
			Code:    400,
			Message: "Gagal memuat Request Body",
		})
		return
	}

	authId := helpers.AuthSession(c.GetHeader("Authorization"))

	var sysParam struct {
		SysValue1 int
	}

	configs.DB.Raw("SELECT sys_value1 FROM system_parameter WHERE sys_param = 'RESI' AND sys_value = ?", authId).Scan(&sysParam)

	trxDate := time.Now().Format("2006-01-02 15:04:05")
	clientIp := c.ClientIP()
	trxSerial := sysParam.SysValue1

	tx := configs.DB.Begin()

	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	if err := tx.Error; err != nil {
		c.JSON(http.StatusInternalServerError, models.ResponseOnlyMessage{
			Code:    500,
			Message: "Terjadi kesalahan saat menyimpan Pembayaran",
		})
		return
	}

	for i := 0; i < len(body.Bills); i++ {

		payment := tx.Exec("INSERT INTO tm_pembayaran (byr_no, byr_tgl, byr_serial, rek_nomor, kar_id, lok_ip, byr_loket, byr_total, byr_cetak, byr_upd_sts, byr_sts) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)", body.Id, trxDate, trxSerial, body.Bills[i].Id, authId, clientIp, "N", body.Bills[i].Amount, 0, trxDate, 1)

		if payment.Error != nil {
			tx.Rollback()
			c.JSON(http.StatusInternalServerError, models.ResponseOnlyMessage{
				Code:    500,
				Message: "Terjadi kesalahan saat menyimpan Pembayaran",
			})
			return
		}

		bill := tx.Exec("UPDATE tm_rekening SET rek_denda = getDenda(rek_total, rek_bln, rek_thn), rek_byr_sts = 1 WHERE rek_nomor = ? AND rek_sts = 1 AND rek_byr_sts = 0", body.Bills[i].Id)

		if bill.Error != nil {
			tx.Rollback()
			c.JSON(http.StatusInternalServerError, models.ResponseOnlyMessage{
				Code:    500,
				Message: "Terjadi kesalahan saat menyimpan Pembayaran",
			})
			return
		}

		trxSerial++

	}

	param := tx.Exec("UPDATE system_parameter SET sys_value1 = ? WHERE sys_param = 'RESI' AND sys_value = ?", trxSerial, authId)

	if param.Error != nil {
		tx.Rollback()
		c.JSON(http.StatusInternalServerError, models.ResponseOnlyMessage{
			Code:    500,
			Message: "Terjadi kesalahan saat menyimpan Pembayaran",
		})
		return
	}

	tx.Commit()

	for i := 0; i < len(body.Bills); i++ {
		helpers.GenerateReceipt(body.Bills[i].Id)
	}

	c.JSON(http.StatusOK, models.ResponseWithData{
		Code:    200,
		Message: "Pembayaran sukses",
		Data:    body.Bills,
	})

}
