package copy_value

import "testing"

type X struct {
	I  int
	Is []int
	Im map[int]int
	S  string
	Ss []string
	Sm map[string]string
}

type Y struct {
	X  X
	Xs []X
	Xm map[string]X
	Xp *X
}

func TestStruct(t *testing.T) {
	answer := createAnswerX()
	dst := createEmptyX()
	CopyValue(&dst, &answer)
	checkX(&dst, &answer, t)
}

func TestStructInStruct(t *testing.T) {
	answer := createAnswerY()
	dst := createEmptyY()
	CopyValue(&dst, &answer)
	checkY(&dst, &answer, t)
}

func createAnswerX() X {
	return X{
		I:  10,
		Is: []int{1, 2, 3},
		Im: map[int]int{1: 10, 2: 20, 3: 30, 4: 40},
		S:  "hoge",
		Ss: []string{"v", "w", "x", "y", "z"},
		Sm: map[string]string{"a": "A", "b": "B", "c": "C", "d": "D", "e": "E", "f": "F"},
	}
}

func createEmptyX() X {
	return X{
		Is: make([]int, 3),
		Im: make(map[int]int),
		Ss: make([]string, 5),
		Sm: make(map[string]string),
	}
}

func createAnswerY() Y {
	xforp := createAnswerX()
	return Y{
		X: createAnswerX(),
		Xs: []X{
			createAnswerX(),
			createAnswerX(),
			createAnswerX(),
		},
		Xm: map[string]X{
			"a": createAnswerX(),
			"b": createAnswerX(),
			"c": createAnswerX(),
			"d": createAnswerX(),
		},
		Xp: &xforp,
	}
}

func createEmptyY() Y {
	xforp := createEmptyX()
	return Y{
		X: createEmptyX(),
		Xs: []X{
			createEmptyX(),
			createEmptyX(),
			createEmptyX(),
		},
		Xm: make(map[string]X),
		Xp: &xforp,
	}
}

func checkX(dst *X, answer *X, t *testing.T) {
	// I
	if dst.I != answer.I {
		t.Errorf("I: Value mismatch: dst.I=%d, answer.I=%d\n", dst.I, answer.I)
	}

	// Is
	if len(dst.Is) != len(answer.Is) {
		t.Errorf("Is: Length mismatch: len(dst.Is)=%d, len(answer.Is)=%d\n", len(dst.Is), len(answer.Is))
	}
	for i, _ := range dst.Is {
		if dst.Is[i] != answer.Is[i] {
			t.Errorf("Is: Value mismatch: dst.Is[%d]=%d, answer.Is[%d]=%d\n", i, dst.Is[i], i, answer.Is[i])
		}
	}

	// Im
	if len(dst.Im) != len(answer.Im) {
		t.Errorf("Im: Length mismatch: len(dst.Im)=%d, len(answer.Im)=%d\n", len(dst.Im), len(answer.Im))
	}
	for k, _ := range dst.Im {
		if dst.Im[k] != answer.Im[k] {
			t.Errorf("Im: Value mismatch: dst.Im[%d]=%d, answer.Im[%d]=%d\n", k, dst.Im[k], k, answer.Im[k])
		}
	}

	// S
	if dst.S != answer.S {
		t.Errorf("S: Value mismatch: dst.S=%s, answer.S=%s\n", dst.S, answer.S)
	}

	// Ss
	if len(dst.Ss) != len(answer.Ss) {
		t.Errorf("Ss: Length mismatch: len(dst.ss)=%d, len(answer.Ss)=%d\n", len(dst.Ss), len(answer.Ss))
	}
	for i, _ := range dst.Ss {
		if dst.Ss[i] != answer.Ss[i] {
			t.Errorf("Ss: Value mismatch: dst.Ss[%d]=%s, answer.Ss[%d]=%s\n", i, dst.Ss[i], i, answer.Ss[i])
		}
	}

	// Sm
	if len(dst.Sm) != len(answer.Sm) {
		t.Errorf("Sm: Length mismatch: len(dst.Sm)=%d, len(answer.Sm)=%d\n", len(dst.Sm), len(answer.Sm))
	}
	for k, _ := range dst.Sm {
		if dst.Sm[k] != answer.Sm[k] {
			t.Errorf("Sm: Value mismatch: dst.Sm[%k]=%k, answer.Sm[%k]=%k\n", k, dst.Sm[k], k, answer.Sm[k])
		}
	}
}

func checkY(dst *Y, answer *Y, t *testing.T) {
	// X
	checkX(&dst.X, &answer.X, t)

	// Xs
	if len(dst.Xs) != len(answer.Xs) {
		t.Errorf("Xs: Length mismatch: len(dst.Xs)=%d, len(answer.Xs)=%d\n", len(dst.Xs), len(answer.Xs))
	}
	for i, _ := range dst.Xs {
		checkX(&dst.Xs[i], &answer.Xs[i], t)
	}

	// Xm
	if len(dst.Xm) != len(answer.Xm) {
		t.Errorf("Xm: Length mismatch: len(dst.Xm)=%d, len(answer.Xm)=%d\n", len(dst.Xm), len(answer.Xm))
	}
	for k, _ := range dst.Xm {
		dx := dst.Xm[k]
		ax := answer.Xm[k]
		checkX(&dx, &ax, t)
	}

	// Xp
	checkX(dst.Xp, answer.Xp, t)
}
