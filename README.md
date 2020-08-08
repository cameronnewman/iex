# iex

[![Build][1]][2]
[![GoDoc][3]][4]
[![Go Report Card][5]][6]
[![codecov][7]][8]

[1]: https://github.com/cameronnewman/iex/workflows/Continuous%20Integration/badge.svg
[2]: https://github.com/cameronnewman/iex/actions
[3]: https://godoc.org/github.com/cameronnewman/iex?status.svg
[4]: https://godoc.org/github.com/cameronnewman/iex
[5]: https://goreportcard.com/badge/github.com/cameronnewman/iex
[6]: https://goreportcard.com/report/github.com/cameronnewman/iex
[7]: https://codecov.io/gh/cameronnewman/iex/branch/master/graph/badge.svg
[8]: https://codecov.io/gh/cameronnewman/iex

## Purpose

iex is a simple client library for iextrading.com API.

## Usage

```golang
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
