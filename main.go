package main

import (
	"context"
	"flag"
	"fmt"
	"os"
	"time"

	"github.com/tevino/tcp-shaker"
)

var (
	addr  string
	usage = `
Usage: ./tcp-shaker-cli -addr localhost:80
`
)

func main() {
	flag.StringVar(&addr, "addr", "", "addr e.g. localhost:80")

	flag.Parse()

	if addr == "" {
		fmt.Fprintf(os.Stderr, "%s", usage)
		os.Exit(1)
	}

	c := tcp.NewChecker()

	ctx, stopChecker := context.WithCancel(context.Background())
	defer stopChecker()
	go func() {
		if err := c.CheckingLoop(ctx); err != nil {
			fmt.Fprintf(os.Stderr, "%+v\n", err)
		}
	}()

	<-c.WaitReady()

	timeout := time.Second * 1
	err := c.CheckAddr(addr, timeout)
	switch err {
	case tcp.ErrTimeout:
		fmt.Fprintf(os.Stderr, "Connection timeout", err)
		os.Exit(1)
	case nil:
		fmt.Fprintf(os.Stderr, "Connection OK")
		os.Exit(0)
	default:
		fmt.Fprintf(os.Stderr, "%+v", err)
		os.Exit(1)
	}
}
