package main 

import "fmt"

type Punto struct {
  x, y int
}

func (p Punto) String() string {
  return fmt.Sprintf("[%d,%d]", p.x, p.y)
}

func riflessioneY(p Punto) Punto {
  var puntoRiflesso Punto
  puntoRiflesso.x = - p.x
  puntoRiflesso.y = p.y
  return puntoRiflesso
}

func main() {
  var p1 Punto
  var x, y int
  fmt.Scan(&x, &y)
  p1.x, p1.y = x, y
  fmt.Println(p1.String())
  fmt.Println(riflessioneY(p1).String())
}
