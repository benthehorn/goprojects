package server

import (
    "errors"
    "fmt"
    "log"
    "net"
    "net/rpc"
    "time"

    "bitbucket.org/HelgeCPH/si_rpc/contract"
)

const port = 1234

func main() {
    StartServer()
}

func StartServer() {
    log.Printf("Server starting on port %v\n", port)

    rpcFuncHandler := &DayRPCHandler{}
    rpc.Register(rpcFuncHandler)

    l, err := net.Listen("tcp", fmt.Sprintf(":%v", port))
    if err != nil {
        log.Fatal(fmt.Sprintf("Unable to listen on given port: %s", err))
    }
    defer l.Close()

    for {
        conn, _ := l.Accept()
        go rpc.ServeConn(conn)
    }
}

type DayRPCHandler struct{}

func (h *DayRPCHandler) GetDay(args *contract.DayRPCRequest, reply *contract.DayRPCResponse) error {

    daytime := time.Now()

    switch args.Qualifier {
    case "Yesterday":
        reply.Day = daytime.Add(-24 * time.Hour).Weekday().String()
    case "Today":
        reply.Day = daytime.Weekday().String()
    case "Tomorrow":
        reply.Day = daytime.Add(24 * time.Hour).Weekday().String()
    default:
        log.Print("Unrecognized qualifier")
        return errors.New("Unrecognized qualifier")
    }

    return nil
}