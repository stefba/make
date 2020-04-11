package main

import (
	"github.com/otiai10/copy"
	"log"
	"os"
)

func copyToLive() error {
	err := checkCreate("./build/live")
	if err != nil {
		return err
	}
	err = copy.Copy("../site/public", "./build/pre")
	if err != nil {
		return err
	}
	err = os.Rename("./build/live", "./build/bot")
	if err != nil {
		return err
	}
	err = os.Rename("./build/pre", "./build/live")
	if err != nil {
		return err
	}
	log.Println("Build made live.")
	err = os.RemoveAll("./build/bot")
	if err != nil {
		return err
	}
	return nil
}


func checkCreate(path string) error {
	_, err := os.Stat(path)
	if err == nil {
		return nil
	}
	return os.MkdirAll(path, 0775)
}
