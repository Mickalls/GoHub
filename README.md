项目原地址：https://github.com/summerblue/gohub

### 已开发接口

| 请求方法 |            API 地址             |                  说明                   |
| :------: | :-----------------------------: | :-------------------------------------: |
|   POST   | /api/v1/auth/signup/phone/exist | 验证手机号是否已注册(JSON参数名`phone`) |
|   POST   | /api/v1/auth/signup/email/exist |  验证邮箱是否已注册(JSON参数名`email`)  |

### 技术栈/第三方库

- `gin`：Web框架
- `gorm`：ORM框架
- `zap`：高性能日志库
- `viper`：配置环境
- `cast`：类型转换
- `govalidator`：请求验证器