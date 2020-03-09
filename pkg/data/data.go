package data

import (
  "fmt"
  . "github.com/woojiahao/govid-19/pkg/utility"
  "gopkg.in/src-d/go-git.v4"
  "os"
  "reflect"
)

type RepoPath string

const (
  Root                RepoPath = "tmp"
  DailyReports        RepoPath = "tmp/csse_covid_19_data/csse_covid_19_daily_reports"
  TimeSeries          RepoPath = "tmp/csse_covid_19_data/csse_covid_19_time_series"
  ConfirmedTimeSeries RepoPath = TimeSeries + "/time_series_19-covid-Confirmed.csv"
  DeathsTimeSeries    RepoPath = TimeSeries + "/time_series_19-covid-Deaths.csv"
  RecoveredTimeSeries RepoPath = TimeSeries + "/time_series_19-covid-Recovered.csv"
)

func (path RepoPath) AsString() string {
  ref := reflect.ValueOf(path)
  return ref.String()
}

// Load the repository into /tmp.
func Load() {
  _, err := git.PlainClone(Root.AsString(), false, &git.CloneOptions{
    URL:      "https://github.com/CSSEGISandData/COVID-19.git",
    Progress: os.Stdout,
  })
  Check(err)
}

// Update the repository to the latest changes.
func Update() {
  r, err := git.PlainOpen(Root.AsString())
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

// Aggregate the values from the time series to create a grouping for the information.
func GetAll() {
  confirmed := GetTimeSeries(Confirmed)
  for _, record := range confirmed {
    if record.Province == "Hubei" {
      fmt.Println(record)
    }
  }
}
