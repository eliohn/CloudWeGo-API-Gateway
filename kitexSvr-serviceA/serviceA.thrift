namespace go kitex.service

//--------------------A request & response--------------
struct AReq {
    1: string data(api.body='data')
}

struct AResp {
    1: bool success(api.body='success'),
    2: string message(api.body='message'),
}

//----------------------A service-------------------
service AService {
    AResp RequestA(1: AReq req)(api.post = '/A-req')
}
