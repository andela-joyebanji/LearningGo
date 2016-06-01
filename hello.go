package main

import "fmt"

func main() {

  fmt.Println("Hello, World")
}


if !ok {
  if !inRange {
    r := frames.NackRange{
      FirstPacketNumber: i,
      LastPacketNumber:  i,
    }
    ranges = append(ranges, r)
    inRange = true
  } else {
    ranges[len(ranges)-1].FirstPacketNumber--
  }
} else {
  inRange = false
  entropy.Add(i, p.EntropyBit)
}

if ok {
  inRange = false
  entropy.Add(i, p.EntropyBit)
}
if !ok && !inRange {
  r := frames.NackRange{
    FirstPacketNumber: i,
    LastPacketNumber:  i,
  }
  ranges = append(ranges, r)
  inRange = true
}
if !ok && inRange {
  ranges[len(ranges)-1].FirstPacketNumber--
}












// Check to see if the listing we are adding already exists in the list. If so delete it.
for i, d := range(index){
  if d.Name != ld.Name {
    continue
  }
  if len(index) == 1 {
    index = []listingData{}
    break
  }
  index = append(index[:i], index[i + 1:]...)
}


/ Check to see if the listing we are adding already exists in the list. If so delete it.
  for i, d := range(index){
    if d.Name == ld.Name {
      if len(index) == 1 {
        index = []listingData{}
        break
      } else {
        index = append(index[:i], index[i + 1:]...)
      }
    }
  }

