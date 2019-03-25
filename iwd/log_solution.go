package iwd

import (
	"encoding/json"
	"log"
	"os"

	"github.com/vroup/mo-iwd-sa/config"
)

func logSolution(conf *config.Config, archive *Archive) {
	logName := "solution_log_" + conf.LogID + "_" + conf.DataSize
	f, err := os.OpenFile(logName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	for idx := range archive.ElementList {
		{
			solution := archive.ElementList[idx]
			score := solution.Wd.Score
			scoreLog, err := json.Marshal(score)
			if err != nil {
				log.Fatal(err)
			}
			if _, err := f.WriteString(string(scoreLog) + ",\n"); err != nil {
				log.Println(err)
			}
		}
	}
}
