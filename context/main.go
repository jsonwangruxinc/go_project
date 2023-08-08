package main

import (
	"context"
	_case "go_project/context/case"
	"os"
	"os/signal"
)

func main() {
	//_case.CloseCoroutineCase()
	_case.ContextCase()
	ctx, stop := signal.NotifyContext(context.Background(), os.Kill, os.Interrupt)
	defer stop()
	<-ctx.Done()
}
