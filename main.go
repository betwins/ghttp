package main

import (
	"flag"
	"fmt"
	"ghttp/model"
	"github.com/gin-gonic/gin"
	"io"
	"log"
	"net/http"
)

func main() {

	log.Println("version 1.01")

	/*	forkPtr := flag.Bool("fork", false, "a bool")*/
	host := flag.String("h", "", "主机名")
	port := flag.Int("p", 80, "目标端口号，默认80")
	flag.Parse()

	var listenAddr string
	if *host != "" {
		listenAddr = fmt.Sprintf("%s:%d", *host, *port)
	} else {
		listenAddr = fmt.Sprintf("0.0.0.0:%d", *port)
	}

	gin.SetMode(gin.ReleaseMode)
	gin.DefaultWriter = io.Discard
	router := gin.Default()
	router.GET("/albums", getAlbums)
	router.GET("/albums/medium", getAlbumsMedium)
	router.GET("/albums/bigger", getAlbumsBigger)
	router.GET("/albums/huge", getAlbumsHuge)
	router.POST("/albums", creaeAlbums)

	router.GET("/presure", getPresure)
	router.GET("/presure/medium", getPresureMedium)
	router.GET("/presure/bigger", getPresureBigger)
	router.GET("/presure/huge", getPresureHuge)
	router.POST("/presure", creaePresure)
	router.POST("/presure/noecho", creaePreusreNoEcho)

	router.Run(listenAddr)

}

func getPresure(c *gin.Context) {
	_, _ = c.Writer.WriteString(strPresureM)
}

func getPresureMedium(c *gin.Context) {
	_, _ = c.Writer.WriteString(strPresureM)
}

func getPresureBigger(c *gin.Context) {
	_, _ = c.Writer.WriteString(strPresureB)
}

func getPresureHuge(c *gin.Context) {
	_, _ = c.Writer.WriteString(strPresureH)
}

func creaePresure(c *gin.Context) {
	data, _ := c.GetRawData()
	c.JSON(http.StatusOK, data)
}

func creaePreusreNoEcho(c *gin.Context) {
	c.JSON(http.StatusOK, nil)
}

func getAlbums(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, albums)
}

func getAlbumsMedium(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, mediumAlbums)
}

func getAlbumsBigger(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, biggerAlbums)
}

func getAlbumsHuge(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, hugeAlbums)
}

func creaeAlbums(c *gin.Context) {
	data, _ := c.GetRawData()
	c.JSON(http.StatusOK, data)
}

var strPresureS = "[\n    {\n\t\t\"id\": \"3\",\n        \"title\": \"Sarah Vaughan and Clifford Brown\",\n        \"artist\": \"Sarah Vaughan\",\n        \"price\": 39.99\n    },\n    {\n        \"id\": \"1\",\n        \"title\": \"Blue Train\",\n        \"artist\": \"John Coltrane\",\n        \"price\": 56.99\n    },\n    {\n        \"id\": \"2\",\n        \"title\": \"Jeru\",\n        \"artist\": \"Gerry Mulligan\",\n        \"price\": 17.99\n    },\n    {\n        \"id\": \"3\",\n        \"title\": \"Sarah Vaughan and Clifford Brown\",\n        \"artist\": \"Sarah Vaughan\",\n        \"price\": 39.99\n    },\n    {\n        \"id\": \"1\",\n        \"title\": \"Blue Train\",\n        \"artist\": \"John Coltrane\",\n        \"price\": 56.99\n    },\n    {\n        \"id\": \"2\",\n        \"title\": \"Jeru\",\n        \"artist\": \"Gerry Mulligan\",\n        \"price\": 17.99\n    },\n    {\n        \"id\": \"3\",\n        \"title\": \"Sarah Vaughan and Clifford Brown\",\n        \"artist\": \"Sarah Vaughan\",\n        \"price\": 39.99\n    }\n]"
var strPresureM = "[\n    {\n\t\t\"id\": \"3\",\n        \"title\": \"Sarah Vaughan and Clifford Brown\",\n        \"artist\": \"Sarah Vaughan\",\n        \"price\": 39.99\n    },\n    {\n        \"id\": \"1\",\n        \"title\": \"Blue Train\",\n        \"artist\": \"John Coltrane\",\n        \"price\": 56.99\n    },\n    {\n        \"id\": \"2\",\n        \"title\": \"Jeru\",\n        \"artist\": \"Gerry Mulligan\",\n        \"price\": 17.99\n    },\n    {\n        \"id\": \"3\",\n        \"title\": \"Sarah Vaughan and Clifford Brown\",\n        \"artist\": \"Sarah Vaughan\",\n        \"price\": 39.99\n    },\n    {\n        \"id\": \"1\",\n        \"title\": \"Blue Train\",\n        \"artist\": \"John Coltrane\",\n        \"price\": 56.99\n    },\n    {\n        \"id\": \"2\",\n        \"title\": \"Jeru\",\n        \"artist\": \"Gerry Mulligan\",\n        \"price\": 17.99\n    },\n    {\n        \"id\": \"3\",\n        \"title\": \"Sarah Vaughan and Clifford Brown\",\n        \"artist\": \"Sarah Vaughan\",\n        \"price\": 39.99\n    },\n\t    {\n\t\t\"id\": \"3\",\n        \"title\": \"Sarah Vaughan and Clifford Brown\",\n        \"artist\": \"Sarah Vaughan\",\n        \"price\": 39.99\n    },\n    {\n        \"id\": \"1\",\n        \"title\": \"Blue Train\",\n        \"artist\": \"John Coltrane\",\n        \"price\": 56.99\n    },\n    {\n        \"id\": \"2\",\n        \"title\": \"Jeru\",\n        \"artist\": \"Gerry Mulligan\",\n        \"price\": 17.99\n    },\n    {\n        \"id\": \"3\",\n        \"title\": \"Sarah Vaughan and Clifford Brown\",\n        \"artist\": \"Sarah Vaughan\",\n        \"price\": 39.99\n    },\n    {\n        \"id\": \"1\",\n        \"title\": \"Blue Train\",\n        \"artist\": \"John Coltrane\",\n        \"price\": 56.99\n    },\n    {\n        \"id\": \"2\",\n        \"title\": \"Jeru\",\n        \"artist\": \"Gerry Mulligan\",\n        \"price\": 17.99\n    },\n    {\n        \"id\": \"3\",\n        \"title\": \"Sarah Vaughan and Clifford Brown\",\n        \"artist\": \"Sarah Vaughan\",\n        \"price\": 39.99\n    }\n]"
var strPresureB = "[\n    {\n\t\t\"id\": \"3\",\n        \"title\": \"Sarah Vaughan and Clifford Brown\",\n        \"artist\": \"Sarah Vaughan\",\n        \"price\": 39.99\n    },\n    {\n        \"id\": \"1\",\n        \"title\": \"Blue Train\",\n        \"artist\": \"John Coltrane\",\n        \"price\": 56.99\n    },\n    {\n        \"id\": \"2\",\n        \"title\": \"Jeru\",\n        \"artist\": \"Gerry Mulligan\",\n        \"price\": 17.99\n    },\n    {\n        \"id\": \"3\",\n        \"title\": \"Sarah Vaughan and Clifford Brown\",\n        \"artist\": \"Sarah Vaughan\",\n        \"price\": 39.99\n    },\n    {\n        \"id\": \"1\",\n        \"title\": \"Blue Train\",\n        \"artist\": \"John Coltrane\",\n        \"price\": 56.99\n    },\n    {\n        \"id\": \"2\",\n        \"title\": \"Jeru\",\n        \"artist\": \"Gerry Mulligan\",\n        \"price\": 17.99\n    },\n    {\n        \"id\": \"3\",\n        \"title\": \"Sarah Vaughan and Clifford Brown\",\n        \"artist\": \"Sarah Vaughan\",\n        \"price\": 39.99\n    },\n\t    {\n\t\t\"id\": \"3\",\n        \"title\": \"Sarah Vaughan and Clifford Brown\",\n        \"artist\": \"Sarah Vaughan\",\n        \"price\": 39.99\n    },\n    {\n        \"id\": \"1\",\n        \"title\": \"Blue Train\",\n        \"artist\": \"John Coltrane\",\n        \"price\": 56.99\n    },\n    {\n        \"id\": \"2\",\n        \"title\": \"Jeru\",\n        \"artist\": \"Gerry Mulligan\",\n        \"price\": 17.99\n    },\n    {\n        \"id\": \"3\",\n        \"title\": \"Sarah Vaughan and Clifford Brown\",\n        \"artist\": \"Sarah Vaughan\",\n        \"price\": 39.99\n    },\n    {\n        \"id\": \"1\",\n        \"title\": \"Blue Train\",\n        \"artist\": \"John Coltrane\",\n        \"price\": 56.99\n    },\n    {\n        \"id\": \"2\",\n        \"title\": \"Jeru\",\n        \"artist\": \"Gerry Mulligan\",\n        \"price\": 17.99\n    },\n    {\n        \"id\": \"3\",\n        \"title\": \"Sarah Vaughan and Clifford Brown\",\n        \"artist\": \"Sarah Vaughan\",\n        \"price\": 39.99\n    },\n\t   {\n\t\t\"id\": \"3\",\n        \"title\": \"Sarah Vaughan and Clifford Brown\",\n        \"artist\": \"Sarah Vaughan\",\n        \"price\": 39.99\n    },\n    {\n        \"id\": \"1\",\n        \"title\": \"Blue Train\",\n        \"artist\": \"John Coltrane\",\n        \"price\": 56.99\n    },\n    {\n        \"id\": \"2\",\n        \"title\": \"Jeru\",\n        \"artist\": \"Gerry Mulligan\",\n        \"price\": 17.99\n    },\n    {\n        \"id\": \"3\",\n        \"title\": \"Sarah Vaughan and Clifford Brown\",\n        \"artist\": \"Sarah Vaughan\",\n        \"price\": 39.99\n    },\n    {\n        \"id\": \"1\",\n        \"title\": \"Blue Train\",\n        \"artist\": \"John Coltrane\",\n        \"price\": 56.99\n    },\n    {\n        \"id\": \"2\",\n        \"title\": \"Jeru\",\n        \"artist\": \"Gerry Mulligan\",\n        \"price\": 17.99\n    },\n    {\n        \"id\": \"3\",\n        \"title\": \"Sarah Vaughan and Clifford Brown\",\n        \"artist\": \"Sarah Vaughan\",\n        \"price\": 39.99\n    },\n\t    {\n\t\t\"id\": \"3\",\n        \"title\": \"Sarah Vaughan and Clifford Brown\",\n        \"artist\": \"Sarah Vaughan\",\n        \"price\": 39.99\n    },\n    {\n        \"id\": \"1\",\n        \"title\": \"Blue Train\",\n        \"artist\": \"John Coltrane\",\n        \"price\": 56.99\n    },\n    {\n        \"id\": \"2\",\n        \"title\": \"Jeru\",\n        \"artist\": \"Gerry Mulligan\",\n        \"price\": 17.99\n    },\n    {\n        \"id\": \"3\",\n        \"title\": \"Sarah Vaughan and Clifford Brown\",\n        \"artist\": \"Sarah Vaughan\",\n        \"price\": 39.99\n    },\n    {\n        \"id\": \"1\",\n        \"title\": \"Blue Train\",\n        \"artist\": \"John Coltrane\",\n        \"price\": 56.99\n    },\n    {\n        \"id\": \"2\",\n        \"title\": \"Jeru\",\n        \"artist\": \"Gerry Mulligan\",\n        \"price\": 17.99\n    },\n    {\n        \"id\": \"3\",\n        \"title\": \"Sarah Vaughan and Clifford Brown\",\n        \"artist\": \"Sarah Vaughan\",\n        \"price\": 39.99\n    }\n]"
var strPresureH = "[\n    {\n\t\t\"id\": \"3\",\n        \"title\": \"Sarah Vaughan and Clifford Brown\",\n        \"artist\": \"Sarah Vaughan\",\n        \"price\": 39.99\n    },\n    {\n        \"id\": \"1\",\n        \"title\": \"Blue Train\",\n        \"artist\": \"John Coltrane\",\n        \"price\": 56.99\n    },\n    {\n        \"id\": \"2\",\n        \"title\": \"Jeru\",\n        \"artist\": \"Gerry Mulligan\",\n        \"price\": 17.99\n    },\n    {\n        \"id\": \"3\",\n        \"title\": \"Sarah Vaughan and Clifford Brown\",\n        \"artist\": \"Sarah Vaughan\",\n        \"price\": 39.99\n    },\n    {\n        \"id\": \"1\",\n        \"title\": \"Blue Train\",\n        \"artist\": \"John Coltrane\",\n        \"price\": 56.99\n    },\n    {\n        \"id\": \"2\",\n        \"title\": \"Jeru\",\n        \"artist\": \"Gerry Mulligan\",\n        \"price\": 17.99\n    },\n    {\n        \"id\": \"3\",\n        \"title\": \"Sarah Vaughan and Clifford Brown\",\n        \"artist\": \"Sarah Vaughan\",\n        \"price\": 39.99\n    },\n\t    {\n\t\t\"id\": \"3\",\n        \"title\": \"Sarah Vaughan and Clifford Brown\",\n        \"artist\": \"Sarah Vaughan\",\n        \"price\": 39.99\n    },\n    {\n        \"id\": \"1\",\n        \"title\": \"Blue Train\",\n        \"artist\": \"John Coltrane\",\n        \"price\": 56.99\n    },\n    {\n        \"id\": \"2\",\n        \"title\": \"Jeru\",\n        \"artist\": \"Gerry Mulligan\",\n        \"price\": 17.99\n    },\n    {\n        \"id\": \"3\",\n        \"title\": \"Sarah Vaughan and Clifford Brown\",\n        \"artist\": \"Sarah Vaughan\",\n        \"price\": 39.99\n    },\n    {\n        \"id\": \"1\",\n        \"title\": \"Blue Train\",\n        \"artist\": \"John Coltrane\",\n        \"price\": 56.99\n    },\n    {\n        \"id\": \"2\",\n        \"title\": \"Jeru\",\n        \"artist\": \"Gerry Mulligan\",\n        \"price\": 17.99\n    },\n    {\n        \"id\": \"3\",\n        \"title\": \"Sarah Vaughan and Clifford Brown\",\n        \"artist\": \"Sarah Vaughan\",\n        \"price\": 39.99\n    },\n\t   {\n\t\t\"id\": \"3\",\n        \"title\": \"Sarah Vaughan and Clifford Brown\",\n        \"artist\": \"Sarah Vaughan\",\n        \"price\": 39.99\n    },\n    {\n        \"id\": \"1\",\n        \"title\": \"Blue Train\",\n        \"artist\": \"John Coltrane\",\n        \"price\": 56.99\n    },\n    {\n        \"id\": \"2\",\n        \"title\": \"Jeru\",\n        \"artist\": \"Gerry Mulligan\",\n        \"price\": 17.99\n    },\n    {\n        \"id\": \"3\",\n        \"title\": \"Sarah Vaughan and Clifford Brown\",\n        \"artist\": \"Sarah Vaughan\",\n        \"price\": 39.99\n    },\n    {\n        \"id\": \"1\",\n        \"title\": \"Blue Train\",\n        \"artist\": \"John Coltrane\",\n        \"price\": 56.99\n    },\n    {\n        \"id\": \"2\",\n        \"title\": \"Jeru\",\n        \"artist\": \"Gerry Mulligan\",\n        \"price\": 17.99\n    },\n    {\n        \"id\": \"3\",\n        \"title\": \"Sarah Vaughan and Clifford Brown\",\n        \"artist\": \"Sarah Vaughan\",\n        \"price\": 39.99\n    },\n\t    {\n\t\t\"id\": \"3\",\n        \"title\": \"Sarah Vaughan and Clifford Brown\",\n        \"artist\": \"Sarah Vaughan\",\n        \"price\": 39.99\n    },\n    {\n        \"id\": \"1\",\n        \"title\": \"Blue Train\",\n        \"artist\": \"John Coltrane\",\n        \"price\": 56.99\n    },\n    {\n        \"id\": \"2\",\n        \"title\": \"Jeru\",\n        \"artist\": \"Gerry Mulligan\",\n        \"price\": 17.99\n    },\n    {\n        \"id\": \"3\",\n        \"title\": \"Sarah Vaughan and Clifford Brown\",\n        \"artist\": \"Sarah Vaughan\",\n        \"price\": 39.99\n    },\n    {\n        \"id\": \"1\",\n        \"title\": \"Blue Train\",\n        \"artist\": \"John Coltrane\",\n        \"price\": 56.99\n    },\n    {\n        \"id\": \"2\",\n        \"title\": \"Jeru\",\n        \"artist\": \"Gerry Mulligan\",\n        \"price\": 17.99\n    },\n    {\n        \"id\": \"3\",\n        \"title\": \"Sarah Vaughan and Clifford Brown\",\n        \"artist\": \"Sarah Vaughan\",\n        \"price\": 39.99\n    },\n\t  {\n\t\t\"id\": \"3\",\n        \"title\": \"Sarah Vaughan and Clifford Brown\",\n        \"artist\": \"Sarah Vaughan\",\n        \"price\": 39.99\n    },\n    {\n        \"id\": \"1\",\n        \"title\": \"Blue Train\",\n        \"artist\": \"John Coltrane\",\n        \"price\": 56.99\n    },\n    {\n        \"id\": \"2\",\n        \"title\": \"Jeru\",\n        \"artist\": \"Gerry Mulligan\",\n        \"price\": 17.99\n    },\n    {\n        \"id\": \"3\",\n        \"title\": \"Sarah Vaughan and Clifford Brown\",\n        \"artist\": \"Sarah Vaughan\",\n        \"price\": 39.99\n    },\n    {\n        \"id\": \"1\",\n        \"title\": \"Blue Train\",\n        \"artist\": \"John Coltrane\",\n        \"price\": 56.99\n    },\n    {\n        \"id\": \"2\",\n        \"title\": \"Jeru\",\n        \"artist\": \"Gerry Mulligan\",\n        \"price\": 17.99\n    },\n    {\n        \"id\": \"3\",\n        \"title\": \"Sarah Vaughan and Clifford Brown\",\n        \"artist\": \"Sarah Vaughan\",\n        \"price\": 39.99\n    },\n\t    {\n\t\t\"id\": \"3\",\n        \"title\": \"Sarah Vaughan and Clifford Brown\",\n        \"artist\": \"Sarah Vaughan\",\n        \"price\": 39.99\n    },\n    {\n        \"id\": \"1\",\n        \"title\": \"Blue Train\",\n        \"artist\": \"John Coltrane\",\n        \"price\": 56.99\n    },\n    {\n        \"id\": \"2\",\n        \"title\": \"Jeru\",\n        \"artist\": \"Gerry Mulligan\",\n        \"price\": 17.99\n    },\n    {\n        \"id\": \"3\",\n        \"title\": \"Sarah Vaughan and Clifford Brown\",\n        \"artist\": \"Sarah Vaughan\",\n        \"price\": 39.99\n    },\n    {\n        \"id\": \"1\",\n        \"title\": \"Blue Train\",\n        \"artist\": \"John Coltrane\",\n        \"price\": 56.99\n    },\n    {\n        \"id\": \"2\",\n        \"title\": \"Jeru\",\n        \"artist\": \"Gerry Mulligan\",\n        \"price\": 17.99\n    },\n    {\n        \"id\": \"3\",\n        \"title\": \"Sarah Vaughan and Clifford Brown\",\n        \"artist\": \"Sarah Vaughan\",\n        \"price\": 39.99\n    },\n\t   {\n\t\t\"id\": \"3\",\n        \"title\": \"Sarah Vaughan and Clifford Brown\",\n        \"artist\": \"Sarah Vaughan\",\n        \"price\": 39.99\n    },\n    {\n        \"id\": \"1\",\n        \"title\": \"Blue Train\",\n        \"artist\": \"John Coltrane\",\n        \"price\": 56.99\n    },\n    {\n        \"id\": \"2\",\n        \"title\": \"Jeru\",\n        \"artist\": \"Gerry Mulligan\",\n        \"price\": 17.99\n    },\n    {\n        \"id\": \"3\",\n        \"title\": \"Sarah Vaughan and Clifford Brown\",\n        \"artist\": \"Sarah Vaughan\",\n        \"price\": 39.99\n    },\n    {\n        \"id\": \"1\",\n        \"title\": \"Blue Train\",\n        \"artist\": \"John Coltrane\",\n        \"price\": 56.99\n    },\n    {\n        \"id\": \"2\",\n        \"title\": \"Jeru\",\n        \"artist\": \"Gerry Mulligan\",\n        \"price\": 17.99\n    },\n    {\n        \"id\": \"3\",\n        \"title\": \"Sarah Vaughan and Clifford Brown\",\n        \"artist\": \"Sarah Vaughan\",\n        \"price\": 39.99\n    },\n\t    {\n\t\t\"id\": \"3\",\n        \"title\": \"Sarah Vaughan and Clifford Brown\",\n        \"artist\": \"Sarah Vaughan\",\n        \"price\": 39.99\n    },\n    {\n        \"id\": \"1\",\n        \"title\": \"Blue Train\",\n        \"artist\": \"John Coltrane\",\n        \"price\": 56.99\n    },\n    {\n        \"id\": \"2\",\n        \"title\": \"Jeru\",\n        \"artist\": \"Gerry Mulligan\",\n        \"price\": 17.99\n    },\n    {\n        \"id\": \"3\",\n        \"title\": \"Sarah Vaughan and Clifford Brown\",\n        \"artist\": \"Sarah Vaughan\",\n        \"price\": 39.99\n    },\n    {\n        \"id\": \"1\",\n        \"title\": \"Blue Train\",\n        \"artist\": \"John Coltrane\",\n        \"price\": 56.99\n    },\n    {\n        \"id\": \"2\",\n        \"title\": \"Jeru\",\n        \"artist\": \"Gerry Mulligan\",\n        \"price\": 17.99\n    },\n    {\n        \"id\": \"3\",\n        \"title\": \"Sarah Vaughan and Clifford Brown\",\n        \"artist\": \"Sarah Vaughan\",\n        \"price\": 39.99\n    }\n]"

var albums = []model.Album{
	{ID: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
	{ID: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
	{ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
	{ID: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
	{ID: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
	{ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
	{ID: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
	{ID: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
	{ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
	{ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
	{ID: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
	{ID: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
	{ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
}

var mediumAlbums = []model.Album{
	{ID: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
	{ID: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
	{ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
	{ID: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
	{ID: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
	{ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
	{ID: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
	{ID: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
	{ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
	{ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
	{ID: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
	{ID: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
	{ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
	{ID: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
	{ID: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
	{ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
	{ID: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
	{ID: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
	{ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
	{ID: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
	{ID: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
	{ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
	{ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
	{ID: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
	{ID: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
	{ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
	{ID: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
	{ID: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
	{ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
	{ID: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
	{ID: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
	{ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
	{ID: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
	{ID: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
	{ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
	{ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
	{ID: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
	{ID: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
	{ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
	{ID: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
	{ID: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
	{ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
	{ID: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
	{ID: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
	{ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
	{ID: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
	{ID: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
	{ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
	{ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
	{ID: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
	{ID: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
	{ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
}

var biggerAlbums = []model.Album{
	{ID: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
	{ID: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
	{ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
	{ID: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
	{ID: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
	{ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
	{ID: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
	{ID: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
	{ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
	{ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
	{ID: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
	{ID: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
	{ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
	{ID: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
	{ID: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
	{ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
	{ID: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
	{ID: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
	{ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
	{ID: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
	{ID: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
	{ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
	{ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
	{ID: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
	{ID: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
	{ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
	{ID: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
	{ID: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
	{ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
	{ID: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
	{ID: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
	{ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
	{ID: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
	{ID: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
	{ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
	{ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
	{ID: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
	{ID: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
	{ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
	{ID: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
	{ID: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
	{ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
	{ID: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
	{ID: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
	{ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
	{ID: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
	{ID: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
	{ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
	{ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
	{ID: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
	{ID: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
	{ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
	{ID: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
	{ID: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
	{ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
	{ID: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
	{ID: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
	{ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
	{ID: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
	{ID: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
	{ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
	{ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
	{ID: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
	{ID: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
	{ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
	{ID: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
	{ID: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
	{ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
	{ID: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
	{ID: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
	{ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
	{ID: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
	{ID: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
	{ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
	{ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
	{ID: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
	{ID: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
	{ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
	{ID: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
	{ID: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
	{ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
	{ID: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
	{ID: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
	{ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
	{ID: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
	{ID: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
	{ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
	{ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
	{ID: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
	{ID: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
	{ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
	{ID: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
	{ID: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
	{ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
	{ID: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
	{ID: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
	{ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
	{ID: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
	{ID: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
	{ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
	{ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
	{ID: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
	{ID: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
	{ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
}

var hugeAlbums = []model.Album{
	{ID: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
	{ID: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
	{ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
	{ID: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
	{ID: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
	{ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
	{ID: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
	{ID: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
	{ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
	{ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
	{ID: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
	{ID: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
	{ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
	{ID: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
	{ID: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
	{ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
	{ID: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
	{ID: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
	{ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
	{ID: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
	{ID: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
	{ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
	{ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
	{ID: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
	{ID: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
	{ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
	{ID: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
	{ID: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
	{ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
	{ID: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
	{ID: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
	{ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
	{ID: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
	{ID: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
	{ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
	{ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
	{ID: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
	{ID: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
	{ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
	{ID: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
	{ID: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
	{ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
	{ID: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
	{ID: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
	{ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
	{ID: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
	{ID: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
	{ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
	{ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
	{ID: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
	{ID: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
	{ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
	{ID: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
	{ID: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
	{ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
	{ID: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
	{ID: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
	{ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
	{ID: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
	{ID: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
	{ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
	{ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
	{ID: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
	{ID: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
	{ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
	{ID: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
	{ID: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
	{ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
	{ID: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
	{ID: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
	{ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
	{ID: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
	{ID: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
	{ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
	{ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
	{ID: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
	{ID: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
	{ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
	{ID: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
	{ID: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
	{ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
	{ID: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
	{ID: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
	{ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
	{ID: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
	{ID: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
	{ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
	{ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
	{ID: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
	{ID: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
	{ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
	{ID: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
	{ID: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
	{ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
	{ID: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
	{ID: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
	{ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
	{ID: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
	{ID: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
	{ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
	{ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
	{ID: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
	{ID: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
	{ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
	{ID: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
	{ID: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
	{ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
	{ID: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
	{ID: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
	{ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
	{ID: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
	{ID: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
	{ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
	{ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
	{ID: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
	{ID: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
	{ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
	{ID: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
	{ID: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
	{ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
	{ID: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
	{ID: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
	{ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
	{ID: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
	{ID: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
	{ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
	{ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
	{ID: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
	{ID: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
	{ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
	{ID: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
	{ID: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
	{ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
	{ID: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
	{ID: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
	{ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
	{ID: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
	{ID: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
	{ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
	{ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
	{ID: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
	{ID: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
	{ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
	{ID: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
	{ID: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
	{ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
	{ID: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
	{ID: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
	{ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
	{ID: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
	{ID: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
	{ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
	{ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
	{ID: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
	{ID: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
	{ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
	{ID: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
	{ID: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
	{ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
	{ID: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
	{ID: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
	{ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
	{ID: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
	{ID: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
	{ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
	{ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
	{ID: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
	{ID: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
	{ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
	{ID: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
	{ID: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
	{ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
	{ID: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
	{ID: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
	{ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
	{ID: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
	{ID: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
	{ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
	{ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
	{ID: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
	{ID: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
	{ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
	{ID: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
	{ID: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
	{ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
	{ID: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
	{ID: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
	{ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
	{ID: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
	{ID: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
	{ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
	{ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
	{ID: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
	{ID: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
	{ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
	{ID: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
	{ID: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
	{ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
	{ID: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
	{ID: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
	{ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
	{ID: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
	{ID: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
	{ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
	{ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
	{ID: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
	{ID: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
	{ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
}
