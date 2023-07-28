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

### 1.2. 业务场景
为便于展示，本项目使用业务场景为——不同类型计算功能服务API的调用，其中共三种计算服务：
1. FirstLevelCalService 一级计算服务。包括加法和减法运算
   - 此服务有两个运行实例，分别占用端口号9990，9991
2. SecondLevelCalService 二级计算服务。包括乘法和除法运算
   - 占用端口号9992
3. AdvancedCalService 高级计算服务。包括阶乘运算和斐波那契数列计算
   - 占用端口号9993

网关服务占用端口号8888

IDL管理平台占用端口号8889

API具体调用方式请参考接口文档

### 1.3. 项目结构
```text
.
├── README.md
├── hertzSvr-Gateway
├── kitexSvr-AdvancedCalService
├── kitexSvr-FirstLevelCalService-1
├── kitexSvr-FirstLevelCalService-2
└── kitexSvr-SecondLevelCalService
```

## 2. Deploy
### 2.1. 启动etcd注册中心
``` bash
etcd --log-level debug
```

### 2.2. 启动IDL管理平台
- in directory `hertzSvr-IDLManagement`
``` bash
sh ./build.sh && sh ./output/bootstrap.sh 
```

### 2.2. 启动网关
- in directory `hertzSvr-Gateway`
``` bash
sh ./build.sh && sh ./output/bootstrap.sh 
```

### 2.3. 启动所有服务
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

### 2.4. 发送请求示例
#### 2.4.1. 调用一级计算服务的加法API，计算1+2
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

#### 2.4.2. 调用二级计算服务的乘法API，计算114×514
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
1. 用户需要调用API时，通过`/gateway/:serviceName/:methodName`向网关发送HTTP请求
2. 网关接收到HTTP请求后进行路由，通过url中的服务名和方法名决定要将请求转发到哪个RPC服务
3. 根据服务名获取到对应的泛化调用客户端，进行HTTP泛化调用，转发至RPC服务
4. 网关收到Response，将其返回给用户

## Test
- in directory `hertzSvr`
``` bash
go test -bench=. main_test.go
```