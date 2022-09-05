package calipower

import (
	"encoding/json"
	"net/http"
)

type ISOData struct {
	CurrentNetDemandWithoutPumpTS               string   `json:"CurrentNetDemandWithoutPumpTS"`
	SlotDate                                    string   `json:"slotDate"`
	RenewablesPercentwoBatteries                int      `json:"renewablesPercentwoBatteries"`
	HistDemand                                  string   `json:"histDemand"`
	CurrentSystemDemandWithoutPumpTS            string   `json:"CurrentSystemDemandWithoutPumpTS"`
	UnloadedCapacityDeratedMax                  int      `json:"Unloaded_capacity_derated_max"`
	HistDemandDate                              string   `json:"histDemandDate"`
	RequiredReserve                             int      `json:"Required_reserve"`
	CurrentSystemDemandTS                       string   `json:"CurrentSystemDemandTS"`
	TodaysPeakDemandTS                          string   `json:"TodaysPeakDemandTS"`
	UnloadedCapacity4H                          int      `json:"Unloaded_capacity_4H"`
	UnloadedCapacity10M                         int      `json:"Unloaded_capacity_10M"`
	CurrentRenewables                           int      `json:"currentRenewables"`
	Limitation                                  string   `json:"Limitation"`
	PSVERSION                                   string   `json:"PS_VERSION"`
	RenewablesPercent                           int      `json:"renewablesPercent"`
	CurrentAdjustedUnloadedGenerationCapacityTS string   `json:"CurrentAdjustedUnloadedGenerationCapacityTS"`
	TodayForecastPeakDemandTS                   string   `json:"todayForecastPeakDemandTS"`
	CurrentCo2Intensity                         float64  `json:"currentCo2intensity"`
	TodayForecastPeakDemand                     int      `json:"todayForecastPeakDemand"`
	CurrentSolar                                int      `json:"currentSolar"`
	CurrentSystemDemand                         int      `json:"CurrentSystemDemand"`
	CurrentCo2                                  int      `json:"currentCo2"`
	CurrentAdjustedUnloadedGenerationCapacity   int      `json:"CurrentAdjustedUnloadedGenerationCapacity"`
	OutageDTS                                   string   `json:"OutageDTS"`
	HistoryFirstDate                            string   `json:"HistoryFirstDate"`
	AvailableCapacity                           int      `json:"Available_capacity"`
	CurrentNetDemandWithoutPump                 int      `json:"CurrentNetDemandWithoutPump"`
	CurrentSystemDemandPlusUnloaded4H           int      `json:"CurrentSystemDemandPlusUnloaded4h"`
	Shutdown                                    string   `json:"Shutdown"`
	UnloadedCapacity1H                          int      `json:"Unloaded_capacity_1H"`
	TodaysPeakDemand                            int      `json:"TodaysPeakDemand"`
	CurrentWind                                 int      `json:"currentWind"`
	TomorrowsForecastPeakDemand                 float64  `json:"tomorrowsForecastPeakDemand"`
	CurrentNetDemandTS                          string   `json:"CurrentNetDemandTS"`
	CurrentNetDemand                            int      `json:"CurrentNetDemand"`
	CurrentRenewableswoBatteries                int      `json:"currentRenewableswoBatteries"`
	Gridstatus                                  []string `json:"gridstatus"`
	CurrentReserve                              int      `json:"Current_reserve"`
	CurrentSystemDemandWithoutPump              int      `json:"CurrentSystemDemandWithoutPump"`
}

func GetPowerData() (*ISOData, error) {
	var err error
	data := &ISOData{}

	// Fetch JSON data from FlexAlert.org
	res, err := http.Get("https://wwwmobile.caiso.com/outlook/SP/stats.txt")
	if err != nil {
		return data, err
	}

	decoder := json.NewDecoder(res.Body)
	err = decoder.Decode(data)
	if err != nil {
		return data, err
	}
	return data, nil
}

func GetCurDemand() (int, error) {
	data, err := GetPowerData()
	if err != nil {
		return 0, err
	}
	return data.CurrentSystemDemand, nil
}
func GetCurReserves() (int, error) {
	data, err := GetPowerData()
	if err != nil {
		return 0, err
	}
	return data.CurrentReserve, nil
}
func GetCurCapacity() (int, error) {
	data, err := GetPowerData()
	if err != nil {
		return 0, err
	}
	return data.CurrentSystemDemandPlusUnloaded4H, nil
}
func GetCurAvailability1H() (int, error) {
	data, err := GetPowerData()
	if err != nil {
		return 0, err
	}
	return data.UnloadedCapacity1H, nil
}
func GetCurAvailability4H() (int, error) {
	data, err := GetPowerData()
	if err != nil {
		return 0, err
	}
	return data.UnloadedCapacity4H, nil
}
func GetCurAvailability() (string, error) {
	data, err := GetPowerData()
	if err != nil {
		return "", err
	}
	return string(data.UnloadedCapacity1H) + " — " + string(data.UnloadedCapacity4H), nil
}
