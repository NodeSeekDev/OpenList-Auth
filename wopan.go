// Package auth wopan.go by xhofe
package auth

import (
	"github.com/gin-gonic/gin"
	"github.com/twoonefour/alist-auth/common"
	"github.com/twoonefour/alist-auth/utils"
	"github.com/xhofe/wopan-sdk-go"
	"net/http"
)

type SmsReq struct {
	Func     string `json:"func"`
	ClientId string `json:"clientId"`
	Param    string `json:"param"`
}
type SmsRes struct {
	STATUS string      `json:"STATUS"`
	MSG    string      `json:"MSG"`
	LOGID  interface{} `json:"LOGID"`
	RSP    struct {
		RSPCODE string      `json:"RSP_CODE"`
		RSPDESC string      `json:"RSP_DESC"`
		DATA    interface{} `json:"DATA"`
	} `json:"RSP"`
}

func wopanLogin(c *gin.Context) {
	req := struct {
		Phone    string `json:"phone"`
		Password string `json:"password"`
	}{}
	if err := c.ShouldBind(&req); err != nil {
		common.Error(c, err)
		return
	}
	w := wopan.Default()
	res, err := w.PcWebLogin(req.Phone, req.Password)
	if err != nil {
		common.Error(c, err)
		return
	}
	c.JSON(200, res)
}

func wopanVerifyCode(c *gin.Context) {
	req := struct {
		Phone      string `json:"phone"`
		Password   string `json:"password"`
		VerifyCode string `json:"verify_code"`
	}{}
	if err := c.ShouldBind(&req); err != nil {
		common.Error(c, err)
		return
	}
	if req.VerifyCode == "" && req.Password == "" && req.Phone != "" {
		wopanSendSmsCode(c, req)
		return
	}
	w := wopan.Default()
	res, err := w.PcLoginVerifyCode(req.Phone, req.Password, req.VerifyCode)
	if err != nil {
		common.Error(c, err)
		return
	}
	c.JSON(200, res)
}

func wopanSendSmsCode(c *gin.Context, req struct {
	Phone      string `json:"phone"`
	Password   string `json:"password"`
	VerifyCode string `json:"verify_code"`
}) {
	w := wopan.Default()
	p, err := w.EncryptParam(wopan.ChannelAPIUser, wopan.Json{
		"operateType": "1",
		"phone":       req.Phone,
		"uuid":        "",
		"verifyCode":  "",
	})
	if err != nil {
		common.Error(c, err)
		return
	}
	smsReq := SmsReq{
		Func:     "pc_send",
		ClientId: wopan.DefaultClientID,
		Param:    p,
	}
	var smsRes SmsRes
	client := utils.RestyClient.R().SetBody(&smsReq).SetResult(&smsRes).SetHeaders(map[string]string{
		"Content-Type": "application/json",
		"Origin":       "https://pan.wo.cn",
		"Referer":      "https://pan.wo.cn/",
		"User-Agent":   "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/115.0.0.0 Safari/537.36",
	})
	_, err = client.Execute("POST", "https://panservice.mail.wo.cn/api-user/sendMessageCodeBase")
	if err != nil {
		common.Error(c, err)
		return
	}
	if smsRes.RSP.RSPCODE != "0000" {
		common.ErrorStr(c, smsRes.RSP.RSPDESC)
		return
	}
	c.Status(http.StatusNoContent)
}
