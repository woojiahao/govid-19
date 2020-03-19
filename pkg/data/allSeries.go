package data

type AllSeries struct {
  confirmed Series
  deaths    Series
  recovered Series
}

func (as *AllSeries) ToJSON() map[string]interface{} {
  return map[string]interface{}{
    "confirmed": as.confirmed.Records,
    "deaths":    as.deaths.Records,
    "recovered": as.recovered.Records,
  }
}

func NewAllSeries(confirmed, deaths, recovered Series) *AllSeries {
  return &AllSeries{
    confirmed: confirmed,
    deaths:    deaths,
    recovered: recovered,
  }
}
