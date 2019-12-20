package cache

import (
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"net/http"
)

func Router(r *gin.Engine) {
	r.GET("/:key", GetCache)
	r.DELETE("/:key", DeleteCache)
	r.PUT("/:key", SetCache)
}

func SetCache(c *gin.Context) {
	body,err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		// TODO error Handler
	}

	k := c.Param("key")

	err = cacheHandler.Set(k, []byte(body))
	if err != nil {
		// TODO error Handler
	}

	c.JSON(http.StatusNoContent, nil)
}

func GetCache(c *gin.Context) {
	k := c.Param("key")

	b, err := cacheHandler.Get(k)

	if err != nil {
		// TODO error Handler
	}

	c.JSON(http.StatusOK, gin.H{
		"value": string(b),
	})
}

func DeleteCache(c *gin.Context) {
	k := c.Param("key")
	err := cacheHandler.Del(k)
	if err != nil {
		// TODO error Handler
	}

	c.JSON(http.StatusNoContent, nil)
}
