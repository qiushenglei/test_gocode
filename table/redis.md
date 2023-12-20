# 注解
- qid: 消息列id（也是群聊里的messageId[mid]）
- userid: 发送人userid或接收者userid
- seqid: 序列号id，是一个通过redis递增的值


# 消息列 message queue
- `chat_message_queue_user_xqid` : 消息列中的用户集合 
  - `type`: zset有序集合
  - `val`: `{score: @timestamp, val:@userid}`

- `chat_message_queue_sender_xqid` : 消息列中的已发送用户的集合,群聊消息里消息列用户的展示用到
  - `type`: zset有序集合
  - `val`: `{score: @timestamp, val:@userid}`

- `chat_user_mesage_queue_xuserid`: 用户参加的消息列集合（用户A参加了 x,y,z消息列）
  - `type`: zset有序集合
  - `val`: `{score: @timestamp, val:@qid}`

- `MemberReadGroupMessageQueue_xqid_xuserid`: 消息列中用户的最大已读seqid
  - `type`: int
  - `val`: seqid

- `MemberQueueChat_xqid`: 消息列最新会话消息
  - `type`: string  
  - `val`: mdata.msg

- `MemberQueueChatCount_xqid`: 消息列 回复次数
  - `type`: int
  - `val`: incr

- `IsMemberQueueChat_xqid`: 标记当前消息是否有消息列
  - `type`: string  
  - `val`: "queue"
    

# 私聊

- `MemberReadSingleMessage_xchatid_xluserid`: 私聊最大已读seqid
  - `type`: int
  - `val`: incr

- `MemberSingleChat_xchatid`: 私聊会话最新消息
  - `type`: string
  - `val`: mdata.msg

- `CHAT_USER_CHAT_LIST_X_xuserid_xchatid`: 私聊会话表行数据
  - `type`: string
  - `val`: mysql的chat_user_list.user_chat_list_x表结构体


# 群聊

- `MemberGroupChat_xgroupid`: 群聊会话最新消息
  - `type`: string
  - `val`: mdata.msg

- `group_info_xgroupid`: 群聊会话表行数据
  - `type`: string
  - `val`: mysql的chat_group.user_group表结构体

- `MemberReadGroupMessage_xchatid_xluserid`: 群聊最大已读seqid
  - `type`: int
  - `val`: incr

- `MemberBeReadGroupMessage_xchatid_xluserid`: 群聊自己以外最大已读seqid
  - `type`: 未知
  - `val`: 未知

- `CHAT_UserGroupUser_X_xuserid_xgroupid`: 未知
  - `type`: 未知
  - `val`: 未知



# 不分会话

- `MemberChatSeqId_xchattype_xchatid`: 会话最大seqid
  - `type`: int
  - `val`: incr


- `MemberOnlineTime_xuserid`: 用户在线登录时间
  - `type`: string
  - `val`: timestamp
  - 更新入口: chatjob的MessageConsumerListenJob定时任务 处理消息入口

- `MemberDeviceOnlineTime_xuserid`: 用户多个设备在线登录时间
  - `type`: Set集合
  - `val`: {xdeviceid: timestamp}
  - 更新入口: chatjob的MessageConsumerListenJob定时任务 处理消息入口

