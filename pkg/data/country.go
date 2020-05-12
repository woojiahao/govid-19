package data

type State struct {
  Name string  `json:"state"`
  Long float32 `json:"long"`
  Lat  float32 `json:"lat"`
}

type Country struct {
  Name   string  `json:"country"`
  States []State `json:"states"`
}

// TODO Might be able to collapse this along with the processing of the other data
func getCountriesFromTimeSeries(s Series) []Country {
  stateMapping := make(map[string][]State, 0)
  for _, r := range s.Records {
    stateMapping[r.Country] = append(stateMapping[r.Country], State{r.State, r.Longitude, r.Latitude})
  }

  countries := make([]Country, 0, len(stateMapping))
  for country, states := range stateMapping {
    countries = append(countries, Country{country, states})
  }

  return countries
}
