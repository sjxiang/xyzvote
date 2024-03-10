


# 需求描述



# 错误处理
https://github.com/cloudwego/kitex/discussions/248

逻辑错误和系统错误的错误信息应该通过什么样子的方式返回？



# model 处理
1. db2struct
db2struct --gorm --no-json -H 127.0.0.1 -d xyz_vote -t user --package db --struct UserM -u root -p 'my-secret-pw'

2. sql 转 gorm 在线平台
http://www.gotool.top/handlesql/sql2gorm



# cookie 与 session


# gorm

First 不会报错

用户注册，先查询是否注册过，日志告警 record not found，即使错误处理的漂亮，也没卵用

Find 

优缺点


约束 unique 冲突告警，还是通过查询判重


# crud

很多工作本身就是 crud，考量
1. 这个业务能不能赚钱？（亏钱的业务不会继续投入，迟早被砍掉）怎么赚钱的，是否提高了公司效率
2. 当前业务是否可以进一步优化，这个优化是否可以进一步提升效率（减少人工或者时间）
3. 业务中有没有典型的问题，或者比较复杂的场景，没有，是自己不知道，或者知道，但没有解决办法

上述三条，已经非常清楚了，可以换工作或者调岗



# 验证码

1. 用处？风控，判断脚本还是人工
2. 常见验证码？图片、拖拽 ...
3. 原理？
后端生成 captcha_id 和 answer 以及基于 answer 生成的 base64 字符串；
前端拿到 captcha_id 和 base64 字符串，渲染成图片；
用户输入。

