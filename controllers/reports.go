package controllers

import (
	"net/http"

	"github.com/atozpw/go-billing-mobile-api/configs"
	"github.com/atozpw/go-billing-mobile-api/helpers"
	"github.com/atozpw/go-billing-mobile-api/models"
	"github.com/gin-gonic/gin"
)

func ReportIndex(c *gin.Context) {

	authId := helpers.AuthSession(c.GetHeader("Authorization"))
	date := c.Query("date")

	var reports []struct {
		ByrNo    string `json:"id"`
		PelNo    string `json:"customerNo"`
		PelNama  string `json:"customerName"`
		ByrJam   string `json:"trxTime"`
		ByrTotal string `json:"trxTotal"`
	}

	result := configs.DB.Raw("SELECT b.byr_no, a.pel_no, a.pel_nama, DATE_FORMAT(b.byr_tgl, '%H:%i:%s') AS byr_jam, c.byr_total FROM tm_rekening a JOIN tm_pembayaran b ON b.rek_nomor = a.rek_nomor JOIN tm_bayar_mobile c ON c.byr_no = b.byr_no WHERE DATE_FORMAT(b.byr_tgl, '%Y-%m-%d') = ? AND b.kar_id = ? AND b.byr_sts > 0 GROUP BY a.pel_no ORDER BY b.byr_tgl DESC", date, authId).Scan(&reports)

	if result.RowsAffected == 0 {
		c.JSON(http.StatusNotFound, models.ResponseWithData{
			Code:    404,
			Message: "Data tidak ditemukan",
			Data:    []int{},
		})
	} else {
		c.JSON(http.StatusOK, models.ResponseWithData{
			Code:    200,
			Message: "Data Penerimaan",
			Data:    reports,
		})
	}

}
