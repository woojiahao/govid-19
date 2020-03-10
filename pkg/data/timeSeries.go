package data

import (
  "encoding/csv"
  . "github.com/woojiahao/govid-19/pkg/utility"
  "io"
  "os"
)

type TimeSeriesType string

// Type of time series row.
const (
  Confirmed TimeSeriesType = "Confirmed"
  Deaths    TimeSeriesType = "Deaths"
  Recovered TimeSeriesType = "Recovered"
)

var timeSeriesPaths = map[TimeSeriesType]RepoPath{
  Confirmed: ConfirmedTimeSeries,
  Deaths:    DeathsTimeSeries,
  Recovered: RecoveredTimeSeries,
}

// Single row in the time series.
type TimeSeriesRecord struct {
  TimeSeriesType TimeSeriesType `json:"-"`
  State          string         `json:"state"`
  Country        string         `json:"country"`
  Longitude      float32        `json:"long"`
  Latitude       float32        `json:"lat"`
  Data           map[string]int `json:"data"`
}

type Series struct {
  TimeSeriesType TimeSeriesType
  Records        []TimeSeriesRecord
}

func (s *Series) GetCountry(country string) TimeSeriesRecord {
  for _, record := range s.Records {
    // TODO Allow fuzzy searching
    if record.Country == country {
      return record
    }
  }

  return TimeSeriesRecord{}
}

func (s *Series) GetState(state string) TimeSeriesRecord {
  for _, record := range s.Records {
    // TODO Allow fuzzy searching
    if record.State == state {
      return record
    }
  }

  return TimeSeriesRecord{}
}

func GetTimeSeries(seriesType TimeSeriesType) Series {
  file, err := os.Open(timeSeriesPaths[seriesType].AsString())
  Check(err)

  r := csv.NewReader(file)
  idx, headers, records := 0, make([]string, 0), make([]TimeSeriesRecord, 0)
  for {
    idx++
    record, err := r.Read()

    if err == io.EOF {
      break
    }

    if idx == 1 {
      for _, header := range record {
        headers = append(headers, header)
      }
    } else {
      rawData, timeHeaders, data := record[4:], headers[4:], make(map[string]int)
      for i, d := range rawData {
        data[timeHeaders[i]] = ToInt(d)
      }
      timeSeriesRecord := TimeSeriesRecord{
        TimeSeriesType: seriesType,
        State:          record[0],
        Country:        record[1],
        Longitude:      ToFloat32(record[2]),
        Latitude:       ToFloat32(record[3]),
        Data:           data,
      }
      records = append(records, timeSeriesRecord)
    }
    Check(err)
  }

  return Series{
    TimeSeriesType: seriesType,
    Records:        records,
  }
}
