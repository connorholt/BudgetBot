package app

import (
	"context"
	"github.com/MaxKut3/BudgetBot/config"
	"github.com/MaxKut3/BudgetBot/internal/controller"
	"github.com/MaxKut3/BudgetBot/internal/usecase"
	"github.com/MaxKut3/BudgetBot/pkg/cache"
	"log"
	"os"
	"os/signal"
	"syscall"
)

func Run(cfg *config.Config) {

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// 1. все завязано на интерфейсы буква D в SOLID
	// 2. легко поменять логику внутри хендлера и легко покрыть тестами
	// 3. в будущем кеш можно поменять на redis и добавить базу
	// db := pgx.Init(...)
	// repository := repository.NewBudget(db)
	useCaseCurrency := usecase.NewCurrencyConverter(ctx, cfg)
	useCaseBudget := usecase.NewBudget(
		ctx,
		useCaseCurrency,
		cache.NewSimple(),
	)

	// 4. создаем клиента для бота
	client := controller.NewTelegramClient(ctx, cfg, useCaseBudget)
	go client.Run()

	// 5. можем создать любого другого еще клиента
	// httpServer := http.Server()
	// router := new Router(httpServer)
	// router.Get("/list", useCaseBudget.List)
	// router.Post("/new", useCaseBudget.New)

	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, syscall.SIGINT, syscall.SIGTERM)
	sig := <-signalChan

	log.Printf("%s signal caught", sig)
	client.Close()
}
