package main

import (
	"fmt"

	"github.com/libp2p/go-libp2p"
)

func main() {
	// To construct a simple host with all the default settings, just use `New`
	h, err := libp2p.New()
	if err != nil {
		panic(err)
	}
	defer h.Close()

	fmt.Printf("Hello world, my hosts Id is %s\n", h.ID())
	fmt.Printf("Hello world, my addresses are %s\n", h.Addrs())
}