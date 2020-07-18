package main

import (
	"github.com/AlexRipoll/enchante_technical_interview/config/mysql"
	"github.com/AlexRipoll/enchante_technical_interview/handler/rest"
)

func main() {
	mysql.InitSession()
	rest.Handler()
}