package controllers

import (
	"cashbag-me-mini/models"
	"cashbag-me-mini/modules/zookeeper"
	"cashbag-me-mini/services"
	"net/http"
	"strings"

	"github.com/labstack/echo"
)

//CreateTransaction  ...
func CreateTransaction(c echo.Context) error {
	body := c.Get("body").(*models.PostTransaction)
	userZoo := zookeeper.GetValueFromZoo("/Users")
	users := strings.Split(userZoo, ",")
	//userBody := (*body).User
	// checktran, _ := redis.RDB.Get("user").Result()
	// fmt.Println(checktran)
	// if checktran == userBody{
	// 	return c.String(http.StatusBadRequest, "User dang request")
	// }
	for _, user := range users {
		if user == body.User {
			// err := redis.RDB.Set("user", body.User, 30000000000)
			// if err != nil {
			// 	log.Println(err)
			// }
			result := services.CreateTransaction(*body)
			return c.JSON(http.StatusCreated, result)
		}
	}
	return c.String(http.StatusBadRequest, "User khong nam trong danh sach hoan tien")

}
