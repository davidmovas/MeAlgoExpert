package ClassPhotos

import (
	"reflect"
	"testing"
)

func Test_ClassPhotos(t *testing.T) {
	testFunctions := map[string]func(red, blue []int) bool{
		"ClassPhotos_V1": ClassPhotos_V1,
	}

	type args struct {
		red  []int
		blue []int
	}

	tests := []struct {
		args args
		want bool
	}{
		{
			args: args{
				red:  []int{5, 8, 1, 3, 4},
				blue: []int{6, 9, 2, 4, 5},
			},
			want: true,
		},
	}

	for name, fn := range testFunctions {
		t.Run(name, func(t *testing.T) {
			for _, tt := range tests {
				if got := fn(tt.args.red, tt.args.blue); !reflect.DeepEqual(got, tt.want) {
					t.Errorf("%s = %v, want %v", name, got, tt.want)
				}
			}
		})
	}
}
