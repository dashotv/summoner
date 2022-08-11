package server

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"github.com/robfig/cron/v3"
	"github.com/sirupsen/logrus"

	"github.com/dashotv/summoner/application"
	"github.com/dashotv/summoner/config"
)

var cfg *config.Config
var app *application.App

type Server struct {
	Router *gin.Engine
	Log    *logrus.Entry
}

func New() (*Server, error) {
	cfg = config.Instance()
	app = application.Instance()
	log := app.Log.WithField("prefix", "server")
	s := &Server{Log: log, Router: app.Router}

	return s, nil
}

func (s *Server) Start() error {
	s.Log.Info("starting summoner...")

	s.Routes()

	c := cron.New(cron.WithSeconds())
	if _, err := c.AddFunc("*/30 * * * * *", s.process); err != nil {
		return errors.Wrap(err, "adding cron function")
	}

	go func() {
		s.Log.Info("starting cron...")
		c.Start()
	}()

	s.Log.Info("starting web...")
	if err := s.Router.Run(fmt.Sprintf(":%d", cfg.Port)); err != nil {
		return errors.Wrap(err, "starting router")
	}

	return nil
}

func (s *Server) process() {
	s.Log.Info("processing downloads")
	// create new downloads
	// search / load matching torrents
	// manage torrents
	// download torrents
	// move torrents
}
