package data

import (
  "encoding/csv"
  . "github.com/woojiahao/govid-19/pkg/utility"
  "io"
  "os"
  "sort"
  "strings"
)

var TimeSeriesPaths = map[TimeSeriesType]RepoPath{
  Confirmed: ConfirmedTimeSeries,
  Deaths:    DeathsTimeSeries,
  Recovered: RecoveredTimeSeries,
}

type TimeSeriesRecordData struct {
  Date  string `json:"date"`
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

// Returns the sum of data for a record.
func (r *TimeSeriesRecord) SumData() int {
  sum := 0
  for _, data := range r.Data {
    sum += data.Value
  }
  return sum
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
    if strings.Contains(record.Country, country) {
      results = append(results, record)
    }
  }

  return *s.Clone(results)
}

// TODO Case sensitive requests
func (s *Series) GetByState(state string) Series {
  results := make([]TimeSeriesRecord, 0)
  for _, record := range s.Records {
    if strings.Contains(record.State, state) {
      results = append(results, record)
    }
  }

  return *s.Clone(results)
}

// Sorts results by the records
func (s Series) SortRecords(order SortOrder) Series {
  for _, record := range s.Records {
    sort.Slice(record.Data, func(i, j int) bool {
      switch order {
      case Ascending:
        return record.Data[i].Value < record.Data[j].Value
      case Descending:
        return record.Data[i].Value > record.Data[j].Value
      default:
        panic("invalid sort order")
      }
    })
  }
  return s
}

// Sorts results by the total
func (s Series) SortData(order SortOrder) Series {
  sort.Slice(s.Records, func(i, j int) bool {
    switch order {
    case Ascending:
      return s.Records[i].Total < s.Records[j].Total
    case Descending:
      return s.Records[i].Total > s.Records[j].Total
    default:
      panic("invalid sort order")
    }
  })
  return s
}

// Retrieves the first [num] (exclusive) of records in the series
func (s Series) First(num int) Series {
  return *s.Clone(s.Records[:num])
}

// Retrieves the last [num] (inclusive) of records in the series
func (s Series) Last(num int) Series {
  return *s.Clone(s.Records[len(s.Records)-num-1:])
}

func GetTimeSeries(seriesType TimeSeriesType) Series {
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
        data = append(data, TimeSeriesRecordData{
          Date:  timeHeaders[i],
          Value: ToInt(d) - prev,
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
