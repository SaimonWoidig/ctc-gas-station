package output

import "time"

type SimulationOutput struct {
	Stations  StationsOutput  `yaml:"stations"`
	Registers RegistersOutput `yaml:"registers"`
}

type StationsOutput struct {
	Gas      StationOutput `yaml:"gas"`
	Diesel   StationOutput `yaml:"diesel"`
	Lpg      StationOutput `yaml:"lpg"`
	Electric StationOutput `yaml:"electric"`
}

type StationOutput struct {
	TotalCars        int           `yaml:"total_cars"`
	TotalTime        time.Duration `yaml:"total_time"`
	AverageQueueTime time.Duration `yaml:"avg_queue_time"`
	MaxQueueTime     time.Duration `yaml:"max_queue_time"`
}

type RegistersOutput struct {
	TotalCars        int           `yaml:"total_cars"`
	TotalTime        time.Duration `yaml:"total_time"`
	AverageQueueTime time.Duration `yaml:"avg_queue_time"`
	MaxQueueTime     time.Duration `yaml:"max_queue_time"`
}
