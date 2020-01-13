package cmd

import (
	"fmt"
	"strconv"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(uptimeCmd)
}

const (
	numberOfArguments = 3
)

var uptimeCmd = &cobra.Command{
	Use:   "uptime <software name> <monitoring duration in seconds> <downtime duration in seconds>",
	Short: "Print the uptime",
	Long:  `All software has uptimes. What's yours?`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("\nCalculating the uptime")
		if len(args) > numberOfArguments || len(args) < numberOfArguments {
			fmt.Println("invalid arguments!\nuptime <software name> <monitoring duration in seconds> <downtime duration in seconds>")
			return
		}

		monitoringDurationInSeconds, err := strconv.ParseFloat(args[1], 64)
		if err != nil {
			fmt.Println("Inavlid monitoring duration!")
			return
		}

		downtimeDurationInSeconds, err := strconv.ParseFloat(args[2], 64)
		if err != nil {
			fmt.Println("Inavlid monitoring duration!")
			return
		}

		downtimePercentage := (float64(downtimeDurationInSeconds) * 100) / float64(monitoringDurationInSeconds)

		uptimePercentage := float64(100) - downtimePercentage

		fmt.Printf("%s uptime is = %.3f\n", args[0], uptimePercentage)
	},
}
