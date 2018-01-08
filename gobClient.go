package client

import (
    "fmt"
    "log"
    "net/rpc"

    "bitbucket.org/HelgeCPH/si_rpc/contract"
)

const port = 1234

func CreateClient() *rpc.Client {
    client, err := rpc.Dial("tcp", fmt.Sprintf("192.168.20.2:%v", port))
    if err != nil {
        log.Fatal("Cannot call server:", err)
    }

    return client
}

func PerformRequest(client *rpc.Client) contract.DayRPCResponse {
    args := &contract.DayRPCRequest{Qualifier: "Yesterday"}
    var reply contract.DayRPCResponse

    err := client.Call("DayRPCHandler.GetDay", args, &reply)
    if err != nil {
        log.Fatal("error:", err)
    }

    return reply
}