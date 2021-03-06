package server

import (
	"context"
	"net/http"
	"os"
	"time"

	"github.com/Selahattinn/todo-app-back/pkg/api"
	"github.com/Selahattinn/todo-app-back/pkg/repository"
	"github.com/Selahattinn/todo-app-back/pkg/service"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
)

type Config struct {
	ListenAddress string `yaml:"host"`
	CertFile      string `yaml:"cert_file"`
	KeyFile       string `yaml:"key_file"`

	API     *api.Config             `yaml:"api"`
	Service *service.Config         `yaml:"service"`
	DB      *repository.MySQLConfig `yaml:"database"`
}

// Instance represents an instance of the server
type Instance struct {
	Config     *Config
	API        *api.API
	Service    service.Service
	httpServer *http.Server
}

// NewInstance returns an new instance of our server
func NewInstance(cfg *Config) *Instance {
	return &Instance{
		Config: cfg,
	}
}

// Start starts the server
func (i *Instance) Start() {
	var err error
	var router = mux.NewRouter()

	// Establish database connection
	repo, err := repository.NewMySQLRepository(i.Config.DB)
	if err != nil {
		logrus.WithError(err).Fatal("Could not create mysql repository")
	}

	i.Service, err = service.NewProvider(i.Config.Service, repo)
	if err != nil {
		logrus.WithError(err).Fatal("Could not create service provider")
	}

	// Initialize API
	i.API, err = api.New(i.Config.API, router, i.Service)
	if err != nil {
		logrus.WithError(err).Fatal("Could not create API instance")
	}
	// Startup the HTTP Server in a way that we can gracefully shut it down again
	i.httpServer = &http.Server{
		Addr:    i.Config.ListenAddress,
		Handler: router,
	}

	err = i.httpServer.ListenAndServe()
	if err != http.ErrServerClosed {
		logrus.WithError(err).Error("HTTP Server stopped unexpected")
		i.Shutdown()
	} else {
		logrus.WithError(err).Info("HTTP Server stopped")
	}
}

// Shutdown stops the server
func (i *Instance) Shutdown() {
	// Shutdown all dependencies
	//i.DB.CloseConnection()

	// Shutdown HTTP server
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	err := i.httpServer.Shutdown(ctx)
	if err != nil {
		logrus.WithError(err).Error("Failed to shutdown HTTP server gracefully")
		os.Exit(1)
	}

	logrus.Info("Shutdown HTTP server...")
	os.Exit(0)
}
