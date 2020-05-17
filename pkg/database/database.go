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
  "strconv"
)

// TODO Add a check if the latest changes are uploaded, if they are, don't need to do the db operations
var (
  user     = os.Getenv("POSTGRES_USER")
  pass     = os.Getenv("POSTGRES_PASSWORD")
  database = os.Getenv("POSTGRES_DB")
  host     = os.Getenv("HOST")
  hasLog   = os.Getenv("HAS_LOG")
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

// Configure GORM database settings
func (manager *Manager) configure() {
  logMode, err := strconv.ParseBool(hasLog)
  if err != nil {
    logMode = false
  }
  manager.DB.LogMode(logMode)
  manager.DB.SingularTable(true)
}

// Create database structure
func (manager *Manager) setupTables() {
  manager.DB.DropTableIfExists(&Record{})
  manager.DB.DropTableIfExists(&Location{})

  manager.DB.CreateTable(&Location{})
  manager.DB.
    CreateTable(&Record{}).
    AddForeignKey("location_id", "location(id)", "RESTRICT", "RESTRICT")
}

// Parse .csv data into format that the database can receive
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

// Upload .csv data to database
func (manager *Manager) UploadData(confirmedCases, recoveredCases, deathCases data.Series) {
  manager.setupTables()
  dataRows := manager.createData(confirmedCases, recoveredCases, deathCases)
  manager.loadData(dataRows)
}

// Return if the database is up-to-date aka latest changes have been uploaded
func (manager *Manager) IsUpToDate() bool {
  tables := []string{"location", "record"}
  for _, table := range tables {
    var result []interface{}
    manager.
      DB.
      Table(table).
      Select("distinct created_at::date").
      Where("current_timestamp::date <> created_at::date").
      Find(&result)

    rowsAffected := len(result)

    if rowsAffected != 0 {
      return false
    }
  }

  return true
}

// Configure GORM
func Setup() *Manager {
  db := connect()
  manager := Manager{db}
  manager.configure()
  return &manager
}
