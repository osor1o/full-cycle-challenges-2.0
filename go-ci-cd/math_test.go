package main

import "testing"

func testSum(t *testing.T) {
  total := Sum(15, 15)

  if total != 30 {
    t.Errorf("Invalid result %d-%d", total, 30)
  }
}
