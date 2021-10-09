package srp

import (
	"reflect"
	"testing"
)

func TestJournal_AddEntry(t *testing.T) {
	type fields struct {
		entries []string
	}
	type args struct {
		text string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   int
	}{
		{
			name: "add one entry",
			fields: fields{
				entries: []string{"first entry"},
			},
			args: args{
				text: "first entry",
			},
			want: 1,
		},
		{
			name: "add multiple entry",
			fields: fields{
				entries: []string{"1st entry", "2nd entry", "3rd entry", "4th entry", "5th entry"},
			},
			args: args{
				text: "first entry",
			},
			want: 5,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			j := &Journal{}
			for _, text := range tt.fields.entries {
				j.AddEntry(text)
			}

			if got := j.entryCount; got != tt.want {
				t.Errorf("AddEntry() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestJournal_RemoveEntry(t *testing.T) {
	type fields struct {
		entries []string
	}
	type args struct {
		index int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   []string
	}{
		{
			name: "remove entry",
			fields: fields{
				entries: []string{"first entry", "second entry"},
			},
			args: args{0},
			want: []string{"second entry"},
		},
		{
			name: "remove entry from middle",
			fields: fields{
				entries: []string{"first entry", "second entry", "last entry"},
			},
			args: args{1},
			want: []string{"first entry", "last entry"},
		},
		{
			name: "remove entry from last",
			fields: fields{
				entries: []string{"first entry", "second entry", "last entry"},
			},
			args: args{2},
			want: []string{"first entry", "second entry"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			j := &Journal{
				entries: tt.fields.entries,
			}
			j.RemoveEntry(tt.args.index)
			if !reflect.DeepEqual(j.entries, tt.want) {
				t.Errorf("DeleteEntry() = %v, want %v", j.entries, tt.want)
			}
		})
	}
}
