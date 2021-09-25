package dices

import (
	"fmt"
	"math/rand"
	"time"
)

const (
	// MDefault is default dimension of dice.
	MDefault = 6
)

// nolint:gochecknoinits
func init() {
	rand.Seed(time.Now().UnixNano())
}

// Result is information about dice roll.
type Result struct {
	Dices []int
	Base  int
	// Threshold means which values of dice dimension is suitable
	Threshold int
	Suitable  int
	Sum       int
}

// NewWithValues returns Result with filled dices values.
func NewWithValues(dd []int) *Result {
	m := len(dd)
	r := &Result{Dices: make([]int, m), Threshold: 0, Sum: 0, Suitable: m}

	copy(r.Dices, dd)

	for i := range r.Dices {
		r.Sum += r.Dices[i]
	}

	return r
}

// NewWithValuesAndThreshold returns Result with filled dices values.
func NewWithValuesAndThreshold(dd []int, t int) *Result {
	m := len(dd)
	r := &Result{Dices: make([]int, m), Threshold: t, Sum: 0, Suitable: 0}

	copy(r.Dices, dd)

	r.suitable()

	return r
}

// Md6 returns Result of M dice roll with default dimension.
func Md6(m int) *Result {
	return MdN(m, MDefault)
}

// Md6WithThreshold returns Result of M dice roll with default dimension using threshold T.
func Md6WithThreshold(m, t int) *Result {
	return MdNWithThreshold(m, MDefault, t)
}

// MdN returns Result of M dice roll with dimension N.
func MdN(m, n int) *Result {
	vv := make([]int, m)
	s := 0

	for i := m - 1; i >= 0; i-- {
		// nolint:gosec
		vv[i] = rand.Intn(n) + 1

		s += vv[i]
	}

	return &Result{
		Dices:     vv,
		Base:      n,
		Threshold: 0,
		Suitable:  m,
		Sum:       s,
	}
}

// MdNWithThreshold returns Result of M dice roll with dimension N using threshold T.
func MdNWithThreshold(m, n, t int) *Result {
	r := MdN(m, n)
	r.Threshold = t

	r.suitable()

	return r
}

// PrettySum returns string with pretty representation of Result with sum.
func (a *Result) PrettySum() string {
	return fmt.Sprintf("Dice roll %v (sum: %d).", a.Dices, a.Sum)
}

// PrettySuitable returns string with pretty representation of Result with count of suitable dices.
func (a *Result) PrettySuitable() string {
	return fmt.Sprintf("Dice roll: %v (suitable: %d).", a.Dices, a.Suitable)
}

// PrettyThreshold returns string with pretty representation of Result with threshold.
func (a *Result) PrettyThreshold() string {
	return fmt.Sprintf("%d of %d with suitable value %d or less.", len(a.Dices), a.Base, a.Threshold)
}

// Eq returns true if this result of dice roll equivalents specified one.
func (a *Result) Eq(b *Result) bool {
	if a == nil {
		return true
	}

	if b == nil {
		return false
	}

	if a.Threshold != b.Threshold || a.Suitable != b.Suitable || a.Sum != b.Sum || len(a.Dices) != len(b.Dices) {
		return false
	}

	for i := range a.Dices {
		if a.Dices[i] != b.Dices[i] {
			return false
		}
	}

	return true
}

func (a *Result) suitable() {
	a.Suitable, a.Sum = 0, 0

	for i := range a.Dices {
		if a.Dices[i] > a.Threshold {
			continue
		}

		a.Suitable++
		a.Sum += a.Dices[i]
	}
}
