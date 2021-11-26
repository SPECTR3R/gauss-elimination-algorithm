# Trace commands

```bash
go build -o gauss
time ./gauss > m.trace
go tool trace m.trace
```
