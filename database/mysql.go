package database

import (
	"database/sql"
	"io/ioutil"
	"log"

	yaml "gopkg.in/yaml.v2"
)

var Db *sql.DB

type conf struct {
	DBUsername string `yaml:"db_username"`
	DBPassword string `yaml:"db_password"`
	DBHost     string `yaml:"db_host"`
	DBName     string `yaml:"db_name"`
}

func initMysql() {
	var err error
	//打开数据库
	Config := conf{}
	yamlFile, err := ioutil.ReadFile("config.yaml")
	if err != nil {
		log.Printf("yamlFile.Get err   #%v ", err)
	}
	err = yaml.Unmarshal(yamlFile, &Config)
	if err != nil {
		log.Fatalf("Unmarshal: %v", err)
	}
	//account:password@tcp(host:3306)/dbname?charset=utf8
	var dbURL = Config.DBUsername + ":" + Config.DBPassword + "@tcp(" + Config.DBHost + ":3306)/" + Config.DBName + "?charset=utf8"
	Db, err = sql.Open("mysql", dbURL)
	log.Printf("db_url:%s", dbURL)
	if err != nil {
		log.Fatalf("Open database error: %s\n", err)
	}

	//连接数据库
	err = Db.Ping()
	if err != nil {
		log.Fatal(err)
	}
}
