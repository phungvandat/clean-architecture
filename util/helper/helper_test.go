package helper

import (
	"reflect"
	"testing"
)

func TestConvertTypeArrayToInterfaceArray(t *testing.T) {
	type input struct {
		A string
	}
	var (
		i1 = input{A: "a1"}
		i2 = input{A: "a2"}
	)

	type args struct {
		input interface{}
	}
	tests := []struct {
		name    string
		args    args
		want    []interface{}
		wantErr error
	}{
		{
			name: "Convert type array to interface array success",
			args: args{
				input: []input{i1, i2},
			},
			want:    []interface{}{i1, i2},
			wantErr: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ConvertTypeArrayToInterfaceArray(tt.args.input)
			if (err != nil) && err != tt.wantErr {
				t.Errorf("ConvertTypeArrayToInterfaceArray() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ConvertTypeArrayToInterfaceArray() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCheckArrIncludeItem(t *testing.T) {
	type args struct {
		arr  []interface{}
		item interface{}
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "Check array include item success",
			args: args{
				arr:  []interface{}{"1", "2"},
				item: "1",
			},
			want: true,
		},
		{
			name: "Check array include item fail",
			args: args{
				arr:  []interface{}{"1", "3"},
				item: "2",
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CheckArrIncludeItem(tt.args.arr, tt.args.item); got != tt.want {
				t.Errorf("CheckArrIncludeItem() = %v, want %v", got, tt.want)
			}
		})
	}
}
