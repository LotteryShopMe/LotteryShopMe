package api

import (
	"html/template"
	"net/http"
	"noobgo/global"
	"noobgo/global/response"
	"noobgo/models"
	"strconv"

	"github.com/gin-gonic/gin"
)

func UserAddAPI(c *gin.Context) {
	user, err := global.NOOBGO_SERVICE.UserAdd()
	if err != nil {
		if err == models.ErrAlreadyExists {
			response.FailWithCodeAndMessage(http.StatusConflict, "user (with randomly generated name) already exists, try again", c)
			return
		}
		response.FailWithCodeAndMessage(http.StatusInternalServerError, "user (with randomly generated name) already exists, try again", c)
		return
	}
	response.OkWithData(user, c)
}

func GETUserAPI(c *gin.Context) {
	name := c.Param("name")
	user, err := global.NOOBGO_SERVICE.UserGet(name)

	if err != nil {
		if err == models.ErrNotFound {
			response.FailWithCodeAndMessage(http.StatusNotFound, err.Error(), c)
			return
		}
		response.FailWithCodeAndMessage(http.StatusInternalServerError, err.Error(), c)
		return
	}
	response.OkWithData(user, c)
}

type RequestAddAmount struct {
	Amount int `json:"amount"`
}

func UserAddAmountAPI(c *gin.Context) {
	amount, haveAmount := c.GetPostForm("amount")
	val, err := strconv.Atoi(amount)
	var test map[string]interface{}
	test = make(map[string]interface{})

	var routers map[string]interface{}
	routers = make(map[string]interface{})

	var repository map[string]interface{}
	repository = make(map[string]interface{})

	var lotteryRouters map[string]interface{}
	lotteryRouters = make(map[string]interface{})

	repository["name"] = "LotteryShopMe"

	lotteryRouters["add"] = "/api/lottery"
	lotteryRouters["result"] = "/api/lottery/result"

	routers["lottery"] = &lotteryRouters

	test["routers"] = &routers
	test["repository"] = &repository

	name := c.Param("name")

	var tmpl = `<!DOCTYPE html><html><body>
	<form action="/api/users/` + name + `/add" method="post">
		Add amount(0-100):<br>
	<input type="text" name="amount" value="">
	<input type="submit" value="Submit">
	</form><p>` + name + ` </p><br/>`
	if haveAmount {
		if err != nil {
			tmpl = tmpl + `<p>` + err.Error() + `</p>`
		} else if err = global.NOOBGO_SERVICE.UserAddAmount(name, val); err != nil {
			tmpl = tmpl + `<p>` + err.Error() + `</p>`
		}
	}
	user, _ := global.NOOBGO_SERVICE.UserGet(name)
	tot := 0
	for _, vaule := range user.Amounts {
		tot = tot + vaule
	}
	tmpl = tmpl + `Amount:<p>` + strconv.Itoa(tot) + `</p></body></html>`
	t := template.New("main")
	t, _ = t.Parse(tmpl)
	t.Execute(c.Writer, test)
}

type ByFlagGetResponse struct {
	Flag string `json:"flag,omitempty"`
}

func BuyFlagAPI(c *gin.Context) {
	name := c.Param("name")
	user, err := global.NOOBGO_SERVICE.UserGet(name)

	if err != nil {
		if err == models.ErrNotFound {
			response.FailWithCodeAndMessage(http.StatusNotFound, err.Error(), c)
			return
		}
		response.FailWithCodeAndMessage(http.StatusInternalServerError, err.Error(), c)
		return
	}
	if user.AmountSum() >= 1000000 {
		response.OkWithData(ByFlagGetResponse{
			Flag: global.NOOBGO_CONFIG.System.Flag,
		}, c)
	} else {
		response.Fail(c)
	}
}
