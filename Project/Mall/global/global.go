package global

import (
	"github.com/jinzhu/gorm"
	oplogging "github.com/op/go-logging"
	"github.com/spf13/viper"
)

type Server struct {
	Mysql  Mysql  `mapstructure:"mysql" json:"mysql"`
	System System `mapstructure:"system" json:"system"`
	Log    Log    `mapstructure:"log" json:"log"`
}
type Mysql struct {
	Username     string `mapstructure:"username" json:"username"`
	Password     string `mapstructure:"password" json:"password"`
	Path         string `mapstructure:"path" json:"path"`
	Dbname       string `mapstructure:"db-name" json:"dbname"`
	Config       string `mapstructure:"config" json:"config"`
	MaxIdleConns int    `mapstructure:"max-idle-conns" json:"maxIdleConns"`
	MaxOpenConns int    `mapstructure:"max-open-conns" json:"maxOpenConns"`
	LogMode      bool   `mapstructure:"log-mode" json:"logMode"`
}
type System struct {
	UseMultipoint bool   `mapstructure:"use-multipoint" json:"useMultipoint"`
	Env           string `mapstructure:"env" json:"env"`
	Addr          int    `mapstructure:"addr" json:"addr"`
	DbType        string `mapstructure:"db-type" json:"dbType"`
}
type Log struct {
	Prefix  string `mapstructure:"prefix" json:"prefix"`
	LogFile bool   `mapstructure:"log-file" json:"logFile"`
	Stdout  string `mapstructure:"stdout" json:"stdout"`
	File    string `mapstructure:"file" json:"file"`
}

var G_DB *gorm.DB
var G_VP *viper.Viper
var G_CONFIG Server
var G_LOG *oplogging.Logger
