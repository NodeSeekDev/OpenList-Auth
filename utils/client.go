package utils

import (
	"github.com/go-resty/resty/v2"
)

var (
	RestyClient = resty.New()
)
