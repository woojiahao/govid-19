package database

import (
  "fmt"
  "github.com/jinzhu/gorm"
  _ "github.com/jinzhu/gorm/dialects/postgres"
  gormbulk "github.com/t-tiger/gorm-bulk-insert/v2"
  "github.com/woojiahao/govid-19/pkg/data"
  . "github.com/woojiahao/govid-19/pkg/utility"
  "net/url"
  "os"
  "time"
)

var (
  user     = os.Getenv("POSTGRES_USER")
  pass     = os.Getenv("POSTGRES_PASSWORD")
  database = os.Getenv("POSTGRES_DB")
  host     = os.Getenv("HOST")
)

const (
  port = 5432
)

type Location struct {
  gorm.Model
  Country string `gorm:"not null;type:text"`
  State   string `gorm:"type:text"`
  Long    float32
  Lat     float32
  Records []Record
}

// Single record for a day in a country/state
type Record struct {
  gorm.Model
  Date       *time.Time
  Confirmed  int `gorm:"not null"`
  Recovered  int `gorm:"not null"`
  Deaths     int `gorm:"not null"`
  Location   Location
  LocationID uint
}

func connect() *gorm.DB {
  databaseUrl := url.URL{
    Scheme:   "postgres",
    User:     url.UserPassword(user, pass),
    Host:     fmt.Sprintf("%s:%d", host, port),
    Path:     database,
    RawQuery: (&url.Values{"sslmode": []string{"disable"}}).Encode(),
  }
  db, err := gorm.Open("postgres", databaseUrl.String())
  Check(err)
  return db
}

func configure(db *gorm.DB) {
  db.LogMode(true)
  db.SingularTable(true)
}

func setupTables(db *gorm.DB) {
  db.DropTableIfExists(&Record{})
  db.DropTableIfExists(&Location{})

  db.CreateTable(&Location{})
  db.
    CreateTable(&Record{}).
    AddForeignKey("location_id", "location(id)", "RESTRICT", "RESTRICT")
}

// Load the database with the time series information
func loadData(db *gorm.DB, confirmedCases, recoveredCases, deathCases data.Series) {
  counter := 0

  type dataRow struct {
    location Location
    records  []interface{}
  }
  data := make([]dataRow, 0)
  for _, r := range confirmedCases.Records {
    counter++
    row := dataRow{records: make([]interface{}, 0)}
    location := Location{
      Country: r.Country,
      State:   r.State,
      Long:    r.Longitude,
      Lat:     r.Latitude,
    }
    row.location = location

    for _, d := range r.Data {
      record := Record{
        Date:       &d.Date,
        Confirmed:  d.Value,
        Recovered:  recoveredCases.GetValueOfDate(r.Country, r.State, d.Date),
        Deaths:     deathCases.GetValueOfDate(r.Country, r.State, d.Date),
        LocationID: uint(counter),
      }
      row.records = append(row.records, record)
    }

    data = append(data, row)
  }

  for _, record := range data {
    db.Create(&record.location)
    gormbulk.BulkInsert(db, record.records, 3000)
  }
}

// Configures the database to upload the data to
func Setup(confirmedCases, recoveredCases, deathCases data.Series) {
  db := connect()
  defer func() {
    _ = db.Close()
  }()

  configure(db)
  setupTables(db)
  loadData(db, confirmedCases, recoveredCases, deathCases)
}
