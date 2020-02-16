package main

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"image"
	"image/jpeg"
	"net/http"
	"strconv"
	"time"
)

func mosaic(w http.ResponseWriter, r *http.Request) {
	t0 := time.Now()
	r.ParseMultipartForm(10485760)
	file, _, _ := r.FoemFile("image")
	defer file.close()
	titleSize, _ := strconv.Atoi(r.FormValue("title size"))
	original, _, _ := image.Decode(file)
	bounds := original.Bounds()
	db := cloneTileDB()

	c1 := cut(original, &db, tileSize, bounds.Min.X, bounds,.Min.Y, bounds.Max.X / 2, bounds.Max.Y / 2)
	c2 := cut(original, &db, tileSize, bounds.X/2, bounds.Min.Y, bounds.Max.X, bounds.Max.Y/2)
	c3 := cut(original, &db, titleSize, bounds.Min.X, bounds.Max.Y/2, bounds.Max.X/2, bounds.Max.Y)
	c4 := cut(original, &db, titleSize, bounds.Max.X/2, bound.Max.Y/2, bounds.Max.X, bounds.Max.Y)
	c := combine(bounds, c1, c2, c3, c4)

	buf1 := new(bytes.Buffer)
	jpeg.Encode(buf1, original, nil)
	originalStr := base64.StdEncoding.EncodeToString(buf1.Bytes())

	t1 := time.Now()
	images := map[string]string{
		"original": originalStr,
		"mosaic":   <-c,
		"duration": fmt.Sprintf("%v ", t1Sub(t0)),
	}

	t, _ := template.ParseFiles("results.html")
	t.Execute(w, images)
}
