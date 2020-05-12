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

type SeriesGrouping map[string][]TimeSeriesRecord

var (
  confirmedCases SeriesGrouping
  recoveredCases SeriesGrouping
  deathCases     SeriesGrouping
  Countries      []Country
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
func GetAll() (Series, Series, Series) {
  confirmed, deaths, recovered := GetTimeSeries(Confirmed), GetTimeSeries(Deaths), GetTimeSeries(Recovered)
  return confirmed, deaths, recovered
}

// Loads all the data that is returned by each endpoint
// Data loading is performed implicitly as the data set is huge and static for the most part
// The data is updated once per day and remains untouched for the rest of the day meaning that
// it's easier to just process everything at once and when the data is updated, re-process it
// to improve performance
func Process() {
  confirmedCases, deathCases, recoveredCases = getCases(Confirmed), getCases(Deaths), getCases(Recovered)
  Countries = getCountriesFromTimeSeries(GetTimeSeries(Confirmed))
}

func groupByCountry(s Series) SeriesGrouping {
  data := make(map[string][]TimeSeriesRecord)
  for _, r := range s.Records {
    data[r.Country] = append(data[r.Country], r)
  }

  return data
}

func getCases(st TimeSeriesType) SeriesGrouping {
  return groupByCountry(GetTimeSeries(st))
}
