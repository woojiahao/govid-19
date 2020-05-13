package database

import (
  "fmt"
  "github.com/jinzhu/gorm"
  _ "github.com/jinzhu/gorm/dialects/postgres"
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

// Configures the database to upload the data to
func Setup() {
  databaseUrl := url.URL{
    Scheme:   "postgres",
    User:     url.UserPassword(user, pass),
    Host:     fmt.Sprintf("%s:%d", host, port),
    Path:     database,
    RawQuery: (&url.Values{"sslmode": []string{"disable"}}).Encode(),
  }
  db, err := gorm.Open("postgres", databaseUrl.String())
  Check(err)
  defer func() {
    _ = db.Close()
  }()

  // Use singlular table naming convention
  db.SingularTable(true)

  // As the entire data structure may change after an update, it's best to drop all
  // necessary tables and re-create them again
  db.DropTableIfExists(&Record{})
  db.DropTableIfExists(&Location{})

  // Automatically generates the database tables necessary
  db.CreateTable(&Location{})
  db.
    CreateTable(&Record{}).
    AddUniqueIndex("idx_unique_record", "date", "location_id").
    AddForeignKey("location_id", "location(id)", "RESTRICT", "RESTRICT")
}
