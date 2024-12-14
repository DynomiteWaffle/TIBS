package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

var list = []string{}

func main() {

	// list = toggle("banana", list)

	// fmt.Println(fmtlist(list))
	router := gin.Default()
	router.GET("/toggle", toggle)
	router.GET("/list", fmtlist)
	router.GET("/gen", gen)
	router.GET("/", home)

	router.Run("localhost:8080")
}
func gen(c *gin.Context) {
	// TODO generate svg

	c.String(http.StatusOK, "wasd")
}
func home(c *gin.Context) {
	// TODO send html

	// c.String(http.StatusOK, "wasd")
	c.Data(http.StatusOK, "text/html; charset=utf-8", []byte(`<!DOCTYPE html>
<textarea name="" id=\"\"></textarea>
<textarea name=\"\" id=\"\"></textarea>
<textarea name=\"\" id=\"\"></textarea>
<button>Gen</button>`))
}

func toggle(c *gin.Context) {
	item := c.Query("item")
	// fmt.Println(item)
	if item == "" {
		c.IndentedJSON(http.StatusBadRequest, "No Item")
		return
	}

	var inlist bool = false
	pos := 0
	// check
	for i := 0; i < len(list); i++ {
		if item == list[i] {
			inlist = true
			pos = i
		}
	}
	// toggle
	if inlist {
		// remove
		list = append(list[:pos], list[pos+1:]...)
		c.IndentedJSON(http.StatusOK, "Removed Item")

	} else {
		// add
		list = append(list, item)
		c.IndentedJSON(http.StatusOK, "Added Item")

	}
	// return list

}
func fmtlist(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, list)
	// c.IndentedJSON(http.StatusOK, strings.Join(list, "\n"))
	// return strings.Join(list, "\n")
}
