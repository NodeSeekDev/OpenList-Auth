package common

import (
	"github.com/gin-gonic/gin"
)

func JsonBytes(c *gin.Context, jsonBytes []byte) error {
	c.Header("Content-Type", "application/json; charset=utf-8")
	c.Header("Access-Control-Allow-Origin", "*")
	c.Header("Access-Control-Allow-Methods", "POST, OPTIONS")
	c.Header("Access-Control-Allow-Headers", "Content-Type, Authorization")
	c.Writer.WriteHeaderNow()
	if _, err := c.Writer.Write(jsonBytes); err != nil {
		return err
	}
	return nil
}
