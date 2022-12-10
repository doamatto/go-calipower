package main

import (
	"fmt"
	"log"
	"math"
	"strconv"
	"strings"

	"github.com/doamatto/go-calipower"
)

func main () {
	data, err := calipower.GetPowerData() 
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("\033[0m\033[34mCurrent demand: \033[0m%s\n", strconv.Itoa(data.CurrentSystemDemand))
	fmt.Printf("\033[0m\033[31mCurrent reserves: \033[0m%s\n", strconv.Itoa(data.CurrentReserve))
	fmt.Printf("Current capacity: %s\n", strconv.Itoa(data.CurrentSystemDemandPlusUnloaded4H))
	// Get percentages of capacity
	demandPercent := (float64(data.CurrentSystemDemand) / float64(data.CurrentSystemDemandPlusUnloaded4H)) * 100
	reservesPercent := (float64(data.CurrentReserve) / float64(data.CurrentSystemDemandPlusUnloaded4H)) * 100
	remainderPercent := float64(100)-(float64(demandPercent) + float64(reservesPercent))
	// Draw capacity
	demandBars := strings.Repeat("|", int(math.Round(demandPercent/5)))
	remainBars := strings.Repeat("|", int(math.Round(remainderPercent/5)))
	reserveBars := strings.Repeat("|", int(math.Round(reservesPercent/5)))
	fmt.Printf("\033[0m[\033[34m%s\033[37m%s\033[31m%s\033[0m]\n", demandBars, remainBars, reserveBars)

	fmt.Printf("Predicted availability in next hour: %s\n", strconv.Itoa(data.UnloadedCapacity1H))
	// Generate estimate percentages
	hrPercent := (float64(data.CurrentSystemDemandPlusUnloaded4H) - float64(data.UnloadedCapacity1H)) / float64(data.CurrentSystemDemandPlusUnloaded4H) * 100
	hrUsed := 100-hrPercent
	hrPercentBars := strings.Repeat("|", int(math.Round(hrPercent/5)))
	hrUsedBars := strings.Repeat("|", int(math.Round(hrUsed/5)))
	// Draw estimate
	fmt.Printf("\033[0m[\033[34m%s\033[37m%s\033[0m]\n", hrUsedBars, hrPercentBars)

	fmt.Printf("Predicted availability in four hours: %s\n", strconv.Itoa(data.UnloadedCapacity4H))
	// Generate estimate percentages
	four_hrPercent := (float64(data.CurrentSystemDemandPlusUnloaded4H) - float64(data.UnloadedCapacity4H)) / float64(data.CurrentSystemDemandPlusUnloaded4H) * 100
	four_hrUsed := 100-four_hrPercent
	four_hrPercentBars := strings.Repeat("|", int(math.Round(four_hrPercent/5)))
	four_hrUsedBars := strings.Repeat("|", int(math.Round(four_hrUsed/5)))
	// Draw estimate
	fmt.Printf("\033[0m[\033[34m%s\033[37m%s\033[0m]\n", four_hrUsedBars, four_hrPercentBars)

	fmt.Printf("Percentage of renewables: %s\n", strconv.Itoa(data.RenewablesPercent))
	// Generate percentages
	dirtyPercent := (float64(100)-float64(data.RenewablesPercent))
	renewBars := strings.Repeat("|", int(math.Round(float64(data.RenewablesPercent)/5)))
	dirtyBars := strings.Repeat("|", int(math.Round(float64(dirtyPercent)/5)))
	// Draw renewable percent
	fmt.Printf("\033[0m[\033[32m%s\033[37m%s\033[0m]\n", renewBars, dirtyBars)
}
