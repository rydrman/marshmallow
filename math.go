package mm

import "math"

const (
    Deg2Rad = math.Pi / 180.0
    Rad2Deg = 180.0 / math.Pi
)

func Max(a, b float64, vals ...float64) float64 {

    m := math.Max(a, b)

    for _, v := range vals {

        m = math.Max(m, v)

    }

    return m

}

func Round(v float64) float64 {

    return math.Floor(v + 0.5)

}

func Clamp(value, min, max float64) float64 {

    return math.Max(min, math.Min(max, value))

}
