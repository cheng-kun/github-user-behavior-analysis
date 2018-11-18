package github_archive

import (
	"errors"
	"fmt"
	"github.com/github-user-behavior-analysis/logs"
	"io"
	"net/http"
	"os"
	"path"
	"time"
)

func DownloadFile (basedir string, year int, month int, day int, hour int) error {
	dir := path.Join(basedir, fmt.Sprintf("%04d", year), fmt.Sprintf("%02d", month), fmt.Sprintf("%02d", day))

	stat, err := os.Stat(dir)
	if err != nil {
		if os.IsNotExist(err) {
			err = os.MkdirAll(dir, os.ModePerm)
				if err != nil {
					logs.PrintLogger().Errorf("Failed to create '%s' ",dir)
					return err
				}
		} else {
			logs.PrintLogger().Error(err)
			return err
		}
	} else if !stat.IsDir() {
		errStmt := fmt.Sprintf("'%s' is not a directory", dir)
		logs.PrintLogger().Error(errStmt)
		return errors.New(errStmt)
	}

	fileName := path.Join(dir, fmt.Sprintf("%d.json.gz",hour))

	_,err = os.Stat(fileName)
	if err != nil {
		if os.IsNotExist(err) {
			url := fmt.Sprintf("http://data.githubarchive.org/%04d-%02d-%02d-%d.json.gz", year, month, day, hour)

			resp, err := http.Get(url)
			if err != nil {
				logs.PrintLogger().Error(err)
				return err
			}
			defer resp.Body.Close()

			file, err := os.Create(fileName)
			if err != nil {
				logs.PrintLogger().Error(err)
				return err
			}
			defer file.Close()

			bytes, err := io.Copy(file, resp.Body)
			if err != nil {
				logs.PrintLogger().Error(err)
				return err
			}

			logs.PrintLogger().Infof("Downloaded %d bytes to '%s'", bytes, fileName)

		} else {
			logs.PrintLogger().Error(err)
			return err
		}
	}

	return nil
}

func daysInMonth(year int, m time.Month) int {
	return time.Date(year, m+1, 0, 0, 0, 0, 0, time.UTC).Day()
}

func SaveFiles(pathName string) error  {
	now := time.Now()

	logs.PrintLogger().Infof("year %d month %d day %d hour %d", now.Year(), now.Month(), now.Day(), now.Hour())

	for year := now.Year(); year >= 2018; year-- {
		endMonth := time.December
		if year == now.Year() {
			endMonth = now.Month()
		}

		for month := time.January; month <= endMonth; month++ {
			endDay := daysInMonth(year, month)
			if year == now.Year() && month == now.Month() {
				endDay = now.Day() - 1
			}

			for day := 1; day <= endDay; day++ {
				if day%4 != 0{
					continue
				}
				for hour := 0; hour < 24 ; hour++ {
					if hour%3 != 0{
						continue
					}
					DownloadFile(pathName, year, int(month), day, hour)
				}
			}
		}
	}
	return nil
}