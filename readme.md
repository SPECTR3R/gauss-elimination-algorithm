# An attempt to apply golang concurrency to the gauss elimination algorithm.
There are two implementations of the gauss elimination algorithm. The first is a sequential implementation and the second is a concurrent implementation.

# Trace commands

```bash
cd conc
go build -o gauss
time ./gauss > m.trace
go tool trace m.trace
```

```bash
cd seq
go build -o gauss
time ./gauss > m.trace
go tool trace m.trace
```