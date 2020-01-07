package structAndInterface

import "testing"

func TestPerimeter(t *testing.T) {
	type args struct {
		rectangle Rectangle
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{
			name: "obj",
			args: args{rectangle: Rectangle{
				Width:  10.0,
				Height: 10.0,
			}},
			want: 40,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Perimeter(tt.args.rectangle); got != tt.want {
				t.Errorf("Perimeter() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestArea(t *testing.T) {

	areaTest := []struct {
		shape Shape
		want  float64
	}{
		{shape: Rectangle{Width: 12, Height: 6}, want: 72.0},
		{shape: Circle{Radius: 10}, want: 314.1592653589793},
		{shape: Triangle{Base: 12, Height: 6}, want: 36.0},
	}

	for _, tt := range areaTest {
		got := tt.shape.Area()
		if got != tt.want {
			t.Errorf("%#v got %.2f want %.2f", tt.shape, got, tt.want)
		}
	}
}
