package main

import (
	"fmt"
	"os"
)

func main() {
	reader, err := os.Open("./debian.torrent")
	if err != nil {
		panic(err)
	}
	bto, err := torrentfile.Open(reader)
	if err != nil {
		panic(err)
	}

	t, err := bto.toTorrentFile()
	if err != nil {
		panic(err)
	}
	fmt.Println(t)
}
