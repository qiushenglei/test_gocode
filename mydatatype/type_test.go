package mydatatype

import (
	"reflect"
	"testing"
)

func TestSplit(t *testing.T) {
	test := map[string]struct {
		s    string
		seq  string
		want []string
	}{
		"test1": {
			s:    "我爱你",
			seq:  "爱你",
			want: []string{"我"},
		},
		"test2": {
			s:    "这里是一个很长长的string，而且还有中文",
			seq:  "，",
			want: []string{"这里是一个很长长的string", "而且还有中文"},
		},
	}

	for k, v := range test {
		t.Run(k, func(t *testing.T) {
			res := Split(v.s, v.seq)
			if !reflect.DeepEqual(res, v.want) {
				t.Errorf("want: %v, res: %v", v.want, res)
			}
		})
	}
}

func BenchmarkSplit(b *testing.B) {
	test := map[string]struct {
		s    string
		seq  string
		want []string
	}{
		"test1": {
			s:    "我爱你",
			seq:  "爱你",
			want: []string{"我"},
		},
		"test2": {
			s:    "这里是一个很长长的string，而且还有中文",
			seq:  "，",
			want: []string{"这里是一个很长长的string", "而且还有中文"},
		},
	}

	for i := 0; i < b.N; i++ {
		for k, v := range test {
			b.Run(k, func(b *testing.B) {
				res := Split(v.s, v.seq)
				if !reflect.DeepEqual(res, v.want) {
					b.Errorf("want: %v, res: %v", v.want, res)
				}
			})
		}
	}
}
