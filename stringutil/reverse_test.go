package stringutil
import (
	"testing"
	"strconv"
)

func TestReverse(t *testing.T) {
	cases := []struct {
		in, want string
	}{
		{"Hello, world", "dlrow ,olleH"},
		{"Hello, 世界", "界世 ,olleH"},
		{"", ""},
	}
	for _, c := range cases {
		got := Reverse(c.in)
		if got != c.want {
			t.Errorf("Reverse(%q) == %q, wang %q", c.in, got, c.want)
		}
	}
}

func TestReverse2(t *testing.T){
	s1 := "5556"
	s2 := "6555"

	sr := Reverse(s1)
	if sr != s2 {
		t.Error("测试不通过")
	} else {
		t.Log("测试通过")
	}
}

func BenchmarkReverse1(b *testing.B){
	for i := 0; i < b.N; i++{
		Reverse(PrefixString(strconv.Itoa(100+i), "012"))
	}
}

func BenchmarkTimeConsume(b *testing.B){
	b.StopTimer()
	b.Log("Starting test Benchmark...\n")
	b.StartTimer()

	for i := 0; i < b.N; i++{
		Reverse(PrefixString(strconv.Itoa(100+i), "012"))
	}
}
