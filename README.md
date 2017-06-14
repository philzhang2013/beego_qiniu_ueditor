# beego_qiniu_ueditor
使用beego框架和百度ueditor实现七牛的图片上传和视频上传

# 引入七牛golang sdk

运行如下命令：</br>
go get -u qiniupkg.com/api.v7

七牛sdk文档地址</br>
https://developer.qiniu.com/kodo/sdk/1238/go

# 配置


conf/app.conf
配置七牛 accesskey，secretkey，对象存储图片和视频的bucket名称和域名</br>

+ qiniuaccesskey = Your qiniu accessKey
+ qiniusercetkey = Your qiniu secretKey
+ qiniuimagebucketname = your qiniu image bucket name
+ qiniuimagehost =your qiniu image bucket host
+ qiniuvideobucketname = your qiniu video bucket name
+ qiniuvideohost = your qiniu video host

# 运行
bee run beego_qiniu_ueditor



