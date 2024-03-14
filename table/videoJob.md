# checkStreamServer (go loop)

统计cdn_log表，日志统计10分钟内错误20次的流，标记成应该下架的流。redis_key:`cdn_error_xorder`


# CreateAnimationBMSStream(go loop) 创建动画的播控自选流

创建当前时间前后一小时的动画播控流(video_animations type=2的播控自选流，其他的是普通流，BMS系统默认播放 播控流)

# Analysis(go loop)

给上线的视频流(video_streams online=1)做分析，通过ffmpeg去读取视频流的声音 和 分辨率

# go loop剩余部分 创建赛事的播控自选源

1. 如果该赛事没有播控自选源，创建 类型为(flv[live_path_encode=2],m3u8[live_path_encode=1]) 的播控自选源(source_type=1)视频流 行数据(video_stream)
2. 调用healthCheck 检测 flv流和m3u8流

# healthCheck

## flv
读了flv头后，第一个tag是否返回err， 递增redis： "pull_stream_err_count:xstreamid:flv"
如果检测到是不健康的，并且超过阈值(连续3次都是err)，设置流(视频，动画或主播)下线

## m3u8 (HLS协议，下载文件本地缓存的方式)

同上


# dealFailStream

如果go loop内的healthCheck检测失败，并且需要重新检测，这里重新调用healthCheck

# pipe

## initVideoSources

请求每个视频网站的赛事列表，把赛事详情拉取下来

## matchLoop 创建场馆的赛事

1. 根据3方网站来的信息来创建赛事记录(video_venus)，有可能多个网站有同一个赛事，所以有可能创建1主N副的情况
2. 检车已经完赛的赛事，更新状态

## pullStreams 拉取视频流url，并保存创建 视频流普通源到数据库

1. 去一个网站，拉去赛事的源。
2. 根据返回的id类型，去视频流表查看是否存在普通源(source_type=2)，不存在则添加视频流普通源数据(video_stream.source_type=2)并且type永远等于0(流类型:1视频直播哦 2动画直播 3主播视频主播 4主播动画直播)，也是flv和m3u8l两种】
> 直播开播的时候会插入type=3,4的主播源行(1,2暂时没有用到)

## pullAnimation  

请求每个视频网站的动画赛事列表，把赛事详情拉取下来



当前项目进度：

3.01号 pc前端 rocket大佬介入，其他端待定

兼容问题： pc端出现了一个取消置顶出现一个空白消息（如上图）
原因:取消置顶(对端是无法看见),消息添加seqid并加入到了chat_message index里，发送下一条消息时，客户端判断seqid不连续，接着会请求getMessage拉取数据。此时会获得两条(第一条取消置顶消息 第二条普通消息)，pc端出现白框
如果现在改成getMessage不返回"取消置顶消息"，客户端要求seqid连续，否则会出现其他问题

解决：
因为只有取消置顶操作有问题，这一版是否可以只把 1.删除消息、编辑消息添加seqid 2.添加取消个人置顶置顶flag 3.es替代mysql 这3个功能 和 UI重构那个项目2期一起提测。并在UI重构这里兼容好取消置顶这个问题。等强更App包的时候再把 取消置顶添加seqid这个功能提测上线



