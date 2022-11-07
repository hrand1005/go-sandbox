package main

import (
  "flag"
  "fmt"
  "strings"
  "strconv"
  // "regexp"
)

// For custom flags it's necessary to implement the flag.Value interface:
//
// type Value interface {
//   Value() string
//   Set(string) error
// }
//
// I'll also implement the sring interface for convenience.

type Rect struct {
  Length int
  Width int
}

func (r *Rect) String() string {
  rowString := strings.Repeat("*", r.Width) + "\n"
  rectString := strings.Repeat(rowString, r.Length)
  return fmt.Sprintf("Length: %v, Width: %v\n\n%s", r.Length, r.Width, rectString)
}

func (r *Rect) Value() string {
  return fmt.Sprintf("%s\n\nLength: %v\nWidth: %v\n", r, r.Length, r.Width)
}

func (r *Rect) Set(s string) error {
  // intRe := `^[1-9]\d*$`
  // rectRe := fmt.Sprintf(`%vx%v`, intRe, intRe)
  // re, err := regexp.Compile(rectRe)
  // if err != nil {
  //   return err
  // }
  //
  // matches := re.FindAllString(s, -1)
  // if matches == nil {
  //   return fmt.Errorf("rectangle dimensions not parse-able: %s", s)
  // }
  // if len(matches) > 1 {
  //   return fmt.Errorf("too many dimensions specified: %s", s)
  // }
  //
  // dimensions := strings.Split(matches[0], "x")
  dimensions := strings.Split(s, "x")
  length, err := strconv.Atoi(dimensions[0])
  if err != nil {
    return fmt.Errorf("could not convert to int: %v", dimensions[0])
  }
  width, err := strconv.Atoi(dimensions[1])
  if err != nil {
    return fmt.Errorf("could not convert to int: %v", dimensions[1])
  }

  r.Length = length
  r.Width = width

  return nil
}

func RectFlag(name string, value Rect, usage string) *Rect {
  r := Rect{
    Length: value.Length,
    Width: value.Width,
  }
  flag.CommandLine.Var(&r, name, usage)
  return &r
}
