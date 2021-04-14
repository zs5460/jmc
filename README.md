# jmc

[![Build Status](https://travis-ci.com/zs5460/jmc.svg?branch=master)](https://travis-ci.com/zs5460/jmc)
[![Go Report Card](https://goreportcard.com/badge/github.com/zs5460/jmc)](https://goreportcard.com/report/github.com/zs5460/jmc)
[![codecov](https://codecov.io/gh/zs5460/jmc/branch/master/graph/badge.svg)](https://codecov.io/gh/zs5460/jmc)
[![GoDoc](https://godoc.org/github.com/zs5460/jmc?status.svg)](https://godoc.org/github.com/zs5460/jmc)

jmc is a Go implementation of the encrypted-config-value library.

## Install

```shell
go install github.com/zs5460/jmc/cmd/jmc
```

## Usage

To modify the profile, the values that need to be encrypted are wrapped in "${enc" and "}"

```javascript
//config.json
{
  "password":"123456"
}
```

Change to

```javascript
{
  "password":"${enc:123456}"
}
```

Use the app to encrypt the profile

```shell
jmc config.json
```

You can see that the values in the profile are encrypted

```javascript
{
    "password":"${enc:bPbmRmA8PvGhcby5LEgBuw==}"
}
```

Refer to the jmc package in the program to call the MustLoadconfig method to load and decrypt the profile

```go
package main

import (
    "fmt"

    "github.com/zs5460/jmc"
)

func main() {
    var cfg map[string]string
    jmc.MustLoadConfig("config.json", &cfg)

    fmt.Println(cfg["password"])
    //output: 123456
}
```

In a production environment, you should set the environment variable "JMC_K" to change the encryption key

## License

Released under MIT license, see [LICENSE](LICENSE) for details.
