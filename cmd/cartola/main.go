package main

import (
	"context"
	"database/sql"
	"net/http"

	"github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
	_ "github.com/go-sql-driver/mysql"
	"github.com/johnkgs/imersao11-consolidation/internal/infra/config/envs"
	"github.com/johnkgs/imersao11-consolidation/internal/infra/db"
	httpHandler "github.com/johnkgs/imersao11-consolidation/internal/infra/http"
	"github.com/johnkgs/imersao11-consolidation/internal/infra/kafka/consumer"
	"github.com/johnkgs/imersao11-consolidation/internal/infra/repository"
	uow "github.com/johnkgs/imersao11-consolidation/pkg"
)

func main() {
	envs.SetupEnvs()

	ctx := context.Background()
	dtb, err := sql.Open(envs.GetEnvs().DB.DriverName, envs.GetEnvs().DB.SourceName)
	if err != nil {
		panic(err)
	}
	defer dtb.Close()
	uow, err := uow.NewUow(ctx, dtb)
	if err != nil {
		panic(err)
	}

	registerRepositories(uow)
	initServer(ctx, dtb)
	initKafka(ctx, uow)
}

func initServer(ctx context.Context, dtb *sql.DB) {
	router := chi.NewRouter()
	router.Use(cors.Handler(cors.Options{
		// AllowedOrigins:   []string{"https://foo.com"}, // Use this to allow specific origin hosts
		AllowedOrigins: []string{"https://*", "http://*"},
		// AllowOriginFunc:  func(r *http.Request, origin string) bool { return true },
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300, // Maximum value not ignored by any of major browsers
	}))
	router.Get("/players", httpHandler.ListPlayersHandler(ctx, *db.New(dtb)))
	router.Get("/my-teams/{teamID}/players", httpHandler.ListMyTeamPlayersHandler(ctx, *db.New(dtb)))
	router.Get("/my-teams/{teamID}/balance", httpHandler.GetMyTeamBalanceHandler(ctx, *db.New(dtb)))
	router.Get("/matches", httpHandler.ListMatchesHandler(ctx, repository.NewMatchRepository(dtb)))
	router.Get("/matches/{matchID}", httpHandler.ListMatchByIDHandler(ctx, repository.NewMatchRepository(dtb)))

	go http.ListenAndServe(":8080", router)
}

func initKafka(ctx context.Context, uow *uow.Uow) {
	var topics = []string{
		"newMatch",
		"chooseTeam",
		"newPlayer",
		"matchUpdateResult",
		"newAction",
	}

	msgChannel := make(chan *kafka.Message)
	go consumer.Consume(topics, msgChannel)
	consumer.ProcessEvents(ctx, msgChannel, uow)
}

func registerRepositories(uow *uow.Uow) {
	uow.Register("PlayerRepository", func(tx *sql.Tx) interface{} {
		repo := repository.NewPlayerRepository(uow.Db)
		repo.Queries = db.New(tx)
		return repo
	})

	uow.Register("MatchRepository", func(tx *sql.Tx) interface{} {
		repo := repository.NewMatchRepository(uow.Db)
		repo.Queries = db.New(tx)
		return repo
	})

	uow.Register("TeamRepository", func(tx *sql.Tx) interface{} {
		repo := repository.NewTeamRepository(uow.Db)
		repo.Queries = db.New(tx)
		return repo
	})

	uow.Register("MyTeamRepository", func(tx *sql.Tx) interface{} {
		repo := repository.NewMyTeamRepository(uow.Db)
		repo.Queries = db.New(tx)
		return repo
	})
}
