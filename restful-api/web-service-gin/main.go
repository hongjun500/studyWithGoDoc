package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	_ "net/http"
)

// 结构体，不用json:将使用结构字段的大写字段名称
type album struct {
	ID     string  `json:"id"`
	Title  string  `json:"title"`
	Artist string  `json:"artist"`
	Price  float64 `json:"price"`
}

//切片
var albums = []album{
	{ID: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
	{ID: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
	{ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
}

//gin.Context是gin最重要的部分。它携带请求详细信息、验证和序列化 JSON 等
func getAlbums(c *gin.Context) {
	//以将结构序列化为 JSON 并将其添加到响应中
	c.IndentedJSON(http.StatusOK, albums)
}

func getAlbumById(c *gin.Context) {
	//获取参数"id"
	id := c.Param("id")
	for _, obj := range albums {
		if id == obj.ID {
			//如果入参的id=albums中数据ID匹配则返回
			c.IndentedJSON(http.StatusOK, obj)
			return
		}
	}
	//找不到则返回404
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "album not found"})
}

func postAlbums(c *gin.Context) {
	var newAlbum album
	//Context.BindJSON 将请求正文绑定到newAlbum
	if err := c.BindJSON(&newAlbum); err != nil {
		//如果有错误直接返回
		return
	}
	//将newAlbum数据追加到切片albums
	albums = append(albums, newAlbum)
	c.IndentedJSON(http.StatusCreated, newAlbum)
}

func main() {

	//初始化一个路由
	router := gin.Default()
	//使用GET函数将路径/albums和getAlbums名称相关联
	router.GET("/albums", getAlbums)
	router.GET("/albums/:id", getAlbumById)
	router.POST("/postAlbums", postAlbums)
	router.Run("localhost:8889")
}
