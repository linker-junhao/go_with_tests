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
	checkShape := func(t *testing.T, shape Shape, want float64) {
		t.Helper()
		got := shape.Area()

		if got != want {
			t.Errorf("got %.2f want %.2f", got, want)
		}
	}

	t.Run("rectangle", func(t *testing.T) {
		rectangle := Rectangle{
			Width:  12,
			Height: 6,
		}
		want := 72.0

		checkShape(t, rectangle, want)
	})

	t.Run("circles", func(t *testing.T) {
		circle := Circle{Radius: 10}
		want := 314.1592653589793

		checkShape(t, circle, want)
	})
}
