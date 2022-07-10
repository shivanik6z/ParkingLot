package parkinglot

import (
	"ParkingLot/pkg/car"
	"ParkingLot/pkg/util"
	"container/heap"
)

type ParkingLot struct {
	ParkingLotService
	Size               int
	Lots               *map[int]car.Car
	AvailableSlots     *util.IntHeap
	SlotsByAgeMap      *map[int][]int
	SlotByRegNumberMap *map[string]int
}

type ParkingLotService interface {
	//Updates lots, SlotsByAgeMap, SlotByRegNumberMap and AvailableSlots. Parks the car at minimum(nearest) available slot
	Park(vehicleRegistrationNumber string, age int) (slot int, err error)

	//Returns value from SlotsByAgeMap
	GetSlotsForAge(age int) (slots []int, err error)

	//Returns value from SlotByRegNumberMap
	GetSlotForVehicleRegistrationNumber(vehicleRegistrationNumber string) (slot int, err error)

	//Deletes car from maps and pushes available slot to AvailableSlots(min heap)
	Leave(slot int) (car car.Car, err error)

	//Uses GetSlotsForAge() to return registration numbers
	GetRegNumbersForAge(age int) (regNumbers []string, err error)
}

// Creates a new Parking Lot
func New(parkingLot ParkingLot) ParkingLotService {

	size := parkingLot.Size

	lots := make(map[int]car.Car, 0)

	availableSlots := &util.IntHeap{}
	heap.Init(availableSlots)
	for i := 1; i <= size; i++ {
		heap.Push(availableSlots, i)
	}

	slotsByAgeMap := make(map[int][]int)

	slotByRegNumberMap := make(map[string]int)

	return &ParkingLot{
		Size:               size,
		Lots:               &lots,
		AvailableSlots:     availableSlots,
		SlotsByAgeMap:      &slotsByAgeMap,
		SlotByRegNumberMap: &slotByRegNumberMap,
	}
}
