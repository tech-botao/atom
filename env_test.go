package lib

import (
	"fmt"
	"testing"
)

func TestSetEnvPath(t *testing.T) {
	type args struct {
		path string
	}
	tests := []struct {
		name string
		args args
	}{
		{".", args{"."}},
		{"~", args{"~"}},
		{"../", args{"../"}},
		{"log", args{"log"}},
		{"/etc", args{"/etc"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			SetEnvPath(tt.args.path)
			fmt.Print(&tt.args.path, GetEnvPath())
		})
	}
}