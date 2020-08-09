package controllers

import (
	"cashbag-me-mini/models"
	//"cashbag-me-mini/modules/zookeeper"
	"cashbag-me-mini/services"
	//"net/http"
	//"strings"
	"cashbag-me-mini/ultis"
	"github.com/labstack/echo"
)

// TransactionCreate  ...
func TransactionCreate(c echo.Context) error {
	body := c.Get("body").(*models.TransactionCreatePayload)
	
	
	rawData,err := services.TransactionCreate(*body)
	//if err
	if err != nil {
		return ultis.Response400(c, nil, err.Error())
	}

	//Success
	return ultis.Response200(c, rawData, "")

}

