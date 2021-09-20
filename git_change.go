package main

import (
  "fmt"
  "os"
  "path"
)

type GitChange_t string

const (
  Add    GitChange_t = "#99FF33"
  Modify GitChange_t = "#FFFF66"
  Delete GitChange_t = "#FF0000"
)

type GitStatusItem struct {
  staged bool
  change GitChange_t
  file string
}

func (gsi GitStatusItem) Drawio(id, x, y int) string {
  staged := ""
  if gsi.staged {
    staged = "arcSize=50;"
  }
  return fmt.Sprintf(`
                <mxCell id="%d" value="%s" style="rounded=1;whiteSpace=wrap;html=1;strokeColor=%s;%s" parent="1" vertex="1">
                    <mxGeometry x="%d" y="%d" width="120" height="60" as="geometry"/>
                </mxCell>
`, id, path.Base(gsi.file), gsi.change, staged, x, y)
}

func GitStatusItem_New(s string) GitStatusItem {
  file := s[3:]
  if s[0] == '?' {
    return GitStatusItem{false, Add, file}
  }
  var staged bool
  var change GitChange_t
  if s[0] == ' ' {
    staged = false
    switch s[1] {
    case 'M':
      change = Modify
    case 'D':
      change = Delete
    default:
      fmt.Println("Unrecognized git change:")
      fmt.Println(s)
      os.Exit(1)
    }
  } else {
    staged = true
    switch s[0] {
    case 'A':
      change = Add
    case 'M':
      change = Modify
    case 'D':
      change = Delete
    default:
      fmt.Println("Unrecognized git change:")
      fmt.Println(s)
      os.Exit(1)
    }
  }
  return GitStatusItem{staged, change, file}
}
