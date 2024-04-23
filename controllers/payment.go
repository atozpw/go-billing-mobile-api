package controllers

import (
	"net/http"

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

}
