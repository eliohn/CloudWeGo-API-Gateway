# CloudWeGo-API-Gateway
CloudWeGo API Gateway

## Deploy
### run etcd
``` bash
etcd --log-level debug
```

### run hertz client
- in directory `hertzSvr`
``` bash
go run .
```

### run all servers
- in directory `kitexSvr-serviceA`
``` bash
go run .
```
- in directory `kitexSvr-serviceB`
``` bash
go run .
```
- in directory `kitexSvr-serviceC`
``` bash
go run .
```
- in directory `kitexSvr-serviceD`
``` bash
go run .
```

## Send Request
```typescript
interface Request{
    svrName: string
    bizParams: string // 后续会使用map
}
```
