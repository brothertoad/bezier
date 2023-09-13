package main

import (
  "fmt"
  "image"
  "os"
  "strings"
  "github.com/brothertoad/bezier"
  "github.com/brothertoad/btu"
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
  .empty {
    fill: none;
    stroke: blue;
    stroke-width: 2;
  }
  .knot {
    fill: blue;
    stroke: blue;
  }
  .control {
    fill: green;
    stroke: green;
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
  // Get the beziers separately so we can plot the control points.
  beziers := bezier.GetControlPointsI(points)
  segments := bezier.ControlPointsToSvgI(beziers)
  fmt.Print(prefix)
  for _, s := range(segments) {
    fmt.Println(`<path d="`, s, `"/>`)
  }
  for _, p := range(points) {
    fmt.Printf(`<circle cx="%d" cy="%d" r="%d" class="knot"/>` + "\n", p.X, p.Y, radius)
  }
  for _, bz := range(beziers) {
    fmt.Printf(`<circle cx="%d" cy="%d" r="%d" class="control"/>` + "\n", bz.P1.X, bz.P1.Y, radius)
    fmt.Printf(`<circle cx="%d" cy="%d" r="%d" class="control"/>` + "\n", bz.P2.X, bz.P2.Y, radius)
  }
  fmt.Println("</svg>")
}

func argsToPoints(args []string) []image.Point {
  points := make([]image.Point, len(args))
  for j, arg := range(args) {
    coords := strings.Split(arg, ",")
    points[j].X = btu.Atoi(coords[0])
    points[j].Y = btu.Atoi(coords[1])
  }
  return points
}
