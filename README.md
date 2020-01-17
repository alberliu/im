### 简要介绍
im是一个即时通讯服务器，代码全部使用golang完成。主要功能  
1.支持tcp，websocket接入  
2.离线消息同步  
3.单用户多设备同时在线    
4.单聊，群聊，以及超大群聊天场景  
5.支持服务水平扩展
### 使用技术：
数据库：Mysql+Redis  
通讯框架：Grpc  
长连接通讯协议：Protocol Buffers  
日志框架：Zap  
### rpc接口简介
项目所有的proto协议在im/public/proto/目录下
1.tcp.proto  
长连接通讯协议  
2.logic_client.ext.proto  
对客户端（Android设备，IOS设备）提供的rpc协议  
3.logic_server.ext.proto    
对业务服务器提供的rpc协议  
4.logic.int.proto  
对conn服务层提供的rpc协议  
5.conn.int.proto  
对logic服务层提供的rpc协议  
### 项目目录简介
项目结构遵循 https://github.com/golang-standards/project-layout
```
api:          服务对外提供的grpc接口
cmd:          服务启动入口
config:       服务配置
internal:     每个服务私有代码
pkg:          服务共有代码
sql:          项目sql文件
test:         长连接测试脚本
```
### 服务简介
1.tcp_conn  
维持与客户端的TCP长连接，心跳，以及TCP拆包粘包，消息编解码  
2.ws_conn  
维持与客户端的WebSocket长连接，心跳，消息编解码  
3.logic  
设备信息，用户信息，群组信息管理，消息转发逻辑  
