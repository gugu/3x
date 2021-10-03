package main
import (
    "fmt"
    "time"
    "runtime"
    "lukechampine.com/uint128"
)

const THREAD_CHUNK = 100_000_000

func checkNumber(iterated uint128.Uint128, guard chan struct{}) {
    if (iterated.Mod64(2) == 0) {
        iterated = iterated.Add64(1)
    }
    for i := iterated; i.Cmp(iterated.Add64(THREAD_CHUNK)) == -1; i = i.Add64(2) {
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
    // iterated := uint128.From64(uint64(2))
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
