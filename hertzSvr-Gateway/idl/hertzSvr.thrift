namespace go hertzSvr.service

struct IDLInfo{
    1: string name (api.body="name")
    2: string idl (api.body="idl")
    3: string version(api.body="version")
}

struct IDLMessage{
    1: string name (api.body="name")
    2: string version(api.body="version")
}

struct RegisterReq {
    1: string svcName (api.body="svcName")
    2: string IDLName (api.body="IDLName")
    3: string version (api.body="version")
}

struct IDLResponse{
    1: bool success (api.body="success")
    2: string message (api.body="message")
}

service IDLService {
    IDLResponse AddIDL(1: IDLInfo idl)(api.post="/idl/add")
    IDLResponse DeleteIDL(1: IDLMessage idl)(api.post="/idl/delete")
    IDLResponse UpdateIDL(1: IDLInfo idl)(api.post="/idl/update")
    IDLInfo QueryIDL(1: IDLMessage idl)(api.get="/idl/query")
    IDLResponse RegisterIDL(1: RegisterReq req)(api.post="/idl/register")
}