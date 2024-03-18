package main

import (
	"bufio"
	"fmt"
	"log/slog"
	"os"
	"sync"
	"time"

	"github.com/SaimonWoidig/ctc-gas-station/pkg/config"
	"github.com/SaimonWoidig/ctc-gas-station/pkg/station"
	"github.com/SaimonWoidig/ctc-gas-station/pkg/utils"
)

const LogsFileName = "logs.txt"
const LoggingEnabled = true

func main() {
	if LoggingEnabled {

		file, err := os.OpenFile(LogsFileName, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0600)
		if err != nil {
			panic(fmt.Sprintf("failed to create log file %q: %s", LogsFileName, err.Error()))
		}
		defer file.Close()

		buf := bufio.NewWriter(file)
		defer buf.Flush()
		slog.SetDefault(slog.New(slog.NewJSONHandler(buf, nil)))
	} else {
		slog.SetDefault(slog.New(utils.NewNoopHandler()))
	}

	slog.Info("starting")
	defer func(start time.Time) {
		took := time.Since(start)
		slog.Info("done", "durMs", took.Milliseconds())
		fmt.Printf("done in %v\n", took)
	}(time.Now())

	c, err := config.NewConfig()
	if err != nil {
		panic(fmt.Sprintf("failed to load config: %s", err.Error()))
	}

	gasStation := station.NewStation(c.Stations.Gas.Count, c.Stations.Gas.ServeTimeMin, c.Stations.Gas.ServeTimeMax, "gas")
	dieselStation := station.NewStation(c.Stations.Diesel.Count, c.Stations.Diesel.ServeTimeMin, c.Stations.Diesel.ServeTimeMax, "diesel")
	lpgStation := station.NewStation(c.Stations.Lpg.Count, c.Stations.Lpg.ServeTimeMin, c.Stations.Lpg.ServeTimeMax, "lpg")
	electricStation := station.NewStation(c.Stations.Electric.Count, c.Stations.Electric.ServeTimeMin, c.Stations.Electric.ServeTimeMax, "electric")
	register := station.NewStation(c.Registers.Count, c.Registers.HandleTimeMin, c.Registers.HandleTimeMax, "register")

	var wg sync.WaitGroup
	for i := 1; i <= c.Cars.Count; i++ {
		arriveTime := utils.RandDurationBetween(c.Cars.ArrivalTimeMin, c.Cars.ArrivalTimeMax)
		time.Sleep(arriveTime)
		wg.Add(1)
		go func(num int) {
			// car arrives
			defer wg.Done()
			slog.Info("car arrived", "num", num, "durMs", arriveTime.Milliseconds())
			// serve fuel to car
			var serveTime time.Duration
			for {
				var st *station.Station
				switch {
				case !gasStation.IsFull():
					st = gasStation
				case !dieselStation.IsFull():
					st = dieselStation
				case !lpgStation.IsFull():
					st = lpgStation
				case !electricStation.IsFull():
					st = electricStation
				default:
					continue
				}
				slog.Info("car serving", "num", num, "station", st.FuelType)
				serveTime, err = st.Serve()
				if err != nil {
					slog.Error("failed to serve car", "num", num, "err", err)
				}
				break
			}
			slog.Info("car served", "num", num, "durMs", serveTime.Milliseconds())
			// car pays
			var handleTime time.Duration
			for {
				var st *station.Station
				switch {
				case !register.IsFull():
					st = register
				default:
					continue
				}
				slog.Info("car paying", "num", num, "station", st.FuelType)
				handleTime, err = st.Serve()
				if err != nil {
					slog.Error("failed to handle car", "num", num, "err", err)
				}
				break
			}
			slog.Info("car paid", "num", num, "durMs", handleTime.Milliseconds())
		}(i)
	}

	wg.Wait()
}
