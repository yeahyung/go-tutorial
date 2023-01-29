package config

import "go-tutorial/gin-tutorial/utils"

// Gin
var (
	GinMode = utils.GetEnvOr("GIN_MODE", "debug")
)

// database
var (
	DsnFormat  = utils.GetEnvOr("DSN_FORMAT", "%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local")
	DbUsername = utils.GetEnvOr("DB_USERNAME", "root")
	DbPassword = utils.GetEnvOr("DB_PASSWORD", "")
	DbHost     = utils.GetEnvOr("DB_HOST", "localhost")
	DbPort     = utils.GetEnvOr("DB_PORT", "3306")
	DbName     = utils.GetEnvOr("DB_NAME", "my_database")
)
