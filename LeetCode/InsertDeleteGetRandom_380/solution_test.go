package InsertDeleteGetRandom_380

import (
	"reflect"
	"testing"
)

func Test_InsertDeleteGetRandom(t *testing.T) {
	testFunctions := map[string]func() []any{
		"InsertDeleteGetRandom_V1": v1,
	}

	tests := []struct {
		want []any
	}{
		{
			want: []any{true, false, true, 1, true, false, 2},
		},
	}

	for name, fn := range testFunctions {
		t.Run(name, func(t *testing.T) {
			for _, tt := range tests {
				if got := fn(); !reflect.DeepEqual(got, tt.want) {
					t.Errorf("%s = %v, want %v", name, got, tt.want)
				}
			}
		})
	}
}
