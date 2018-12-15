package common

import (
	"github.com/afocus/captcha"
	"image/color"
)

var Captcha *captcha.Captcha

func init() {

	Captcha = captcha.New()

	if err := Captcha.SetFont("./comic.ttf"); err != nil {
		panic(err.Error())
	}

	Captcha.SetSize(128, 64)
	Captcha.SetDisturbance(captcha.NORMAL)
	Captcha.SetFrontColor(color.RGBA{0, 0, 255, 255})
	Captcha.SetBkgColor(color.RGBA{255, 255, 255, 255})

	//http.HandleFunc("/r", func(w http.ResponseWriter, r *http.Request) {
	//	img, str := Captcha.Create(4, captcha.ALL)
	//	png.Encode(w, img)
	//	println(str)
	//})
	//
	//http.HandleFunc("/c", func(w http.ResponseWriter, r *http.Request) {
	//	str := r.URL.RawQuery
	//	img := Captcha.CreateCustom(str)
	//	png.Encode(w, img)
	//})
	//
	//http.ListenAndServe(":8085", nil)

}
