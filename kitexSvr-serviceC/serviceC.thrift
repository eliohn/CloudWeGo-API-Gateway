namespace go kitex.service

//--------------------C request & response--------------
struct CReq {
    1: string data(api.body='data')
}

struct CResp {
    1: bool success(api.body='success'),
    2: string message(api.body='message'),
}

//----------------------A service-------------------
service CService {
    CResp requestA(1: CReq req)(api.post = '/C/req')
}
