package main

import (
	"fmt"
	"flag"
  //"image"
	//"image/color"
)

func getPowerOfTwo(n uint8)uint8 {
  switch n {
    case 0:
      return 1
    case 1:
      return 2
    case 2:
      return 4
    case 3:
      return 8
    case 4:
      return 16
    case 5:
      return 32
    case 6:
      return 64
    case 7:
      return 128
  }
  return 0
}

func main() {

  var (
    argRule uint
    size uint
    rows uint
  )

  flag.UintVar(&argRule, "rule", 18, "a positive integer between 0 and 255 naming the rule to be applied")
  flag.UintVar(&size, "width", 64, "width in cells")
  flag.UintVar(&rows, "height", 32, "height in cells")
  flag.Parse()

  var rule = uint8(argRule)
  var cells = getStart(size)
  for i := uint(0); i < rows; i++ {
    fmt.Println(cells)
    cells = iterate_cells(rule,cells)
  }
}

func iterate_cells(rule uint8, states []byte)[]byte {
  var answer = make([]byte, len(states))
  var neighborhood byte
  var width = len(states)

  for idx, v := range states {
    neighborhood = 0
    if idx == 0 {
      if states[len(states)-1] == 1 {
        neighborhood += 1;
      }
    } else {
      if states[idx - 1] == 1 {
        neighborhood += 1;
      }
    }

    if v == 1 {
      neighborhood += 2;
    }

    if idx == width - 1 {
      if states[0] == 1 {
        neighborhood += 4;
      }
    } else {
      if states[idx + 1] == 1 {
        neighborhood += 4;
      }
    }
     if rule & getPowerOfTwo(neighborhood) > 0 {
       answer[idx] = 1
     } else {
       answer[idx] = 0
     }
  }

  return answer
}

func getStart(size uint)[]byte {
  var answer = make([]byte, size)
  answer[size/2] = 1
  return answer
}
