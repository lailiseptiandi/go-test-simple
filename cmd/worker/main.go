package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/lailiseptiandi/go-test-simple/internal/events"
	"github.com/lailiseptiandi/go-test-simple/internal/kafka"
	"golang.org/x/sync/errgroup"
)

func main() {
	ctx, stop := signal.NotifyContext(
		context.Background(),
		os.Interrupt,
		syscall.SIGTERM,
	)
	defer stop()

	broker := os.Getenv("KAFKA_BROKER")

	paymentConsumer := kafka.NewConsumer(
		broker,
		events.TopicPayment,
		"payment-worker",
	)

	orderConsumer := kafka.NewConsumer(
		broker,
		events.TopicOrder,
		"order-worker",
	)

	g, ctx := errgroup.WithContext(ctx)

	g.Go(func() error {
		return paymentConsumer.Start(ctx, events.RouteEvent)
	})

	g.Go(func() error {
		return orderConsumer.Start(ctx, events.RouteEvent)
	})

	if err := g.Wait(); err != nil {
		if ctx.Err() == nil {
			log.Println("Worker error:", err)
		}
	}
}
