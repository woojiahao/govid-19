package utility

import (
  "fmt"
  "strconv"
)

func Check(e error) {
  if e != nil {
    panic(e)
  }
}

func ToFloat32(s string) float32 {
  f, err := strconv.ParseFloat(s, 32)
  Check(err)
  return float32(f)
}

func ToInt(s string) int {
  if s == "" {
    return 0
  }
  f, err := strconv.ParseInt(s, 10, 64)
  Check(err)
  return int(f)
}

func CheckInt(value, prop string, min, max int) (num int, status bool, errMsg string) {
  f, err := strconv.ParseInt(value, 10, 64)
  if err != nil {
    status = false
    errMsg = fmt.Sprintf("Invalid input for '%s'. Must be int.", prop)
    return
  }

  num = int(f)

  if num < min {
    status = false
    errMsg = fmt.Sprintf("Invalid input for '%s'. Must be greater than %d", prop, min)
    return
  }

  if num > max {
    status = false
    errMsg = fmt.Sprintf("Invalid input for '%s'. Must be less than %d", prop, max)
    return
  }

  status = true

  return
}
