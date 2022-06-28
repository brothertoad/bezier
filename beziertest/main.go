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

const prefix = `<?xml version="1.0" standalone="no"?>
<!DOCTYPE svg PUBLIC "-//W3C//DTD SVG 1.1//EN" "http://www.w3.org/Graphics/SVG/1.1/DTD/svg11.dtd">
<svg width="1800px" height="1000px" viewBox="0 0 1800 1000" version="1.1" xmlns="http://www.w3.org/2000/svg" xmlns:xlink= "http://www.w3.org/1999/xlink">
  <style>
  * {
    fill: none;
    stroke: black;
    stroke-width: 3;
  }
  </style>
`

const radius = 4

func main() {
  var points []image.Point
  if len(os.Args) > 1 {
    points = argsToPoints(os.Args[1:])
  } else {
    points = argsToPoints([]string{ "300,200", "350,344", "388,411", "440,477", "500,500" })
  }
  segments := bezier.SvgControlPointsI(points)
  fmt.Print(prefix)
  for _, s := range(segments) {
    fmt.Println(`<path d="`, s, `"/>`)
  }
  for _, p := range(points) {
    fmt.Printf(`<circle cx="%d" cy="%d" r="%d"/>` + "\n", p.X, p.Y, radius)
  }
  fmt.Println("</svg>")
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
