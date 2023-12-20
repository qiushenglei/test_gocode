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
- 作用: 每月更新`chat_group.user_group_user_xuserid(用户和群聊的关联会话表)`的 、`read_max_id(本人已读最大seqid)`、`other_read_max_id(除自己以外的会话最大seqid)`

## 工作4

- crontab: 0 0 4 * * ?(每个月4号00:00执行一次)
- 作用: 每月更新用户在线最新时间`chat_edge.user_login_device`
