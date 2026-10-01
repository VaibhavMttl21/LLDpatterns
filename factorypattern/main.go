// Question is to make a vehicle library , that client say create vehicle and we will give object
// first thinking simple class and inherit two classes
// package main

// import (
// 	"designpatterns/factory"
// 	// "designpatterns/vehicle"
// )

// func main(){
// 	// car := vehicle.Car{}
// 	// car.Create()

// 	// bike := vehicle.Bike{}
// 	// bike.Create()

// 	car := factory.CreateVehicle("car")
// 	car.Create()
// }

// PRODUCT SIDE

//              Vehicle
//              interface
//              Drive()
//                 ▲
//                 │
//         ┌───────┴───────┐
//         │               │
//        Car             Bike
//       Drive()          Drive()


// FACTORY SIDE

//           VehicleFactory
//              interface
//        CreateVehicle()
//                 ▲
//                 │
//         ┌───────┴───────┐
//         │               │
//    CarFactory       BikeFactory
//  CreateVehicle()   CreateVehicle()
package main

import (
    "fmt"
    "factorypattern/factory"
)

func main() {
   // Create a CarFactory, but store/access it through the VehicleFactory interface.
    var carFactory factory.VehicleFactory = factory.CarFactory{}

    car := carFactory.CreateVehicle()

    car.Create()

    fmt.Println("-----")

    var bikeFactory factory.VehicleFactory = factory.BikeFactory{}

    bike := bikeFactory.CreateVehicle()

    bike.Create()
}