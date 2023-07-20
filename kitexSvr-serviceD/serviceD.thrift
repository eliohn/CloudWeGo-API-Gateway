namespace go kitex.service

//--------------------D request & response--------------
struct DReq {
    1: string data(api.body='data')
}

struct DResp {
    1: bool success(api.body='success'),
    2: string message(api.body='message'),
}

//----------------------D service-------------------
service DService {
    DResp requestA(1: DReq req)(api.post = '/D/req')
}
