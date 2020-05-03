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
