package main

import (
	demo "github.com/SchrodingerwithCat/rpcsvr/kitex_gen/demo/studentservice"
	"log"
)


////////////////////
// add global var
// var stuSvr = NewStuService()


func main() {
	////////////////////////////////

	stuSvr := NewStuService()
	stuSvr.InitStuService("foo.db")


	///////////////////////////////

	svr := demo.NewServer(new(StudentServiceImpl))

	err := svr.Run()

	if err != nil {
		log.Println(err.Error())
	}

	

}



