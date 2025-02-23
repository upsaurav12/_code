package main

import "os"

func GetEnv() string {
	return os.Getenv("CONFIG_PATH")
}
