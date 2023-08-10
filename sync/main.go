package main

import (
	"context"
	_case "go_project/sync/case"
	"os"
	"os/signal"
)

func main() {
	//_case.AtomicCase()
	//_case.AtomicCase1()
	//_case.AtomicCase2()
	//_case.AtomicCase3()
	//_case.MutexCase()
	//_case.MapCase()
	//_case.MapCase1()
	//_case.PoolCase()
	//_case.OnceCase()
	//_case.WaitGroupCase()
	_case.WaitGroupCase1()
	ctx, stop := signal.NotifyContext(context.Background(), os.Kill, os.Interrupt)
	defer stop()
	<-ctx.Done()
}
