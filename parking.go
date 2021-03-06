package main

import (
    "fmt"
    "os"
    "sync"
    "bufio"
    "strings"
    "strconv"
)

type mymap map[int]*spot
var spots = make(mymap)
var mutex = &sync.Mutex{}

func get_file() *os.File {
    if len(os.Args) > 1 {  
        arg := os.Args[1]
        file, err := os.Open(arg)
        if err != nil {
            panic(err)
        }
        return file
    } else {
        return os.Stdin
    }
}

func main() {
    file := get_file()
    defer file.Close()
    reader := bufio.NewReader(file)
    for {
        line, err := reader.ReadString('\n')
        read_line := strings.TrimSuffix(line, "\n")
        // Process the line here.
        split := strings.Split(read_line, " ")
        process(split)
        if err != nil {
            break
        }
    }
}

func process(tokens []string) (err error) {
    defer func() {
      if r := recover(); r != nil {
         err = r.(error)
      }
    }()
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
        regs := spots.vehicles_with_color(tokens[1])
        fmt.Println(regs)
    case "slot_numbers_for_cars_with_colour":
        spotIds := spots.spots_with_color(tokens[1])
        fmt.Println(spotIds)
    case "slot_number_for_registration_number":
        spotId, _ := spots.find_spot(tokens[1])
        fmt.Println(spotId)
    default:
        fmt.Println("incorrect operation")
    }
    return
} 
