package main

import (
    "fmt"
)

func main(){
    var min_uint = 0
    var max_uint8 uint8 = ^uint8(0)
    var max_uint16 uint16 = ^uint16(0)
    var max_uint32 uint32 = ^uint32(0)
    var max_uint64 uint64 = ^uint64(0)

    var max_int8 int8 = int8(max_uint8>>1)
    var min_int8 int8 = -max_int8 - 1
    var max_int16 int16 = int16(max_uint16>>1)
    var min_int16 int16 = -max_int16 - 1
    var max_int32 int32 = int32(max_uint32>>1)
    var min_int32 int32 = -max_int32 - 1
    var max_int64 int64 = int64(max_uint64>>1)
    var min_int64 int64 = -max_int64 - 1

    fmt.Println("uint8 -> ", min_uint, " to ", max_uint8)
    fmt.Println("uint16 -> ", min_uint, " to ", max_uint16)
    fmt.Println("uint32 -> ", min_uint, " to ", max_uint32)
    fmt.Println("uint64 -> ", min_uint, " to ", max_uint64)
    fmt.Println("")
    fmt.Println("int8 -> ", min_int8, " to ", max_int8)
    fmt.Println("int16 -> ", min_int16, " to ", max_int16)
    fmt.Println("int32 -> ", min_int32, " to ", max_int32)
    fmt.Println("int64 -> ", min_int64, " to ", max_int64)
}
