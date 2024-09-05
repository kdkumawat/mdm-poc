package agent

import (
	"log"

	"golang.org/x/sys/windows/svc"
)

const (
	serviceBasePath = "http://localhost:3040/"
)

type Service struct{}

func (m *Service) Execute(args []string, r <-chan svc.ChangeRequest, s chan<- svc.Status) (svcSpecificEC bool, exitCode uint32) {
	const cmdsAccepted = svc.AcceptStop | svc.AcceptShutdown

	log.Println("Starting service")
	s <- svc.Status{State: svc.StartPending}

	s <- svc.Status{State: svc.Running, Accepts: cmdsAccepted}
	err := applyPolicies()
	if err != nil {
		log.Fatalln("error appliying policies", err)
		return false, 1
	}

loop:
	for c := range r {
		switch c.Cmd {
		case svc.Stop, svc.Shutdown:
			break loop
		default:
		}
	}

	log.Println("Stopping service")
	s <- svc.Status{State: svc.StopPending}

	return
}
