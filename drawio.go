package main

import (
  "os"
  "bufio"
  "fmt"
  "path"
)

func check(e error) {
  if e != nil { fmt.Println(e); os.Exit(1)}
}

func GenDrawioFile(a_gsi []GitStatusItem) {
  folder_name, err := os.Getwd()
  check(err)
  fmt.Println("Generating drawio file:", folder_name+"/"+path.Base(folder_name)+".dio")

  f, err := os.Create(folder_name+"/"+path.Base(folder_name)+".dio")
  check(err)

  defer f.Close()

  w := bufio.NewWriter(f)

  _, err = w.WriteString(`
<mxfile host="65bd71144e">
    <diagram id="xYE1o5OOA-KSdR3Z0aG8" name="Page-1">
        <mxGraphModel dx="1748" dy="975" grid="1" gridSize="10" guides="1" tooltips="1" connect="1" arrows="1" fold="1" page="1" pageScale="1" pageWidth="1169" pageHeight="827" math="0" shadow="0">
            <root>
                <mxCell id="0"/>
                <mxCell id="1" parent="0"/>
`)
  check(err)

  offsetI:=2
  offsetX:=10
  offsetY:=30
  for i, gsi := range a_gsi {
    x := ((i % 7) * 170) + offsetX
    y := ((i / 7) * 90) + offsetY
    str := gsi.Drawio(i+offsetI, x, y)
    _, err = w.WriteString(str)
    check(err)
  }

  _, err = w.WriteString(`
          </root>
        </mxGraphModel>
    </diagram>
</mxfile>
`)
  check(err)

  w.Flush()
}
