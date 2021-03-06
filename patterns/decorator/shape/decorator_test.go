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
			name: "Square 3.00 radius which is blue",
			fields: fields{
				Shape: &Square{
					Side: 3,
				},
				Colour: "blue",
			},
			want: "Square with side 3.00 has the colour blue",
		},
		{
			name: "Square 5.44 radius which is green",
			fields: fields{
				Shape: &Square{
					Side: 5.44,
				},
				Colour: "green",
			},
			want: "Square with side 5.44 has the colour green",
		},
		{
			name: "Square 4 radius which is red",
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

func TestTransparentShapeDecorator_Render(t1 *testing.T) {
	type fields struct {
		Shape        Shape
		Transparency float32
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name: "Circle 3.00 radius which is blue",
			fields: fields{
				Shape: &Circle{
					Radius: 3,
				},
				Transparency: 4.00,
			},
			want: "Circle with radius 3.00 has 4.00% transparency",
		},
		{
			name: "Circle 5.44 radius which is green",
			fields: fields{
				Shape: &ColorShapeDecorator{
					Shape: &Circle{
						Radius: 5.44,
					},
					Colour: "green",
				},

				Transparency: 4.00,
			},
			want: "Circle with radius 5.44 has the colour green has 4.00% transparency",
		},
		{
			name: "Square 3.00 radius which is blue",
			fields: fields{
				Shape: &Square{
					Side: 3,
				},
				Transparency: 4.00,
			},
			want: "Square with side 3.00 has 4.00% transparency",
		},
		{
			name: "Square 5.44 radius which is green",
			fields: fields{
				Shape: &ColorShapeDecorator{
					Shape: &Square{
						Side: 5.44,
					},
					Colour: "green",
				},

				Transparency: 4.00,
			},
			want: "Square with side 5.44 has the colour green has 4.00% transparency",
		},
	}
	for _, tt := range tests {
		t1.Run(tt.name, func(t1 *testing.T) {
			t := TransparentShapeDecorator{
				Shape:        tt.fields.Shape,
				Transparency: tt.fields.Transparency,
			}
			if got := t.Render(); got != tt.want {
				t1.Errorf("Render() = %v, want %v", got, tt.want)
			}
		})
	}
}
