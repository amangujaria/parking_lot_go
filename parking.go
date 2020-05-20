package main

import (
    "os"
    "bufio"
    "strings"
    "strconv"
)


func main() {
    arg := os.Args[1]
    file, err := os.Open(arg)
    defer file.Close()
    if err != nil {
        return
    }
    reader := bufio.NewReader(file)

    var line string
    for {
        line, err = reader.ReadString('\n')

        read_line := strings.TrimSuffix(line, "\n")
        // Process the line here.
        split := strings.Split(read_line, " ")
        process(split)

        if err != nil {
            break
        }
    }
}

func process(tokens []string) {
    switch tokens[0] {
    case "create_parking_lot":
        size, _ := strconv.Atoi(tokens[1])
        spots.create_parking_lot(size)
    case "park":
        registration := tokens[1]
        color := tokens[2]
        spots.park(registration, color)
    case "leave":
        slot, _ := strconv.Atoi(tokens[1])
        spots.leave(slot)
    case "status":
        status()
    case "registration_numbers_for_cars_with_colour":
        spots.vehicles_with_color(tokens[1])
    case "spot_numbers_for_cars_with_colour":
        spots.spots_with_color(tokens[1])
    }
} 
