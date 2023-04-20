package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"strings"
)

func main() {

	r := gin.Default()

	//➜ curl -X POST -d 'name=飞雪' http://localhost:8080/users
	r.POST("/users", createUser)

	//➜ curl http://localhost:8080/users/99
	r.GET("/users/:id", getUser)

	r.DELETE("/users/:id", deleteUser)

	r.PATCH("/users/:id", updateUserName)

	r.GET("/users", listUser)

	r.Run(":8080")

}

func listUser(c *gin.Context) {

	c.JSON(200, ginUsers)

}

func getUser(c *gin.Context) {

	id := c.Param("id")

	var user ginUser

	found := false

	//类似于数据库的SQL查询

	for _, u := range ginUsers {

		if strings.EqualFold(id, strconv.Itoa(u.ID)) {

			user = u

			found = true

			break

		}

	}

	if found {

		c.JSON(200, user)

	} else {

		c.JSON(404, gin.H{

			"message": "用户不存在",
		})

	}

}

func createUser(c *gin.Context) {

	name := c.DefaultPostForm("name", "")

	if name != "" {

		u := ginUser{ID: len(ginUsers) + 1, Name: name}

		ginUsers = append(ginUsers, u)

		c.JSON(http.StatusCreated, u)

	} else {

		c.JSON(http.StatusOK, gin.H{

			"message": "请输入用户名称",
		})

	}

}

func deleteUser(c *gin.Context) {

	id := c.Param("id")

	i := -1

	//类似于数据库的SQL查询

	for index, u := range ginUsers {

		if strings.EqualFold(id, strconv.Itoa(u.ID)) {

			i = index

			break

		}

	}

	if i >= 0 {

		ginUsers = append(ginUsers[:i], ginUsers[i+1:]...)

		c.JSON(http.StatusNoContent, "")

	} else {

		c.JSON(http.StatusNotFound, gin.H{

			"message": "用户不存在",
		})

	}
}

func updateUserName(c *gin.Context) {

	id := c.Param("id")

	i := -1

	//类似于数据库的SQL查询

	for index, u := range ginUsers {

		if strings.EqualFold(id, strconv.Itoa(u.ID)) {

			i = index

			break

		}

	}

	if i >= 0 {

		ginUsers[i].Name = c.DefaultPostForm("name", ginUsers[i].Name)

		c.JSON(http.StatusOK, ginUsers[i])

	} else {

		c.JSON(http.StatusNotFound, gin.H{

			"message": "用户不存在",
		})

	}

}

var ginUsers = []ginUser{

	{ID: 1, Name: "张三"},

	{ID: 2, Name: "李四"},

	{ID: 3, Name: "王五"},
}

type ginUser struct {
	ID int

	Name string
}
