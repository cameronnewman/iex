package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/cameronnewman/iex"
)

func main() {

	iexClient := iex.NewIEX(&http.Client{
		Timeout: time.Second * 10,
	})

	fmt.Println(iexClient.GetDEEP("APPL"))

}
