package main

import (
	"fmt"
	"log"
	"os"

	"github.com/Averagejoestudent/Gator/internal/config"
)

func main() {
    cfg, err := config.Read()
    if err != nil {
    log.Fatal(err)
}
    st := &state{cfg: &cfg}    
    myvar := commands{variable: make(map[string]func(*state, command) error)}
    myvar.register("login",handlerLogin)
    if len(os.Args) < 2 {
        fmt.Println("not enough arguments")
        os.Exit(1)
    }
    cmdName := os.Args[1]
    cmdArgs := os.Args[2:]
    cmd := command{
        name: cmdName,
        args: cmdArgs,
    }
    if err := myvar.run(st, cmd); err != nil {
    log.Fatal(err)
}

}