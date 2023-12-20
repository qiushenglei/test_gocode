# 注解
- qid: 消息列id（也是群聊里的messageId[mid]）
- userid: 发送人userid或接收者userid
- seqid: 序列号id，是一个通过redis递增的值



# 私聊

- `MemberSingleChat_xchatid`: 私聊会话最新消息
  - `type`: string
  - `val`: mdata.msg
  - 更新入口:
    - chatjob MessageConsumerListenJob 入库es，处理**发送**私聊消息方法(handleHistoryMessage)
    - chatjob MessageConsumerListenJob 入库es，处理**编辑**私聊消息方法(handleSingleMessage)

- `MemberReadSingleMessage_xchatid_xluserid`: 私聊最大已读seqid
  - `type`: int
  - `val`: incr
  - 更新入口:
    - chatjob MessageConsumerListenJob 入库es，处理**已读**私聊消息flag:13(handleSingleMessage)
    - chatjob MessageConsumerListenJob 入库es，处理**发送**私聊消息，还把当前用户的最新已读更新(handleSingleMessage)

- `CHAT_USER_CHAT_LIST_X_xuserid_xchatid`: 私聊会话表行数据
  - `type`: string
  - `val`: mysql的chat_user_list.user_chat_list_x表结构体
  - 更新入口:
    - chatapi 创建、删除、获取私聊

# 群聊

- `MemberGroupChat_xgroupid`: 私聊会话最新消息
  - `type`: string
  - `val`: mdata.msg
  - 更新入口: 
    - chatjob MessageConsumerListenJob 入库es，处理群聊消息方法(handleGroupMessage)


- `MemberReadGroupMessage_xchatid_xluserid`: 群聊最大已读seqid
  - `type`: int
  - `val`: seqid incr
  - 更新入口:
    - chatjob MessageConsumerListenJob 入库es，处理**已读**群聊消息flag=13
    - chatjob MessageConsumerListenJob 入库es，处理**发送**群聊消息时，还把当前用户的最新已读更新(handleGroupMessage)


- `MemberBeReadGroupMessage_xchatid_xluserid`: 群聊消息中自己的消息中最大已读seqid
  - `type`: int
  - `val`: seqid incr
  - 更新入口:
    - chatjob MessageConsumerListenJob 入库es，处理已读flag=13，处理发送人的**最大已读**，还处理了被读人的**最大被已读**


- `group_info_xgroupid`: 群聊会话表行数据
  - `type`: string
  - `val`: mysql的chat_group.user_group表结构体
  - 更新入口:
    - chatapi 创建、修改、删除群接口(这里为什么不扔到chatjob里面？？)


- `CHAT_UserGroupUser_X_xuserid_xgroupid`: 
  - `type`: string
  - `val`: mysql chat_group.user_group_user_x表结构体
  - 更新入口:
    - chatapi 创建、删除、编辑、获取群聊消息

# 消息列 message queue

- `chat_message_queue_user_xqid` : 消息列中的用户集合
  - `type`: zset有序集合
  - `val`: `{score: @timestamp, val:@userid}`
  - 更新入口:
    - chatjob MessageConsumerListenJob 入库es，处理**发送**消息列消息(handleGroupHistoryMessageQueue)


- `chat_message_queue_sender_xqid` : 消息列中的已发送用户的集合,群聊消息里消息列用户的展示用到
  - `type`: zset有序集合
  - `val`: `{score: @timestamp, val:@userid}`
  - 更新入口:
    - chatjob MessageConsumerListenJob 入库es，处理**发送**消息列消息(handleGroupHistoryMessageQueue)
    - chatjob MessageConsumerListenJob 入库es，处理**删除**消息列消息(DeleteQueueMsg)


- `chat_user_mesage_queue_xuserid`: 用户参加的消息列集合（用户A参加了 x,y,z消息列）
  - `type`: zset有序集合
  - `val`: `{score: @timestamp, val:@qid}`
  - 更新入口:
    - chatjob MessageConsumerListenJob 入库es，处理**发送**消息列消息(handleGroupHistoryMessageQueue)


- `chat_user_message_queue_sender_count_xqid_xuserid`: 消息列中用户发送的消息数量
  - `type`: int
  - `val`: incr 统计数量
  - 更新入口:
    - chatjob MessageConsumerListenJob 入库es，处理**发送**消息列消息(handleGroupHistoryMessageQueue)
    - chatjob MessageConsumerListenJob 入库es，处理**删除**消息列消息(DeleteQueueMsg)


- `MemberQueueChat_xqid`: 消息列最新会话消息
  - `type`: string
  - `val`: mdata.msg
  - 更新入口:
    - chatjob MessageConsumerListenJob 入库es，**发送**消息列消息(handleGroupHistoryMessageQueue)


- `MemberChatQueueSeqId_xqid_xuserid`: 消息列中最大seqid
  - `type`: int
  - `val`: seqid
  - 更新入口: chatpush.readPump.getGidAndSeqId 生成会话seqid的时候

- `MemberReadGroupMessageQueue_xqid_xuserid`: 消息列中用户的最大已读seqid
  - `type`: int
  - `val`: seqid
  - 更新入口:
    - chatjob MessageConsumerListenJob 入库es，处理**发送**消息列消息(handleGroupHistoryMessageQueue)


- `MemberQueueChatCount_xqid`: 消息列 回复次数
  - `type`: int
  - `val`: incr
  - 更新入口:
    - chatjob MessageConsumerListenJob 入库es，处理**发送**消息列消息(handleGroupHistoryMessageQueue)
    - chatjob MessageConsumerListenJob 入库es，处理**删除**消息列消息(DelteQueueMsg)


- `IsMemberQueueChat_xqid`: 标记当前消息是否有消息列
  - `type`: string
  - `val`: "queue"
  - 更新入口:
    - chatjob MessageConsumerListenJob 入库es，处理**发送**消息列消息(handleGroupHistoryMessageQueue)


# 不分会话

- `MemberChatSeqId_xchattype_xchatid`: 会话最大seqid
  - `type`: int
  - `val`: incr
  - 更新入口: chatpush.readPump.getGidAndSeqId 生成会话seqid的时候


- `MemberOnlineTime_xuserid`: 用户在线登录时间
  - `type`: string
  - `val`: timestamp
  - 更新入口: chatjob的MessageConsumerListenJob定时任务 处理消息入口

- `MemberDeviceOnlineTime_xuserid`: 用户多个设备在线登录时间
  - `type`: Set集合
  - `val`: {xdeviceid: timestamp}
  - 更新入口: chatjob的MessageConsumerListenJob定时任务 处理消息入口

