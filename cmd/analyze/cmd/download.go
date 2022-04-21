package cmd

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"strings"

	"github.com/petershen0307/kdWatchDog/stock-api/alphavantage"
	"github.com/spf13/cobra"
)

func init() {
	var filePath string
	var downloadCmd = &cobra.Command{
		Use:   "download",
		Short: "download stock prices",
		Args: func(cmd *cobra.Command, args []string) error {
			if len(args) < 1 {
				return errors.New("requires a stock id")
			}
			if len(args) > 1 {
				return errors.New("to many stock id")
			}
			return nil
		},
		Run: func(cmd *cobra.Command, args []string) {
			if filePath == "" {
				filePath = args[0] + ".json"
			}
			api := alphavantage.New("...")
			prices, _ := api.GetDailyPriceEx(strings.ToUpper(args[0]), alphavantage.OutputsizeFull)
			jdata, _ := json.Marshal(prices)
			os.WriteFile(filePath, jdata, 0777)
			fmt.Printf("The stock %v output file is %v\n", args[0], filePath)
		},
	}
	downloadCmd.Flags().StringVarP(&filePath, "file", "f", "", "output file path or name only")
	rootCmd.AddCommand(downloadCmd)
}
