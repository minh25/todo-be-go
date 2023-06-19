package config

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"net/url"
	"strings"
	"time"
)

// MySQLOption for MySQL options
type MySQLOption struct {
	Key   string `json:"key" mapstructure:"key"`
	Value string `json:"value" mapstructure:"value"`
}

// MySQLConfig for configuring MySQL
type MySQLConfig struct {
	Host         string        `json:"host" mapstructure:"host"`
	Port         uint16        `json:"port" mapstructure:"port"`
	Database     string        `json:"database" mapstructure:"database"`
	Username     string        `json:"username" mapstructure:"username"`
	Password     string        `json:"password" mapstructure:"password"`
	MaxOpenConns int           `json:"max_open_conns" mapstructure:"max_open_conns"`
	MaxIdleConns int           `json:"max_idle_conns" mapstructure:"max_idle_conns"`
	Options      []MySQLOption `json:"options" mapstructure:"options"`
}

func MySQLDefaultConfig() MySQLConfig {
	return MySQLConfig{
		Host:         "localhost",
		Port:         3306,
		Database:     "mysql_db",
		Username:     "1",
		Password:     "1",
		MaxOpenConns: 30,
		MaxIdleConns: 10,
		Options: []MySQLOption{
			{Key: "parseTime", Value: "true"},
		},
	}
}

// DSN returns data source name
func (c MySQLConfig) DSN() string {
	var opts []string
	for _, o := range c.Options {
		key := url.QueryEscape(o.Key)
		value := url.QueryEscape(o.Value)
		opts = append(opts, key+"="+value)
	}
	optStr := strings.Join(opts, "&")
	return fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?%s", c.Username, c.Password, c.Host, c.Port, c.Database, optStr)
}

// MustConnect connects to database using sqlx
func (c MySQLConfig) MustConnect() *sqlx.DB {
	db := sqlx.MustOpen("mysql", c.DSN())

	db.SetMaxOpenConns(c.MaxOpenConns)
	db.SetMaxIdleConns(c.MaxIdleConns)
	db.SetConnMaxIdleTime(4 * time.Hour)
	return db
}
