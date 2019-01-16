package controllers

import (
	"github.com/kataras/iris"
	"fmt"
	"net/http"
			"net/url"
	"io/ioutil"
	"encoding/json"
	"errors"
)

type AuthorizationController struct {
	Ctx iris.Context
}

func (c *AuthorizationController) GetUserAuthorization()  {
	params := c.Ctx.URLParams()

	fmt.Println(params)

	if _, ok := params["code"]; ok {
		fmt.Println(params["code"])
		AccessTokenUrl := "http://localhost:8082/access_token";

		//secret := []byte("abc123")
		//client_secret,_:= bcrypt.GenerateFromPassword(secret,bcrypt.DefaultCost)
		data := url.Values{}
		data.Add("grant_type","authorization_code")
		data.Add("client_id","myawesomeapp")
		data.Add("client_secret", "abc123") //string(client_secret)
		data.Add("redirect_uri","http://iris.local:8090/auth/user/authorization")
		data.Add("code",params["code"])
		resp,err := http.PostForm(AccessTokenUrl, data)

		if err !=nil {
			fmt.Println(err.Error())
		}
		defer resp.Body.Close()
		bodyBytes, _ := ioutil.ReadAll(resp.Body)

		var respMap map[string]interface{}
		err = json.Unmarshal(bodyBytes, &respMap)
		if err != nil {
			err = errors.New(fmt.Sprintf("%s %s", err.Error(), string(bodyBytes)))
		}
		c.Ctx.JSON(respMap)

	}else {

		redirectUrl := "http://localhost:8082/authorize?response_type=code&client_id=myawesomeapp&redirect_uri=http://iris.local:8090/auth/user/authorization&scope=basic&state=sfsd"
		c.Ctx.Redirect(redirectUrl)
	}
}