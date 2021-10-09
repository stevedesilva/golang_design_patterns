package lsp

import "testing"

func TestUseIt(t *testing.T) {
	type args struct {
		sized Sized
	}
	tests := []struct {
		name     string
		args     args
		wantArea int
	}{
		{
			name: "get area for rectangle",
			args: args{sized: &Rectangle{2,3}},
			wantArea: 20,
		},
		//{
		//	name: "get area for square",
		//	args: args{sized: NewSquare(2)},
		//	wantArea: 20,
		//},

	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotArea := UseIt(tt.args.sized); gotArea != tt.wantArea {
				t.Errorf("UseIt() = %v, want %v", gotArea, tt.wantArea)
			}
		})
	}
}