package env

import (
	"log"
	"os/exec"
	"strings"
	"time"
)

type TimedEnvData struct {
	time int64
	data *EnvData
}

type Env struct {
	history []*TimedEnvData
}

func NewEnv() *Env {
	return &Env{}
}

func (e *Env) GetEnvData() *EnvData {
	data := &EnvData{}
	data.Temperature, data.Humidity = e.readTemperature()
	data.Ch2o = e.readCh2o()
	return data
}

func (e *Env) readTemperature() (string, string) {
	for i := 0; i < 3; i++ {
		out, err := exec.Command("./read_pi3_temperature.sh").Output()
		if err != nil {
			log.Printf("%s", err)
		} else {
			fs := strings.Split(string(out), " ")
			if len(fs) == 4 {
				return strings.TrimSuffix(fs[1], "\n"), strings.TrimSuffix(fs[3], "\n")
			}
		}
		time.Sleep(3 * time.Second)
	}
	log.Printf("failed to read temperature")
	return "", ""
}

func (e *Env) readCh2o() string {
	for i := 0; i < 3; i++ {
		out, err := exec.Command("./read_pi3_ch2o.sh").Output()
		if err != nil {
			log.Printf("%s", err)
		} else {
			fs := strings.Split(string(out), " ")
			if len(fs) == 2 {
				return strings.TrimSuffix(fs[1], "\n")
			}
		}
		time.Sleep(3 * time.Second)
	}
	log.Printf("failed to read ch2o")
	return ""
}
