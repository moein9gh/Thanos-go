package cmd

import (
	"errors"
	"strconv"

	"github.com/thanos-go/config"
	"github.com/thanos-go/log"
	"github.com/thoas/go-funk"

	"github.com/thanos-go/store/mysqlrepo"

	"github.com/thanos-go/store/migrator"

	"github.com/spf13/cobra"
)

func registerMigrateCmd(root *cobra.Command) {

	longDescription, validationMsg, actionKeys := descriptionGenerator(map[string]string{
		"up":     "for applying new migrations",
		"down":   "for rolling back applied migrations",
		"status": "for checking the status of applied & un-applied migrations",
		"fresh":  "drops all the tables and applies all of the migrations",
		"new":    "prints the current date as the prefix of the new migration name",
	})

	cmd := &cobra.Command{
		Use:        "migrate [action] [count]",
		Short:      "Provide database migration functionalities",
		Long:       longDescription,
		SuggestFor: []string{"migration"},
		Args: func(cmd *cobra.Command, args []string) error {
			if len(args) == 2 {
				if !funk.ContainsString(actionKeys, args[0]) {
					return errors.New(validationMsg)
				}
				if args[0] == "fresh" && config.Get().IsProduction() {
					return errors.New("fresh migration is disabled in production environment")
				}

				if _, err := strconv.Atoi(args[1]); err != nil {
					return errors.New("migration count is not valid integer")
				}
			} else {
				return errors.New("you must pass only 2 argument")
			}
			return nil
		},
		Run: func(cmd *cobra.Command, args []string) {
			runMigrateCmd(args[0], args[1])
		},
	}

	root.AddCommand(cmd)
}

func runMigrateCmd(action, count string) {

	mysqlRepo := mysqlrepo.New(&config.Get().Mysql)
	defer mysqlRepo.Close()

	// we ensured that args[1] is valid integer
	countInt, _ := strconv.Atoi(count)

	if err := migrator.New(mysqlRepo).Run(action, countInt); err != nil {
		log.Fatal(err.Error())
	}
}
