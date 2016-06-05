package main

import (
	"image/png"
	"net/http"

	"github.com/boombuler/barcode/qr"
	"github.com/nfnt/resize"
)

func handGenQr(w http.ResponseWriter, r *http.Request) {
	coded := r.URL.Query().Get("data")
	w.Header().Add("Content-Type", "image/png")
	bc, err := qr.Encode(coded, qr.H, qr.Unicode)
	if err != nil {
		panic(err.Error())
	}

	ret := resize.Resize(256, 0, bc, resize.NearestNeighbor)
	png.Encode(w, ret)
}
