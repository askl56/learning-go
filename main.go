package main

import (
  "fmt"
  "flag"
  "time"
  "math/rand"
  "errors"
)

var r = flag.Int("r", 2, "r")
var c = flag.Int("c", 2, "c")
var n = flag.Int("n", 3, "n")

func random(min, max int) int {
  return rand.Intn(max - min) + min
}

func buildGrid(r int, c int, n int) ([][]int, error) {
  grid := [][]int{}
  row := []int{}

  for i := 0; i < r; i++ {
    for j := 0; j < c; j++ {
      row = append(row, random(1, n+1))
      if j > 0 {
        if row[j-1] == row[j] {
          return [][]int{}, errors.New("Bad grid")
        }
      }

      if i > 0 {
        if grid[i-1][j] == row[j] {
          return [][]int{}, errors.New("Bad grid")
        }
      }
    }

    grid = append(grid, row)
    row = []int{}
  }

  return grid, nil
}

func main() {
  flag.Parse()
  rand.Seed(time.Now().Unix())

  grid, err := buildGrid(*r, *c, *n)
  for err != nil {
    grid, err = buildGrid(*r, *c, *n)
  }

  fmt.Println(grid)
}
