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
<svg width="1200px" height="800px" viewBox="0 0 1200 800" version="1.1" xmlns="http://www.w3.org/2000/svg" xmlns:xlink= "http://www.w3.org/1999/xlink">
  <style>
  * {
    fill: none;
    stroke: black;
    stroke-width: 3;
  }
  .knot {
    fill: blue;
    stroke: blue;
  }
  .empty {
    fill: none;
    stroke: blue;
    stroke-width: 2;
  }
  </style>
  <rect x="20" y="20" width="1160" height="760" class="empty"/>
`

const radius = 4

func main() {
  var points []image.Point
  if len(os.Args) > 1 {
    points = argsToPoints(os.Args[1:])
  } else {
    points = argsToPoints([]string{ "200,200", "350,344", "388,431", "460,527", "800,400" })
  }
  segments := bezier.SvgControlPointsI(points)
  fmt.Print(prefix)
  for _, s := range(segments) {
    fmt.Println(`<path d="`, s, `"/>`)
  }
  for _, p := range(points) {
    fmt.Printf(`<circle cx="%d" cy="%d" r="%d" class="knot"/>` + "\n", p.X, p.Y, radius)
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
