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

type dataRow struct {
  location Location
  records  []interface{}
}

type Manager struct {
  DB *gorm.DB
}

func connect() *gorm.DB {
  databaseUrl := url.URL{
    Scheme:   "postgres",
    User:     url.UserPassword(user, pass),
    Host:     fmt.Sprintf("%s:%d", host, port),
    Path:     database,
    RawQuery: (&url.Values{"sslmode": []string{"disable"}}).Encode(),
  }
  DB, err := gorm.Open("postgres", databaseUrl.String())
  Check(err)
  return DB
}

func (manager *Manager) configure() {
  manager.DB.LogMode(true)
  manager.DB.SingularTable(true)
}

func (manager *Manager) setupTables() {
  manager.DB.DropTableIfExists(&Record{})
  manager.DB.DropTableIfExists(&Location{})

  manager.DB.CreateTable(&Location{})
  manager.DB.
    CreateTable(&Record{}).
    AddForeignKey("location_id", "location(id)", "RESTRICT", "RESTRICT")
}

func (manager *Manager) createData(confirmedCases, recoveredCases, deathCases data.Series) []dataRow {
  counter := 0

  dataRows := make([]dataRow, 0)
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
        Date:       d.Date,
        Confirmed:  d.Value,
        Recovered:  recoveredCases.GetValueOfDate(r.Country, r.State, d.Date),
        Deaths:     deathCases.GetValueOfDate(r.Country, r.State, d.Date),
        LocationID: uint(counter),
      }
      row.records = append(row.records, record)
    }

    dataRows = append(dataRows, row)
  }

  return dataRows
}

// Load the database with the time series information
func (manager *Manager) loadData(dataRows []dataRow) {
  for _, record := range dataRows {
    manager.DB.Create(&record.location)
    _ = gormbulk.BulkInsert(manager.DB, record.records, 3000)
  }
}

// Configures the database to upload the data to
func Setup(confirmedCases, recoveredCases, deathCases data.Series) *Manager {
  db := connect()
  manager := Manager{db}
  manager.configure()
  manager.setupTables()
  dataRows := manager.createData(confirmedCases, recoveredCases, deathCases)
  manager.loadData(dataRows)
  return &manager
}
