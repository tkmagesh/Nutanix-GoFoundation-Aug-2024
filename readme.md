# Go Foundation

## Magesh Kuppan
- tkmagesh77@gmail.com

## Schedule
| What | When |
|------|------|
|Commence | 09:30 AM |
|Tea Break | 11:00 AM (15 mins)|
|Lunch Break | 01:00 PM (1 hour) |
|Tea Break  | 03:15 PM (15 mins) |
|Wind up | 04:30 PM |

## Repository
- https://github.com/tkmagesh/nutanix-gofoundation-aug-2024

## Software Requirements
- Go Tools (https://go.dev/dl)
- Any Editor
- Git Client

## Methodology
- No powerpoint
- No dedicate time Q & A
- Pls unmute yourself and speak anytime

## Why Go?
- Simplicity
    - ONLY 25 keywords
    - No access modifiers (private/public/protected etc)
    - No classes (only structs)
    - No inheritance (only composition)
    - No reference types (everything is a value in Go)
    - No pointer arithmatic
    - No exceptions (ONLY errors which are just values)
    - No try..catch..finally construct
    - No implicit type conversion
- Cloud Native
- Performance
    - Close to hardware
    - Compile to native code
    - Support for cross compilation
- Concurrency
    - Managed Concurrency (vs OS Thread based)
    - Cheap
    - Concurrency support is built IN the language
        - "go" keyword, "channel" data type, channel "<-" operator, "range" & "select-case" constructs
        - standard library APIs
            - "sync" package
            - "sync/atomic" package

## Go Lang
### Data Types
- bool
- string
- integers
    - int8
    - int16
    - int32
    - int64
    - int
- unsigned integers
    - uint8
    - uint16
    - uint32
    - uint64
    - uint
- floating points
    - float32
    - float64
- complex numbers
    - complex64 (real[float32] + imaginary[float32])
    - complex128 (real[float64] + imaginary[float64])
- type alias
    - byte
    - rune (unicode code point)

#### Zero values

| type | value |
| -------|-------- |
| int, uint, float | 0 |
| string | "" |
| bool | false |
| func | nil |
| struct | struct instance |
| pointer |nil |
| interface | nil | 

## Build & Execute
```
go run [file_name.go]
```
## Create a build
```
go build [file_name.go]
```
## To list the env variables for the go tool
```
go env
```
```
go env [var_1] [var_2] ....
```
## Env variables for cross compilation
```
go env GOOS GOARCH
```
## TO list the supported platforms (for cross compilation)
```
go tool dist list
```
## Cross compilation
```
GOOS=[supported os] GOARCH=[supported processor arch] go build [file_name.go]
ex:
GOOS=windows GOARCH=amd64 go build 01-hello-world.go
```
