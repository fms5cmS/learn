package some

import "testing"

type T struct {}

func (*T) foo() {

}

func (T) bar()  {

}

type S2 struct {
	*T
}

func TestStructAssemble(t *testing.T) {
	s := S2{}
	t.Logf("%#v",s)
	_ = s.foo
	s.foo()
	_ = s.bar
}
