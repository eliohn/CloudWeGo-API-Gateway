namespace go hertzSvr.service

struct SvrRequest{
    1: map<string, string> bizParams (api.body="bizParams") //应该是JSON，用string代替
}

struct RegisterIDL{
    1: string name (api.body="name")
    2: string version(api.body="version")
    3: string idl (api.body="idl")
}

struct IDLMessage{
    1: string name (api.body="name")
    2: string version(api.body="version")
}

struct SvrResponse{
    1: bool success (api.body="success")
    2: string message (api.body="message") //应该是JSON，用string代替
}

service HertzSvr{
    SvrResponse Request(1: SvrRequest request)(api.post="/gateway/:svc/request")
    SvrResponse AddIDL(1: RegisterIDL idl)(api.post="/AddIDL")
    SvrResponse DeleteIDL(1: IDLMessage idl)(api.post="/DeleteIDL")
    SvrResponse UpdateIDL(1: RegisterIDL idl)(api.post="UpdateIDL")
    SvrResponse QueryIDL(1: IDLMessage idl)(api.post="/QueryIDL")
}