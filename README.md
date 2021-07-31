# [jzodiac] -- Japanese Zodiac

[![check vulns](https://github.com/spiegel-im-spiegel/jzodiac/workflows/vulns/badge.svg)](https://github.com/spiegel-im-spiegel/jzodiac/actions)
[![lint status](https://github.com/spiegel-im-spiegel/jzodiac/workflows/lint/badge.svg)](https://github.com/spiegel-im-spiegel/jzodiac/actions)
[![GitHub license](https://img.shields.io/badge/license-Apache%202-blue.svg)](https://raw.githubusercontent.com/spiegel-im-spiegel/jzodiac/master/LICENSE)
[![GitHub release](http://img.shields.io/github/release/spiegel-im-spiegel/jzodiac.svg)](https://github.com/spiegel-im-spiegel/jzodiac/releases/latest)

## Usage

```go
package jzodiac_test

import (
    "fmt"
    "time"

    "github.com/spiegel-im-spiegel/jzodiac"
)

func ExampleZodiacDayNumber() {
    t := time.Date(2021, time.July, 28, 0, 0, 0, 0, jzodiac.JST)
    kan, shi := jzodiac.ZodiacDayNumber(t)
    fmt.Printf("%v is %v%v\n", t.Format("2006-01-02"), kan, shi)
    // Output:
    // 2021-07-28 is 丁丑
}

func ExampleZodiacYearNumber() {
    kan, shi := jzodiac.ZodiacYearNumber(2021)
    fmt.Printf("Year %v is %v%v\n", 2021, kan, shi)
    // Output:
    // Year 2021 is 辛丑
}
```

[jzodiac]: https://github.com/spiegel-im-spiegel/jzodiac "spiegel-im-spiegel/jzodiac: Japanese Zodiac"
