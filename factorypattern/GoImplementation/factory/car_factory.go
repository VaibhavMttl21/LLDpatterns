package factory

import "factorypattern/vehicle"

type CarFactory struct{}

func (f CarFactory) CreateVehicle() vehicle.Vehicle {
    return vehicle.Car{}
}