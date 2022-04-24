package cmd

import (
	"encoding/json"
	"log"
	"math"
	"os"
	"strconv"
	"strings"

	"github.com/petershen0307/kdWatchDog/models"
	"github.com/spf13/cobra"
)

type dataPrice struct {
	Year  uint16
	Month uint16
	Day   uint16
	Price models.Price
}

func init() {
	var inputFilePath string
	var outputFilePath string
	var analyzeCmd = &cobra.Command{
		Use:   "analyze",
		Short: "analyze stock prices",
		Args: func(cmd *cobra.Command, args []string) error {
			return nil
		},
		Run: func(cmd *cobra.Command, args []string) {
			data, err := os.ReadFile(inputFilePath)
			if err != nil {
				log.Fatalf("Open data(%v) with error=%v\n", inputFilePath, err)
			}
			prices := make(map[string]models.Price)
			json.Unmarshal(data, &prices)
			dataMap := make(map[uint16]map[uint16][]float64)
			for date, price := range prices {
				d := strings.Split(date, "-")
				year, _ := strconv.Atoi(d[0])
				month, _ := strconv.Atoi(d[1])
				date, _ := strconv.Atoi(d[2])
				if _, exist := dataMap[uint16(year)]; !exist {
					dataMap[uint16(year)] = make(map[uint16][]float64)
				}
				if _, exist := dataMap[uint16(year)][uint16(month)]; !exist {
					dataMap[uint16(year)][uint16(month)] = make([]float64, 32)
				}
				dataMap[uint16(year)][uint16(month)][date], _ = strconv.ParseFloat(price.Close, 64)
			}
			result := make([][]int, 13)
			for _, yv := range dataMap {
				for month, mv := range yv {
					if result[month] == nil {
						result[month] = make([]int, 32)
					}
					min := math.MaxFloat64
					theDate := 0
					for date, dv := range mv {
						if min > dv && dv != 0 {
							min = dv
							theDate = date
						}
					}
					result[month][theDate]++
				}
			}
			for month, mv := range result {
				max := 0
				theDate := 0
				for date, dv := range mv {
					if max < dv && dv != 0 {
						max = dv
						theDate = date
					}
				}
				log.Printf("m=%v, d=%v\n", month, theDate)
			}
		},
	}
	analyzeCmd.Flags().StringVarP(&inputFilePath, "file", "f", "", "input file")
	analyzeCmd.Flags().StringVarP(&outputFilePath, "output", "o", "out", "output file")
	rootCmd.AddCommand(analyzeCmd)
}
