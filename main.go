package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/Averagejoestudent/Gator/internal/config"
	"github.com/Averagejoestudent/Gator/internal/database"

	_ "github.com/lib/pq"
)

func main() {
	cfg, err := config.Read()
	if err != nil {
		log.Fatal(err)
	}
	db, err := sql.Open("postgres", cfg.DbURL)
	if err != nil {
		log.Fatal(err)
	}
	dbQueries := database.New(db)

	st := &state{db: dbQueries, cfg: &cfg}
	myvar := commands{registeredCommands: make(map[string]func(*state, command) error)}
	myvar.register("login", handlerLogin)
	myvar.register("register", handlerRegister)
	myvar.register("reset", handlerReset)
	myvar.register("users", handlerListUsers)
	myvar.register("agg", handlerAgg)
	myvar.register("addfeed", handlerAddFeed)
	myvar.register("feeds", handler_Feed)
	if len(os.Args) < 2 {
		fmt.Println("not enough arguments")
		os.Exit(1)
	}
	cmdName := os.Args[1]
	cmdArgs := os.Args[2:]
	cmd := command{
		Name: cmdName,
		Args: cmdArgs,
	}
	if err := myvar.run(st, cmd); err != nil {
		log.Fatal(err)
	}

}
