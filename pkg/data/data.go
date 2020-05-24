package data

import (
  "fmt"
  . "github.com/woojiahao/govid-19/pkg/utility"
  "gopkg.in/src-d/go-git.v4"
  "os"
)

type RepoPath string

const (
  Root                RepoPath = "tmp"
  DailyReports        RepoPath = "tmp/csse_covid_19_data/csse_covid_19_daily_reports"
  TimeSeries          RepoPath = "tmp/csse_covid_19_data/csse_covid_19_time_series"
  ConfirmedTimeSeries          = TimeSeries + "/time_series_covid19_confirmed_global.csv"
  DeathsTimeSeries             = TimeSeries + "/time_series_covid19_deaths_global.csv"
  RecoveredTimeSeries          = TimeSeries + "/time_series_covid19_recovered_global.csv"
)

// Load the repository into /tmp.
func load() {
  _, err := git.PlainClone(string(Root), false, &git.CloneOptions{
    URL:      "https://github.com/CSSEGISandData/COVID-19.git",
    Progress: os.Stdout,
  })
  Check(err)
}

// Update the repository to the latest changes.
func update() {
  r, err := git.PlainOpen(string(Root))
  Check(err)

  w, err := r.Worktree()
  Check(err)

  err = w.Pull(&git.PullOptions{
    RemoteName: "origin",
    Progress:   os.Stdout,
  })
  if err != nil {
    if err == git.NoErrAlreadyUpToDate {
      fmt.Println("Repository is up-to-date")
    } else {
      panic(err)
    }
  }
}

// If the current instance already contains the repository, pull the latest changes from the repository
// If the current instance does not contain the repository, clone the repository
func Update() {
  if _, err := os.Stat(string(Root)); os.IsNotExist(err) {
    tmp := os.Mkdir(string(Root), 0755)
    Check(tmp)

    load()
  } else {
    update()
  }
}

func Load() (Series, Series, Series) {
  return getTimeSeries(Confirmed), getTimeSeries(Recovered), getTimeSeries(Deaths)
}
