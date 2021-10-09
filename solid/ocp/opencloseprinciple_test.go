package ocp

import (
	"reflect"
	"testing"
)

func TestBetterFilter_Filter(t *testing.T) {
	type args struct {
		products       []Product
		specification Specification
	}
	tests := []struct {
		name       string
		args       args
		wantResult []Product
	}{
		{
			name: "Get green products",
			args: args{
				products: []Product{
					{
						name:   "Apple",
						colour: green,
						size:   small,
					},
					{
						name:   "Tree",
						colour: green,
						size:   large,
					},
					{
						name:   "House",
						colour: blue,
						size:   large,
					},
				},
				specification: ColourSpecification{colour: green},
			},
			wantResult: []Product{
				{
					name:   "Apple",
					colour: green,
					size:   small,
				},
				{
					name:   "Tree",
					colour: green,
					size:   large,
				},
			},
		},
		{
			name: "Get large products",
			args: args{
				products: []Product{
					{
						name:   "Apple",
						colour: green,
						size:   small,
					},
					{
						name:   "Tree",
						colour: green,
						size:   large,
					},
					{
						name:   "House",
						colour: blue,
						size:   large,
					},
				},
				specification: SizeSpecification{size: large},
			},
			wantResult: []Product{
				{
					name:   "Tree",
					colour: green,
					size:   large,
				},
				{
					name:   "House",
					colour: blue,
					size:   large,
				},
			},
		},
		{
			name: "Get large & green products",
			args: args{
				products: []Product{
					{
						name:   "Apple",
						colour: green,
						size:   small,
					},
					{
						name:   "Tree",
						colour: green,
						size:   large,
					},
					{
						name:   "House",
						colour: blue,
						size:   large,
					},
				},
				specification: AndSpecification{
					first:  SizeSpecification{size: large},
					second: ColourSpecification{colour: green},
				},
			},
			wantResult: []Product{
				{
					name:   "Tree",
					colour: green,
					size:   large,
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := &BetterFilter{}
			if gotResult := b.Filter(tt.args.products, tt.args.specification); !reflect.DeepEqual(gotResult, tt.wantResult) {
				t.Errorf("Filter() = %v, want %v", gotResult, tt.wantResult)
			}
		})
	}
}
