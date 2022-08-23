# Trace commands

```bash
cd seq
go build -o gauss
time ./gauss > m.trace
go tool trace m.trace
```
