package main

import (
	"fmt"
	"os"

	"github.com/kimonides/bittorent-client/torrentfile"
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

	t, err := bto.ToTorrentFile()
	if err != nil {
		panic(err)
	}
	fmt.Println(t)
}
