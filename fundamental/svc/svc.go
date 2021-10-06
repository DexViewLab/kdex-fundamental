package svc

import (
	"os"
	"os/signal"
	"sync"
	"syscall"
)

// SignalHandleFunc signal handle function
type SignalHandleFunc func(os.Signal) bool

var handlers struct {
	sync.Mutex

	m map[os.Signal][]SignalHandleFunc
}

var signalChan = make(chan os.Signal, 1)

// DefaultSignals default signals to register, and the behavior is to terminate the process
var DefaultSignals = []os.Signal{syscall.SIGINT, syscall.SIGTERM}

// Service interface for Run argument
type Service interface {
	Start() error
	Stop() error
}

// Run start the service and handle signals
func Run(svc Service) error {
	if err := svc.Start(); err != nil {
		return err
	}

	handlers.Lock()
	for _, sig := range DefaultSignals {
		if handlers.m[sig] == nil {
			signal.Notify(signalChan, sig)
		}
	}
	handlers.Unlock()

	for {
		s := <-signalChan

		handlers.Lock()

		if hs := handlers.m[s]; hs != nil {
			quit := false

			for _, h := range hs {
				if h(s) {
					quit = true
				}
			}

			if !quit {
				handlers.Unlock()
				continue
			}

		}

		handlers.Unlock()

		break
	}

	return svc.Stop()
}

// Signal Register a signal handler
func Signal(s os.Signal, f SignalHandleFunc) {
	handlers.Lock()
	defer handlers.Unlock()

	if handlers.m == nil {
		handlers.m = make(map[os.Signal][]SignalHandleFunc)
	}

	handlers.m[s] = append(handlers.m[s], f)

	signal.Notify(signalChan, s)
}
