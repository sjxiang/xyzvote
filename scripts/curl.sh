




# 请求 OTP
curl --location --request POST 'localhost:8080/api/v1/otp/verify' \
--header 'Content-Type: application/json' \
--data-raw '{
    "captcha_id": "HusUu0dFNXMjhjMvaxVz",
    "data": "155218"
}'

# 验证 OTP
curl --location --request GET 'localhost:8080/api/v1/otp/gen' 

# 用户注册
curl --location --request POST 'localhost:8080/api/v1/user/register' \
--header 'Content-Type: application/json' \
--data-raw '{
    "username": "shgqmrf19",
    "password": "123456789qwe",
    "confirm_password": "123456789qwe",
    "email": "1535484943@qq.com"
}'

# 用户通过账号登录
curl --location --request POST 'localhost:8080/api/v1/user/login' \
--header 'Content-Type: application/json' \
--data-raw '{
    "username": "shgqmrf19",
    "password": "123456789qwe"
}'

# 用户详情
curl --location --request GET 'localhost:8080/api/v1/user/me' \
--header 'Cookie: credentials=shgqmrf19; path=/' 

# 用户管理后台
curl --location --request POST 'localhost:8080/api/v1/user/admin' \
--header 'Cookie: credentials=shgqmrf19; path=/' \
--header 'Content-Type: application/json' \
--data-raw '{
    "page_id": 1,
    "page_size": 5
}'

# 用户退出
curl --location --request GET 'localhost:8080/api/v1/user/logout' \
--header 'Cookie: credentials=shgqmrf19; path=/' 

