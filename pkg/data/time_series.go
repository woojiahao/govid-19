package data

import (
  "encoding/csv"
  "fmt"
  . "github.com/woojiahao/govid-19/pkg/utility"
  "io"
  "os"
  "strings"
  "time"
)

var TimeSeriesPaths = map[TimeSeriesType]RepoPath{
  Confirmed: ConfirmedTimeSeries,
  Deaths:    DeathsTimeSeries,
  Recovered: RecoveredTimeSeries,
}

// Single day of data for a specific country/region
type TimeSeriesRecordData struct {
  Date  time.Time `json:"date"`
  Value int    `json:"value"`
}

// Single row in the time series.
type TimeSeriesRecord struct {
  TimeSeriesType TimeSeriesType         `json:"-"`
  State          string                 `json:"state"`
  Country        string                 `json:"country"`
  Longitude      float32                `json:"long"`
  Latitude       float32                `json:"lat"`
  Total          int                    `json:"total"`
  Data           []TimeSeriesRecordData `json:"data"`
}

type Series struct {
  TimeSeriesType TimeSeriesType     `json:"-"`
  Records        []TimeSeriesRecord `json:"records"`
}

func (s *Series) Clone(newRecords []TimeSeriesRecord) *Series {
  return &Series{
    TimeSeriesType: s.TimeSeriesType,
    Records:        newRecords,
  }
}

// TODO Case sensitive requests
func (s *Series) GetByCountry(country string) Series {
  results := make([]TimeSeriesRecord, 0)
  for _, record := range s.Records {
    if strings.ToLower(record.Country) == strings.ToLower(country) {
      results = append(results, record)
    }
  }

  return *s.Clone(results)
}

// TODO Case sensitive requests
func (s *Series) GetByState(state string) Series {
  results := make([]TimeSeriesRecord, 0)
  for _, record := range s.Records {
    if strings.ToLower(record.State) == strings.ToLower(state) {
      results = append(results, record)
    }
  }

  return *s.Clone(results)
}

func (s *Series) GetValueOfDate(country, state string, date time.Time) int {
  location := s.GetByCountry(country)
  location = location.GetByState(state)
  if len(location.Records) <= 0 {
    return 0
  }

  for _, record := range location.Records[0].Data {
    if date == record.Date {
      return record.Value
    }
  }

  return -1
}

func getTimeSeries(seriesType TimeSeriesType) Series {
  file, err := os.Open(TimeSeriesPaths[seriesType].AsString())
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
      rawData, timeHeaders, data := record[4:], headers[4:], make([]TimeSeriesRecordData, 0)
      for i, d := range rawData {
        prev := 0
        if i != 0 {
          prev = ToInt(rawData[i-1])
        }

        increment := ToInt(d) - prev

        date := strings.Split(timeHeaders[i], "/")
        month, day, year := date[0], date[1], date[2]
        const timeLayout = "01/02/2006"
        formattedDate, err := time.Parse(timeLayout, fmt.Sprintf("%02s/%02s/20%s", month, day, year))
        Check(err)

        data = append(data, TimeSeriesRecordData{
          Date:  formattedDate,
          Value: increment,
        })
      }
      timeSeriesRecord := TimeSeriesRecord{
        TimeSeriesType: seriesType,
        State:          record[0],
        Country:        record[1],
        Longitude:      ToFloat32(record[2]),
        Latitude:       ToFloat32(record[3]),
        Total:          ToInt(rawData[len(rawData)-1]),
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