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
  var board [][]uint8

  board = initialize_platten(size, rows)
  print_board(board)

  for i := uint(0); i < rows; i++ {
    print_board(board)
    board = iterate_board(rule, board)
  }
}

func print_board(board [][]uint8) {
  for row := len(board) - 1; row >= 0; row-- {
    fmt.Println(board[row])
  }
}

func iterate_board(rule uint8, board [][]uint8)[][]uint8 {
  var new_board [][]uint8
  
  size := len(board[0])
  rows := uint(len(board))
  for i := uint(1); i < rows; i++ {
    new_board = append(new_board, make([]uint8, size))
    copy(new_board[i - 1], board[i])
  }
  new_board = append(new_board, iterate_cells(rule, board[len(board)-1]))
  return new_board
}

func iterate_cells(rule uint8, states []uint8)[]uint8 {
  var answer = make([]uint8, len(states))
  var neighborhood uint8
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

func initialize_platten(size uint, rows uint)[][]uint8 {

  var board = [][]uint8{}

  for i := uint(0); i < rows - 1; i++ {
    board = append(board, make([]uint8, size))
  }
  var first = make([]uint8, size)
  first[size/2] = 1
  return append(board, first)
}
