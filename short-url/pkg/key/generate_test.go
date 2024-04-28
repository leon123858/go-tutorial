package key

import (
	"math/rand"
	"reflect"
	"testing"
)

func Test_generateKey(t *testing.T) {
	type args struct {
		randGen *rand.Rand
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Test generateKey",
			args: args{
				randGen: rand.New(rand.NewSource(0)),
			},
			want: "mUNERA",
		},
		{
			name: "Test generateKey",
			args: args{
				randGen: rand.New(rand.NewSource(1)),
			},
			want: "BpLnfg",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := generateKey(tt.args.randGen); got != tt.want {
				t.Errorf("generateKey() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGenerateKeys(t *testing.T) {
	type args struct {
		num  int
		seed int64
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			name: "Test GenerateKeys",
			args: args{
				num:  5,
				seed: 0,
			},
			want: []string{"mUNERA", "9rI2cv", "TK4UHo", "mcjcEQ", "vymkzA"},
		}, {
			name: "Test GenerateKeys",
			args: args{
				num:  -1,
				seed: 0,
			},
			want: []string{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GenerateKeys(tt.args.num, tt.args.seed); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GenerateKeys() = %v, want %v", got, tt.want)
			}
		})
	}
}
