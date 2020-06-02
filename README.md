# Mächtiger Aluhut

An alexa skill which gives a random conspiracy theory when invoked.

[Here's the Amazon shop page]((https://www.amazon.de/Moritz-Kammerer-Mächtiger-Aluhut/dp/B01MRCZJ7Z)). At the moment, this skill is german language only.

## Building

```
go build
```

## Running

```
./go-maechtiger-aluhut
```

The skill now runs on `localhost:8192/maechtiger-aluhut`

```
./go-maechtiger-aluhut -address :8080
```

The skill now runs on `*:8080/maechtiger-aluhut`

## Usage

```
Usage of go-maechtiger-aluhut:
  -address string
        address to listen on (default "localhost:8192")
  -help
        prints this help
```

## Testing

```
go test ./...
```

## License

Licensed under [LGPLv3](https://www.gnu.org/licenses/lgpl-3.0.html).
