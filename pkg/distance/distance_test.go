package distance

import "testing"

func TestCalculate(t *testing.T) {
	type args struct {
		x1 float64
		y1 float64
		x2 float64
		y2 float64
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{
			name: "0 0 3 4 then 5",
			args: args{x1: 0, y1: 0, x2: 3, y2: 4},
			want: 5,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Calculate(tt.args.x1, tt.args.y1, tt.args.x2, tt.args.y2); got != tt.want {
				t.Errorf("Calculate() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCalcWithGeo(t *testing.T) {
	type args struct {
		lat1 float64
		lng1 float64
		lat2 float64
		lng2 float64
		unit []string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "0 0 3 4 then 345",
			args: args{lat1: 0, lng1: 0, lat2: 3, lng2: 4},
			want: 345,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := CalcWithGeo(tt.args.lat1, tt.args.lng1, tt.args.lat2, tt.args.lng2, tt.args.unit...)
			if int(got) != tt.want {
				t.Errorf("CalcWithGeo() = %v, want %v", got, tt.want)
			}
		})
	}
}
