package main

import (
    "fmt"
)

func (spots *mymap) create_parking_lot(size int) {
    mutex.Lock()
    defer mutex.Unlock()
    start := len(*spots)
    end := len(*spots) + size
    for i := start; i < end; i++ {
        (*spots)[i] = create_spot()
    }
}

func (spots *mymap) find_vacant() int {
    for i := 0; i < len(*spots); i++ {
        k := (*spots)[i]
        if (*k).fsm.Current() == "unoccupied" {
            return i
        }
    }
    return -1
}

func (spots *mymap) park(registration string, color string) {
    mutex.Lock()
    defer mutex.Unlock()
    pos := spots.find_vacant()
    activity(pos, *spots, "arrive")
    modify_attributes(pos, *spots, registration, color)
}

func (spots *mymap) leave(id int) {
    mutex.Lock()
    defer mutex.Unlock()
    activity(id, *spots, "leave")
    modify_attributes(id, *spots, "", "")
}

func (spots *mymap) find(color string) int {
    total := 0
    for i := 0; i < len(*spots); i++ {
        if (*spots)[i].color == color {
            total += 1
        }
    }
    return total
}

func (spots *mymap) find_spot(registration string) (int, string) {
    for i := 0; i < len(*spots); i++ {
        if (*spots)[i].registration == registration {
            return i, (*spots)[i].color
        }
    }
    return -1, "abnormal"
}

func (spots *mymap) vehicles_with_color(color string) []string {
    var results []string
    for i := 0; i < len(*spots); i++ {
        if (*spots)[i].color == color {
            results = append(results, (*spots)[i].registration)
        }
    }
    return results
}

func (spots *mymap) spots_with_color(color string) []int {
    var results []int
    for i := 0; i < len(*spots); i++ {
        if (*spots)[i].color == color {
            results = append(results, i)
        }
    }
    return results
}

func status() {
    fmt.Println("spot number\tregistration\t\tcolor")
    for i := 0; i < len(spots); i++ {
        if spots[i].fsm.Current() == "occupied" {
            fmt.Printf("%d\t\t%s\t\t%s\n", i, spots[i].registration, spots[i].color)
        }
    }
}
