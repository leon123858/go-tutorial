package url

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"short-url/internal/services/pool"
)

var up pool.IUrlPool

func init() {
	var err error
	up, err = pool.NewUrlPool()
	if err != nil {
		panic("pool init error")
	}
}

// GetLongURL godoc
//
//	@Summary		Get long URL
//	@Description	Get the long URL for a given short URL token
//	@Tags			URLs
//	@Accept			json
//	@Produce		json
//	@Param			token	path		string	true	"Short URL token"
//	@Success		302		{string}	string	"Redirect to long URL"
//	@Failure		404		{string}	string	"URL not found"
//	@Router			/{token} [get]
func GetLongURL(c echo.Context) error {
	// get the token in the path
	token := c.Param("token")
	// get the URL
	url, err := up.GetLongURL(token)
	if err != nil {
		return c.String(http.StatusNotFound, "not found")
	}
	return c.Redirect(http.StatusFound, url)
}

// SetSortURL godoc
//
//	@Summary		Create short URL
//	@Description	Create a new short URL for a given long URL
//	@Tags			URLs
//	@Accept			json
//	@Produce		json
//	@Param			url	body		ShortenRequest	true	"URL to be shortened"
//	@Success		200	{string}	string			"Short URL"
//	@Failure		400	{string}	string			"Invalid request"
//	@Failure		500	{string}	string			"Server error"
//	@Router			/shorten [post]
func SetSortURL(c echo.Context) error {
	// get the Post body
	req := new(ShortenRequest)
	if err := c.Bind(req); err != nil {
		return c.String(http.StatusBadRequest, "invalid request")
	}
	// create short URL
	shortURL, err := up.CreateShortURL(req.URL, req.Password)
	if err != nil {
		return c.String(http.StatusInternalServerError, "create short URL error")
	}
	return c.String(http.StatusOK, shortURL)
}
