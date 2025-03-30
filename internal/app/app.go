package app

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"stakeway/config"
	"stakeway/internal/blockchain"
	"stakeway/internal/handler"
	"stakeway/internal/service"
	"stakeway/internal/store"
	"stakeway/internal/store/pg"
	"stakeway/pkg/logger"
	"syscall"
)

const depositDataFile = "deposit_data.json"

type App struct {
	cfg config.Config
	log *logger.Logger
}

func New(cfg config.Config, log *logger.Logger) *App {
	return &App{
		cfg: cfg,
		log: log,
	}
}

func (app *App) Run() error {
	ctx, cancel := context.WithCancel(context.Background())

	repo, err := pg.NewRepository("database.db")
	if err != nil {
		app.log.Errorf("failed initialize db: %s", err)
		return err
	}

	store.ApplyMigrations(repo.DB.DB)

	s := service.New(ctx, app.cfg, repo, app.log)
	h := handler.New(s)

	if app.cfg.Blockchain.Enabled {
		go func() {
			app.log.Info("Starting blockchain integration...")

			privateKeyHex := app.cfg.Blockchain.PrivateKey
			log.Printf("PRIVATE_KEY: %v", privateKeyHex)

			txHash, err := blockchain.ExecuteDepositTransaction(
				"https://ethereum-holesky.publicnode.com",
				privateKeyHex,
				depositDataFile,
			)

			if err != nil {
				app.log.Errorf("Blockchain error: %v", err)
			} else {
				app.log.Infof("Transaction successful! Hash: %s", txHash)
			}
		}()
	}

	go func() {
		exit := make(chan os.Signal, 1)
		signal.Notify(exit, os.Interrupt, syscall.SIGTERM, syscall.SIGINT)
		<-exit
		h.Shutdown()
		cancel()
	}()

	return h.Listen(fmt.Sprintf(":%s", app.cfg.App.Port))
}
