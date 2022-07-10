package driver

const Create_parking_lot string = "Create_parking_lot"
const Park string = "Park"
const Slot_numbers_for_driver_of_age string = "Slot_numbers_for_driver_of_age"
const Slot_number_for_car_with_number string = "Slot_number_for_car_with_number"
const Leave string = "Leave"
const Vehicle_registration_number_for_driver_of_age string = "Vehicle_registration_number_for_driver_of_age"

type NewParkingLotRequest struct {
	Size int
}

type ParkRequest struct {
	VehicleRegistrationNumber string
	Age                       int
}

type GetSlotsForAgeRequest struct {
	Age int
}

type GetSlotForVehicleRegistrationNumberRequest struct {
	VehicleRegistrationNumber string
}

type LeaveRequest struct {
	SlotNumber int
}

type GetRegNumbersForAgeRequest struct {
	Age int
}
