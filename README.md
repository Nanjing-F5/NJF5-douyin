# 0 进度
## 5.18 20:40
目前数据库中用户密码使用的是明文，后续可进行加密<br/>
调试：
1. /mysql/douyin.sql下
2. 执行建库、建表、插入用户数据语句
3. 执行main.exe
4. 打开app，使用数据库中的用户登录
5. 登录成功会跳转到用户信息页

## 5.24 17:10
1. 修改了下user表结构，需要重新删除表再创建(/mysql/douyin.sql有对应语句)
2. 基本完成注册功能
   1. 调试:
   2. 执行main.exe
   3. 打开app -> 注册 -> 登录

## 5.25 21:00
1. 新建了video表
2. 实现了用户上传视频 douyin/publish/action/
   1. 测试：
   2. 执行main.exe
   3. 打开app -> 登录 -> 点击'+'号上传视频
   4. 返回 发布成功
   5. 查看 video 表视频相关信息是否保存

# 1 文件介绍

![douyin](https://gitee.com/kkite/blogimg/raw/master/202205160931683.png)

# 2 大致逻辑

![大致项目逻辑](https://gitee.com/kkite/blogimg/raw/master/202205160931668.png)