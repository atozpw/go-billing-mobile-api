package controllers

import (
	"net/http"

	"github.com/atozpw/go-billing-mobile-api/configs"
	"github.com/atozpw/go-billing-mobile-api/helpers"
	"github.com/atozpw/go-billing-mobile-api/models"
	"github.com/gin-gonic/gin"
)

func Profile(c *gin.Context) {

	authSession := helpers.AuthSession(c.GetHeader("Authorization"))

	var user models.User

	configs.DB.Raw("SELECT a.kar_id, a.kar_nama, a.kar_pass, b.kp_ket FROM tm_karyawan a JOIN tr_kota_pelayanan b ON a.kp_kode = b.kp_kode WHERE a.kar_id = ? AND a.grup_id = '020'", authSession).Scan(&user)

	var data struct {
		ID    string `json:"id"`
		Name  string `json:"name"`
		Kopel string `json:"kopel"`
	}

	data.ID = user.KarId
	data.Name = user.KarNama
	data.Kopel = user.KpKet

	c.JSON(http.StatusOK, models.ResponseWithData{
		Code:    200,
		Message: "Data Pengguna",
		Data:    data,
	})

}
