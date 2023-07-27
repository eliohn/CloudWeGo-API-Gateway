namespace go kitex.service

struct Request{
    1: i32 operand (api.body="operand")
}

struct Response{
    1: bool success (api.body="success")
    2: string message (api.body="message") //应该是JSON，用string代替
    3: i32 data (api.body="data")
}

service AdvancedCalService{
    Response Fact(1: Request request)(api.post="/gateway/AdvancedCalService/fact")
    Response Fib(1: Request request)(api.post="/gateway/AdvancedCalService/fib")
}