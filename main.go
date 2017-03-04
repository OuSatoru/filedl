package main

import (
	"fmt"
	"image/png"
	"net/http"
	"os"
	"strings"
	"html/template"
	"github.com/boombuler/barcode/qr"
	"github.com/boombuler/barcode"
)

type anApk struct {
	title string
	icon string
	qrCode string
}

func main() {
	http.HandleFunc("/apk", apk)
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("www/static"))))
	http.ListenAndServe(":2333", nil)
}

func qrGen(s string) {
	f, _ := os.Create(fmt.Sprintf("www/static/qr/%s.png", strings.TrimPrefix(s, "http://")))
	defer f.Close()

	qrCode, err := qr.Encode(s, qr.L, qr.Auto)
	if err != nil {
		panic(err)
	}
	qrCode, err = barcode.Scale(qrCode, 100, 100)
	if err != nil {
		panic(err)
	}
	png.Encode(f, qrCode)
}

func apk(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.URL.Path)
	t, _ := template.ParseFiles("www/template/apks.html")
	t.Execute(w, nil)
}
