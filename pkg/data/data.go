package data

import (
  "fmt"
  . "github.com/woojiahao/govid-19/pkg/utility"
  "gopkg.in/src-d/go-git.v4"
  "os"
  "reflect"
  "strings"
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

type CountryInformation map[string]map[string]StateData

func (ci CountryInformation) FilterCountry(c string) CountryInformation {
  result := CountryInformation{}
  for country, stateData := range ci {
    if strings.Contains(strings.ToLower(country), strings.ToLower(c)) {
      result[country] = stateData
    }
  }
  return result
}

func (ci CountryInformation) FilterState(s string) CountryInformation {
  result := CountryInformation{}
  for country, stateData := range ci {
    for state, data := range stateData {
      if strings.Contains(strings.ToLower(state), strings.ToLower(s)) {
        if _, ok := result[country]; !ok {
          result[country] = make(map[string]StateData)
        }
        result[country][state] = data
      }
    }
  }
  return result
}

func (ci CountryInformation) First(n int) CountryInformation {
  result, counter := CountryInformation{}, 0
  for country, stateData := range ci {
    if counter >= n {
      break
    }

    result[country] = stateData
    counter++
  }
  return result
}

func (ci CountryInformation) Last(n int) CountryInformation {
  result, counter := CountryInformation{}, 0
  for country, stateData := range ci {
    if counter < n {
      continue
    }

    result[country] = stateData
    counter++
  }
  return result
}

func (ci CountryInformation) SortTotal(order SortOrder) {

}

var (
  Countries      []Country
  ConfirmedCases Series
  RecoveredCases Series
  DeathCases     Series
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

// Data loaded on startup as changes happen once a day and can be updated once the changes are made
func process() {
  ConfirmedCases, DeathCases, RecoveredCases = getTimeSeries(Confirmed),
    getTimeSeries(Deaths),
    getTimeSeries(Recovered)

  Countries = getCountriesFromTimeSeries(ConfirmedCases)
}

// If the current instance already contains the repository, pull the latest changes from the repository
// If the current instance does not contain the repository, clone the repository
func LoadData() {
  if _, err := os.Stat(Root.AsString()); os.IsNotExist(err) {
    tmp := os.Mkdir(Root.AsString(), 0755)
    Check(tmp)

    load()
  } else {
    update()
  }

  process()
}
