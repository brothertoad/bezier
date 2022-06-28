package main

import (
  "fmt"
  "image"
  "log"
  "os"
  "strconv"
  "strings"
  "github.com/brothertoad/bezier"
)

func main() {
  var points []image.Point
  if len(os.Args) > 1 {
    points = argsToPoints(os.Args[1:])
  } else {
    points = argsToPoints([]string{ "300,200", "417,344", "388,411", "440,477", "500,500" })
  }
  fmt.Printf("%+v\n", points)
  segments := bezier.SvgControlPointsI(points)
  for _, s := range(segments) {
    fmt.Printf("%s\n", s)
  }
}

func argsToPoints(args []string) []image.Point {
  points := make([]image.Point, len(args))
  for j, arg := range(args) {
    coords := strings.Split(arg, ",")
    x, err := strconv.Atoi(coords[0])
    if err != nil {
      log.Fatalf("x value %s is not a number\n", coords[0])
    }
    y, err := strconv.Atoi(coords[1])
    if err != nil {
      log.Fatalf("y value %s is not a number\n", coords[1])
    }
    points[j].X = x
    points[j].Y = y
  }
  return points
}
