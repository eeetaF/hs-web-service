package main

import "hs-web-service/internal/service"

func main() {
	var mySvc service.HearthstoneSvc = service.New()

	err := mySvc.GetPlayerByAccountID()
	if err != nil {
		panic(err)
	}
}
