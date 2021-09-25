package dices

import (
	"testing"
)

func TestMdN(t *testing.T) {
	t.Parallel()
	m := 20
	r := MdN(m, MDefault)

	if len(r.Dices) != m {
		t.Fatalf("len=%d expected=%d", len(r.Dices), m)
	}
	if r.Threshold != 0 {
		t.Fatalf("threshold=%d expected=0", r.Threshold)
	}
	if r.Suitable != m {
		t.Fatalf("suitable=%d expected=%d", r.Suitable, m)
	}
	s := 0
	for _, v := range r.Dices {
		if v < 1 || v > MDefault {
			t.Fatalf("value=%v expected=[1, %d]", v, MDefault)
		}
		s += v
	}
	if r.Sum != s {
		t.Fatalf("sum=%d expected=%d", r.Sum, s)
	}
}

func TestMdNWithThreshold(t *testing.T) {
	t.Parallel()
	m := 20
	th := 4
	r := MdNWithThreshold(m, MDefault, th)

	if len(r.Dices) != m {
		t.Fatalf("len=%d expected=%d", len(r.Dices), m)
	}
	if r.Threshold != th {
		t.Fatalf("threshold=%d expected=%d", r.Threshold, th)
	}

	sui, sum := 0, 0
	for _, v := range r.Dices {
		if v < 1 || v > MDefault {
			t.Fatalf("value=%v expected=[1, %d]", v, MDefault)
		}
		if v > th {
			continue
		}
		sui++
		sum += v
	}
	if r.Suitable != sui {
		t.Fatalf("suitable=%d expected=%d", r.Suitable, sui)
	}
	if r.Sum != sum {
		t.Fatalf("sum=%d expected=%d", r.Sum, sum)
	}
}

func Test0d6(t *testing.T) {
	t.Parallel()
	r := Md6(0)

	if len(r.Dices) != 0 {
		t.Fatalf("len=%d expected=0", len(r.Dices))
	}
	if r.Threshold != 0 {
		t.Fatalf("threshold=%d expected=0", r.Threshold)
	}
	if r.Suitable != 0 {
		t.Fatalf("suitable=%d expected=0", r.Suitable)
	}
	if r.Sum != 0 {
		t.Fatalf("sum=%d expected=0", r.Sum)
	}
}

func TestResult_Equal(t *testing.T) {
	t.Parallel()
	d1 := Md6(1)
	d2 := new(Result)
	*d2 = *d1
	d2.Dices = []int{d1.Dices[0]}

	if !d2.Eq(d1) {
		t.Errorf("d2=%+v expected=%+v", d2, d1)
	}

	for d2.Dices[0] == d1.Dices[0] {
		d2 = Md6(1)
	}

	if d2.Eq(d1) {
		t.Errorf("d2=%+v expectedNot=%+v", d2, d1)
	}
}

func TestNewWithValuesAndThreshold(t *testing.T) {
	t.Parallel()
	d := NewWithValuesAndThreshold([]int{6, 1, 1, 5, 5, 6, 1, 3, 2}, 4)

	if d.Suitable != 5 {
		t.Errorf("Suitable=%v expected=5", d.Suitable)
	}
}
