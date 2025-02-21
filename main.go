package main

import (
	"net/http"
	"github.com/gin-gonic/gin"
)

// album represents data a record album

type album struct {
	ID string `json:"id"`
	Title string `json:"Title"`
	Artist string `json:"artist"`
	Price float64 `json:"price"`
}

func main(){
	router := gin.Default()
	router.GET("/albums",getAlbums)
	router.POST("/albums",Add_Album)
	router.GET("/albums/:id",getAlbumsById)
	router.Run("localhost:8080")
}
// this is the slice which holds the album collection
var albums =[]album{
	{ID:"1",Title:"blue Train", Artist :"Jone Coltrane",Price : 56.32},
	{ID:"2",Title:"Jeru", Artist :"Gerry Mulligan",Price : 17.99},
	{ID:"3",Title:"Sarah Vaughan and lifford brown", Artist :"sarrah vaughan",Price : 76.32},
} 


// getAlbums responds with the list of all albums as Json
func getAlbums(c *gin.Context) {
	c.IndentedJSON(http.StatusOK,albums)
}

// implimenting the logic to add a album to album list
func  Add_Album(c *gin.Context){
	var newalbum album
	
	//try to bind the json to new album
	if err := c.BindJSON(&newalbum); err != nil{
		// return c.JSON(http.StatusBadRequest,gin.H{"massage":"invalid payload"})
		return
	}
	//add the new album to the slice
	albums = append(albums,newalbum)
	c.IndentedJSON(http.StatusCreated,newalbum)

}

func getAlbumsById(c *gin.Context){
	var id string
	id = c.Param("id")

	// var album album
	for _,v := range albums {
		if v.ID == id {
			c.JSON(http.StatusOK,v)
			return
		}
		
	}
	c.JSON(http.StatusNotFound,gin.H{"message":"Album not found"})
	
}


