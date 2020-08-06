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
	check := 0
	for _, user := range users {
		if user == body.User {
			check = 1
		}
	}
	if check == 0 {
		return c.String(http.StatusBadRequest, "User khong nam trong danh sach hoan tien")
	}
	result := services.CreateTransaction(*body)
	return c.JSON(http.StatusCreated, result)

}