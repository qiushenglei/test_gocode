# chat_user库

- `user_info`：用户表
- `user_info_detail`：用户详情表
- `user_friends_x`：用户联系人表

- `user_notice_list_user`：用户组表的用户(目前来看只有@功能用到了)


# chat_user_list库

- `user_chat_list_x`：用户私聊会话列表


# chat_group库

- `user_group`：群聊元数据(群聊id，最大seqid，是否置顶pin，最后一条消息的具体信息等等)
- `user_group_users_x`：群聊里的用户 和 群聊的元数据(当前用户已读最大id，已最大删除最大id，是否有置顶pin等元数据)
- `user_group_id_x`：群聊用户 和 群聊的关联(其实user_group_users_x已经包括了)


# chat_edge

- `chat_pin_record`：置顶消息表
- `chat_del_record`：删除消息的记录表(批量删除时，就会插入多条)
- `user_login_device_x`：用户设备在线登录时间

# chat_edge

- `chat_pin_record`：置顶消息表
- `chat_del_record`：删除消息的记录表(批量删除时，就会插入多条)


# chat_notify

- `sms_send_task`：短信发送任务
- `sms_send_detail`：短信发送任务具体内容