package data

import (
  "encoding/csv"
  . "github.com/woojiahao/govid-19/pkg/utility"
  "io"
  "os"
)

type State struct {
  Name string  `json:"state"`
  Long float32 `json:"long"`
  Lat  float32 `json:"lat"`
}

type Country struct {
  Name   string  `json:"country"`
  States []State `json:"states"`
}

func getCountriesFromTimeSeries(seriesType TimeSeriesType) []Country {
  file, err := os.Open(TimeSeriesPaths[seriesType].AsString())
  Check(err)
  r := csv.NewReader(file)
  idx, rawCountries := 0, make(map[string][]State)
  for {
    idx++
    record, err := r.Read()

    if err == io.EOF {
      break
    }

    if idx == 1 {
      continue
    }

    state := State{
      Name: record[0],
      Long: ToFloat32(record[2]),
      Lat:  ToFloat32(record[3]),
    }

    rawCountries[record[1]] = append(rawCountries[record[1]], state)
  }

  countries := make([]Country, 0, len(rawCountries))
  for country, states := range rawCountries {
    c := Country{country, states}
    countries = append(countries, c)
  }

  return countries
}

func GetCountries() []Country {
  confirmedCountries, deathsCountries, recoveredCountries := getCountriesFromTimeSeries(Confirmed),
    getCountriesFromTimeSeries(Deaths),
    getCountriesFromTimeSeries(Recovered)

  collatedCountries := make([]Country, 0)
  collatedCountries = append(collatedCountries, confirmedCountries...)
  collatedCountries = append(collatedCountries, deathsCountries...)
  collatedCountries = append(collatedCountries, recoveredCountries...)

  countryCheck, countries := make(map[string]bool), make([]Country, 0)
  for _, country := range collatedCountries {
    if !countryCheck[country.Name] {
      countryCheck[country.Name] = true
      countries = append(countries, country)
    }
  }

  return countries
}
