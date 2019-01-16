package tools

import (
	"github.com/kataras/iris"
	"time"
	"crypto/md5"
	"strconv"
	"fmt"
)

type ToolsSignController struct {
	Ctx iris.Context
}

func (c *ToolsSignController) GetSign(){

	t := time.Now()
	sign := make(map[string]interface{})

	sign["timestamp"] = t.Unix()
	md5byte := []byte("BFCB1CFB14A9_" + strconv.FormatInt(t.Unix(),10))
	sign["sign"] = fmt.Sprintf("%x", md5.Sum(md5byte))
	result := make(map[string]interface{})
	result["status_code"]	= "ok"
	result["data"] = sign
	c.Ctx.JSON(result)
}