package driver

import (
	"ParkingLot/pkg/parkinglot"
	"errors"
	"fmt"
	"strconv"
	"strings"
)

var parkingLotService parkinglot.ParkingLotService

func RunQuery(query string) {

	queryType, request, err := parseQuery(query)
	if err != nil {
		fmt.Println(err)
		return
	}

	switch queryType {
	case Create_parking_lot:

		req := request.(NewParkingLotRequest)

		parkingLot := &parkinglot.ParkingLot{
			Size: req.Size,
		}

		parkingLotService = parkinglot.New(*parkingLot)

		fmt.Println("Created parking of " + strconv.Itoa(req.Size) + " slots")

	case Park:

		req := request.(ParkRequest)

		slot, err := parkingLotService.Park(req.VehicleRegistrationNumber, req.Age)

		if err != nil {
			fmt.Println(err)
			return
		}

		fmt.Println("Car with vehicle registration number " + req.VehicleRegistrationNumber + " has been parked at slot number " + strconv.Itoa(slot))

	case Slot_numbers_for_driver_of_age:

		req := request.(GetSlotsForAgeRequest)

		slots, err := parkingLotService.GetSlotsForAge(req.Age)

		if err != nil {
			fmt.Println(err)
			return
		}

		for ind, slot := range slots {
			if ind == len(slots)-1 {
				fmt.Print(strconv.Itoa(slot))
			} else {
				fmt.Print(strconv.Itoa(slot) + ", ")
			}
		}

		fmt.Println()

	case Slot_number_for_car_with_number:

		req := request.(GetSlotForVehicleRegistrationNumberRequest)

		slot, err := parkingLotService.GetSlotForVehicleRegistrationNumber(req.VehicleRegistrationNumber)

		if err != nil {
			fmt.Println(err)
			return
		}

		fmt.Println(strconv.Itoa(slot))

	case Leave:

		req := request.(LeaveRequest)

		car, err := parkingLotService.Leave(req.SlotNumber)

		if err != nil {
			fmt.Println(err)
			return
		}

		message := fmt.Sprintf("Slot number %s vacated, the car with vehicle registration number %s left the space, the driver of the car was of age %s", strconv.Itoa(req.SlotNumber), car.VehicleRegistrationNumber, strconv.Itoa(car.Age))

		fmt.Println(message)

	case Vehicle_registration_number_for_driver_of_age:

		req := request.(GetRegNumbersForAgeRequest)

		regNumbers, err := parkingLotService.GetRegNumbersForAge(req.Age)

		if err != nil {
			fmt.Println(err)
			return
		}

		for ind, regNumber := range regNumbers {

			if ind == len(regNumbers)-1 {
				fmt.Print(regNumber)
			} else {
				fmt.Print(regNumber + ", ")
			}

		}
		fmt.Println()

	default:
		fmt.Println("Invalid command")

	}

}

func parseQuery(query string) (string, interface{}, error) {

	var invalidQueryMessage string = "invalid query"
	var invalidQueryArgumentsMessage string = "invalid query arguments"
	words := strings.Fields(query)

	if len(words) == 0 {
		return "", nil, errors.New(invalidQueryMessage)
	}
	switch words[0] {
	case Create_parking_lot:

		size, err := strconv.Atoi(words[1])

		if err != nil {
			return "", nil, errors.New(invalidQueryArgumentsMessage)
		}

		req := NewParkingLotRequest{
			Size: size,
		}

		return Create_parking_lot, req, nil

	case Park:

		if len(words) < 4 {
			return "", nil, errors.New(invalidQueryArgumentsMessage)
		}

		regNumber := words[1]
		age, err := strconv.Atoi(words[3])

		if err != nil {
			return "", nil, errors.New(invalidQueryArgumentsMessage)
		}

		req := ParkRequest{
			VehicleRegistrationNumber: regNumber,
			Age:                       age,
		}

		return Park, req, nil

	case Slot_numbers_for_driver_of_age:

		if len(words) < 2 {
			return "", nil, errors.New(invalidQueryArgumentsMessage)
		}

		age, err := strconv.Atoi(words[1])

		if err != nil {
			return "", nil, errors.New(invalidQueryArgumentsMessage)
		}

		req := GetSlotsForAgeRequest{
			Age: age,
		}

		return Slot_numbers_for_driver_of_age, req, nil

	case Slot_number_for_car_with_number:

		if len(words) < 2 {
			return "", nil, errors.New(invalidQueryArgumentsMessage)
		}

		regNumber := words[1]

		req := GetSlotForVehicleRegistrationNumberRequest{
			VehicleRegistrationNumber: regNumber,
		}

		return Slot_number_for_car_with_number, req, nil

	case Leave:

		if len(words) < 2 {
			return "", nil, errors.New(invalidQueryArgumentsMessage)
		}

		slotNumber, err := strconv.Atoi(words[1])

		if err != nil {
			return "", nil, errors.New(invalidQueryArgumentsMessage)
		}

		req := LeaveRequest{
			SlotNumber: slotNumber,
		}

		return Leave, req, nil

	case Vehicle_registration_number_for_driver_of_age:

		if len(words) < 2 {
			return "", nil, errors.New(invalidQueryArgumentsMessage)
		}

		age, err := strconv.Atoi(words[1])

		if err != nil {
			return "", nil, errors.New(invalidQueryArgumentsMessage)
		}

		req := GetRegNumbersForAgeRequest{
			Age: age,
		}

		return Vehicle_registration_number_for_driver_of_age, req, nil

	default:
		fmt.Println(invalidQueryMessage)
	}

	return "", nil, nil

}
