package main

import (
    "fmt"
    "github.com/looplab/fsm"
)

type spot struct {
    fsm *fsm.FSM
    channel chan string
    ack chan bool
    registration string
    color string
}

func leave(fsm *fsm.FSM) {
    err := fsm.Event("leave")
    if err != nil {
        fmt.Println(err)
    }
}

func arrive(fsm *fsm.FSM) {
    err := fsm.Event("arrive")
    if err != nil {
        fmt.Println(err)
    }
}

func receive_events(fsm *fsm.FSM, channel <-chan string, ack chan<- bool) {
    for {
        select {
            case event := <- channel :
                if event == "arrive" {
                    arrive(fsm)
                    ack <- true
                } else if event == "leave" {
                    leave(fsm)
                    ack <- true
                }
        }
    }

}

func create_spot() *spot {
    fsm := fsm.NewFSM(
        "unoccupied",
        fsm.Events{
            {Name: "leave", Src: []string{"occupied"}, Dst: "unoccupied"},
            {Name: "arrive", Src: []string{"unoccupied"}, Dst: "occupied"},
        },
        fsm.Callbacks{},
    )
    channel := make(chan string, 1)
    ack := make(chan bool, 1)
    go receive_events(fsm, channel, ack)
    return &spot{fsm: fsm, channel: channel, ack: ack}
}

func activity(id int, spots map[int]*spot, event string) {
    (*spots[id]).channel <- event
    <- (*spots[id]).ack
}

func modify_attributes(id int, spots map[int]*spot, registration string, color string) {
    (*spots[id]).color = color
    (*spots[id]).registration = registration
}
