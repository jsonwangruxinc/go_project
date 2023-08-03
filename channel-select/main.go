package main

import (
	_case "go_project/channel-select/case"
	"os"
	"os/signal"
)

func main() {
	//_case.Communication()
	//_case.ConcurrentSync()
	_case.NoticeAndMultiplexing()
	ch := make(chan os.Signal, 0)
	signal.Notify(ch, os.Interrupt, os.Kill)
	<-ch
}
