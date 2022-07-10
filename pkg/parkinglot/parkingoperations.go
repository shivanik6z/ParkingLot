package parkinglot

import (
	"ParkingLot/pkg/car"
	"container/heap"
	"errors"
)

func (parkingLot ParkingLot) Park(vehicleRegistrationNumber string, age int) (slot int, err error) {

	availableSlots := parkingLot.AvailableSlots
	lots := *parkingLot.Lots
	if len(*availableSlots) < 1 {
		errorMessage := "oops! parking lot has no more space to park your car having registration number " + vehicleRegistrationNumber
		return 0, errors.New(errorMessage)
	}

	availableSlot := heap.Pop(availableSlots).(int)

	car := car.Car{
		Age:                       age,
		VehicleRegistrationNumber: vehicleRegistrationNumber,
	}

	lots[availableSlot] = car

	SlotsByAgeMap := *parkingLot.SlotsByAgeMap
	SlotsByAgeMap[age] = append(SlotsByAgeMap[age], availableSlot)

	SlotByRegNumberMap := *parkingLot.SlotByRegNumberMap
	SlotByRegNumberMap[vehicleRegistrationNumber] = availableSlot

	return availableSlot, nil

}

func (parkingLot *ParkingLot) GetSlotsForAge(age int) (slots []int, err error) {

	SlotsByAgeMap := *parkingLot.SlotsByAgeMap
	slots = SlotsByAgeMap[age]
	return slots, nil
}

func (parkingLot *ParkingLot) GetSlotForVehicleRegistrationNumber(vehicleRegistrationNumber string) (slot int, err error) {

	SlotByRegNumberMap := *parkingLot.SlotByRegNumberMap
	if _, ok := SlotByRegNumberMap[vehicleRegistrationNumber]; !ok {
		errorMessage := "oops! couldn't find this car parked having vehicle number " + vehicleRegistrationNumber
		return 0, errors.New(errorMessage)
	}

	return SlotByRegNumberMap[vehicleRegistrationNumber], nil
}

func (parkingLot *ParkingLot) Leave(slot int) (car car.Car, err error) {

	lots := *parkingLot.Lots
	if _, ok := lots[slot]; !ok {

		errorMessage := "this slot is already empty"
		return car, errors.New(errorMessage)
	}

	car = lots[slot]

	delete(*parkingLot.SlotByRegNumberMap, car.VehicleRegistrationNumber)

	SlotsByAgeMap := *parkingLot.SlotsByAgeMap
	for i, carSlot := range SlotsByAgeMap[car.Age] {

		if carSlot == slot {
			a := SlotsByAgeMap[car.Age]
			a[i] = a[len(a)-1]
			a[len(a)-1] = -1
			SlotsByAgeMap[car.Age] = a[:len(a)-1]
		}
	}

	heap.Push(parkingLot.AvailableSlots, slot)

	return car, nil
}

func (parkingLot *ParkingLot) GetRegNumbersForAge(age int) (regNumbers []string, err error) {

	slots, err := parkingLot.GetSlotsForAge(age)

	if err != nil {
		return regNumbers, err
	}

	lots := *parkingLot.Lots
	for _, slot := range slots {

		regNumbers = append(regNumbers, lots[slot].VehicleRegistrationNumber)
	}

	return regNumbers, nil
}
