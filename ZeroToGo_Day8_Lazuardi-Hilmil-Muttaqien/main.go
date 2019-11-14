package main

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo"
)

type User struct {
	Id      int
	Name    string
	Address string
}

type xUser struct {
	UserId       int    `json:"id" form:"id"`
	UserName     string `json:"name" form:"name"`
	UserEmail    string `json:"email" form:"email"`
	UserPassword string `json:"password" form:"password"`
}

var xUsers []xUser

func main() {
	//coba hello world
	e := echo.New()
	e.GET("/hello/", HelloController)

	//URL Param
	e.GET("/getuser/:id", GetUser)

	//Query string
	e.GET("/getuser", UserSearchController)

	//Query string with filter
	e.GET("/getuser", UserFilter)

	//task8
	e.GET("/users", GetUsersController)
	e.GET("/users/:id", GetUserController)
	e.POST("/users", CreateUserController)
	e.PUT("/users/:id", UpdateController)
	e.DELETE("/users/:id", DeleteController)
	e.Logger.Fatal(e.Start(":8000"))

	//fmt.Println(len(xUsers))
}

func HelloController(c echo.Context) error {
	return c.String(http.StatusOK, "Hello World")
}

func GetUser(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	user := User{Id: id, Name: "Bane", Address: "Simpang Tambora"}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"user": user,
	})
}

func UserSearchController(c echo.Context) error {
	match := c.QueryParam("match")

	return c.JSON(http.StatusOK, map[string]interface{}{
		"match":  match,
		"result": []string{"ada", "kaila", "snif"},
	})
}

func UserFilter(c echo.Context) error {
	id, _ := strconv.Atoi(c.QueryParam("id"))
	user := User{Id: id, Name: "Poby"}
	if id != 3 {
		return c.JSON(http.StatusNotFound, "tidak ada")
	}
	return c.JSON(http.StatusOK, user)
}

//task 8
// get all users
func GetUsersController(c echo.Context) error {
	return c.JSON(http.StatusOK, map[string]interface{}{
		"messages": "success get all users",
		"users":    xUsers,
	})
}

// create data
func CreateUserController(c echo.Context) error {
	xuser := xUser{}
	c.Bind(&xuser)
	if len(xUsers) == 0 {
		xuser.UserId = 1
	} else {
		newId := xUsers[len(xUsers)-1].UserId + 1
		xuser.UserId = newId
	}
	xUsers = append(xUsers, xuser)
	return c.JSON(http.StatusOK, map[string]interface{}{
		"messages": "success create user",
		"user":     xuser,
		//"total":    len(xUsers),
	})
}

// get user by id
func GetUserController(c echo.Context) error {
	xid, _ := strconv.Atoi(c.Param("id"))
	//var xidstring string = c.Param("id")

	for _, xdata := range xUsers {
		//xarray2 := strings.Fields(xdata)
		xarray2 := xdata.UserId
		//fmt.Println("xarray: ", xarray2)
		//fmt.Println("xid: ", xid)
		if xarray2 == xid {
			return c.JSON(http.StatusOK, map[string]interface{}{
				"messages": "Data ditemukan",
				"user":     xdata,
			})
		}

	}
	return c.JSON(http.StatusNotFound, "Data tidak ditemukan")
}

// update user by id
func UpdateController(c echo.Context) error {
	xid, _ := strconv.Atoi(c.Param("id"))
	//var xidstring string = c.Param("id")
	xuserbind := xUser{}
	c.Bind(&xuserbind)

	xUsers[xid-1].UserName = xuserbind.UserName
	xUsers[xid-1].UserPassword = xuserbind.UserPassword
	xUsers[xid-1].UserEmail = xuserbind.UserEmail

	return c.JSON(http.StatusOK, map[string]interface{}{
		"messages": "Data telah terupdate",
		"user":     xUsers[xid-1],
	})

}

// delete user by id
func DeleteController(c echo.Context) error {
	xid, _ := strconv.Atoi(c.Param("id"))

	if xid == len(xUsers) {
		xUsers[xid-1] = xUser{}
		xUsers = xUsers[:len(xUsers)-1]
	} else {
		xUsers[xid-1] = xUsers[len(xUsers)-1]
		xUsers[len(xUsers)-1] = xUser{}
		xUsers = xUsers[:len(xUsers)-1]
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"messages": "Data telah dihapus",
	})
}
