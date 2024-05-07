package url

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

// GetLongURL
// @Summary		獲取用戶擁有的社區列表
// @Description	切換社區, 選擇社區用
// @Param			userId		query	string	true	"用戶 firebase id"
// @Param			communityId	query	string	false	"社區 id, 用來篩選特定社區, 測權限用的, 無權限回傳空陣列"
// @Tags			community
// @Success		200	{object}	[]string
// @Failure		400	{object}	string
// @Router			/community/getCommunityList [get]
func GetLongURL(c echo.Context) error {
	// get the URL
	return c.String(http.StatusOK, "get")
}
