package controllers

import (
	"net/http"

	"github.com/atozpw/go-billing-mobile-api/configs"
	"github.com/atozpw/go-billing-mobile-api/models"
	"github.com/gin-gonic/gin"
)

func CustomerFind(c *gin.Context) {

	var customer struct {
		PelNo     string `json:"id"`
		PelNama   string `json:"name"`
		PelAlamat string `json:"address"`
		DkdKd     string `json:"zone"`
		GolKet    string `json:"group"`
		KpKet     string `json:"unit"`
		KpsKet    string `json:"status"`
	}

	result := configs.DB.Raw("SELECT a.pel_no, a.pel_nama, a.pel_alamat, a.dkd_kd, b.gol_ket, c.kp_ket, d.kps_ket FROM tm_pelanggan a JOIN tr_gol b ON b.gol_kode = a.gol_kode JOIN tr_kota_pelayanan c ON c.kp_kode = a.kp_kode JOIN tr_kondisi_ps d ON d.kps_kode = a.kps_kode WHERE a.pel_no = ?", c.Param("id")).Scan(&customer)

	if result.RowsAffected == 0 {
		c.JSON(http.StatusNotFound, models.ResponseWithData{
			Code:    404,
			Message: "Data tidak ditemukan",
			Data:    []int{},
		})
	} else {
		c.JSON(http.StatusOK, models.ResponseWithData{
			Code:    200,
			Message: "Data Pelanggan",
			Data:    customer,
		})
	}
}

func CustomerBills(c *gin.Context) {

	var bills []struct {
		RekNomor    string `json:"id"`
		RekThn      string `json:"year"`
		RekBln      string `json:"month"`
		RekStanlalu string `json:"lastWm"`
		RekStankini string `json:"currentWm"`
		RekPakai    string `json:"waterUsage"`
		RekUangair  string `json:"amount"`
		RekAdm      string `json:"adminFee"`
		RekMeter    string `json:"meterCost"`
		RekDenda    string `json:"additionalAmount"`
		RekTotal    string `json:"total"`
	}

	result := configs.DB.Raw("SELECT rek_nomor, rek_thn, MONTHNAME(CONCAT(rek_thn, '-', rek_bln, '-1')) as rek_bln, rek_stanlalu, rek_stankini, (rek_stankini - rek_stanlalu) AS rek_pakai, rek_uangair, rek_adm, rek_meter, getDenda(rek_total, rek_bln, rek_thn, rek_gol) AS rek_denda, (getDenda(rek_total, rek_bln, rek_thn, rek_gol) + rek_total) AS rek_total FROM tm_rekening WHERE pel_no = ? AND rek_sts = 1 AND rek_byr_sts = 0", c.Param("id")).Scan(&bills)

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
