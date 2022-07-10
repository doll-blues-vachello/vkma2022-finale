package main

import (
	"bytes"
	"database/sql"
	"fmt"
	"hash/fnv"
	"image/png"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	_ "github.com/mattn/go-sqlite3"
)

const API_URL = "localhost:8080/"

func main() {
	var r = gin.Default()
	var db, err = sql.Open("sqlite3", "storage.db")

	if err != nil {
		panic(err)
	}

	defer db.Close()

	r.POST("/album", func(c *gin.Context) {
		var repo = AlbumRepo{db}
		var userId, _ = strconv.ParseInt(c.Query("user_id"), 10, 64)

		var albumId = repo.Add(Album{
			UserID: userId,
			Title:  c.DefaultQuery("title", "untitled"),
		})
		c.JSON(http.StatusOK, albumId)
	})

	r.GET("/album/:album_id", func(c *gin.Context) {
		var repo = AlbumRepo{db}
		var albumId, err = strconv.ParseInt(c.Param("album_id"), 10, 64)

		if err != nil {
			c.JSON(http.StatusBadRequest, "error")
			return
		}

		var album, err2 = repo.GetById(albumId)

		if err2 != nil {
			c.JSON(http.StatusBadRequest, "error")
			return
		}

		c.JSON(http.StatusOK, album)
	})

	r.POST("/photo", func(c *gin.Context) {
		var repo = PhotoRepo{db}
		var albumId, _ = strconv.ParseInt(c.Query("album_id"), 10, 64)
		var uploaderId, _ = strconv.ParseInt(c.Query("uploader_id"), 10, 64)

		data, e := ioutil.ReadAll(c.Request.Body)
		if e != nil {
			panic(e)
		}

		reader := bytes.NewReader(data)
		im, err := png.Decode(reader)
		if err != nil {
			panic("Bad png")
		}

		fileName := hashName(albumId, uploaderId)
		filePath := "var/photos/" + fileName

		f, err := os.OpenFile(filePath, os.O_WRONLY|os.O_CREATE, 0777)
		if err != nil {
			panic("Cannot open file")
		}

		png.Encode(f, im)
		f.Close()

		url := API_URL + filePath

		var photoId = repo.Add(Photo{
			AlbumID:    albumId,
			UploaderID: uploaderId,
			Url:        url,
		})

		c.JSON(http.StatusOK, PhotoResponse{PhotoID: photoId, Url: url})
	})

	r.GET("/photo/:photo_id", func(c *gin.Context) {
		var repo = PhotoRepo{db}
		var photoId, err = strconv.ParseInt(c.Param("photo_id"), 10, 64)

		if err != nil {
			c.JSON(http.StatusBadRequest, "error")
			return
		}

		var photo, err2 = repo.GetById(photoId)

		if err2 != nil {
			c.JSON(http.StatusBadRequest, "error")
			return
		}

		c.JSON(http.StatusOK, photo)
	})

	r.Run()
}

func hashName(id_1 int64, id_2 int64) string {
	h := fnv.New32a()
	s := fmt.Sprintf("%d-%s-%d", id_1, time.Now(), id_2)
	h.Write([]byte(s))

	name := fmt.Sprintf("%d.png", h.Sum32())
	return name
}
