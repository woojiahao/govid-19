package database

import (
  "github.com/jinzhu/gorm"
  "time"
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
  Date       time.Time `gorm:"type:date"`
  Confirmed  int        `gorm:"not null"`
  Recovered  int        `gorm:"not null"`
  Deaths     int        `gorm:"not null"`
  Location   Location
  LocationID uint
}
