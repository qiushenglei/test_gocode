# checkStreamServer

统计cdn_log表，日志统计10分钟内错误20次的流，标记成应该下架的流。redis_key:`cdn_error_xorder`


# CreateAnimationBMSStream

创建当前时间前后一小时的动画播控流(type=2的流，其他的是普通流，BMS系统默认播放播控流)