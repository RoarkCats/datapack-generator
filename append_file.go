package main

import "os"

func AppendFile(file string, text string) {

	f, err := os.OpenFile(file, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	_, err = f.Write([]byte(text))
	if err != nil {
		panic(err)
	}
}
