**About** 📋 <br/>

A parking lot system that can hold up to ‘n’ cars at any given point in time. Each slot is given a number starting at one increasing with increasing distance from the entry point in steps of one. We want to create an automated ticketing system that allows our customers to use our parking lot without human intervention. <br/> <br/>
When a car enters the parking lot, we want to have a ticket issued to the driver. The ticket issuing process includes:-  <br/> <br/>
-> We are taking note of the number written on the vehicle registration plate and the age of the driver of the car. <br/>
-> And we are allocating an available parking slot to the car before actually handing over a ticket to the driver (we assume that our customers are kind enough to always park in the slots allocated to them). <br/> <br/>
The customer should be allocated a parking slot that is nearest to the entry. At the exit, the customer returns the ticket, marking the slot they were using as being available. <br/> <br/>
Due to government regulation, the system should provide us with the ability to find out:- <br/> <br/>
a. *Vehicle Registration numbers for all cars which are parked by the driver of a certain age.* <br/>
b. *Slot number in which a car with a given vehicle registration plate is parked.* <br/>
c. *Slot numbers of all slots where cars of drivers of a particular age are parked.* <br/>

**Approach** 💡  <br/>

* Used hashmap to keep track of slots and cars parked.
* Used min heap to get nearest available slot everytime in just O(log n) instead of O(n)
* Used hashmaps to get query results in constant/O(1) time instead of O(n) <br/>

**Steps To Run** 🏃🏻‍♀️  <br/>
1. Write commands in `input.txt` file
2. run `go run main.go` from root directory <br/>

**Commands** 🛠 <br/>

Description of each command from the sample input file:-<br/>

`Create_parking_lot 6` <br/>
*Create a parking lot of 6 slots*<br/>

`Park KA-01-HH-1234 driver_age 21`<br/>
*Park car with vehicle registration number “ KA-01-HH-1234”, and the vehicle is driven by the driver of age 21*<br/>

`Park PB-01-HH-1234 driver_age 21`<br/>
*Park car with vehicle registration number “ PB-01-HH-1234”, and the car is driven by the driver of age 21*<br/>

`Slot_numbers_for_driver_of_age 21`<br/>
*Return all Slot Number(Comma-separated) of all cars which have drivers with age==21*<br/>

`Park PB-01-TG-2341 driver_age 40`<br/>
*Park car with vehicle registration number “ PB-01-TG-2341”, and the car is driven by the driver of age 40*<br/>

`Slot_number_for_car_with_number PB-01-HH-1234`<br/>
*Return slot number for the car with registration number “PB-01-HH-1234”*<br/>

`Leave 2`<br/>
*Vacate the slot number 2 from the parking lot, i.e. car which was parked at slot number 2 has left the space if there exists no car at slot number 2, print “Slot already vacant”*<br/>

`Park HR-29-TG-3098 driver_age 39`<br/>
*Park car with vehicle registration number “ HR-29-TG-3098”, and the car is driven by the driver of age 39*<br/>

`Vehicle_registration_number_for_driver_of_age 18`<br/>
*Get all parked vehicle registration number of cars parked by the driver of age 18*<br/>
