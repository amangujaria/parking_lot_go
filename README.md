# parking_lot_go
parking lot functionality in Go


**Usage**
- `./start.sh`
    - for terminal input
- `./start.sh Filename`
    - for file input (files under testfiles/ can be tried)

Exposed utilities (inputs from terminal/file):
- create_parking_lot Size where Size is number of parking spots inside the parking lot
- park Registration Colour
    - triggers a parking event for a vehicle with a certain registration number and colour
- leave SpotId
    - triggers a departure event for a vehicle parked at spot represented by SpotId
- status
    - represents the status of parked vehicles in a tabular form
- registration_numbers_for_cars_with_colour Colour
    - gives regisration numbers of all parked vehicles having a certain colour
- slot_numbers_for_cars_with_colour Colour
    - gives a list containing parking spot ids of all parked vehicles having a certain colour
- slot_number_for_registration_number Registration
    - gives the spot id where a vehicle having a certain registration number is parked at
