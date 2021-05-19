账号系统

#### 主要思想

    - 同一用户下绑定不同的账号，利用三元组进行隔离

#### 功能

- 登录
    - 密码登录
    - 短信登录
    - 微信登录
- 查询详情
- 注销
- 修改密码
- 找回密码
- 上传头像
- 更改昵称和简介

#### 数据库设计

- user表

    字段 | 类型|说明
    ---|---|---
    user_id | string| 用户主键
    password|string|用户密码|
    name | string | 用户昵称
    avatar|string | 用户头像
    status|string|用户状态
    account_list|string|用户账号列表
    extension |string|扩展信息
    create_time|string|创建时间
    modify_time|string|更改时间

- user_bind表

    字段 | 类型|说明
    ---|---|---
    id |string| 主键
    user_id | string|用户id
    channel|string|账号渠道
    identify_type|string|账号类型|
    identifier|string|用户账号
    status|string|账号状态
    extension|string|扩展信息
    create_time|string|创建时间
    modify_time|string|修改时间

- user_action
  字段|类型|说明
  ---|--|--
  id|string|主键
  user_id|string|用户id
  channel|string|账号渠道
  identify_type|string|账号类型|
  identifier|string|账号
  action|string|动作
  extension|string|扩展信息
  create_time|string|创建时间
  modify_time|string|更新时间

- session_record
  
    字段|类型|说明
    --|--|--
    id|string|主键
    user_id|string|用户id
    token|string|
    extension|string|扩展信息
    expire_at|string|创建时间戳
    create_time|string|创建时间
    modify_time|string|更改时间



### 接口设计

- 注册
     - 方法 POST Content-Type : application/json
    - 参数
        字段 | 类型|必填|说明
        ---|---|---|---
        source |string|是| 调用业务方
        user_id | string|非|用户id
        channel|string|是|账号渠道
        identify_type|string|是|账号类型|
        identifier|string|是|用户账号
        register_mode|int|是|注册模式 0 仅注册 1 注册并登录 2 绑定用户

    - 返回参数
        字段|说明|
        -- | -- 
        user_id|注册用户id
        token|生成token

- 登录
    - 方法 POST Content-Type:application/json
    - 参数

        字段 | 类型|说明
        ---|---|---
        user_id | string|否|用户id或者三元组必须填一个
        passport_word|string|登录密码
        identify_type|string|账号类型|
        identifier|string|用户账号
        channel|string|渠道
        otp_code|string|邮箱或者短信验证码
    - 返回
        字段|说明
        --|--
        user_id|用户id
        token|登录凭证   

  - 短信发送
    - 方法 POST Content-Type:application/json
    - 参数
        字段 | 类型|说明
        ---|---|---
        identify_type|string|账号类型|
        identifier|string|用户账号
        channel|string|渠道
        point|string|发送短信节点
    - 返回
        字段|说明
        --|--
        otp_code|验证码     
   
- 查询详情
    - 方法：GET
    - 参数
        字段|类型|必填|说明
        -- |--|--|--
        user_id|string|否|用户id和三元组必须用一个
        channel|string|否|
        identify_type|string|否
        identifier|string|否
    - 返回参数    
        字段|说明
        --| --|
        user_id|用户id
        account_list|账号列表
        name|昵称
        user_introduction|用户简介

- 注销
    - 方法 POST Content-Type:application/json
    - 参数
        字段|类型|必填|说明
        -- |--|--|--
        user_id|string|否|用户id和三元组必须用一个
        channel|string|否|
        identify_type|string|否
        identifier|string|否
    - 返回参数
        字段|说明
        --|--
        result|success


#### 问题
  - 数据库的创建
  - redis的创建
  - log日志的创建
  - 健康检查
  - 日志追踪
  - 设计模式
  - ticket、token、password加密
  - 数据一致性










































































