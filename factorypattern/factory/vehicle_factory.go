// in this still we have to modify this if we want to add more
//Every time we add a new vehicle, we have to modify this factory.
// That is exactly what OCP warns us about.

// soln is to make a factory interface
// package factory

// import "designpatterns/vehicle"

// func CreateVehicle(vehicleType string) vehicle.Vehicle{
// 	switch vehicleType{
// 	case "car":
// 		return vehicle.Car{}
// 	case "bike":
// 		return vehicle.Bike{}
// 	default:
// 		return nil
// 	}
// }

package factory

import "factorypattern/vehicle"

// Every vehicle factory must have a CreateVehicle() method, and that method must return a Vehicle.
// vehicle.Vehicle this means Anything that I call a Vehicle
// must provide Drive()
// var v vehicle.Vehicle v can hold any value whose type satisfies the Vehicle interface.
type VehicleFactory interface {
    CreateVehicle() vehicle.Vehicle
}