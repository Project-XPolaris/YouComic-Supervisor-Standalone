package setup

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"net/http"
)

var HomeHandler gin.HandlerFunc = func(c *gin.Context) {
	c.HTML(http.StatusOK, "index.tmpl", gin.H{})
}

type ConfigForm struct {
	ApiURL string `form:"api_url"`
	Port string `form:"port"`
}


var SubmitConfigHandler gin.HandlerFunc = func(c *gin.Context) {
	var configForm ConfigForm
	err := c.ShouldBind(&configForm)
	if err != nil {
		fmt.Println(err.Error())
		c.Redirect(500, "/error")
	}
	webConfig := viper.New()
	webConfig.AddConfigPath("static")
	webConfig.SetConfigType("json")
	webConfig.SetConfigName("config")
	webConfig.Set("apiURL",configForm.ApiURL)
	err = webConfig.WriteConfig()
	if err != nil {
		fmt.Println(err.Error())
		c.Redirect(500, "/error")
	}
	//server port
	viper.Set("port",configForm.Port)

	//finish
	viper.Set("first",false)
	err = viper.WriteConfig()
	if err != nil { // Handle errors reading the config file
		fmt.Println(err.Error())
		c.Redirect(500, "/error")
	}
	c.Redirect(http.StatusFound,"/success")
}

var ErrorHandler gin.HandlerFunc = func(c *gin.Context) {
	c.HTML(http.StatusOK, "error.tmpl", gin.H{})
}

var SuccessHandler gin.HandlerFunc = func(c *gin.Context) {
	c.HTML(http.StatusOK, "success.tmpl", gin.H{})
}