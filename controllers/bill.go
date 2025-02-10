package controllers

import (
	"net/http"

	"github.com/atozpw/go-billing-mobile-api/configs"
	"github.com/atozpw/go-billing-mobile-api/models"
	"github.com/gin-gonic/gin"
)

func BillIndex(c *gin.Context) {

	customerNo := c.Query("customerNo")

	var bills []struct {
		RekNomor  string `json:"id"`
		RekPeriod string `json:"period"`
		RekTotal  string `json:"total"`
	}

	result := configs.DB.Raw("SELECT rek_nomor, CONCAT(MONTHNAME_ID(rek_bln), ' ', rek_thn) AS rek_period, (getDenda(rek_total, rek_bln, rek_thn) + rek_total) AS rek_total FROM tm_rekening WHERE pel_no = ? AND rek_sts = 1 AND rek_byr_sts = 0", customerNo).Scan(&bills)

	if result.RowsAffected == 0 {
		c.JSON(http.StatusNotFound, models.ResponseWithData{
			Code:    404,
			Message: "Data tidak ditemukan",
			Data:    []int{},
		})
	} else {
		c.JSON(http.StatusOK, models.ResponseWithData{
			Code:    200,
			Message: "Data Tagihan",
			Data:    bills,
		})
	}

}

func BillFind(c *gin.Context) {

	id := c.Param("id")

	var bill struct {
		RekNomor    string `json:"id"`
		RekPeriod   string `json:"period"`
		RekStanlalu string `json:"lastWm"`
		RekStankini string `json:"currentWm"`
		RekPakai    string `json:"waterUsage"`
		RekUangair  string `json:"amount"`
		RekAdm      string `json:"adminFee"`
		RekMeter    string `json:"meterCost"`
		RekDenda    string `json:"additionalAmount"`
		RekTotal    string `json:"total"`
	}

	result := configs.DB.Raw("SELECT rek_nomor, CONCAT(MONTHNAME_ID(rek_bln), ' ', rek_thn) AS rek_period, rek_stanlalu, rek_stankini, (rek_stankini - rek_stanlalu) AS rek_pakai, rek_uangair, rek_adm, rek_meter, getDenda(rek_total, rek_bln, rek_thn) AS rek_denda, (getDenda(rek_total, rek_bln, rek_thn) + rek_total) AS rek_total FROM tm_rekening WHERE rek_nomor = ? AND rek_sts = 1 AND rek_byr_sts = 0", id).Scan(&bill)

	if result.RowsAffected == 0 {
		c.JSON(http.StatusNotFound, models.ResponseWithData{
			Code:    404,
			Message: "Data tidak ditemukan",
			Data:    []int{},
		})
	} else {
		c.JSON(http.StatusOK, models.ResponseWithData{
			Code:    200,
			Message: "Detail Tagihan",
			Data:    bill,
		})
	}

}
