package main

import (
  "fmt"
  "math"
  "sort"
  "strconv"
  "strings"
  "time"
)

func findMinDifference(timePoints []string) int {

  times := make([]time.Time, 0, len(timePoints))
  now := time.Now()

  for _, tP := range timePoints {
    strParsed := strings.Split(tP, ":")
    hour, _ := strconv.Atoi(strParsed[0])
    minutes, _ := strconv.Atoi(strParsed[1])
    year, month, day := now.Date()

    tme := time.Date(year, month, day, hour, minutes, 0, 0, time.UTC)

    times = append(times, tme)
  }

  minDif := math.MaxFloat64

  sort.Slice(times, func(i, j int) bool {
    dur := times[i].Sub(times[j])
    abs := math.Abs(dur.Minutes())
    if abs < minDif {
      minDif = abs
    }
    return times[i].Before(times[j])
  })

  times[0] = times[0].Add(time.Hour*24)

  if sub := times[0].Sub(times[len(times) -1]); sub.Minutes() < minDif {
    minDif = sub.Minutes()
  }

  return int(minDif)
}

func main() {
  fmt.Println(findMinDifference([]string{"23:59", "00:00"}))
}
