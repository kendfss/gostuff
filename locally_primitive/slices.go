package main
type Slice []int
func (self Slice) Clone(item int) Slice {
    other = Slice{}
    for i := range self {
        other = append(other, self[i])
    }
    return other
}
func (self Slice) Index(item int) int {
    for i, e := range self {
        if e == item {
            return i
        }
    }
    return -1
}
func (self *Slice) Remove(item int) {
    new :=
}
