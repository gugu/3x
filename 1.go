package main
import (
    "fmt"
    "time"
    "lukechampine.com/uint128"
)

func main() {
    iterated := uint128.From64(^uint64(0))
    printLimit := 0;
    t := time.Now()
    for {
        x := iterated
        for {
            if (x.Cmp(iterated) == -1) {
                break;
            }
            if (x.Mod64(2) == 1 ) {
                x = x.Mul64(3).Add64(1);
            }
            x = x.Rsh(1)
        }
        iterated = iterated.Add64(1)
        printLimit += 1;
        if (printLimit == 100_000_000) {
            fmt.Println(iterated, time.Now().Sub(t), time.Now())
            t = time.Now()
            printLimit = 0;
        }
    }
    // = 9789690303392599179035;
    /*
    let c = 9789690303392599179035n;
    let i = 0;
    while (true) {
        c = c + BigInt((Math.random() * 4000000000).toFixed());
        x = c;
        i ++;
        if (i % 500000 === 0)
            console.log(x, new Date());
        let l = 0;
        while (true) {
            if (x <= MAX_CHECKED) {
                break;
            }
            l ++;
            if (x === 0n) {
                console.log('0!!!!');
                process.exit(1);
            }
            if (x % 2n === 1n) {
                x = x * 3n + 1n;
            }
            if (x % 2n === 0n) {
                x = x / 2n;
            }
            if (l > 50000) {
                console.log('INF!', c, x, l)
            }
        }
    }*/
}
