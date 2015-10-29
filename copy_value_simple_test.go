package copy_value

import "testing"

func TestInt(t *testing.T) {
	var dst int
	src := 10
	CopyValue(&dst, &src)
	if dst != 10 {
		t.Fail()
	}
}

func TestString(t *testing.T) {
	var dst string
	src := "hoge"
	CopyValue(&dst, &src)
	if dst != "hoge" {
		t.Fail()
	}
}

func TestSliceShorterToLonger(t *testing.T) {
	dst := make([]string, 5)
	src := []string{"a", "b", "c"}
	CopyValue(&dst, &src)

	answer := []string{"a", "b", "c", "", ""}
	if len(dst) != len(answer) {
		t.Fail()
	}
	for i := 0; i < len(answer); i++ {
		if dst[i] != answer[i] {
			t.Fail()
		}
	}
}

func TestSliceLongerToShorter(t *testing.T) {
	dst := make([]string, 3)
	src := []string{"a", "b", "c", "d", "e"}
	CopyValue(&dst, &src)

	answer := []string{"a", "b", "c"}
	if len(dst) != len(answer) {
		t.Fail()
	}
	for i := 0; i < len(answer); i++ {
		if dst[i] != answer[i] {
			t.Fail()
		}
	}
}

func TestMap(t *testing.T) {
	dst := map[string]string{
		"d": "D",
	}
	src := map[string]string{
		"a": "A",
		"b": "B",
		"c": "C",
	}
	CopyValue(&dst, &src)

	answer := map[string]string{
		"a": "A",
		"b": "B",
		"c": "C",
		"d": "D",
	}
	if len(dst) != len(answer) {
		t.Errorf("Length mismatch: len(dst)=%d, len(answer)=%d\n", len(dst), len(answer))
	}
	for k := range answer {
		if dst[k] != answer[k] {
			t.Errorf("Value mismatch: dst[%s]=%s, answer[%s]=%s\n", k, dst[k], k, answer[k])
		}
	}
}
