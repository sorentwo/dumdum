package main

import (
	"crypto/rand"
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

var (
	SignalHandlers = map[os.Signal]func(){
		syscall.SIGWINCH: restart,
		syscall.SIGHUP:   reload,
		syscall.SIGTTIN:  increase,
		syscall.SIGTTOU:  decrease,
		syscall.SIGUSR1:  bloat,
		syscall.SIGUSR2:  work,
	}

	Garbage []byte = make([]byte, 0)
)

func main() {
	fmt.Println("Dumdum Started")

	handleSignals()
}

func handleSignals() {
	signals := make(chan os.Signal)

	for key, _ := range SignalHandlers {
		signal.Notify(signals, key)
	}

	for {
		sig := <-signals
		fun := SignalHandlers[sig]
		fun()
	}
}

func reload() {
	fmt.Println("reload")
}

func restart() {
	fmt.Println("restart")
}

func increase() {
	fmt.Println("increase")
}

func decrease() {
	fmt.Println("decrease")
}

// Artificially bloat resident memory through useless allocation
func bloat() {
	fmt.Println("bloating")
	bytes := make([]byte, 100000)
	rand.Read(bytes)
	Garbage = append(Garbage, bytes...)
}

func work() {
	fmt.Println("working")
}
