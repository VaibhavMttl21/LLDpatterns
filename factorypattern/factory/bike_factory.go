package factory

import "factorypattern/vehicle"

type BikeFactory struct{}

func (f BikeFactory) CreateVehicle() vehicle.Vehicle {
    return vehicle.Bike{}
}