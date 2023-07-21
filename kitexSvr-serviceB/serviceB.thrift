namespace go kitex.service

//--------------------B request & response--------------
struct BReq {
    1: string data(api.body='data')
}

struct BResp {
    1: bool success(api.body='success'),
    2: string message(api.body='message'),
}

//----------------------A service-------------------
service BService {
    BResp RequestB(1: BReq req)(api.post = '/B-req')
}
