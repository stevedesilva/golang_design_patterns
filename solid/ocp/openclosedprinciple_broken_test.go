package ocp

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"reflect"
	"testing"
)

func TestFilter_FilterByColour(t *testing.T) {
	type args struct {
		products []Product
		color    Colour
	}
	tests := []struct {
		name string
		args args
		want []Product
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
				color: green,
			},
			want: []Product{
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
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f := &Filter{}
			if got := f.FilterByColour(tt.args.products, tt.args.color); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FilterByColour() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFilter_FilterBySize(t *testing.T) {
	type args struct {
		products []Product
		size    Size
	}
	tests := []struct {
		name string
		args args
		want []Product
	}{
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
				size: large,
			},
			want: []Product{
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
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f := &Filter{}
			if got := f.FilterBySize(tt.args.products, tt.args.size); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FilterByColour() = %v, want %v", got, tt.want)
			}
		})
	}
}


func testColour(colour Colour) Colour {
	fmt.Sprintf("Input type %T %[1]d",colour)
	return colour
}

func TestColour(t *testing.T) {
	// why no check her
	assert.Equal(t,red,testColour(0))
	assert.NotEqual(t,0,testColour(0))
	//assert.NotEqual(t,red,0)
	//
	//assert.Equal(t,green,testColour(1))
	//assert.Equal(t,blue,testColour(2))
	//assert.Equal(t,blue,testColour(2))


	//assert.Equal(t,testColour(red),"0")
	//assert.Equal(t,testColour(blue),"2")
	//assert.Equal(t,testColour(99),"99")

	//assert.Equal(t,testColour(red),"0")
	//assert.Equal(t,testColour(0),0)

}