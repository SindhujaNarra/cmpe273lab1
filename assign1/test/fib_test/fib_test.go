package fib
import ( "testing")

func Test_fib (t *testing.T) {
      x := fib(4)
      if x != 3 {
        t.Error("expected is 3, but actual is ",x)
      }
      if fib(0) != 0 {
         t.Error("failed")
      }

}
