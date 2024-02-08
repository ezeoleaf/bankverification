# Bankverification
The bankverification package validates swedish bank accounts.

A Go port of [getkickback/bankverify](https://github.com/getkickback/bankverify)

### Install

This package requires Go modules.

```shell
go get github.com/ezeoleaf/bankverification
```

### Usage

#### Personal accounts

```go
// Valid bank account.
be := bankverification.NewBankAccount("9420 - 4172385")
fmt.Println(be.IsValid())
// true
fmt.Println(be.ValidationErrors())
// []
fmt.Println(be.GetBank())
// Forex Bank
fmt.Println(be.GetClearingNumber())
// 9420

// Invalid bank account.
be = bankverification.NewBankAccount("8032540213939")
fmt.Println(be.IsValid())
// false
fmt.Println(be.ValidationErrors())
// [invalid checksum]
fmt.Println(be.GetBank())
// Swedbank
fmt.Println(be.GetClearingNumber())
// 8032
```

#### Bankgiro
In progress

#### Plusgiro
In progress

### Documentation 
[![Go Reference](https://pkg.go.dev/badge/github.com/ezeoleaf/bankverification.svg)](https://pkg.go.dev/github.com/ezeoleaf/bankverification)

Full `go doc` style documentation for the package can be viewed online without
installing this package by using the GoDoc site here: 
http://pkg.go.dev/github.com/ezeoleaf/bankverification
