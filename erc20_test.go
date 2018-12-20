package ethrpc

import (
	"strings"
	"testing"
)

func Test_paddingAddress(t *testing.T) {
	type args struct {
		address string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{name: "0", args: args{"0x0"}, want: strings.Repeat("0", 25)},
		{name: "1", args: args{"0"}, want: strings.Repeat("0", 25)},
		{name: "2", args: args{"1"}, want: strings.Repeat("0", 24) + "1"},
		{name: "3", args: args{"0x3f5ce5fbfe3e9af3971dd833d26ba9b5c936f0be"}, want: "0000000000000000000000003f5ce5fbfe3e9af3971dd833d26ba9b5c936f0be"},
		{name: "4", args: args{"3f5ce5fbfe3e9af3971dd833d26ba9b5c936f0be"}, want: "0000000000000000000000003f5ce5fbfe3e9af3971dd833d26ba9b5c936f0be"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := paddingAddress(tt.args.address); got != tt.want {
				t.Errorf("paddingAddress() = %v, want %v", got, tt.want)
			}
		})
	}
}
