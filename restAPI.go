package main

import (
    "fmt"
    "io"
    "io/ioutil"
    "log"
    "net/http"
    "strconv"
)

type Stack struct {
    top  *Node
    size int
}

type Node struct {
    value int
    next  *Node
}

func (s *Stack) Peek() int {
    return s.top.value
}

func (s *Stack) Push(val int) {
    s.top = &Node{val, s.top}
    s.size++
}

func (s *Stack) Pop() (val int) {
    if s.size > 0 {
        val, s.top = s.top.value, s.top.next
        s.size--
        return val
    }
    return 0
}

func getNumberFrom(requestBody io.ReadCloser) int {
    body, err := ioutil.ReadAll(requestBody)
    if err != nil {
        log.Println(err.Error())
    }
    number, err := strconv.Atoi(string(body))
    if err != nil {
        log.Println("Ups... cannot get the number")
    }
    return number
}

func pushNumber(res http.ResponseWriter, req *http.Request) {
    if req.Method == "POST" {
        // Test for example with:
        // curl -X POST http://192.168.20.2:8080/push -d '5'
        number := getNumberFrom(req.Body)
        stack.Push(number)
        fmt.Fprint(res, stack.Peek())
    }
}

func addNumber(res http.ResponseWriter, req *http.Request) {
    number := getNumberFrom(req.Body)
    stack.Push(number)
    result := stack.Pop() + stack.Pop()
    stack.Push(result)
    fmt.Fprint(res, result)
}

func subNumber(res http.ResponseWriter, req *http.Request) {
    number := getNumberFrom(req.Body)
    stack.Push(number)
    result := stack.Pop() - stack.Pop()
    stack.Push(result)
    fmt.Fprint(res, result)
}

var stack *Stack

func main() {
    stack = new(Stack)
    stack.Push(0)

    http.HandleFunc("/push", pushNumber)
    http.HandleFunc("/add", addNumber)
    http.HandleFunc("/sub", subNumber)

    http.ListenAndServe(":8080", nil)
}