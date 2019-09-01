package mdservice

import (
	"fmt"

	"short/fw"
)

type Service struct {
	name   string
	server fw.Server
	logger fw.Logger
}

func (s Service) Start(port int) {
	defer s.logger.Info(fmt.Sprintf("%s started", s.name))

	go func() {
		err := s.server.ListenAndServe(port)

		if err != nil {
			s.logger.Error(err)
		}
	}()
}

func (s Service) StartAndWait(port int) {
	s.Start(port)
	select {}
}

func (s Service) Stop() {
	defer s.logger.Info(fmt.Sprintf("%s stopped", s.name))

	err := s.server.Shutdown()
	if err != nil {
		s.logger.Error(err)
	}
}

func New(name string, server fw.Server, logger fw.Logger) Service {
	return Service{
		name:   name,
		server: server,
		logger: logger,
	}
}
