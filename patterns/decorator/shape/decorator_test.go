package shape

import "testing"

func TestColorShapeDecorator_Render(t *testing.T) {
	type fields struct {
		Shape  Shape
		Colour string
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name: "Circle 3.00 radius which is blue",
			fields: fields{
				Shape: &Square{
					Side: 3,
				},
				Colour: "blue",
			},
			want: "Square with side 3.00 has the colour blue",
		},
		{
			name: "Circle 5.44 radius which is green",
			fields: fields{
				Shape: &Square{
					Side: 5.44,
				},
				Colour: "green",
			},
			want: "Square with side 5.44 has the colour green",
		},
		{
			name: "Circle 4 radius which is red",
			fields: fields{
				Shape: &Square{
					Side: 4,
				},
				Colour: "red",
			},
			want: "Square with side 4.00 has the colour red",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &ColorShapeDecorator{
				Shape:  tt.fields.Shape,
				Colour: tt.fields.Colour,
			}
			if got := c.Render(); got != tt.want {
				t.Errorf("Render() = %v, want %v", got, tt.want)
			}
		})
	}
}
