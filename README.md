# [jzodiac] -- Japanese Zodiac

[![check vulns](https://github.com/spiegel-im-spiegel/jzodiac/workflows/vulns/badge.svg)](https://github.com/spiegel-im-spiegel/jzodiac/actions)
[![lint status](https://github.com/spiegel-im-spiegel/jzodiac/workflows/lint/badge.svg)](https://github.com/spiegel-im-spiegel/jzodiac/actions)
[![GitHub license](https://img.shields.io/badge/license-Apache%202-blue.svg)](https://raw.githubusercontent.com/spiegel-im-spiegel/jzodiac/master/LICENSE)
[![GitHub release](http://img.shields.io/github/release/spiegel-im-spiegel/jzodiac.svg)](https://github.com/spiegel-im-spiegel/jzodiac/releases/latest)

## Usage

```go
// +build run

package main

import (
    "flag"
    "fmt"
    "os"
    "time"

    "github.com/spiegel-im-spiegel/jzodiac"
)

func main() {
    flag.Parse()
    args := flag.Args()
    if len(args) < 1 {
        fmt.Fprintln(os.Stderr, os.ErrInvalid)
        return
    }
    for _, s := range args {
        t, err := time.Parse("2006-01-02", s)
        if err != nil {
            fmt.Fprintln(os.Stderr, err)
            continue
        }
        kan, shi := jzodiac.ZodiacYearNumber(t.Year())
        fmt.Printf("Year %v is %v%v\n", t.Year(), kan, shi)
        kan, shi = jzodiac.ZodiacDayNumber(t)
        fmt.Printf("Day %v is %v%v\n", t.Format("2006-01-02"), kan, shi)
    }
}
```

[jzodiac]: https://github.com/spiegel-im-spiegel/jzodiac "spiegel-im-spiegel/jzodiac: Japanese Zodiac"
