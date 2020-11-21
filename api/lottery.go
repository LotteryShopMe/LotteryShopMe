package api

import (
	"crypto/md5"
	"io"
	"net/http"
	"noobgo/global"
	"noobgo/global/response"
	"noobgo/models"

	"github.com/gin-gonic/gin"
)

type RequestLotteryAdd struct {
	Name string `json:"account"`
}

func LotteryAddAPI(c *gin.Context) {
	var R RequestLotteryAdd
	c.ShouldBindJSON(&R)
	if R.Name == "" {
		response.FailWithCodeAndMessage(http.StatusUnprocessableEntity, "'account' must be filled", c)
		return
	}
	if err := global.NOOBGO_SERVICE.LotteryAdd(R.Name); err != nil {
		if err == models.ErrNotFound {
			response.FailWithCodeAndMessage(http.StatusNotFound, err.Error(), c)
			return
		}
		response.FailWithCodeAndMessage(http.StatusInternalServerError, err.Error(), c)
		return
	}
	response.Ok(c)
}

type ResponseLotteryResults struct {
	Winners []string `json:"winners"`
}

func LotteryResultsAPI(c *gin.Context) {
	winners := global.NOOBGO_SERVICE.LotteryResults()
	resp := ResponseLotteryResults{
		Winners: make([]string, 0, 0),
	}
	for _, winner := range winners {
		h := md5.New()
		if _, err := io.WriteString(h, winner); err != nil {
			response.FailWithCodeAndMessage(http.StatusInternalServerError, err.Error(), c)
			return
		}
		h.Sum(nil)
		resp.Winners = append(resp.Winners, string(h.Sum(nil)))
	}
	response.OkWithData(resp, c)
}
