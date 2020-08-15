package controllers

import(
	"cashbag-me-mini/services"
	"cashbag-me-mini/models"

)

// UserProgramCreate ...
func UserProgramCreate(c echo.Context) error{
	var (
		body = c.Get("body").(*models.UserProgramCreatePayload)
	)

	// Process data
	rawData, err := services.UserProgramCreate(*body)

	// if err
	if err != nil {
		return ultis.Response400(c, nil, err.Error())
	}

	//Success
	return ultis.Response200(c, rawData, "")
}