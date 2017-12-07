package cmd

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/ddspog/doni/experiment"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	dir     string
	verbose bool
)

// startCmd represents the start command
var startCmd = &cobra.Command{
	Use:   "start <experiment> <outputfile>",
	Short: "Executes a experiment following a configuration file.",
	Long: `Reads a configuration file, and use the information to setup
a experiment. Configuration file must be in a YAML format.`,
	Args: cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Reading experiment configuration...")
		viper.SetConfigName("Expfile")
		viper.SetConfigType("yml")
		viper.AddConfigPath(".")

		if err := viper.ReadInConfig(); err != nil {
			panic(fmt.Errorf("problem on config file: %s", err))
		}

		e := experiment.Executor{
			Dir:     dir,
			Verbose: verbose,
			Output:  args[1],

			Context: getSignalContext(),

			Stdin:  os.Stdin,
			Stdout: os.Stdout,
			Stderr: os.Stderr,
		}

		if err := e.ParseExpFile(); err != nil {
			panic(fmt.Errorf("problem parsing config file: %s", err))
		}

		_ = e.Run(args[0])
	},
}

func init() {
	rootCmd.AddCommand(startCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	startCmd.Flags().BoolVarP(&verbose, "verbose", "v", false, "Enables verbose mode")
	startCmd.Flags().StringVarP(&dir, "dir", "d", "", "Sets directory of execution")
}

func getSignalContext() context.Context {
	ch := make(chan os.Signal, 1)
	signal.Notify(ch, os.Interrupt, os.Kill, syscall.SIGTERM)
	ctx, cancel := context.WithCancel(context.Background())
	go func() {
		sig := <-ch
		log.Printf("doni: signal received: %s", sig)
		cancel()
	}()
	return ctx
}
