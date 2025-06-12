package auth

import (
	"fmt"
	"github.com/twoonefour/alist-auth/common"
	"github.com/twoonefour/alist-auth/utils"

	"github.com/gin-gonic/gin"
)

var (
	baiduClientId     string
	baiduClientSecret string
	baiduCallbackUri  = FrontEndBaseUrl + "/tool/baidu/callback"
)

func baiduToken(c *gin.Context) {
	code := c.Query("code")
	if code == "" {
		common.ErrorStr(c, "no code")
		return
	}
	res, err := utils.RestyClient.R().
		Get(fmt.Sprintf(
			"https://openapi.baidu.com/oauth/2.0/token?grant_type=authorization_code&code=%s&client_id=%s&client_secret=%s&redirect_uri=%s",
			code, baiduClientId, baiduClientSecret, baiduCallbackUri))
	if err != nil {
		common.Error(c, err)
		return
	}
	common.JsonBytes(c, res.Body())
}
