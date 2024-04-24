package controllers

import (
	"net/http"
	"time"

	"github.com/atozpw/go-billing-mobile-api/configs"
	"github.com/atozpw/go-billing-mobile-api/models"
	"github.com/gin-gonic/gin"
)

func PaymentIndex(c *gin.Context) {

	var payments []struct {
		ByrNo      string `json:"id"`
		ByrTgl     string `json:"trxDate"`
		RekThn     string `json:"billYear"`
		RekBln     string `json:"billMonth"`
		RekPakai   string `json:"waterUsage"`
		RekUangair string `json:"amount"`
		RekAdm     string `json:"adminFee"`
		RekMeter   string `json:"meterCost"`
		RekDenda   string `json:"additionalAmount"`
		ByrTotal   string `json:"trxTotal"`
	}

	result := configs.DB.Raw("SELECT b.byr_no, DATE_FORMAT(b.byr_tgl, '%Y-%m-%d %H:%i:%s') AS byr_tgl, a.rek_thn, MONTHNAME(CONCAT(a.rek_thn, '-', a.rek_bln, '-1')) as rek_bln, (a.rek_stankini - a.rek_stanlalu) AS rek_pakai, a.rek_uangair, a.rek_adm, a.rek_meter, a.rek_denda, b.byr_total FROM tm_rekening a JOIN tm_pembayaran b ON b.rek_nomor = a.rek_nomor WHERE DATE_FORMAT(b.byr_tgl, '%Y-%m-%d') = ? AND b.kar_id = ? AND b.byr_sts > 0 ORDER BY b.byr_tgl", "2024-03-08", "VSI").Scan(&payments)

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

func PaymentStore(c *gin.Context) {

	type bill struct {
		Id     string
		Amount string
	}

	var body struct {
		Id     string
		Amount string
		UserId string
		Bills  []bill
	}

	if c.Bind(&body) != nil {
		c.JSON(http.StatusBadRequest, models.ResponseOnlyMessage{
			Code:    400,
			Message: "Gagal memuat Request Body",
		})
		return
	}

	trxDate := time.Now().Format("2006-01-02 15:04:05")

	for i := 0; i < len(body.Bills); i++ {

		result := configs.DB.Exec("INSERT INTO tm_pembayaran (byr_no, byr_tgl, byr_serial, rek_nomor, kar_id, lok_ip, byr_loket, byr_total, byr_cetak, byr_upd_sts, byr_sts) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)", body.Id, trxDate, 1, body.Bills[i].Id, "test", c.ClientIP(), "N", body.Amount, 0, trxDate, 1)

		if result.Error != nil {
			c.JSON(http.StatusInternalServerError, models.ResponseOnlyMessage{
				Code:    500,
				Message: "Terjadi kesalahan saat menyimpan Pembayaran",
			})
			return
		}

	}

	c.JSON(http.StatusOK, models.ResponseWithData{
		Code:    200,
		Message: "Pembayaran sukses",
		Data:    body.Bills,
	})

}
