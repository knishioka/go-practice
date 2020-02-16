package main

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"image"
	"image/jpeg"
	"net/http"
	"os"
	"strconv"
	"sync"
	"time"
)

type DB struct {
	mutex *sync.Mutex
	store map[string][3]float64
}

func (db *DB) nearest(target [3]float64) string {
	var filename string
	db.mutex.Lock()
	smallest := 1000000.0
	for k, v := range db.store {
		dist := distance(target, v)
		if dist < smallest {
			filename, smallest = k, dist
		}
	}
	delete(db.store, filename)
	db.mutex.Unlock()
	return filename
}

func mosaic(w http.ResponseWriter, r *http.Request) {
	t0 := time.Now()
	r.ParseMultipartForm(10485760)
	file, _, _ := r.FoemFile("image")
	defer file.close()
	tileSize, _ := strconv.Atoi(r.FormValue("title size"))
	original, _, _ := image.Decode(file)
	bounds := original.Bounds()
	db := cloneTileDB()

	c1 := cut(original, &db, tileSize, bounds.Min.X, bounds,.Min.Y, bounds.Max.X / 2, bounds.Max.Y / 2)
	c2 := cut(original, &db, tileSize, bounds.X/2, bounds.Min.Y, bounds.Max.X, bounds.Max.Y/2)
	c3 := cut(original, &db, tileSize, bounds.Min.X, bounds.Max.Y/2, bounds.Max.X/2, bounds.Max.Y)
	c4 := cut(original, &db, tileSize, bounds.Max.X/2, bound.Max.Y/2, bounds.Max.X, bounds.Max.Y)
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

func cut(original image.Image, db *DB, tileSize, x1, y1, x2, y2 int) <-chan image.Image {
	c := make(chan image.Image)
	sp := image.Point{0, 0}
	go func() {
		newimage := image.NewNRGBA(image.Rect(x1, y1, x2, y2))
		for y := y1; y < y2; y = y + tileSize {
			for x := x1; x < x2; x = x + tileSize {
				r, g, b, _ := original.At(x, y).RGBA()
				color := [3]float64{float64(r), float64(g), float64(b)}
				nearest := db.nearest(color)
				file, err := os.Open(nearest)
				if err == nil {
					t := resize(img, tileSize)
					tile := t.SubImage(t.Bounds())
					tileBounds := image.Rect(x, y, x+tileSize, y+tileSize)
				} else {
					fmt.Println("error:", err)
				}
			} else {
				fmt.Println("error:", nearest)
			}
		}
	}
}
