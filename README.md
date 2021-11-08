# SQL Builder

[![PkgGoDev](https://pkg.go.dev/badge/gitlab.com/osamai/go-sqlbuilder)](https://pkg.go.dev/gitlab.com/osamai/go-sqlbuilder)
[![Go Report Card](https://goreportcard.com/badge/gitlab.com/osamai/go-sqlbuilder)](https://goreportcard.com/report/gitlab.com/osamai/go-sqlbuilder)

> Simple utilities to make working with SQL queries easier and more readable.

## Install

```sh
go get -u gitlab.com/osamai/go-sqlbuilder
```

## Example

```go
package main

import (
	"fmt"

	"gitlab.com/osamai/go-sqlbuilder"
)

func main() {
	q := sqlbuilder.NewQuery("users").Select("name", "email").Where("id = ?", 1)

	fmt.Println(q.String()) // "SELECT name,email FROM users WHERE id = $1"
	fmt.Println(q.Args())   // [1]
}
```
