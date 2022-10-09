package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"os"

	"github.com/Igusaya/igusaya_blog/api/config"
)

func main() {
	if err := run(context.Background()); err != nil {
		log.Printf("failed to terminate server: %v", err)
		os.Exit(1)
	}
}

func run(ctx context.Context) error {
	cfg, err := config.New()
	if err != nil {
		return err
	}
	l, err := net.Listen("tcp", fmt.Sprintf(":%d", cfg.Port))
	if err != nil {
		log.Fatalf("failed to listen port %d: %v", cfg.Port, err)
	}
	url := fmt.Sprintf("http://%s", l.Addr().String())
	log.Printf("start with: %v", url)
	mux, cleanup, err := NewRouter(ctx, cfg)
	if err != nil {
		return err
	}
	defer cleanup()
	s := NewServer(l, mux)
	return s.Run(ctx)
}
