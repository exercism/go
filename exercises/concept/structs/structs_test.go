package structs

import "testing"

func TestNewCar(t *testing.T){
	tests := []struct{
		name string
		car Car
		expected Car
	}{
		{
			name: "Create a new car.",
			car: Car{
				speed: 5,
				batteryDrain: 2,
				battery: 100, 
			},
			expected: Car{
				speed: 5,
				batteryDrain: 2,
				battery: 100,
			},

		},
	}

	for _,tt :=range tests{
		t.Run(tt.name, func(t *testing.T){
			got := NewCar(tt.car.speed, tt.car.batteryDrain)

			if got!=tt.expected{
				t.Errorf("NewCar(%v,%v) = %v; expected %v", tt.car.speed, tt.car.batteryDrain,got, tt.expected)
			}
		})
	}

}


func TestNewTrack(t *testing.T){
	tests :=[]struct{
		name string
		track Track
		expected Track
	}{
		{
			name: "Create a new track.",
			track: Track{
				distance: 800,
			},
			expected: Track{
				distance: 800,
			},
		},
		{
			name: "Create a new track.",
			track: Track{
				distance: 360,
			},
			expected: Track{
				distance: 360,
			},
		},
	}


	for _,tt :=range tests{
		t.Run(tt.name, func(t *testing.T){
			got := NewTrack(tt.track.distance)

			if got!=tt.expected{
				t.Errorf("NewTrack(%v) = %v; expected %v", tt.track.distance, tt.track, tt.expected)
			}
		})
	}
}

func TestDrive(t *testing.T){
	tests := []struct{
		name string
		car Car
		expected Car
	}{
		{
			name: "Drive the car with full battery.",
			car: Car{
				speed: 5,
				batteryDrain: 2,
				battery: 100, 
				distance: 0,
			},
			expected: Car{
				speed: 5,
				batteryDrain: 2,
				battery: 98,
				distance: 5,
			},

		},
		{
			name: "Drive the car with battery drained.",
			car: Car{
				speed: 5,
				batteryDrain: 5,
				battery: 4, 
				distance: 0,
			},
			expected: Car{
				speed: 5,
				batteryDrain: 5,
				battery: 4,
				distance: 0,
			},

		},
	}

	for _,tt :=range tests{
		t.Run(tt.name, func(t *testing.T){
			got := Drive(tt.car)

			if got!=tt.expected{
				t.Errorf("NewCar(%v) = %v; expected %v", tt.car, got, tt.expected)
			}
		})
	}
}


func TestCanFinish(t *testing.T){
	tests := []struct{
		name string
		car Car
		track Track
		expected bool
	}{
		{
			name: "Car can easily finish the track distance.",
			car: Car{
				speed: 5,
				batteryDrain: 2,
				battery: 100, 
				distance: 20,
			},
			track: Track{
				distance: 30,
			},
			expected: true,
		},
		{
			name: "Car can finish the track distance just before battery drained.",
			car: Car{
				speed: 5,
				batteryDrain: 5,
				battery: 100, 
				distance: 20,
			},
			track: Track{
				distance: 20,
			},
			expected: true,
		},
		{
			name: "Car cannot finish the track distrance.",
			car: Car{
				speed: 2,
				batteryDrain: 10,
				battery: 100, 
				distance: 5,
			},
			track: Track{
				distance: 80,
			},
			expected: false,
		},
	}

	for _,tt :=range tests{
		t.Run(tt.name, func(t *testing.T){
			got := CanFinish(tt.car, tt.track)

			if got!=tt.expected{
				t.Errorf("NewCar(%v, %v) = %v; expected %v", tt.car,tt.track, got, tt.expected)
			}
		})
	}
}
