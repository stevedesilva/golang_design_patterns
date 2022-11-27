package generator

import (
	"reflect"
	"testing"
)

func TestNewEmployeeFactory(t *testing.T) {
	type args struct {
		position     string
		annualIncome int
	}
	tests := []struct {
		name string
		args args
		want func(name string) *Employee
	}{
		{},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewEmployeeFactory(tt.args.position, tt.args.annualIncome); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewEmployeeFactory() = %v, want %v", got, tt.want)
			}
		})
	}
}
