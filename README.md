# CloudWeGo-API-Gateway
This is the CloudWeGo API Gateway Project, implementing an API Gateway which can transmit different HTTP requests to different services.

## 0. Members
| 姓名  | 学号        | 分工             |
|-----|-----------|----------------|
| 薛瑞宸 | 211250142 | 代码框架设计、业务代码编写  |
| 金煦东 | 211250208 | IDL管理接口        |
| 刘克典 | 211230043 | 测试代码、测试报告与测试文档 |

## 1. Intro
### 1.1. 概述
本项目实现了API网关的基本功能，包括服务注册、服务发现、路由服务、负载均衡等功能，并支持idl的热更新。

- 技术方案选型
  - 注册中心：选用etcd
  - idl持久化数据库：选用sqlite
  - 负载均衡策略：选用加权轮询

项目开发过程使用飞书进行管理。知识库链接已公开：[知识库](https://qiyaq6lc2gg.feishu.cn/wiki/space/7255927814334054404?ccm_open_type=lark_wiki_spaceLink)

### 1.2. 业务场景
为便于展示，本项目使用业务场景为——不同类型计算功能服务API的调用，其中共三种计算服务：
1. FirstLevelCalService 一级计算服务。包括加法和减法运算
   - 此服务有两个运行实例，分别占用端口号9990，9991
2. SecondLevelCalService 二级计算服务。包括乘法和除法运算
   - 占用端口号9992
3. AdvancedCalService 高级计算服务。包括阶乘运算和斐波那契数列计算
   - 占用端口号9993

网关服务Gateway占用端口号8888

IDL管理平台IDLManagement占用端口号8889

API具体调用方式请参考[接口文档](https://qiyaq6lc2gg.feishu.cn/wiki/VFgfwJLPmisKd8kZh55cvWpUnyf)

### 1.3. 项目结构
```text
.
├── README.md
├── hertzSvr-Gateway                //网关服务
├── hertzSvr-IDLManagement          //IDL管理平台
├── kitexSvr-AdvancedCalService     //高级计算服务
├── kitexSvr-FirstLevelCalService-1 //一级计算服务-实例1
├── kitexSvr-FirstLevelCalService-2 //一级计算服务-实例2
└── kitexSvr-SecondLevelCalService  //二级计算服务

```

## 2. Deploy
### 2.1. 启动etcd注册中心
``` bash
etcd --log-level debug
```

> 以下服务启动命令均可使用 go run . 代替
### 2.2. 启动IDL管理平台
- in directory `hertzSvr-IDLManagement`
``` bash
sh ./build.sh && sh ./output/bootstrap.sh 
```

### 2.3. 启动网关
- in directory `hertzSvr-Gateway`
``` bash
sh ./build.sh && sh ./output/bootstrap.sh 
```

### 2.4. 启动所有服务
- in directory `kitexSvr-FirstLevelCalService-1`
``` bash
sh ./build.sh && sh ./output/bootstrap.sh 
```
- in directory `kitexSvr-FirstLevelCalService-2`
``` bash
sh ./build.sh && sh ./output/bootstrap.sh 
```
- in directory `kitexSvr-SecondLevelCalService`
``` bash
sh ./build.sh && sh ./output/bootstrap.sh 
```
- in directory `kitexSvr-AdvancedCalService`
``` bash
sh ./build.sh && sh ./output/bootstrap.sh 
```

### 2.5. 发送请求示例
#### 2.4.1. 向IDL管理平台添加三个服务的idl
> 在添加idl时，idl的字符串必须为转义后的文本
> 
> 为方便测试，三个服务的idl转义字符串已经保存在各个服务项目下的./idl/idl-transferred-meaning.txt文件中
- 一级计算服务
```bash
curl --location 'http://localhost:8889/idl/add' \
--header 'Content-Type: application/json' \
--data '{
    "name": "FirstLevelCalService",
    "idl":"namespace go kitex.service\n\nstruct Request{\n    1: i32 operand_1 (api.body=\"operand_1\")\n    2: i32 operand_2 (api.body=\"operand_2\")\n}\n\nstruct Response{\n    1: bool success (api.body=\"success\")\n    2: string message (api.body=\"message\")\n    3: i32 data (api.body=\"data\")\n}\n\nservice FirstLevelCalService{\n    Response Add(1: Request request)(api.post=\"/gateway/FirstLevelCalService/add\")\n    Response Sub(1: Request request)(api.post=\"/gateway/FirstLevelCalService/sub\")\n}"
}'
```
- 二级计算服务
```bash
curl --location 'http://localhost:8889/idl/add' \
--header 'Content-Type: application/json' \
--data '{
    "name": "SecondLevelCalService",
    "idl":"namespace go kitex.service\n\nstruct Request{\n    1: i32 operand_1 (api.body=\"operand_1\")\n    2: i32 operand_2 (api.body=\"operand_2\")\n}\n\nstruct Response{\n    1: bool success (api.body=\"success\")\n    2: string message (api.body=\"message\")\n    3: i32 data (api.body=\"data\")\n}\n\nservice SecondLevelCalService{\n    Response Mul(1: Request request)(api.post=\"/gateway/SecondLevelCalService/mul\")\n    Response Div(1: Request request)(api.post=\"/gateway/SecondLevelCalService/div\")\n}"
}'
```
- 高级计算服务
```bash
curl --location 'http://localhost:8889/idl/add' \
--header 'Content-Type: application/json' \
--data '{
    "name": "AdvancedCalService",
    "idl":"namespace go kitex.service\n\nstruct Request{\n    1: i32 operand (api.body=\"operand\")\n}\n\nstruct Response{\n    1: bool success (api.body=\"success\")\n    2: string message (api.body=\"message\")\n    3: i32 data (api.body=\"data\")\n}\n\nservice AdvancedCalService{\n    Response Fact(1: Request request)(api.post=\"/gateway/AdvancedCalService/fact\")\n    Response Fib(1: Request request)(api.post=\"/gateway/AdvancedCalService/fib\")\n}"
}'
```

#### 2.4.2. 调用一级计算服务的加法API，计算1+2
- 发送request
```bash
curl --location 'http://localhost:8888/gateway/FirstLevelCalService/add' \
--header 'Content-Type: application/json' \
--data '{
    "operand_1": 1,
    "operand_2": 2
}'
```
- 收到response
```json
{
  "data": 3,
  "message": "FirstLevelCal Service Add method success",
  "success": true
}
```

#### 2.4.3. 调用二级计算服务的乘法API，计算114×514
- 发送request
```bash
curl --location 'http://localhost:8888/gateway/SecondLevelCalService/mul' \
--header 'Content-Type: application/json' \
--data '{   
    "operand_1" : 114,
    "operand_2" : 514
}'
```

- 收到response
```json
{
  "data": 58596,
  "message": "SecondLevelCal Service Mul method success", 
  "success": true
}
```

#### 2.4.3. 调用高级计算服务的阶乘API，计算10的阶乘
- 发送request
```bash
curl --location 'http://localhost:8888/gateway/AdvancedCalService/fact' \
--header 'Content-Type: application/json' \
--data '{   
    "operand" : 10
}'
```

- 收到response
```json
{
  "data": 3628800,
  "message": "AdvancedCal Service Fact method success",
  "success": true
}
```

## 3. 项目原理
### 3.1. API网关
1. 用户需要调用API时，通过`/gateway/:serviceName/:methodName`向网关发送HTTP请求
2. 网关接收到HTTP请求后进行路由，通过url中的服务名和方法名决定要将请求转发到哪个RPC服务
3. 根据服务名获取到对应的泛化调用客户端，进行HTTP泛化调用，转发至RPC服务
4. 网关收到Response，将其返回给用户

### 3.2. IDL管理
1. 用户向IDL管理平台添加对应服务的idl
2. 在创建client时，网关会自动向平台请求对应服务的idl，进行创建client
3. 若网关服务的idl需要更新（以一级计算服务为例），先调用IDL管理平台的update接口进行idl更新，再调用网关接口`http://localhost:8888/idl/update?svcName=FirstLevelCalService` ，此时网关会令相关服务的provider使用最新的idl进行热更新
4. 若用户直接调用网关的idl更新接口，而相关client此时还未被创建，则网关会直接使用管理平台中最新的idl进行client创建

## Test
- in directory `hertzSvr`
``` bash
go test -bench=. main_test.go
```