package main

import (
	"errors"
	"fmt"
	"log"
	"os"
)

const (
	usage = `
usage:
	create [pkgName] [version]

`
)

var (
	err error

	gopath  string
	argsLen int
)

func init() {
	gopath = get_gopath()
}

func argsLengthCheck(n int) error {
	if argsLen < n {
		fmt.Print(usage)
		return errors.New("not enough arguments")
	}
	return nil
}

func main() {
	argsLen = len(os.Args)
	if argsLengthCheck(2) != nil {
		return
	}
	cmd := os.Args[1]
	switch cmd {
	case "create":
		if argsLengthCheck(4) != nil {
			return
		}
		if err = create(os.Args[2], os.Args[3]); err != nil {
			log.Fatal(err)
			return
		}
	default:
		fmt.Println("unknown command")
		fmt.Print(usage)
	}
}

func create(pkgName string, version string) error {
	dir := fmt.Sprintf("%s\\src\\custom\\%s\\%s\\%s", gopath, pkgName, version, pkgName)
	if err := os.MkdirAll(dir, os.ModePerm); err != nil {
		return err
	}
	return nil
}
