namespace go hertzSvr.service

//--------------------update request & response--------------
struct UpdateReq {
    1: string idl(api.body = 'idl')
}

struct UpdateResp {
    1: bool success(api.body='success'),
    2: string message(api.body='message'),
}

//--------------------A request & response--------------
struct AReq {
    1: string data(api.body='data')
}

struct AResp {
    1: bool success(api.body='success'),
    2: string message(api.body='message'),
}

//--------------------B request & response--------------
struct BReq {
    1: string data(api.body='data')
}

struct BResp {
    1: bool success(api.body='success'),
    2: string message(api.body='message'),
}

//--------------------C request & response--------------
struct CReq {
    1: string data(api.body='data')
}

struct CResp {
    1: bool success(api.body='success'),
    2: string message(api.body='message'),
}

//--------------------D request & response--------------
struct DReq {
    1: string data(api.body='data')
}

struct DResp {
    1: bool success(api.body='success'),
    2: string message(api.body='message'),
}

//----------------------update service-------------------
service UpdateService {
    UpdateResp Update(1: UpdateReq req)(api.get = '/update')
}

//----------------------A service-------------------
service AService {
    AResp RequestA(1: AReq req)(api.post = '/A-req')
}

//----------------------B service-------------------
service BService {
    BResp RequestB(1: BReq req)(api.post = '/B-req')
}

//----------------------C service-------------------
service CService {
    CResp RequestC(1: CReq req)(api.post = '/C-req')
}

//----------------------D service-------------------
service DService {
    DResp RequestD(1: DReq req)(api.post = '/D-req')
}