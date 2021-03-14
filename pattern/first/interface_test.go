package first

import (
	"fmt"
	"testing"
)


type Shape interface {
	Sides() int
	Area() int
}
type Square struct {
	len int
}
func (s* Square) Sides() int {
	return 4
}
func TestInterface(t *testing.T) {
	// var _ Shape = (*Square)(nil)
	s := Square{len: 5}
	fmt.Printf("%d\n",s.Sides())
}