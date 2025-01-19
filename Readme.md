# CHX通行证

## 1.简介

CHX同通行证是CHX开发的身份认证方案，通过CHX同通行证，用户可以获得CHX的账号，并使用CHX的账号进行登录。用户可以添加自定义配置完成自己的一些需求。

## 2.Api用法

### /login 登录

> 方法 POST

> 请求体 application/json
```json
{
    "username": "string",
    "password": "string"
}
```

> 响应体 application/json

```json
{
    "message": "string",
    "code": "string",
    "data": {
        "refresh_token": "string",
        "access_token": "string"
    }
}
```

### /register 注册

> 方法 POST

> 请求体 application/json
```json
{
    "username": "string",
    "password": "string",
    "email": "string",
    "custom_config": "string"
}
```
> 响应体 application/json

```json
{
    "message": "string",
    "code": "string",
    "data": {
        "refresh_token": "string",
        "access_token": "string"
    }
}
```

### /refresh 刷新token

> 方法 POST

> 请求体 application/json

```json
{
    "refresh_token": "string"
}
```

> 响应体 application/json

```json
{
    "message": "string",
    "code": "string",
    "data": {
        "access_token": "string"
    }
}
```

### /user/info 获取用户信息

> 方法 GET

> 请求头 Authorization: Bearer {access_token}

> 响应体 application/json

```json
{
    "message": "string",
    "code": "string",
    "data": {
        "userinfo": {}
    }
}
```

### /user/change_info 修改用户信息

> 方法 POST

> 请求头 Authorization: Bearer {access_token}

> 请求体 application/json

```json
{
    "username": "string",
    "email": "string",
    "custom_config": "string",
    "change_pwd_new":: "string"
}
```