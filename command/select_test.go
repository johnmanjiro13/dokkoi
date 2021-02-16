package command

import (
	"math/rand"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestSelectCmd_Exec(t *testing.T) {
	rand.Seed(time.Now().UnixNano())
	cmd := &selectCmd{
		elements: []string{"john", "manji", "ro"},
	}
	var cnt1, cnt2, cnt3 int
	for i := 0; i < 5000; i++ {
		res, err := cmd.Exec()
		if err != nil {
			t.Fatal(err)
		}

		switch res {
		case "john":
			cnt1++
		case "manji":
			cnt2++
		case "ro":
			cnt3++
		default:
			t.Fatalf("unexpected result. %s", res)
		}
	}
	assert.Equal(t, true, equalRoughly(cnt1, cnt2))
	assert.Equal(t, true, equalRoughly(cnt2, cnt3))
}

func equalRoughly(a, b int) bool {
	acceptRange := 0.08
	return int(float64(a)*(1-acceptRange)) < b && b < int(float64(a)*(1+acceptRange))
}
