package cmd

import (
	"fmt"
	"strings"

	"github.com/thanos-go/config"
	"github.com/thanos-go/log"
	"go.uber.org/zap/zapcore"

	"github.com/spf13/cobra"
	_ "github.com/thanos-go/log" // important for capturing the std logger
	"github.com/thanos-go/pkg/logger/zap_logger"
	_ "go.uber.org/automaxprocs" // prevent cpu throttling
)

func Execute() {

	var rootCmd = &cobra.Command{
		Use:   "Thanos",
		Short: "the Thanos",
		Long:  `the Thanos`,
		PersistentPreRun: func(cmd *cobra.Command, args []string) {
			boot()
		},
	}

	registerMigrateCmd(rootCmd)
	registerServeCmd(rootCmd)
	if err := rootCmd.Execute(); err != nil {
		log.Fatal("root cmd execution", err)
	}
}

func boot() {
	// getting the config
	cfg := config.Get()

	// update the global log handler
	zapper, zapCloser := zap_logger.New(config.Get().IsProduction(), cfg.ZapLogger, zapcore.InfoLevel)
	defer zapCloser()
	log.SetAdapter(zapper)
}

func descriptionGenerator(actions map[string]string) (string, string, []string) {
	size := len(actions)
	keys := make([]string, 0)

	description := fmt.Sprintf("there are %v actions that you can use:\n", size)
	if size == 1 {
		description = "there is only 1 action that you can use:\n"
	}

	for cmd, desc := range actions {
		description += fmt.Sprintf("- %s: %s \n", cmd, desc)
		keys = append(keys, cmd)
	}

	msg := fmt.Sprintf("the action must be one of: %s", strings.Join(keys, ", "))
	if size == 1 {
		msg = fmt.Sprintf("the action must be: %s", strings.Join(keys, ", "))
	}

	return description, msg, keys
}
