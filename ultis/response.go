package ultis

import (
	"net/http"

	"github.com/labstack/echo"
)

// Response ...
type Response map[string]interface{}

func generateResponse(data interface{}, message string) Response {
	return Response{
		"data":    data,
		"message": message,
	}
}

// Response200 .....
func Response200(c echo.Context, data interface{}, message string) error {
	if message == "" {
		message = " Thanh Cong!"
	}
	return c.JSON(http.StatusOK, generateResponse(data, message))
}
// Response400 ...
func Response400(c echo.Context, data interface{}, message string) error {
	if message == "" {
		message = "Du lieu khong hop le "
	}
	return c.JSON(http.StatusBadRequest, generateResponse(data, message))

}