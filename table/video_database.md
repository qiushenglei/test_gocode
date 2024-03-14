# video_core

`cdn_log`: cdn资源的错误日志

`video`: 视频列表

`video_venues`: 场馆(三方赛事源就是一个场馆，例如PM，IM)
> 一个视频可以再多个场馆下出现，因为他们两个源都能获取到这一场比赛，由于源不同，即使是同一场比赛，返回的赛事id(eid)也不一样
![在这里插入图片描述](https://img-blog.csdnimg.cn/direct/b883cd50f09043d99bee42120ecba8eb.png)

`video_streams`: 视频流
> - 一个视频一般会有 播控源 + 原始源(普通源)
> - 播控源一般等于2：flv + m3u8
> - 原始源一般等于2*n : (flv + m3u8) * 场馆
>   ![在这里插入图片描述](https://img-blog.csdnimg.cn/direct/b499f9639e1946c7a2eef53e9a770569.png)
> - 播控自选源(source_type=2): 通常都是flv + m3u8，一个场馆下的视频(flv+m3u8)都同时被当作自选源或者普通源

`video_animations`: 动画流

