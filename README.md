# gin-plugin-captcha
## 模式
model == "library"
model == "original"

form == "image": 生成图片, base64
from == "email": 生成邮件html内容

返回结果 {"code": 0, "data": {"captcha_id": "xxx-xxx-xxx-xxx-xxx", "src": "base64 | html"}, "msg": "操作成功"}