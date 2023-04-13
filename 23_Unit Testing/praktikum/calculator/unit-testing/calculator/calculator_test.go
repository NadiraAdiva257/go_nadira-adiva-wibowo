package calculator

import (
	"testing"
)

func Test_addition(t *testing.T) {
	type args struct {
		angka1 int
		angka2 int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "success",
			args: args{
				angka1: 1,
				angka2: 2,
			},
			want: 3,
		},
		{
			name: "fail",
			args: args{
				angka1: 1,
				angka2: 2,
			},
			want: 4,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := addition(tt.args.angka1, tt.args.angka2); got != tt.want {
				t.Errorf("addition() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_subtraction(t *testing.T) {
	type args struct {
		angka1 int
		angka2 int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "success",
			args: args{
				angka1: 5,
				angka2: 4,
			},
			want: 1,
		},
		{
			name: "fail",
			args: args{
				angka1: 5,
				angka2: 4,
			},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := subtraction(tt.args.angka1, tt.args.angka2); got != tt.want {
				t.Errorf("subtraction() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_multiplication(t *testing.T) {
	type args struct {
		angka1 int
		angka2 int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "success",
			args: args{
				angka1: 3,
				angka2: 7,
			},
			want: 21,
		},
		{
			name: "fail",
			args: args{
				angka1: 3,
				angka2: 7,
			},
			want: 18,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := multiplication(tt.args.angka1, tt.args.angka2); got != tt.want {
				t.Errorf("multiplication() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_division(t *testing.T) {
	type args struct {
		angka1 int
		angka2 int
	}
	tests := []struct {
		name    string
		args    args
		want    int
		wantErr bool
	}{
		{
			name: "success",
			args: args{
				angka1: 8,
				angka2: 2,
			},
			want: 4,
		},
		{
			name: "fail",
			args: args{
				angka1: 8,
				angka2: 2,
			},
			want: 5,
		},
		{
			name: "success",
			args: args{
				angka1: 10,
				angka2: 4,
			},
			want: 2,
		},
		{
			name: "fail",
			args: args{
				angka1: 10,
				angka2: 4,
			},
			want: 3,
		},
		{
			name: "fail divide by zero",
			args: args{
				angka1: 3,
				angka2: 0,
			},
			want:    0,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := division(tt.args.angka1, tt.args.angka2)
			if (err != nil) != tt.wantErr {
				t.Errorf("division() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("division() = %v, want %v", got, tt.want)
			}
		})
	}
}
