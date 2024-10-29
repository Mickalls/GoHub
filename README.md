项目原地址：https://github.com/summerblue/gohub

### 已开发接口

| 请求方法 |              API 地址               |                   说明                   |
| :------: |:---------------------------------:| :--------------------------------------: |
|   `POST`   | `/api/v1/auth/signup/phone/exist` | 验证手机号是否已注册(JSON参数名`phone`)  |
|   `POST`   |  `/api/v1/auth/signup/email/exist`  |  验证邮箱是否已注册(JSON参数名`email`)   |
|   `POST`   | `/api/v1/auth/verify-codes/captcha` | 生成图像验证码并`base64`编码后返回给客户端 |



### 技术栈/第三方库

- `gin`：Web框架
- `gorm`：ORM框架
- `zap`：高性能日志库
- `viper`：配置环境
- `cast`：类型转换
- `govalidator`：请求验证器
- `Redis`：缓存
- `captcha`：图片验证码库
- `aliyum`的`SMS`服务`SDK`：发送短信验证