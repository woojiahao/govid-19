package api

import "strings"

type sortOrder string

const (
  ascending  sortOrder = "asc"
  descending sortOrder = "desc"
)

func (so *sortOrder) isValid() bool {
  order := sortOrder(strings.ToLower(string(*so)))

  switch order {
  case ascending:
    fallthrough
  case descending:
    return true
  default:
    return false
  }
}
