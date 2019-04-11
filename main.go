package main

import "GoTravis/api"

/**
 * :=  created by:  Shuza
 * :=  create date:  11-Apr-2019
 * :=  (C) CopyRight Shuza
 * :=  www.shuza.ninja
 * :=  shuza.sa@gmail.com
 * :=  Fun  :  Coffee  :  Code
 **/

func main() {
	r := api.GetRoutes()
	r.Run(":8000")
}
