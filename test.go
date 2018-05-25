package main

import (
	"io/ioutil"
	"log"

	"gopkg.in/src-d/go-git.v4/plumbing/format/packfile"
)

func main() {
	delta, err := ioutil.ReadFile("delta.bin")
	if err != nil {
		log.Fatal(err)
	}

	src, err := ioutil.ReadFile("src.bin")
	if err != nil {
		log.Fatal(err)
	}

	_, err = packfile.PatchDelta(src, delta)
	if err != nil {
		log.Fatal(err)
	}
}
