# iex

[![GoDoc][1]][2]
[![GoCard][3]][4]
[![Build][5]][6]

[1]: https://godoc.org/github.com/cameronnewman/iex?status.svg
[2]: https://godoc.org/github.com/cameronnewman/iex
[3]: https://goreportcard.com/badge/github.com/cameronnewman/iex
[4]: https://goreportcard.com/report/github.com/cameronnewman/iex
[5]: https://github.com/cameronnewman/iex/workflows/pipeline/badge.svg
[6]: https://github.com/cameronnewman/iex/actions


## Purpose ##

iex is a simple client library for iextrading.com API.


## Usage

```
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
```

## Issues
* None

## License
MIT Licensed