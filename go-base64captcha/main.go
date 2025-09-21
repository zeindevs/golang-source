package main

import (
	"encoding/base64"
	"encoding/json"
	"log"
	"net/http"
	"strings"

	"github.com/mojocn/base64Captcha"
)

var store = base64Captcha.DefaultMemStore

type CaptchaResponse struct {
	CaptchaID string `json:"captcha_id"`
	Image     string `json:"image"`
}

func generateCaptcha(w http.ResponseWriter, _ *http.Request) {
	driver := base64Captcha.NewDriverDigit(80, 240, 5, 0.7, 80)
	c := base64Captcha.NewCaptcha(driver, store)
	id, b64s, _, err := c.Generate()
	if err != nil {
		http.Error(w, "Captcha generation failed!", http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(CaptchaResponse{
		CaptchaID: id,
		Image:     b64s,
	})
}

func imageCaptcha(w http.ResponseWriter, r *http.Request) {
	driver := base64Captcha.NewDriverDigit(80, 240, 5, 0.7, 80)
	c := base64Captcha.NewCaptcha(driver, store)
	id, b64s, _, err := c.Generate()
	if err != nil {
		http.Error(w, "Captcha generation failed!", http.StatusInternalServerError)
		return
	}
	if strings.HasPrefix(b64s, "data:image") {
		parts := strings.SplitN(b64s, ",", 2)
		b64s = parts[1]
	}
	img, err := base64.StdEncoding.DecodeString(b64s)
	if err != nil {
		http.Error(w, "Decode image failed!", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "image/png")
	w.Header().Set("Captcha-Id", id)
	w.Write(img)
}

func verifyCaptcha(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	id := r.FormValue("captcha_id")
	answer := r.FormValue("captcha_solution")
	if store.Verify(id, answer, true) {
		w.Write([]byte("OK"))
	} else {
		w.Write([]byte("FAILED"))
	}
}

func main() {
	http.HandleFunc("/captcha/json", generateCaptcha)
	http.HandleFunc("/captcha/image", imageCaptcha)
	http.HandleFunc("/captcha/verify", verifyCaptcha)
	log.Println("server listening on :5001")
	http.ListenAndServe(":5001", nil)
}
