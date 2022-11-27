package struct

import (
	"reflect"
	"testing"
)

func TestNewPerson(t *testing.T) {
	type args struct {
		name string
		age  int
	}
	tests := []struct {
		name string
		args args
		want Person
	}{
		{
			name: "basic person",
			args: args{
				name: "Steve",
				age:  33,
			},
			want: &Person{
				Name:     "Steve",
				Age:      33,
				EyeCount: 2,
			},
		},
		{
			name: "basic person",
			args: args{
				name: "Mike Moon",
				age:  333,
			},
			want: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got, _ := NewPerson(tt.args.name, tt.args.age); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewPerson() = %v, want %v", got, tt.want)
			}
		})
	}
}
