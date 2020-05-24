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
  ConfirmedTimeSeries          = TimeSeries + "/time_series_covid19_confirmed_global.csv"
  DeathsTimeSeries             = TimeSeries + "/time_series_covid19_deaths_global.csv"
  RecoveredTimeSeries          = TimeSeries + "/time_series_covid19_recovered_global.csv"
)

func (path RepoPath) AsString() string {
  ref := reflect.ValueOf(path)
  return ref.String()
}

// Load the repository into /tmp.
func load() {
  _, err := git.PlainClone(Root.AsString(), false, &git.CloneOptions{
    URL:      "https://github.com/CSSEGISandData/COVID-19.git",
    Progress: os.Stdout,
  })
  Check(err)
}

// Update the repository to the latest changes.
func update() {
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

// If the current instance already contains the repository, pull the latest changes from the repository
// If the current instance does not contain the repository, clone the repository
func UpdateData() {
  if _, err := os.Stat(Root.AsString()); os.IsNotExist(err) {
    tmp := os.Mkdir(Root.AsString(), 0755)
    Check(tmp)

    load()
  } else {
    update()
  }
}

func LoadData() (Series, Series, Series) {
  return getTimeSeries(Confirmed), getTimeSeries(Recovered), getTimeSeries(Deaths)
}
