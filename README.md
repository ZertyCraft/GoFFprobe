# GoFFprobe

GoFFprobe is a Go package providing a simple and efficient way to interact with `ffprobe` for retrieving media file information.

## Installation

To use GoFFprobe, you need to have `ffprobe` installed on your system. `ffprobe` is part of the `ffmpeg` package.

After ensuring that `ffprobe` is installed, you can add GoFFprobe to your project by running:

```
go get github.com/ZertyCraft/GoFFprobe
```

## Usage

Here's a simple example of how to use GoFFprobe:

```go
package main

import (
    "fmt"
    "github.com/ZertyCraft/GoFFprobe"
)

func main() {
    // Example file path
    filePath := "path/to/your/media/file"

    // Get media information
    info, err := GoFFprobe.Execute(filePath, GoFFprobe.Options{
        ShowFormat: true,
        ShowStreams: true,
    })

    if err != nil {
        fmt.Println("Error:", err)
        return
    }

    fmt.Printf("Media Information: %+v\n", info)
}
```

## Features

- Retrieve media file information using `ffprobe`.
- Flexible options to customize the `ffprobe` command.
- Parses `ffprobe` output to a user-friendly format.
