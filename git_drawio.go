package main

import (
	"fmt"
	"os/exec"
	"sort"
	"strings"
)


func main()  {
  fmt.Println("Setting up git repo...")
  cmd := exec.Command("git", "status", "-s")
  b_result, err := cmd.Output()

  if (err != nil) {
    fmt.Println(err.Error())
    return
  }

  // Parse to string
  str_result := string(b_result[:])

  fmt.Println("Identifying changes...")
  var a_gsi []GitStatusItem // changes stored here
  // For each line of result
  for _, s := range strings.Split(str_result, "\n") {
    if (len(s) == 0) { // Last line is empty
      continue
    }
    if (s[len(s) - 1] == '/') { // ignore folders
      continue
    }
    gsi := GitStatusItem_New(s) // parse GSI
    a_gsi = append(a_gsi, gsi) // Add to changes array
  }

  // Sort changes alphabetically by file name
  sort.Slice(a_gsi, func(l, r int) bool {
    return a_gsi[l].file < a_gsi[r].file
  })

  GenDrawioFile(a_gsi)
  fmt.Println("Done!")
}
