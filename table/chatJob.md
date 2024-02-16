# BotSendMessageJob

- 作用：发送机器人消息，就是一个http请求(开发环境配置项缺失，启动不了，jenkins也没有这个工程)

# deviceReportJob
- 作用：用户
- 

# LoginLogConsumerListenJob
- 作用: 消费kafka消息 给用户登录客户端记入es日志落库
- 生产入口: chatapi登录接口
- topic: chat_api_login_log

# MessageConsumerListenJob
- 作用: 消费kafka消息， 给客户端发送的消息落库es
- 生产入口: chatpush的readPump
- topic: chat_push_xnum, group_chat_push_xnum, bot_chat_push_xnum

# MessageToDatabaseJob

## 工作1

- crontab: 1 * * * * ?(每分钟执行一次)
- 作用: 每分钟更新`chat_user_list.user_chat_list_x(私聊会话表)`的 `read_max_id(发送人最大已读seqid)`、`other_read_max_id(接收人最大已读seqid)`、`latest_msg(会话最新消息)`、`send_max_id(会话最大seqid)`

## 工作2

- crontab: 0 0 2 * * ?(每个月2号00:00执行一次)
- 作用: 每月更新`chat_group.user_group(群聊会话表)`的 、`latest_msg(会话最新消息)`、`send_max_id(会话最大seqid)`

## 工作3

- crontab: 0 0 3 * * ?(每个月3号00:00执行一次)
- 作用: 每月更新`chat_group.user_group_user_xuserid(用户和群聊的关联会话表)`的 、`read_max_id(本人已读最大seqid)`、`other_read_max_id(自己的消息中，最大被已读seqid)`

## 工作4

- crontab: 0 0 4 * * ?(每个月4号00:00执行一次)
- 作用: 每月更新用户在线最新时间`chat_edge.user_login_device`

# SmsSendJob 短信发送

- crontab: 每10s扫一次 sms_send_task表
- 作用: 异步发送短信

# UserChatHistoryRemoveJob

## 工作1

- command: `cmd -task clearUserData -delUserNames username1,username2,username3`
- 作用：清除用户信息

## 工作2

- go(updateDeleteTimeOfChatGroup)
- 作用：里面写的mysql条件很奇怪，不确定。清除群聊会话信息，`chat_group.user_group`表有个清除会话策略字段`delete_strategy`，根据它清除mysql

## 工作3

- go(DeleteUploadFile)
- 作用：清除上传文件 `chat_edge.up_file_info`

## 工作4
- go(trimUserQueue)
- 作用：把用户3个月前参加过的消息列集合清除，只是清除了redis的key:`chat_user_message_queue_xuserid`


# UserPasswordGenerator

- 作用：给输入文件中的用户名 创建密码
- 一次性脚本，非常驻