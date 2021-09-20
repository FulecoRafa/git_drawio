package main

import (
	"fmt"
	"os/exec"
	"strings"
)


func main()  {
  cmd := exec.Command("git", "status", "-s")
  b_result, err := cmd.Output()

  if (err != nil) {
    fmt.Println(err.Error())
    return
  }

  str_result := string(b_result[:])

  for i, s := range strings.Split(str_result, "\n") {
    if (len(s) == 0) {
      continue
    }
    gsi := git_status_item_New(s)
    fmt.Println(gsi.Drawio(i, 120, 120))
  }
}
