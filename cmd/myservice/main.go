package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"os/signal"
	"syscall"

	"github.com/Shin-Thant/service-test/internal"
	"github.com/Shin-Thant/service-test/internal/application"
	"github.com/Shin-Thant/service-test/internal/config"
	"github.com/Shin-Thant/service-test/internal/pb/feed/v1"
	"google.golang.org/grpc"
)

func main() {
	cfg := config.Load()

	application.New(&cfg)

	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGTERM, syscall.SIGQUIT, syscall.SIGINT)
	defer stop()

	grpcServer, listener, err := setupGRPCServer(internal.NewFeedServiceHandler(), cfg.Port)
	if err != nil {
		log.Fatalln(err)
	}

	go func() {
		log.Println("LOG: GRPC server listening on port", cfg.Port)
		err := grpcServer.Serve(listener)
		if err != nil {
			log.Println("ERROR: grpc server serve error:", err)
		}
	}()

	<-ctx.Done()
	log.Println("GRPC server stopped.")
	grpcServer.GracefulStop()

	log.Println("GRPC listener closed.")
	listener.Close()

	log.Println("Service shutdown successfully.")
}

func setupGRPCServer(feedServiceServer feed.FeedServiceServer, port int) (*grpc.Server, net.Listener, error) {
	grpcServer := grpc.NewServer()
	feed.RegisterFeedServiceServer(grpcServer, feedServiceServer)

	listener, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", port))
	if err != nil {
		return nil, nil, fmt.Errorf("couldn't start listener on port %d: %v", port, err)
	}

	return grpcServer, listener, nil
}
