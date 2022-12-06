package main

import (
	"context"
	"database/sql"
	"net/http"

	"github.com/go-chi/chi"
	_ "github.com/go-sql-driver/mysql"
	"github.com/johnkgs/imersao11-consolidation/internal/infra/db"
	httpHandler "github.com/johnkgs/imersao11-consolidation/internal/infra/http"
	"github.com/johnkgs/imersao11-consolidation/internal/infra/repository"
	uow "github.com/johnkgs/imersao11-consolidation/pkg"
)

func main() {
	ctx := context.Background()
	dtb, err := sql.Open("mysql", "root:root@tcp(localhost:3306)/cartola?parseTime=true")
	if err != nil {
		panic(err)
	}
	defer dtb.Close()
	uow, err := uow.NewUow(ctx, dtb)
	if err != nil {
		panic(err)
	}
	registerRepositories(uow)

	router := chi.NewRouter()
	router.Get("/players", httpHandler.ListPlayersHandler(ctx, *db.New(dtb)))
	router.Get("/my-teams/{teamID}/players", httpHandler.ListMyTeamPlayersHandler(ctx, *db.New(dtb)))
	router.Get("/my-teams/{teamID}/balance", httpHandler.GetMyTeamBalanceHandler(ctx, *db.New(dtb)))
	router.Get("/matches", httpHandler.ListMatchesHandler(ctx, repository.NewMatchRepository(dtb)))
	router.Get("/matches/{matchID}", httpHandler.ListMatchByIDHandler(ctx, repository.NewMatchRepository(dtb)))

	if err = http.ListenAndServe(":8080", router); err != nil {
		panic(err)
	}
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
