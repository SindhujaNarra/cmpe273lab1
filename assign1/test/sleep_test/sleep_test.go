package sleep

import ("testing";"time")

func Test_sleep(t *testing.T) {

    x := time.Now()
    sleep()
    y := time.Now()
    if x == y {
      t.Error("failed !")
    }
}
