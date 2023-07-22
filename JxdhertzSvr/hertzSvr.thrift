namespace go hertzSvr.service

struct SvrRequest{
    1: required string svrName (api.body="svrName")
    2: required string methodName (api.body="methodName")
    3: required string version (api.body="version")
    4: string data (api.body="data") //应该是JSON，用string代替
}

struct RegisterIDL{
    1: string name (api.body="name")
    2: string version(api.body="version")
    3: string idl (api.body="idl")
}

struct SvrResponse{
    1: bool isSuccess (api.body="isSuccess")
    2: string message (api.body="message") //应该是JSON，用string代替
}

service HertzSvr{
    SvrResponse Request(1: SvrRequest request)(api.post="/Request")
    SvrResponse RegisterIDL(1: RegisterIDL idl)(api.post="/RegisterIDL")
}