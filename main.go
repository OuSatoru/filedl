package main

import (
	"fmt"
	"github.com/boombuler/barcode"
	"github.com/boombuler/barcode/qr"
	"image/png"
	"net/http"
	"os"
	"strings"
)

func main() {

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
