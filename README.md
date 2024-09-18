# Time Parser

`timep` is a Go package for parsing and working with extended time durations. It supports common time units like years (`y`), months (`M`), weeks (`w`), days (`d`), hours (`h`), and so on. Additionally, it allows defining custom time units for more flexible duration parsing.

## Installation

To install the `timep` package, run:

```bash
go get github.com/go-shafaq/timep
```

## Usage

### Parsing Common Durations

Here's how to parse a duration string with common units:

```go
package main

import (
	"fmt"
	"github.com/go-shafaq/timep"
)

func main() {
	dur, _ := timep.ParseDuration("1y2M3w4d5h6m")
	fmt.Println(dur)
}

```

**Input**: `"1y2M3w4d5h6m"`  
**Output**: `8975h5m0s`

### Defining Custom Units

You can define custom time units to suit your needs. For example, to define a "season" as 90 days:

```go
package main

import (
	"fmt"
	"time"
	"github.com/go-shafaq/timep"
)

func main() {
	const day = uint64(24 * time.Hour)

	um := timep.BaseUnitMap() // Clone a base unit map
	um["season"] = 90 * day   // Define custom unit "season" as 90 days

	dp, _ := timep.DurationParser(um) // Create a duration parser with the custom units

	dur, _ := dp.ParseDuration("3season")
	fmt.Println("Parsed Custom Duration:", dur)
}

```

**Input**: `"3season"`  
**Output**: `6480h0m0s`

## Time Unit Definitions

The supported standard time units are:

- `y` = Year (365 days)
- `M` = Month (30 days)
- `w` = Week (7 days)
- `d` = Day (24 hours)
- `h` = Hour
- `m` = Minute
- `s` = Second


If you have any issues or need further assistance, just let me know!
