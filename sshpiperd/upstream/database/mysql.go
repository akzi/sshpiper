package database

import (
	"fmt"

	mysqldriver "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql" // gorm dialiect

	upstreamprovider "github.com/tg123/sshpiper/sshpiperd/upstream"
)

type mysqlplugin struct {
	plugin

	Config struct {
		Host     string `long:"upstream-mysql-host" default:"127.0.0.1" description:"MySQL host" env:"SSHPIPERD_UPSTREAM_MYSQL_HOST" ini-name:"upstream-mysql-host"`
		User     string `long:"upstream-mysql-user" default:"root" description:"MySQL user" env:"SSHPIPERD_UPSTREAM_MYSQL_USER" ini-name:"upstream-mysql-user"`
		Password string `long:"upstream-mysql-password" default:"" description:"MySQL password" env:"SSHPIPERD_UPSTREAM_MYSQL_PASSWORD" ini-name:"upstream-mysql-password"`
		Port     uint   `long:"upstream-mysql-port" default:"3306" description:"MySQL port" env:"SSHPIPERD_UPSTREAM_MYSQL_PORT" ini-name:"upstream-mysql-port"`
		Dbname   string `long:"upstream-mysql-dbname" default:"sshpiper" description:"MySQL database name" env:"SSHPIPERD_UPSTREAM_MYSQL_DBNAME" ini-name:"upstream-mysql-dbname"`
	}
}

func (p *mysqlplugin) create() (*gorm.DB, error) {

	config := mysqldriver.Config{
		User:      p.Config.User,
		Passwd:    p.Config.Password,
		Net:       "tcp",
		Addr:      fmt.Sprintf("%v:%v", p.Config.Host, p.Config.Port),
		DBName:    p.Config.Dbname,
		ParseTime: true,
	}

	db, err := gorm.Open("mysql", config.FormatDSN())
	if err != nil {
		return nil, err
	}

	return db, nil
}

func (mysqlplugin) GetName() string {
	return "mysql"
}

func (p *mysqlplugin) GetOpts() interface{} {
	return &p.Config
}

func init() {
	p := &mysqlplugin{}
	p.createdb = p
	upstreamprovider.Register("mysql", p)
}
