package core

import (
	"context"
	"fmt"
	"net/http"
	"noobgo/global"
	"noobgo/initialize"
	"noobgo/service"
	"os"
	"time"
)

var (
	DefaultLotteryPeriod      = time.Minute * 1
	DefaultAccountRemoveAfter = time.Minute * 10
)

func RunServer() {
	Router := initialize.Routers()
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	lotteryPeriod := DefaultLotteryPeriod
	if d, err := time.ParseDuration(os.Getenv("LOTTERY_PERIOD")); err == nil {
		lotteryPeriod = d
	}
	deleteAccountAfter := DefaultAccountRemoveAfter
	if d, err := time.ParseDuration(os.Getenv("DELETE_ACCOUNT_AFTER")); err == nil {
		deleteAccountAfter = d
	}

	// Application service.
	global.NOOBGO_SERVICE = service.NewService(ctx, lotteryPeriod, deleteAccountAfter)
	address := fmt.Sprintf(":%d", global.NOOBGO_CONFIG.System.Addr)
	s := &http.Server{
		Addr:           address,
		Handler:        Router,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	// 保证文本顺序输出
	// In order to ensure that the text order output can be deleted
	time.Sleep(10 * time.Microsecond)
	global.NOOBGO_LOG.Debug("server run success on ", address)

	global.NOOBGO_LOG.Error(s.ListenAndServe())
}
