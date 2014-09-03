package main

import "code.google.com/p/go-tour/pic"

func Pic(dx, dy int) [][]uint8 {
  pic := make([][]uint8, dy)
  for y := range pic {
    s := make([]uint8, dx)
    for x := range s {
      //s[x] = uint8(x)^uint8(y)
      //s[x] = (uint8(x)+uint8(y))/2
      s[x] = uint8(x)*uint8(y)
    }
    pic[y] = s
  }
  return pic

}

func main() {
    pic.Show(Pic)
}
