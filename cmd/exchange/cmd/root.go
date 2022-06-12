package cmd

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"github.com/vstdy/xt_test_project/api"
	"github.com/vstdy/xt_test_project/cmd/exchange/cmd/common"
)

const (
	flagConfigPath              = "config"
	flagLogLevel                = "log_level"
	flagTimeout                 = "timeout"
	flagRunAddress              = "run_address"
	flagDatabaseURI             = "database_uri"
	flagStorageType             = "storage_type"
	envUpdaterTimeout           = "updater_timeout"
	envBtcUsdtRateCheckInterval = "btc_usdt_rate_check_interval"
	envCurRubRateCheckInterval  = "cur_rub_rate_check_interval"
)

// Execute prepares cobra.Command context and executes root cmd.
func Execute() error {
	return newRootCmd().ExecuteContext(common.NewBaseCmdCtx())
}

// newRootCmd creates a new root cmd.
func newRootCmd() *cobra.Command {
	cmd := &cobra.Command{
		PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
			if err := setupLogger(cmd); err != nil {
				return fmt.Errorf("app initialization: %w", err)
			}

			if err := setupConfig(cmd); err != nil {
				return fmt.Errorf("app initialization: %w", err)
			}

			return nil
		},

		RunE: func(cmd *cobra.Command, args []string) error {
			config := common.GetConfigFromCmdCtx(cmd)
			svcCtx, svcCancel := context.WithCancel(context.Background())
			defer svcCancel()

			svc, err := config.BuildService(svcCtx)
			if err != nil {
				return fmt.Errorf("app initialization: service building: %w", err)
			}

			srv := api.NewServer(svc, config.HTTPServer)

			go func() {
				if err = srv.ListenAndServe(); err != http.ErrServerClosed {
					log.Error().Err(err).Msg("HTTP server ListenAndServe")
				}
			}()

			stop := make(chan os.Signal, 1)
			signal.Notify(stop, syscall.SIGINT, syscall.SIGTERM)
			<-stop

			svcCancel()

			shutdownCtx, shutdownCancel := context.WithTimeout(context.Background(), 5*time.Second)
			defer shutdownCancel()

			if err = srv.Shutdown(shutdownCtx); err != nil {
				return fmt.Errorf("shutting shutdown server: %w", err)
			}

			if err = svc.Close(); err != nil {
				log.Error().Err(err).Msg("shutting down service")
			}

			log.Info().Msg("server stopped")

			return nil
		},
	}

	config := common.BuildDefaultConfig()
	cmd.PersistentFlags().String(flagConfigPath, "./config.toml", "Config file path")
	cmd.PersistentFlags().String(flagLogLevel, "info", "Logger level [debug,info,warn,error,fatal]")
	cmd.PersistentFlags().Duration(flagTimeout, config.Timeout, "Request timeout")
	cmd.PersistentFlags().StringP(flagDatabaseURI, "d", config.PSQLStorage.URI, "Database URI")
	cmd.Flags().StringP(flagRunAddress, "a", config.HTTPServer.RunAddress, "Run address")
	cmd.Flags().StringP(flagStorageType, "s", config.StorageType, "Storage type [psql]")

	cmd.AddCommand(newMigrateCmd())

	return cmd
}

// setupLogger configures global logger.
func setupLogger(cmd *cobra.Command) error {
	if err := viper.BindPFlag(flagLogLevel, cmd.Flag(flagLogLevel)); err != nil {
		return fmt.Errorf("%s flag binding: %w", flagLogLevel, err)
	}
	if err := viper.BindEnv(flagLogLevel); err != nil {
		return fmt.Errorf("%s env binding: %w", flagLogLevel, err)
	}
	logLevelBz := viper.GetString(flagLogLevel)
	logLevel, err := zerolog.ParseLevel(logLevelBz)
	if err != nil {
		return fmt.Errorf("%s flag parsing: %w", flagLogLevel, err)
	}

	logWriter := zerolog.ConsoleWriter{
		Out:        os.Stdout,
		TimeFormat: time.RFC3339,
	}
	log.Logger = log.Output(logWriter).Level(logLevel)

	return nil
}

// setupConfig reads app config and stores it to cobra.Command context.
func setupConfig(cmd *cobra.Command) error {
	if err := viper.BindPFlags(cmd.Flags()); err != nil {
		return fmt.Errorf("flags binding: %w", err)
	}

	if err := viper.BindEnv(envUpdaterTimeout); err != nil {
		return fmt.Errorf("%s env binding: %w", envUpdaterTimeout, err)
	}
	if err := viper.BindEnv(envBtcUsdtRateCheckInterval); err != nil {
		return fmt.Errorf("%s env binding: %w", envBtcUsdtRateCheckInterval, err)
	}
	if err := viper.BindEnv(envCurRubRateCheckInterval); err != nil {
		return fmt.Errorf("%s env binding: %w", envCurRubRateCheckInterval, err)
	}

	configPath := viper.GetString(flagConfigPath)
	viper.SetConfigFile(configPath)
	if err := viper.ReadInConfig(); err != nil {
		log.Warn().Err(err).Msg("reading config file")
	}

	viper.AutomaticEnv()
	for _, key := range viper.AllKeys() {
		val := viper.Get(key)
		viper.Set(key, val)
	}

	config := common.BuildDefaultConfig()
	if err := viper.Unmarshal(&config); err != nil {
		return fmt.Errorf("config unmarshal: %w", err)
	}

	common.SetConfigToCmdCtx(cmd, config)

	return nil
}
