package controllers

import (
	"net/http"
	"os"
	"time"

	"github.com/atozpw/go-billing-mobile-api/configs"
	"github.com/atozpw/go-billing-mobile-api/helpers"
	"github.com/atozpw/go-billing-mobile-api/models"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

func Login(c *gin.Context) {

	var body struct {
		Username string
		Password string
		DeviceId string
	}

	if c.Bind(&body) != nil {
		c.JSON(http.StatusBadRequest, models.ResponseOnlyMessage{
			Code:    400,
			Message: "Gagal memuat Request Body",
		})
		return
	}

	if body.Username == "" {
		c.JSON(http.StatusBadRequest, models.ResponseOnlyMessage{
			Code:    400,
			Message: "Username harus diisi",
		})
		return
	}

	if body.Password == "" {
		c.JSON(http.StatusBadRequest, models.ResponseOnlyMessage{
			Code:    400,
			Message: "Password harus diisi",
		})
		return
	}

	if body.DeviceId == "" {
		c.JSON(http.StatusBadRequest, models.ResponseOnlyMessage{
			Code:    400,
			Message: "Device ID harus diisi",
		})
		return
	}

	var user models.User

	configs.DB.Raw("SELECT a.kar_id, a.kar_nama, a.kar_pass, b.kp_ket, a.device_id FROM tm_karyawan a JOIN tr_kota_pelayanan b ON a.kp_kode = b.kp_kode WHERE a.kar_id = ? AND a.grup_id = '020'", body.Username).Scan(&user)

	if user.KarId == "" {
		c.JSON(http.StatusBadRequest, models.ResponseOnlyMessage{
			Code:    400,
			Message: "Username tidak terdaftar",
		})
		return
	}

	err := bcrypt.CompareHashAndPassword([]byte(user.KarPass), []byte(body.Password))

	if err != nil {
		c.JSON(http.StatusBadRequest, models.ResponseOnlyMessage{
			Code:    400,
			Message: "Password salah",
		})
		return
	}

	duration := time.Duration(24 - time.Now().Hour())

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": user.KarId,
		"exp": time.Now().Add(time.Hour * duration).Unix(),
	})

	tokenString, err := token.SignedString([]byte(os.Getenv("SECRET_KEY")))

	if err != nil {
		c.JSON(http.StatusInternalServerError, models.ResponseOnlyMessage{
			Code:    500,
			Message: "Terjadi kesalahan saat membuat Token",
		})
		return
	}

	var data struct {
		ID    string `json:"id"`
		Name  string `json:"name"`
		Kopel string `json:"kopel"`
		Token string `json:"token"`
	}

	data.ID = user.KarId
	data.Name = user.KarNama
	data.Kopel = user.KpKet
	data.Token = tokenString

	if user.DeviceId != body.DeviceId {

		randomCode := helpers.RandomCode(6)

		configs.DB.Exec("UPDATE tm_karyawan SET device_id = ? WHERE kar_id = ? AND grup_id = '020'", randomCode, body.Username)

		c.JSON(http.StatusUnauthorized, models.ResponseWithData{
			Code:    401,
			Message: "Device tidak terdaftar",
			Data:    data,
		})
		return

	}

	c.JSON(http.StatusOK, models.ResponseWithData{
		Code:    200,
		Message: "Login sukses",
		Data:    data,
	})

}

func Register(c *gin.Context) {

	authSession := helpers.AuthSession(c.GetHeader("Authorization"))

	var body struct {
		DeviceId         string
		VerificationCode string
	}

	if c.Bind(&body) != nil {
		c.JSON(http.StatusBadRequest, models.ResponseOnlyMessage{
			Code:    400,
			Message: "Gagal memuat Request Body",
		})
		return
	}

	if body.DeviceId == "" {
		c.JSON(http.StatusBadRequest, models.ResponseOnlyMessage{
			Code:    400,
			Message: "Device ID harus diisi",
		})
		return
	}

	if body.VerificationCode == "" {
		c.JSON(http.StatusBadRequest, models.ResponseOnlyMessage{
			Code:    400,
			Message: "Kode Verifikasi harus diisi",
		})
		return
	}

	var user struct {
		KarId string
	}

	configs.DB.Raw("SELECT kar_id FROM tm_karyawan WHERE kar_id = ? AND grup_id = '020' AND device_id = ?", authSession, body.VerificationCode).Scan(&user)

	if user.KarId == "" {
		c.JSON(http.StatusBadRequest, models.ResponseOnlyMessage{
			Code:    400,
			Message: "Kode Verifikasi salah",
		})
		return
	}

	result := configs.DB.Exec("UPDATE tm_karyawan SET device_id = ? WHERE kar_id = ? AND grup_id = '020'", body.DeviceId, authSession)

	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, models.ResponseOnlyMessage{
			Code:    500,
			Message: "Terjadi kesalahan saat memperbaharui Device ID",
		})
		return
	}

	var data struct {
		DeviceId string `json:"deviceId"`
	}

	data.DeviceId = body.DeviceId

	c.JSON(http.StatusOK, models.ResponseWithData{
		Code:    200,
		Message: "Perangkat berhasil diregistrasi",
		Data:    data,
	})

}
