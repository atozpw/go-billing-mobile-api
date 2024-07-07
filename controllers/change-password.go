package controllers

import (
	"net/http"

	"github.com/atozpw/go-billing-mobile-api/configs"
	"github.com/atozpw/go-billing-mobile-api/helpers"
	"github.com/atozpw/go-billing-mobile-api/models"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func ChangePassword(c *gin.Context) {

	authSession := helpers.AuthSession(c.GetHeader("Authorization"))

	var body struct {
		CurrentPassword      string
		NewPassword          string
		ConfirmationPassword string
	}

	if c.Bind(&body) != nil {
		c.JSON(http.StatusBadRequest, models.ResponseOnlyMessage{
			Code:    400,
			Message: "Gagal memuat Request Body",
		})
		return
	}

	if body.CurrentPassword == "" {
		c.JSON(http.StatusBadRequest, models.ResponseOnlyMessage{
			Code:    400,
			Message: "Password saat ini harus diisi",
		})
		return
	}

	if body.NewPassword == "" {
		c.JSON(http.StatusBadRequest, models.ResponseOnlyMessage{
			Code:    400,
			Message: "Password baru harus diisi",
		})
		return
	}

	if body.ConfirmationPassword == "" {
		c.JSON(http.StatusBadRequest, models.ResponseOnlyMessage{
			Code:    400,
			Message: "Konfirmasi Password baru harus diisi",
		})
		return
	}

	var user models.User

	configs.DB.Raw("SELECT a.kar_id, a.kar_nama, a.kar_pass, b.kp_ket FROM tm_karyawan a JOIN tr_kota_pelayanan b ON a.kp_kode = b.kp_kode WHERE a.kar_id = ? AND a.grup_id = '020'", authSession).Scan(&user)

	err := bcrypt.CompareHashAndPassword([]byte(user.KarPass), []byte(body.CurrentPassword))

	if err != nil {
		c.JSON(http.StatusBadRequest, models.ResponseOnlyMessage{
			Code:    400,
			Message: "Password saat ini salah",
		})
		return
	}

	if body.ConfirmationPassword != body.NewPassword {
		c.JSON(http.StatusBadRequest, models.ResponseOnlyMessage{
			Code:    400,
			Message: "Konfirmasi Password tidak sesuai",
		})
		return
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(body.NewPassword), 10)

	if err != nil {
		c.JSON(http.StatusInternalServerError, models.ResponseOnlyMessage{
			Code:    500,
			Message: "Terjadi kesalahan saat membuat Password",
		})
		return
	}

	result := configs.DB.Exec("UPDATE tm_karyawan SET kar_pass = ? WHERE kar_id = ? AND grup_id = '020'", string(hash), authSession)

	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, models.ResponseOnlyMessage{
			Code:    500,
			Message: "Terjadi kesalahan saat memperbaharui Password",
		})
		return
	}

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
		Message: "Kata sandi berhasil diganti",
		Data:    data,
	})

}
