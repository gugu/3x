package main
import (
    "fmt"
    "time"
    "runtime"
    "lukechampine.com/uint128"
)

const THREAD_CHUNK = 10_000_000

func checkNumber(iterated uint128.Uint128, guard chan struct{}) {
    for i := iterated; i.Cmp(iterated.Add64(10_000_000)) == -1; i = i.Add64(1) {
        x := i
        for {
            if (x.Cmp(iterated) == -1) {
                break;
            }
            if (x.Mod64(2) == 1 ) {
                x = x.Mul64(3).Add64(1);
            }
            x = x.Rsh(1)
        }
    }
    <- guard
}

func main() {
    iterated := uint128.From64(^uint64(0))
    printLimit := 0;
    t := time.Now()
    guard := make(chan struct{}, runtime.NumCPU())

    for {
        guard <- struct {}{}
        go checkNumber(iterated, guard)
        iterated = iterated.Add64(THREAD_CHUNK)
        printLimit += THREAD_CHUNK;
        if (printLimit == 1000_000_000) {
            fmt.Println(iterated, time.Now().Sub(t), time.Now())
            t = time.Now()
            printLimit = 0;
        }
    }
}
