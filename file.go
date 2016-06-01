package main

import (
  "fmt"
  "os"
  "io"
  "log"
  "bufio"
)

func errCheck(e error) {
    if e != nil {
        log.Fatal(e)
        panic(e)
    }
}
func main() {

  file, err := os.Create("pyjac.txt");
  defer file.Close()
  errCheck(err)

  file.WriteString("Pyjac is here again\n")
  file.WriteString("Pyjac is here again 2\n")
  file.WriteString("Pyjac is here again 3")

  fmt.Println("Done")

  file, err = os.Open("pyjac.txt")
  errCheck(err)
  scanner := bufio.NewScanner(file)
  for scanner.Scan() {
      fmt.Println(scanner.Text())
  }

  if err := scanner.Err(); err != nil {
      log.Fatal(err)
  }

  // Reference: https://www.rosettacode.org/wiki/Read_a_file_line_by_line#Go
  f, err := os.Open("pyjac.txt")
    if err != nil {
        log.Fatal(err)
    }
    bf := bufio.NewReader(f)
    for {
        switch line, err := bf.ReadString('\n'); err {
        case nil:
            // valid line, echo it.  note that line contains trailing \n.
            fmt.Print(line)
        case io.EOF:
            if line > "" {
                // last line of file missing \n, but still valid
                fmt.Println(line)
            }
            return
        default:
            log.Fatal(err)
        }
    }
}