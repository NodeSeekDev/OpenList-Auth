package auth

import (
	"github.com/gin-gonic/gin"
	"github.com/twoonefour/115-sdk-go"
	"github.com/twoonefour/alist-auth/utils"
	"net/http"
)

var (
	clientID string
)

type TokenReq struct {
	Uid          string `json:"uid" binding:"required"`
	CodeVerifier string `json:"code_verifier" binding:"required"`
}

func Open115Qrcode(c *gin.Context) {
	client := sdk.New()
	var cv string
	var resp *sdk.AuthDeviceCodeResp
	cv, err := utils.GenerateCodeVerifier(64)
	if err != nil {
		c.Error(err)
		return
	}
	resp, err = client.AuthDeviceCode(c, clientID, cv)
	if err != nil {
		c.Error(err)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code_verifier": cv,
		"resp":          resp,
	})
}

func Open115Token(c *gin.Context) {
	client := sdk.New()
	var req TokenReq
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	var codeToTokenResp *sdk.CodeToTokenResp
	codeToTokenResp, err := client.CodeToToken(c, req.Uid, req.CodeVerifier)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"resp": codeToTokenResp})
}
