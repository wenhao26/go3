package main

import (
	"fmt"
)

type DBConfig struct {
	Host     string `json:"host"`
	Port     int    `json:"port"`
	Username string `json:"username"`
	Password string `json:"password"`
}

var SysConfig = struct {
	Name string
	Mode int64
}{"开发者环境", 9978}

type Paging struct {
	Page  int
	Limit int
}

func (p *Paging) setPage(page int) {
	p.Page = page
}

func (p *Paging) setLimit(limit int) {
	p.Limit = limit
}

func main() {
	cnf := DBConfig{}
	cnf.Host = "127.0.0.1"
	cnf.Port = 113011
	cnf.Username = "admin"
	cnf.Password = "admin"

	fmt.Println(cnf, SysConfig)

	paging := Paging{}
	paging.setPage(1)
	paging.setLimit(20)
	fmt.Println(paging)

}
