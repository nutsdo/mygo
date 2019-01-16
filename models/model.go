package models

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

type DatabaseConfig struct {
	Driver DriverConfig `yaml:"driver" json:"driver"`
}

type DriverConfig struct {
	Default string      `yaml:"default" json:"default"`
	Mysql   MysqlConfig `yaml:"mysql" json:"mysql"`
}

type MysqlConfig struct {
	Username string `yaml:"username" json:"username"`
	Password string `yaml:"password" json:"password"`
	Host     string `yaml:"host" json:"host"`
	Port     string `yaml:"port" json:"port"`
	Dbname   string `yaml:"dbname" json:"dbname"`
}

var DB *xorm.Engine

func init() {

	//读取配置文件
	config, rerr := ioutil.ReadFile("configs/database.yaml")

	if rerr != nil {
		fmt.Println(rerr)
	}
	dbconf := new(DatabaseConfig)
	cerr := yaml.Unmarshal(config, dbconf)

	if cerr != nil {
		fmt.Println("数据库配置文件解析错误:")
		fmt.Println(cerr)
	}
	username := dbconf.Driver.Mysql.Username
	password := dbconf.Driver.Mysql.Password
	protocol := "tcp"
	dbport := dbconf.Driver.Mysql.Port
	dbhost := dbconf.Driver.Mysql.Host
	dbname := dbconf.Driver.Mysql.Dbname

	if dbport != "" {
		dbhost = dbconf.Driver.Mysql.Host + ":" + dbport
	} else {
		dbhost = dbconf.Driver.Mysql.Host + ":3306"
	}

	var err error
	DB, err = xorm.NewEngine(dbconf.Driver.Default, fmt.Sprintf("%s:%s@%s(%s)/%s?charset=utf8mb4", username, password, protocol, dbhost, dbname))

	if err != nil {
		fmt.Println(err)
	}
	DB.ShowSQL(true)
}
