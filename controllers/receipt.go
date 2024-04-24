package controllers

import (
	"bytes"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"os"

	"github.com/atozpw/go-billing-mobile-api/models"
	"github.com/gin-gonic/gin"
)

func ReceiptToWhatsapp(c *gin.Context) {

	url := os.Getenv("WHATSAPP_ENDPOINT") + "/send-file-from-url"
	method := "POST"

	payload := &bytes.Buffer{}
	writer := multipart.NewWriter(payload)
	_ = writer.WriteField("number", "089608036785")
	_ = writer.WriteField("file", "https://qris.tirtaraharja.co.id/storage/INV-202303610002961.pdf")
	_ = writer.WriteField("filename", "test")
	err := writer.Close()

	if err != nil {
		c.JSON(http.StatusBadRequest, models.ResponseOnlyMessage{
			Code:    400,
			Message: "Gagal memuat Request Body",
		})
		return
	}

	client := &http.Client{}
	req, err := http.NewRequest(method, url, payload)

	if err != nil {
		c.JSON(http.StatusInternalServerError, models.ResponseOnlyMessage{
			Code:    500,
			Message: "Terjadi kesalahan saat membuat HTTP Client",
		})
		return
	}

	req.Header.Set("Content-Type", writer.FormDataContentType())
	res, err := client.Do(req)

	if err != nil {
		c.JSON(http.StatusInternalServerError, models.ResponseOnlyMessage{
			Code:    500,
			Message: "Terjadi kesalahan saat membuat request ke Server Whatsapp",
		})
		return
	}

	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)

	if err != nil {
		c.JSON(http.StatusInternalServerError, models.ResponseOnlyMessage{
			Code:    500,
			Message: "Terjadi kesalahan saat mendapatkan response dari Server Whatsapp",
		})
		return
	}

	fmt.Println(string(body))

	c.JSON(http.StatusOK, models.ResponseOnlyMessage{
		Code:    200,
		Message: "Berhasil mengirim Bukti Bayar",
	})

}
