
package shape
import "testing"

func Test_shape(t *testing.T) {
r := Rectangle{2, 4}
sin := r.perimeter()
if sin != 12 {
t.Error("Expected is 12,but actual is ", sin)
}

c := Circle{5}
x := c.area()
if x != 78.53981633974483 {
  t.Error("Expected is 78.53981633974483, but actual is ",x)
}
}
