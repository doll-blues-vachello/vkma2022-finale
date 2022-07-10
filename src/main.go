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

	r.DELETE("/album/:id", func(c *gin.Context) {
		var repo = AlbumRepo{db}
		var id, _ = strconv.ParseInt(c.Param("id"), 10, 64)

		repo.DeleteByID(id)
		c.JSON(http.StatusOK, "")
	})

	r.PUT("/album/:album_id", func(c *gin.Context) {
		var repo = AlbumRepo{db}
		var albumId, _ = strconv.ParseInt(c.Param("album_id"), 10, 64)
		var userId, _ = strconv.ParseInt(c.Query("user_id"), 10, 64)
		var title = c.Query("title")

		repo.Update(Album{ID: albumId, UserID: userId, Title: title})
	})

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

		var album, err2 = repo.GetByID(albumId)

		if err2 != nil {
			c.JSON(http.StatusBadRequest, "error")
			return
		}

		c.JSON(http.StatusOK, album)
	})

	r.DELETE("/photo/:id", func(c *gin.Context) {
		var repo = PhotoRepo{db}
		var id, _ = strconv.ParseInt(c.Param("id"), 10, 64)

		repo.DeleteByID(id)
		c.JSON(http.StatusOK, "")
	})

	r.PUT("/photo/:photo_id", func(c *gin.Context) {
		var repo = PhotoRepo{db}
		var photoId, _ = strconv.ParseInt(c.Param("photo_id"), 10, 64)
		var albumId, _ = strconv.ParseInt(c.Query("album_id"), 10, 64)

		repo.Update(Photo{ID: photoId, AlbumID: albumId})
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

		fileName := hashName(albumId, uploaderId, "png")
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

		var photo, err2 = repo.GetByID(photoId)

		if err2 != nil {
			c.JSON(http.StatusBadRequest, "error")
			return
		}

		c.JSON(http.StatusOK, photo)
	})

	r.DELETE("/autograph/svg/:id", func(c *gin.Context) {
		var repo = ImageAutographRepo{db}
		var id, _ = strconv.ParseInt(c.Param("id"), 10, 64)

		repo.DeleteByID(id)
		c.JSON(http.StatusOK, "")
	})

	r.POST("/autograph/svg", func(c *gin.Context) {
		var repo = ImageAutographRepo{db}
		var photoId, _ = strconv.ParseInt(c.Query("photo_id"), 10, 64)
		var authorId, _ = strconv.ParseInt(c.Query("author_id"), 10, 64)

		data, e := ioutil.ReadAll(c.Request.Body)
		if e != nil {
			panic(e)
		}

		fileName := hashName(photoId, authorId, "svg")
		filePath := "var/autographs/" + fileName

		f, err := os.OpenFile(filePath, os.O_WRONLY|os.O_CREATE, 0777)
		if err != nil {
			panic("Cannot open file")
		}

		f.Write(data)
		f.Close()

		url := API_URL + filePath

		var autographId = repo.Add(ImageAutograph{
			PhotoID:  photoId,
			AuthorID: authorId,
			Url:      url,
		})

		c.JSON(http.StatusOK, ImageAutographResponse{
			AutographID: autographId,
			Url:         url,
		})
	})

	r.GET("/autograph/svg/:autograph_id", func(c *gin.Context) {
		var repo = ImageAutographRepo{db}
		var autographId, e = strconv.ParseInt(c.Param("autograph_id"), 10, 64)

		if e != nil {
			c.JSON(http.StatusBadRequest, "error")
			fmt.Println(e)
			return
		}

		var autograph, err2 = repo.GetByID(autographId)

		if err2 != nil {
			c.JSON(http.StatusNotFound, "error")
			fmt.Println(err2)
			return
		}

		c.JSON(http.StatusOK, autograph)
	})

	r.PUT("/autograph/text/:id", func(c *gin.Context) {
		var repo = TextAutographRepo{db}
		var id, _ = strconv.ParseInt(c.Param("id"), 10, 64)
		var text = c.Query("text")

		repo.Update(TextAutograph{ID: id, Text: text})
		c.JSON(http.StatusOK, "")
	})

	r.DELETE("/autograph/text/:id", func(c *gin.Context) {
		var repo = TextAutographRepo{db}
		var id, _ = strconv.ParseInt(c.Param("id"), 10, 64)

		repo.DeleteByID(id)
		c.JSON(http.StatusOK, "")
	})

	r.POST("/autograph/text", func(c *gin.Context) {
		var repo = TextAutographRepo{db}
		var photoId, _ = strconv.ParseInt(c.Query("photo_id"), 10, 64)
		var authorId, _ = strconv.ParseInt(c.Query("author_id"), 10, 64)
		var text = c.Query("text")

		var autographId = repo.Add(TextAutograph{
			PhotoID:  photoId,
			AuthorID: authorId,
			Text:     text,
		})

		c.JSON(http.StatusOK, autographId)
	})

	r.GET("/autograph/text/:autograph_id", func(c *gin.Context) {
		var repo = TextAutographRepo{db}
		var autographId, e = strconv.ParseInt(c.Param("autograph_id"), 10, 64)

		if e != nil {
			c.JSON(http.StatusBadRequest, "error")
			fmt.Println(e)
			return
		}

		var autograph, err2 = repo.GetByID(autographId)

		if err2 != nil {
			c.JSON(http.StatusNotFound, "error")
			fmt.Println(err2)
			return
		}

		c.JSON(http.StatusOK, autograph)
	})

	r.Run(":4567")
}

func hashName(id_1 int64, id_2 int64, ext string) string {
	h := fnv.New32a()
	s := fmt.Sprintf("%d-%s-%d", id_1, time.Now(), id_2)
	h.Write([]byte(s))

	name := fmt.Sprintf("%d.%s", h.Sum32(), ext)
	return name
}
