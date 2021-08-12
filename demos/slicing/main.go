package main
import "github.com/leaanthony/slicer"
import "fmt"

func main() {
  s := slicer.String()
  s.Add("one")
  s.Add("two")
  s.AddSlice([]string{"three","four"})
  fmt.Printf("My slice = %+v\n", s.AsSlice())
  
  t := slicer.String()
  t.Add("zero")
  t.AddSlicer(s)
  fmt.Printf("My slice = %+v\n", t.AsSlice())
}

