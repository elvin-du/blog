package main

import (
	_ "blog/routers"
	"github.com/astaxie/beego"
	"fmt"
)

func main() {
	fmt.Println("listening")	
	beego.Run()
}

