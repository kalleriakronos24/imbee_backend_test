package utils

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"time"

	"github.com/claudiu/gocron"
	"github.com/kalleriakronos24/imbee-backend/config"
)

func backupDatabase() {
	dateToday := time.Now()
	day := dateToday.Day()
	month := dateToday.Month()
	year := dateToday.Year()
	hour := dateToday.Hour()
	minute := dateToday.Minute()
	fileName := fmt.Sprintf("DB-%d%d%d%d-%d.sql", day, month, year, hour, minute)

	workdir, err := os.Getwd()

	if err != nil {
		fmt.Printf("failed to get os workdir : %s", err)
	}

	dbBackupPath := filepath.Join(workdir, fmt.Sprintf("../database-backup/%s", fileName))
	backupCommand1 := fmt.Sprintf("docker exec -u %s %s pg_dump %s > %s", config.AppConfig.DBUsername, config.AppConfig.DBContainerName, config.AppConfig.DBDatabase, dbBackupPath)
	cmd, err := exec.Command("/bin/bash", "-c", backupCommand1).CombinedOutput()

	if err != nil {
		fmt.Println(fmt.Sprint(err) + ": " + string(cmd))
		return
	}

	// Print the output
	fmt.Printf("database backed up : %s \n", fileName)
	fmt.Printf("database backed up every %d minutes", config.AppConfig.DBBackupTimerInMinutes)
}

func DatabaseBackupCron() {
	gocron.Start()
	gocron.Every(config.AppConfig.DBBackupTimerInMinutes).Minutes().Do(backupDatabase)
	//time.Sleep(20 * time.Second)
	//gocron.Clear()
	fmt.Println("All task removed")
}
