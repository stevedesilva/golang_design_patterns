package constructorinterface

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
			name: "old person",
			args: args{
				name: "mike",
				age:  100,
			},
			want: &oldPerson{
				name: "mike",
				age:  100,
			},
		},
		{
			name: "younger person",
			args: args{
				name: "ben",
				age:  10,
			},
			want: &person{
				name: "ben",
				age:  10,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewPerson(tt.args.name, tt.args.age); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewPerson() = %v, want %v", got, tt.want)
			}
		})
	}
}
