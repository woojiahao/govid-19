package data

type SortOrder string

const (
  Ascending  SortOrder = "asc"
  Descending SortOrder = "desc"
)

func CheckSortOrder(raw string) (order SortOrder, status bool, errMsg string) {
  // TODO Better way to check against the constants
  switch raw {
  case "asc":
    order = Ascending
    status = true
  case "desc":
    order = Descending
    status = true
  default:
    order = ""
    status = false
    errMsg = "Invalid sort order. Available values are [ asc, desc ]"
  }
  return
}
