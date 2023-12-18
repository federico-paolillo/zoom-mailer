package main

import (
	"os"
	"zoo-mailer/internal/zoo"

	"github.com/spf13/viper"
)

const ZOO_ENV_PREFIX = "ZOO__"

const ZOO_DIR_CFG_KEY = "folder"
const ZOO_DIR_CFG_DEFAULT = "/var/zoo/files"

const ZOO_MAIL_TMPL_CFG_KEY = "mail_tmpl"
const ZOO_MAIL_TMPL_CFG_DEFAULT = "/etc/zoo/templates/mail.tmpl"

const ZOO_SENDLIST_FILE_CFG_KEY = "sendlist"
const ZOO_SENDLIST_FILE_CFG_DEFAULT = "/etc/zoo/send.list"

func main() {

	viper.SetEnvPrefix(ZOO_ENV_PREFIX)
	viper.AllowEmptyEnv(false)
	viper.AutomaticEnv()

	viper.SetDefault(ZOO_DIR_CFG_KEY, ZOO_DIR_CFG_DEFAULT)
	viper.SetDefault(ZOO_MAIL_TMPL_CFG_KEY, ZOO_MAIL_TMPL_CFG_DEFAULT)
	viper.SetDefault(ZOO_SENDLIST_FILE_CFG_KEY, ZOO_SENDLIST_FILE_CFG_DEFAULT)

	zooDirPath := viper.GetString(ZOO_DIR_CFG_KEY)
	emailTemplateFilePath := viper.GetString(ZOO_MAIL_TMPL_CFG_KEY)
	sendlistFilePath := viper.GetString(ZOO_SENDLIST_FILE_CFG_KEY)

	zooConfig := zoo.NewZooAppConfig(
		zooDirPath,
		emailTemplateFilePath,
		sendlistFilePath,
	)

	statusCode := zoo.Run(zooConfig)

	os.Exit(statusCode)
}
