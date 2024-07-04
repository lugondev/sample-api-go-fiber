package main

import (
	"fmt"
	"os"
	api "sample/api/cmd/app"
	"sample/api/config"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "api",
	Short: "API service",
}

var envFile string // config file

func init() {
	cobra.OnInitialize(initConfig)

	defaultEnvFile := config.GetDefaultEnvPath()
	rootCmd.PersistentFlags().StringVar(
		&envFile,
		"config", "",
		fmt.Sprintf("config file (default is %s)", defaultEnvFile),
	)
	rootCmd.AddCommand(api.Cmd)
}

func initConfig() {
	if envFile != "" {
		config.LoadEnvFile(envFile)
	}
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
