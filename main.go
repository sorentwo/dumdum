package main

import (
	"crypto/rand"
	"flag"
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

var (
	isChild = flag.Bool("child", false, "parent pid")

	SignalHandlers = map[os.Signal]func(){
		syscall.SIGWINCH: restart,
		syscall.SIGHUP:   reload,
		syscall.SIGTTIN:  increase,
		syscall.SIGTTOU:  decrease,
		syscall.SIGUSR1:  bloat,
		syscall.SIGUSR2:  work,
	}

	Garbage = make([]byte, 0)
)

func main() {
	flag.Parse()

	fmt.Println("Dumdum Started")

	if !*isChild {
		handleParentSignals()
	} else {
		handleChildSignals()

		ppid := os.Getppid()
		fmt.Println("ppid", ppid)
	}
}

func handleChildSignals() {
	for {
		// Don't do anything yet
	}
}

func handleParentSignals() {
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

	os.StartProcess(os.Args[0], append(os.Args, "-child"), &os.ProcAttr{})
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
