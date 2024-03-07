## 嘉为蓝鲸IBMMQ插件使用说明

## 使用说明

### 插件功能

### 版本支持

操作系统支持: linux

是否支持arm: 不支持

**是否支持远程采集:**

是

### 参数说明


| **参数名**                       | **含义**                               | **是否必填** | **使用举例**        |
|-------------------------------|--------------------------------------|----------|-----------------|
| --ibmmq.httpListenHost        | 插件监听地址(下发请保持默认)                      | 是        | 127.0.0.1       |
| --ibmmq.httpListenPort        | 插件监听端口(下发请保持默认)                      | 是        | 9601            |
| --ibmmq.client                | 以客户端形式连接(开关参数), 默认关闭                 | 否        |                 |
| IBMMQ_CONNECTION_CONNNAME     | ibmmq服务连接地址(环境变量), 注意填写形式 `ip(port)` | 是        | 127.0.0.1(1414) |
| IBMMQ_CONNECTION_QUEUEMANAGER | ibmmq队列管理器名称(环境变量)                   | 是        | QM1             |
| IBMMQ_CONNECTION_CHANNEL      | ibmmq连接通道名称(环境变量)                    | 是        | SERVER          |
| IBMMQ_CONNECTION_USER         | ibmmq连接账户名称(环境变量)                    | 否        | admin           |
| IBMMQ_CONNECTION_PASSWORD     | ibmmq连接密码(环境变量)                      | 否        |                 |
| --log.level                   | 日志级别                                 | 否        | info            |

### 使用指引
1. 配置IBM MQ redistributable client
使用该探针的前提条件必须在下发探针的机器配置IBM MQ redistributable client。
> 如果没有配置客户端, 可以下载安装包。  
根据需要的版本下载x.x.x.x-IBM-MQC-Redist-LinuxX64.tar.gz  
下载地址: [IBM MQ redistributable client](https://public.dhe.ibm.com/ibmdl/export/pub/software/websphere/messaging/mqdev/redist)   
指定版本下载地址: [9.3.4.1-IBM-MQC-Redist-LinuxX64.tar.gz](https://public.dhe.ibm.com/ibmdl/export/pub/software/websphere/messaging/mqdev/redist/9.3.4.1-IBM-MQC-Redist-LinuxX64.tar.gz)   

> 下载后解压到对应目录   
linux在目录`/opt/mqm`下解压, 如果没有该目录请新建一个即可    

2. 配置账户授权 
```
# 进入名称为QM1的队列管理器管理界面
runmqsc QM1

# 查看ccsid
display qmgr ccsid

# 修改队列管理器的ccsid, 建议修改为819
alter qmgr ccsid(819)
 
# 定义类型为SVRCONN名称为SERVER的远程连接通道，同时授权mqm账号进行远程连接
def chl(SERVER) chltype(SVRCONN) trptype(tcp) mcauser('mqm') replace

# 定义名称为TCP端口为1414的tcp监听端口
def listener(TCP) trptype(tcp) port(1414)

# 启动监听端口
start listener(TCP)
```

定义通道时 `SERVER` 是填入探针参数 IBMMQ_CONNECTION_CHANNEL的通道名称, `mqm` 是填入探针参数 IBMMQ_CONNECTION_USER 已存在的账号名称

可以选择验证方式, 一般不建议取消验证  
- 正常验证
    ```
    # 关闭对mqm的禁用  
    SET CHLAUTH(*) TYPE(BLOCKUSER) USERLIST(*MQADMIN) ACTION(REMOVE)
    ```


- 取消验证
    ```
    # 禁用通道校验
    ALTER QMGR CHLAUTH(DISABLED)
    
    # 禁用连接权限认证
    ALTER QMGR CONNAUTH('')
    
    # 刷新安全策略：  
    REFRESH SECURITY TYPE(CONNAUTH)
    ```



### 指标简介

| **指标ID**                                                         | **指标中文名**               | **维度ID**                              | **维度含义**             | **单位**  |
|------------------------------------------------------------------|-------------------------|---------------------------------------|----------------------|---------|
| ibmmq_qmgr_uptime                                                | IBMMQ队列管理器已运行时长         | description, hostname, platform, qmgr | 描述, 主机名, 平台, 队列管理器名称 | s       |
| ibmmq_qmgr_active_listeners                                      | IBMMQ活动监听器数量            | description, hostname, platform, qmgr | 描述, 主机名, 平台, 队列管理器名称 | -       |
| ibmmq_qmgr_channel_initiator_status                              | IBMMQ通道初始化状态            | description, hostname, platform, qmgr | 描述, 主机名, 平台, 队列管理器名称 | -       |
| ibmmq_qmgr_command_server_status                                 | IBMMQ命令服务器状态            | description, hostname, platform, qmgr | 描述, 主机名, 平台, 队列管理器名称 | -       |
| ibmmq_qmgr_create_durable_subscription_count                     | IBMMQ创建持久订阅计数           | description, hostname, platform, qmgr | 描述, 主机名, 平台, 队列管理器名称 | -       |
| ibmmq_qmgr_create_non_durable_subscription_count                 | IBMMQ创建非持久订阅计数          | description, hostname, platform, qmgr | 描述, 主机名, 平台, 队列管理器名称 | -       |
| ibmmq_qmgr_delete_durable_subscription_count                     | IBMMQ删除持久订阅计数           | description, hostname, platform, qmgr | 描述, 主机名, 平台, 队列管理器名称 | -       |
| ibmmq_qmgr_delete_non_durable_subscription_count                 | IBMMQ删除非持久订阅计数          | description, hostname, platform, qmgr | 描述, 主机名, 平台, 队列管理器名称 | -       |
| ibmmq_qmgr_durable_subscriber_high_water_mark                    | IBMMQ持久订阅者高水位标记         | description, hostname, platform, qmgr | 描述, 主机名, 平台, 队列管理器名称 | -       |
| ibmmq_qmgr_durable_subscriber_low_water_mark                     | IBMMQ持久订阅者低水位标记         | description, hostname, platform, qmgr | 描述, 主机名, 平台, 队列管理器名称 | -       |
| ibmmq_qmgr_subscription_delete_failure_count                     | IBMMQ订阅删除失败计数           | description, hostname, platform, qmgr | 描述, 主机名, 平台, 队列管理器名称 | -       |
| ibmmq_qmgr_exporter_collection_time                              | IBMMQ导出器收集时间            | platform, qmgr                        | 平台, 队列管理器名称          | -       |
| ibmmq_qmgr_exporter_publications                                 | IBMMQ导出器发布数量            | platform, qmgr                        | 平台, 队列管理器名称          | -       |
| ibmmq_qmgr_failed_browse_count                                   | IBMMQ失败浏览计数             | description, hostname, platform, qmgr | 描述, 主机名, 平台, 队列管理器名称 | -       |
| ibmmq_qmgr_failed_create_alter_resume_subscription_count         | IBMMQ失败创建/更改/恢复订阅计数     | description, hostname, platform, qmgr | 描述, 主机名, 平台, 队列管理器名称 | -       |
| ibmmq_qmgr_failed_mqcb_count                                     | IBMMQ失败MQCB计数           | description, hostname, platform, qmgr | 描述, 主机名, 平台, 队列管理器名称 | -       |
| ibmmq_qmgr_failed_mqclose_count                                  | IBMMQ失败MQCLOSE计数        | description, hostname, platform, qmgr | 描述, 主机名, 平台, 队列管理器名称 | -       |
| ibmmq_qmgr_failed_mqconn_mqconnx_count                           | IBMMQ失败MQCONN/MQCONNX计数 | description, hostname, platform, qmgr | 描述, 主机名, 平台, 队列管理器名称 | -       |
| ibmmq_qmgr_failed_mqget_count                                    | IBMMQ失败MQGET计数          | description, hostname, platform, qmgr | 描述, 主机名, 平台, 队列管理器名称 | -       |
| ibmmq_qmgr_failed_mqinq_count                                    | IBMMQ失败MQINQ计数          | description, hostname, platform, qmgr | 描述, 主机名, 平台, 队列管理器名称 | -       |
| ibmmq_qmgr_failed_mqopen_count                                   | IBMMQ失败MQOPEN计数         | description, hostname, platform, qmgr | 描述, 主机名, 平台, 队列管理器名称 | -       |
| ibmmq_qmgr_failed_mqput1_count                                   | IBMMQ失败MQPUT1计数         | description, hostname, platform, qmgr | 描述, 主机名, 平台, 队列管理器名称 | -       |
| ibmmq_qmgr_failed_mqput_count                                    | IBMMQ失败MQPUT计数          | description, hostname, platform, qmgr | 描述, 主机名, 平台, 队列管理器名称 | -       |
| ibmmq_qmgr_failed_mqset_count                                    | IBMMQ失败MQSET计数          | description, hostname, platform, qmgr | 描述, 主机名, 平台, 队列管理器名称 | -       |
| ibmmq_qmgr_failed_mqsubrq_count                                  | IBMMQ失败MQSUBRQ计数        | description, hostname, platform, qmgr | 描述, 主机名, 平台, 队列管理器名称 | -       |
| ibmmq_qmgr_log_current_primary_space_in_use_percentage           | IBMMQ日志当前主空间使用百分比       | description, hostname, platform, qmgr | 描述, 主机名, 平台, 队列管理器名称 | percent |
| ibmmq_qmgr_log_file_system_in_use_bytes                          | IBMMQ日志文件系统使用空间大小       | description, hostname, platform, qmgr | 描述, 主机名, 平台, 队列管理器名称 | bytes   |
| ibmmq_qmgr_log_file_system_max_bytes                             | IBMMQ日志文件系统最大空间         | description, hostname, platform, qmgr | 描述, 主机名, 平台, 队列管理器名称 | bytes   |
| ibmmq_qmgr_log_in_use_bytes                                      | IBMMQ日志使用空间大小           | description, hostname, platform, qmgr | 描述, 主机名, 平台, 队列管理器名称 | bytes   |
| ibmmq_qmgr_log_logical_written_bytes                             | IBMMQ逻辑写入日志大小           | description, hostname, platform, qmgr | 描述, 主机名, 平台, 队列管理器名称 | bytes   |
| ibmmq_qmgr_log_max_bytes                                         | IBMMQ日志字节容量             | description, hostname, platform, qmgr | 描述, 主机名, 平台, 队列管理器名称 | bytes   |
| ibmmq_qmgr_log_physical_written_bytes                            | IBMMQ物理写入空间日志大小         | description, hostname, platform, qmgr | 描述, 主机名, 平台, 队列管理器名称 | bytes   |
| ibmmq_qmgr_log_size_archive                                      | IBMMQ存档日志大小             | description, hostname, platform, qmgr | 描述, 主机名, 平台, 队列管理器名称 | bytes   |
| ibmmq_qmgr_log_size_media                                        | IBMMQ媒体日志大小             | description, hostname, platform, qmgr | 描述, 主机名, 平台, 队列管理器名称 | bytes   |
| ibmmq_qmgr_log_size_restart                                      | IBMMQ重新启动恢复日志大小         | description, hostname, platform, qmgr | 描述, 主机名, 平台, 队列管理器名称 | bytes   |
| ibmmq_qmgr_log_size_reusable                                     | IBMMQ可重用日志大小            | description, hostname, platform, qmgr | 描述, 主机名, 平台, 队列管理器名称 | bytes   |
| ibmmq_qmgr_log_workload_primary_space_utilization_percentage     | IBMMQ日志工作负载主空间利用率       | description, hostname, platform, qmgr | 描述, 主机名, 平台, 队列管理器名称 | percent |
| ibmmq_qmgr_log_write_latency_seconds                             | IBMMQ日志写入延迟             | description, hostname, platform, qmgr | 描述, 主机名, 平台, 队列管理器名称 | s       |
| ibmmq_qmgr_log_write_size_bytes                                  | IBMMQ写入日志大小             | description, hostname, platform, qmgr | 描述, 主机名, 平台, 队列管理器名称 | bytes   |
| ibmmq_qmgr_alter_durable_subscription_count                      | IBMMQ修改持久订阅计数           | description, hostname, platform, qmgr | 描述, 主机名, 平台, 队列管理器名称 | -       |
| ibmmq_qmgr_expired_message_count                                 | IBMMQ过期消息计数             | description, hostname, platform, qmgr | 描述, 主机名, 平台, 队列管理器名称 | -       |
| ibmmq_qmgr_connection_count                                      | IBMMQ连接数                | description, hostname, platform, qmgr | 描述, 主机名, 平台, 队列管理器名称 | -       |
| ibmmq_qmgr_commit_count                                          | IBMMQ提交计数               | description, hostname, platform, qmgr | 描述, 主机名, 平台, 队列管理器名称 | -       |
| ibmmq_qmgr_mqcb_count                                            | IBMMQ MQCB操作计数          | description, hostname, platform, qmgr | 描述, 主机名, 平台, 队列管理器名称 | -       |
| ibmmq_qmgr_mqclose_count                                         | IBMMQ MQCLOSE操作数量       | description, hostname, platform, qmgr | 描述, 主机名, 平台, 队列管理器名称 | -       |
| ibmmq_qmgr_mqctl_count                                           | IBMMQ MQCTL操作数量         | description, hostname, platform, qmgr | 描述, 主机名, 平台, 队列管理器名称 | -       |
| ibmmq_qmgr_mqdisc_count                                          | IBMMQ MQDISC操作数量        | description, hostname, platform, qmgr | 描述, 主机名, 平台, 队列管理器名称 | -       |
| ibmmq_qmgr_mqinq_count                                           | IBMMQ MQINQ操作数量         | description, hostname, platform, qmgr | 描述, 主机名, 平台, 队列管理器名称 | -       |
| ibmmq_qmgr_mqopen_count                                          | IBMMQ MQOPEN操作数量        | description, hostname, platform, qmgr | 描述, 主机名, 平台, 队列管理器名称 | -       |
| ibmmq_qmgr_mqset_count                                           | IBMMQ MQSET操作数量         | description, hostname, platform, qmgr | 描述, 主机名, 平台, 队列管理器名称 | -       |
| ibmmq_qmgr_mqstat_count                                          | IBMMQ MQSTAT操作数量        | description, hostname, platform, qmgr | 描述, 主机名, 平台, 队列管理器名称 | -       |
| ibmmq_qmgr_mqsubrq_count                                         | IBMMQ MQSUBRQ操作数量       | description, hostname, platform, qmgr | 描述, 主机名, 平台, 队列管理器名称 | -       |
| ibmmq_qmgr_rollback_count                                        | IBMMQ回滚操作计数             | description, hostname, platform, qmgr | 描述, 主机名, 平台, 队列管理器名称 | -       |
| ibmmq_qmgr_persistent_message_browse_bytes                       | IBMMQ持久消息浏览大小           | description, hostname, platform, qmgr | 描述, 主机名, 平台, 队列管理器名称 | bytes   |
| ibmmq_qmgr_persistent_message_browse_count                       | IBMMQ持久消息浏览计数           | description, hostname, platform, qmgr | 描述, 主机名, 平台, 队列管理器名称 | -       |
| ibmmq_qmgr_persistent_message_destructive_get_count              | IBMMQ持久消息破坏性获取计数        | description, hostname, platform, qmgr | 描述, 主机名, 平台, 队列管理器名称 | -       |
| ibmmq_qmgr_persistent_message_mqput1_count                       | IBMMQ持久消息MQPUT1操作数量     | description, hostname, platform, qmgr | 描述, 主机名, 平台, 队列管理器名称 | -       |
| ibmmq_qmgr_persistent_message_mqput_count                        | IBMMQ持久消息MQPUT操作数量      | description, hostname, platform, qmgr | 描述, 主机名, 平台, 队列管理器名称 | -       |
| ibmmq_qmgr_put_persistent_messages_bytes                         | IBMMQ放置持久消息大小           | description, hostname, platform, qmgr | 描述, 主机名, 平台, 队列管理器名称 | bytes   |
| ibmmq_qmgr_queue_manager_file_system_free_space_percentage       | IBMMQ队列管理器文件系统空闲空间百分比   | description, hostname, platform, qmgr | 描述, 主机名, 平台, 队列管理器名称 | percent |
| ibmmq_qmgr_queue_manager_file_system_in_use_bytes                | IBMMQ队列管理器文件系统使用中的大小    | description, hostname, platform, qmgr | 描述, 主机名, 平台, 队列管理器名称 | bytes   |
| ibmmq_qmgr_cpu_load_fifteen_minute_average_percentage            | IBMMQ十五分钟CPU平均负载百分比     | description, hostname, platform, qmgr | 描述, 主机名, 平台, 队列管理器名称 | percent |
| ibmmq_qmgr_cpu_load_five_minute_average_percentage               | IBMMQ五分钟CPU平均负载百分比      | description, hostname, platform, qmgr | 描述, 主机名, 平台, 队列管理器名称 | percent |
| ibmmq_qmgr_cpu_load_one_minute_average_percentage                | IBMMQ一分钟CPU平均负载百分比      | description, hostname, platform, qmgr | 描述, 主机名, 平台, 队列管理器名称 | percent |
| ibmmq_qmgr_ram_free_percentage                                   | IBMMQ RAM空闲百分比          | description, hostname, platform, qmgr | 描述, 主机名, 平台, 队列管理器名称 | percent |
| ibmmq_qmgr_ram_total_bytes                                       | IBMMQ RAM总字节数           | description, hostname, platform, qmgr | 描述, 主机名, 平台, 队列管理器名称 | bytes   |
| ibmmq_qmgr_ram_total_estimate_for_queue_manager_bytes            | IBMMQ RAM队列管理器总字节数      | description, hostname, platform, qmgr | 描述, 主机名, 平台, 队列管理器名称 | bytes   |
| ibmmq_qmgr_system_cpu_time_estimate_for_queue_manager_percentage | IBMMQ系统队列管理器CPU时间百分比    | description, hostname, platform, qmgr | 描述, 主机名, 平台, 队列管理器名称 | percent |
| ibmmq_qmgr_system_cpu_time_percentage                            | IBMMQ系统CPU时间百分比         | description, hostname, platform, qmgr | 描述, 主机名, 平台, 队列管理器名称 | percent |
| ibmmq_qmgr_user_cpu_time_estimate_for_queue_manager_percentage   | IBMMQ用户队列管理器CPU时间百分比    | description, hostname, platform, qmgr | 描述, 主机名, 平台, 队列管理器名称 | percent |
| ibmmq_qmgr_user_cpu_time_percentage                              | IBMMQ用户CPU时间百分比         | description, hostname, platform, qmgr | 描述, 主机名, 平台, 队列管理器名称 | percent |
| ibmmq_qmgr_purged_queue_count                                    | IBMMQ已清除队列数量            | description, hostname, platform, qmgr | 描述, 主机名, 平台, 队列管理器名称 | -       |
| ibmmq_qmgr_concurrent_connections_high_water_mark                | IBMMQ并发连接数高水位标记         | description, hostname, platform, qmgr | 描述, 主机名, 平台, 队列管理器名称 | -       |
| ibmmq_qmgr_got_non_persistent_messages_bytes                     | IBMMQ获取非持久性消息大小         | description, hostname, platform, qmgr | 描述, 主机名, 平台, 队列管理器名称 | bytes   |


### 版本日志

#### weops_IBMMQ_exporter 5.5.2

- weops调整

#### weops_IBMMQ_exporter 5.5.2

- 优化文档配置用户说明

添加“小嘉”微信即可获取IBMMQ监控指标最佳实践礼包，其他更多问题欢迎咨询

<img src="https://wedoc.canway.net/imgs/img/小嘉.jpg" width="50%" height="50%">
