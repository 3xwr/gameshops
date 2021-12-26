package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"gameservice/internal/client"
	"gameservice/internal/handler"
	"gameservice/internal/repository"

	"github.com/go-chi/chi"
	"github.com/rs/zerolog"
)

func main() {
	cache := repository.NewCache(10 * time.Minute)
	c := client.New(cache)

	logger := zerolog.New(os.Stdout).With().Timestamp().Logger()
	h := handler.New(&logger, c)

	r := chi.NewRouter()

	r.Route("/", func(r chi.Router) {
		r.Get("/steamprice", h.SteamPriceHandler)
		r.Get("/steampayprice", h.SteamPayPriceHandler)
		r.Get("/gogprice", h.GOGPriceHandler)
		r.Get("/platiruprice", h.PlatiruHandler)
		r.Get("/compareprice", h.CompareHandler)
		r.Get("/getname", h.GetNameHandler)
	})

	srv := http.Server{
		Addr:    fmt.Sprintf(":%s", "8080"),
		Handler: r,
	}

	shutdown := make(chan os.Signal, 1)
	signal.Notify(shutdown, os.Interrupt, syscall.SIGTERM)
	defer signal.Stop(shutdown)

	go func() {
		logger.Info().Msgf("Server is listening on :%d", "8080")
		err := srv.ListenAndServe()
		if err != nil && err != http.ErrServerClosed {
			logger.Fatal().Err(err).Msg("Server error")
		}
	}()

	<-shutdown

	logger.Info().Msg("Shutdown signal received")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Minute)
	defer func() {
		cancel()
	}()

	if err := srv.Shutdown(ctx); err != nil {
		logger.Fatal().Err(err).Msg("Server shutdown error")
	}

	logger.Info().Msg("Server stopped gracefully")

}
