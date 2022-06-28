package catmullrom

import (
  "image"
)

// The "F" suffix on the struct names means float64.

type PointF struct {
  X, Y float64
}

type BezierF struct {
  P0, P1, P2, P3 PointF
}

type Bezier struct {
  P0, P1, P2, P3 image.Point
}

func GetControlPointsI(p []image.Point) []Bezier {
  // Convert integer values to floats, calcalute control points, and then
  // convert back to integers.
  pf := make([]PointF, len(p), len(p))
  for j, source := range(p) {
    pf[j].X = float64(source.X)
    pf[j].Y = float64(source.Y)
  }
  bezierF := GetControlPointsF(pf)
  bezier := make([]Bezier, len(bezierF), len(bezierF))
  for j, bz := range(bezierF) {
    bezier[j].P0.X = int(bz.P0.X)
    bezier[j].P0.Y = int(bz.P0.Y)
    bezier[j].P1.X = int(bz.P1.X)
    bezier[j].P1.Y = int(bz.P1.Y)
    bezier[j].P2.X = int(bz.P2.X)
    bezier[j].P2.Y = int(bz.P2.Y)
    bezier[j].P3.X = int(bz.P3.X)
    bezier[j].P3.Y = int(bz.P3.Y)
  }
  return bezier
}

// Algorithm ported from https://www.particleincell.com/wp-content/uploads/2012/06/bezier-spline.js
func GetControlPointsF(p []PointF) []BezierF {
  n := len(p) - 1 // number of segments
  bezier := make([]BezierF, n, n)

  // Fill in the end points of the Bezier curves.
  for j, bz := range(bezier) {
    bz.P0 = p[j]
    bz.P3 = p[j+1]
  }

  // Slices needed for calculations, listed as "rhs vector" on page listed above.
  a := make([]PointF, n, n)
  b := make([]PointF, n, n)
  c := make([]PointF, n, n)
  r := make([]PointF, n, n)

  a[0].X, a[0].Y = 0.0, 0.0
  b[0].X, b[0].Y = 2.0, 2.0
  c[0].X, c[0].Y = 1.0, 1.0
  r[0].X = p[0].X + 2.0 * p[1].X
  r[0].Y = p[0].Y + 2.0 * p[1].Y

  for j := 1; j < (n-1); j++ {
    a[j].X, a[j].Y = 1.0, 1.0
    b[j].X, b[j].Y = 4.0, 4.0
    c[j].X, c[j].Y = 1.0, 1.0
    r[j].X = 4.0 * p[j].X + 2.0 * p[j+1].X
    r[j].Y = 4.0 * p[j].Y + 2.0 * p[j+1].Y
  }

  a[n-1].X, a[n-1].Y = 2.0, 2.0
  b[n-1].X, b[n-1].Y = 7.0, 7.0
  c[n-1].X, c[n-1].Y = 0.0, 0.0
  r[n-1].X = 8.0 * p[n-1].X + p[n].X
  r[n-1].Y = 8.0 * p[n-1].Y + p[n].Y

  for j := 1; j < n; j++ {
    m := a[j].X / b[j-1].X
    b[j].X = b[j].X - m * c[j-1].X
    r[j].X = r[j].X - m * r[j-1].X
    m = a[j].Y / b[j-1].Y
    b[j].Y = b[j].Y - m * c[j-1].Y
    r[j].Y = r[j].Y - m * r[j-1].Y
  }

  // Compute first control point (P1 in BeziefF), from last to first.
  bezier[n-1].P1.X = r[n-1].X / b[n-1].X
  bezier[n-1].P1.Y = r[n-1].Y / b[n-1].Y
  for j := n - 2; j >= 0; j-- {
    bezier[j].P1.X = (r[j].X - c[j].X * bezier[j+1].P1.X) / b[j].X
    bezier[j].P1.Y = (r[j].Y - c[j].Y * bezier[j+1].P1.Y) / b[j].Y
  }

  // Compute second control point (P2 in BezierF), from first to last.
  for j := 0; j < n - 1; j++ {
    bezier[j].P2.X = 2 * p[j].X - bezier[j].P1.X
    bezier[j].P2.Y = 2 * p[j].Y - bezier[j].P1.Y
  }
  bezier[n-1].P2.X = 0.5 * (p[n].X + bezier[n-1].P1.X)
  bezier[n-1].P2.Y = 0.5 * (p[n].Y + bezier[n-1].P1.Y)

  return bezier
}
