package global

import (
	"noobgo/config"

	"noobgo/service"

	oplogging "github.com/op/go-logging"
	"github.com/spf13/viper"
)

var (
	NOOBGO_CONFIG  config.Server
	NOOBGO_VP      *viper.Viper
	NOOBGO_LOG     *oplogging.Logger
	NOOBGO_SERVICE *service.Service
)
