# local environment
> go build

then run it.

# docker
1. Windows PowerShell
> $env:GOOS="linux"

> go build -o accountservice-linux-amd64

2. Linux
> go build -o accountservice-linux-amd64

2. Build image
> docker build -t xxx/accountservice .

> docker run --rm -p 8989:8989 xxx/accountservice

# test
> curl http://localhost:8989/accounts/10001
