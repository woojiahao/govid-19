package utility

import "strconv"

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
  f, err := strconv.ParseInt(s, 10, 64)
  Check(err)
  return int(f)
}
