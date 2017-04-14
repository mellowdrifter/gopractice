package main

import "testing"

type testPair struct {
  values []int32
  sum int32
}

var tests = []testPair{
  { []int32{5,5}, 10},
  { []int32{0,0}, 0},
  { []int32{-1,9}, 8},
  { []int32{40,5}, 45},
  { []int32{-5,-5}, -10},
}

func TestAdd(t *testing.T) {
  for _, pair := range tests {
    sum := Add(pair.values[0], pair.values[1])
    if sum != pair.sum {
      t.Error("Expected ", pair.sum, ", Got", sum)
    }
  }
}
