package cmd

import (
	"fmt"

	"github.com/thanos-go/config"
	"github.com/thanos-go/interactor/account"
	"github.com/thanos-go/interactor/validation"
	"github.com/thanos-go/log"

	"github.com/thanos-go/interactor/authentication"
	"github.com/thanos-go/store/migrator"

	"github.com/labstack/echo-contrib/prometheus"
	handler2 "github.com/thanos-go/delivery/http/handler"
	"github.com/thanos-go/store/mysqlrepo"

	"github.com/spf13/cobra"
	mw "github.com/thanos-go/delivery/http/middleware"
	"github.com/thanos-go/delivery/http/server"
)

func registerServeCmd(root *cobra.Command) {

	applyMigrations := true
	listenPost := config.Get().App.Port
	cmd := &cobra.Command{
		Use:        "serve",
		Short:      fmt.Sprintf("Start listening to incoming requests on port %v", listenPost),
		SuggestFor: []string{"run", "start"},
		Run: func(cmd *cobra.Command, args []string) {
			runServeCmd(applyMigrations, listenPost)
		},
	}
	cmd.Flags().BoolVarP(&applyMigrations, "migrate", "m", false, "apply new migrations, if exists")
	cmd.Flags().IntVarP(&listenPost, "port", "p", listenPost, "set the listening port")

	root.AddCommand(cmd)
}

func runServeCmd(applyMigrations bool, listenPort int) {

	cfg := config.Get()

	log.Named("serve").Debug("Trying to connect to mysql")
	mysqlRepo := mysqlrepo.New(&cfg.Mysql)
	defer mysqlRepo.Close()

	if applyMigrations {
		if err := migrator.New(mysqlRepo).Run("up", 0); err != nil {
			fmt.Println("plannedItems")
			log.Named("serve").Fatal("migration error: %v", err)
		}
	}

	validationService := validation.NewValidation(*cfg)
	authenticationService := authentication.New(&cfg.Authentication)
	accountService := account.New(*cfg, mysqlRepo, authenticationService, validationService)

	log.Named("serve").Debug("Starting stream manager")

	router := server.NewRouter(cfg)

	// add the prometheus
	prometheus.NewPrometheus("Thanos", nil).Use(router)

	// check the services status
	router.Use(mw.HealthCheck(mysqlRepo))

	handler := handler2.NewHandlers(&accountService, authenticationService, *cfg, mysqlRepo)

	// registering the routes
	server.RegisterRoutes(router, handler)

	log.Named("serve").Debug("starting http server...")
	server.NewServer(router, listenPort, cfg.App.GracefulShutdown).StartListening()

}
