# gin-plugin-captcha
基于Redis缓存的gin框架插件，适用于获取和构建验证码，现支持字符串、base64图片的生成

## 说明
验证码现有6种模式，默认保存时间10分钟

### 原生字符串
mode=original 使用 字母+数字+部分特殊字符的混合，默认长度为6， 例如 "4I0KYX"、"NiqGyg"、"IbZYt@"等

### 语料库
mode=library 使用redis内缓存的语料库进行获取验证码，可存放成语等，长度限制在6个字符以内， 例如 "三羊开泰"、"六六大顺"、"福如东海"等

### 数字
mode=digit 使用0~9的数字进行随机组合，长度由用户指定，默认长度为6 例如：1777、25894、654289

### 语料库内容base64图片
mode=original_image 使用随机生成的原生字符串进行绘制图片，并制成base64返回
例如：

![语料库内容base64图片](https://github.com/iddxc/gin-plugin-captcha/blob/main/images/base64-library.png)


### 原生字符串base64图片
mode=library_image 使用随机获取语料库中的内容进行绘制图片，并制成base64返回
例如：

![原生字符串base64图片](https://github.com/iddxc/gin-plugin-captcha/blob/main/images/base64-org.png)


### 数字base64图片
mode=original_image 使用随机生成的数字进行绘制图片，并制成base64返回
例如：

![数字base64图片](https://github.com/iddxc/gin-plugin-captcha/blob/main/images/base64-digit.png)

## 使用

### 载入插件
运行指令`go get github.com/GPorter-t/gin-plugin-captcha`进行安装。

在 gin 项目的初始化过程中加入 plugin的注册项，如果已经设置好该步骤，则直接调用即可。

**示例参考[gin-vue-admin](https://github.com/flipped-aurora/gin-vue-admin/blob/main/server/initialize/plugin.go)的插件安装**

```go
package initialize

import (
	"Noteus/global"
	"Noteus/plugin/email"
	"Noteus/utils/plugin"
	"fmt"

	"github.com/gin-gonic/gin"
	captcha "github.com/iddxc/gin-plugin-captcha"
	chatroom "github.com/iddxc/gin-plugin-chatroom"
)

func PluginInit(group *gin.RouterGroup, Plugin ...plugin.Plugin) {
	for _, item := range Plugin {
		PluginGroup := group.Group(item.RouterPath())
		item.Register(PluginGroup)
	}
}

func InstallPlugin(Router *gin.Engine) {
	PublicGroup := Router.Group("")
	fmt.Println("无鉴权插件安装==>", PublicGroup)
	PrivateGroup := Router.Group("")
	fmt.Println("鉴权插件安装==>", PrivateGroup)
	// PrivateGroup.User()
	PluginInit(PrivateGroup, email.CreateEmailPlug(
		global.GVA_CONFIG.Email.To,
		global.GVA_CONFIG.Email.From,
		global.GVA_CONFIG.Email.Host,
		global.GVA_CONFIG.Email.Secret,
		global.GVA_CONFIG.Email.Nickname,
		global.GVA_CONFIG.Email.Port,
		global.GVA_CONFIG.Email.IsSSL,
	))

    // 验证码插件，需要传入redis配置信息 和 字体文件路径
	PluginInit(PrivateGroup, captcha.CreateCaptchaPlugin(
		global.GVA_CONFIG.Redis.Addr,
		global.GVA_CONFIG.Redis.Password,
		"D:\\FZSTK.TTF",
		global.GVA_CONFIG.Redis.DB,
	))
}

```

### 调用
#### 获取验证码
```python
import requests

def test_get_captcha(mode):
    url = f"http://localhost:8081/captcha?mode={mode}&length=4"

    # url = f"http://localhost:8081/captcha?mode={mode}&width={width}&height={height}&fontsize={fontsize}"
    rsp = requests.get(url)
    print(mode, rsp.text)

if __name__ == "__main__":
    modes = ["library", "original", "digit", "library_image", "original_image", "digit_image"]
    for mode in modes:
        test_get_captcha(mode)

"""
library {"code":0,"data":{"key_id":"f0377c8c-e48c-4770-b0bd-290f9bead925","src":""},"msg":"查询成功"}
original {"code":0,"data":{"key_id":"72ef2c93-6a10-4c4c-b47e-548e07ef9d64","src":""},"msg":"查询成功"}
digit {"code":0,"data":{"key_id":"6d638f7d-7496-41fb-81d7-feb6a1d1b565","src":""},"msg":"查询成功"}
library_image {"code":0,"data":{"key_id":"12f3d886-2116-4e0c-ab77-ff4417a15a32","src":"data:image/png;base64,iVBORw0KGgoAAAANSUhEUgAAAMgAAABQCAMAAABrs8qCAAAAP1BMVEUAAADMBpLMBpLoIq70LrrSDJjXEZ3RC5faFKDeGKTzLbnRC5fOCJTLBZHuKLT5M7/KBJDJA4/aFKDmIKzoIq5CikRGAAAAAXRSTlMAQObYZgAAAj5JREFUeJzsmu12oyAQhnfrqMiJTt7m/u91D/hRyTZBPtRp5fnRH3iI75MZIE38w7+EIiKNa4rQfjmSCRGhkR3TJHDJijh85s2RzLPIfeO8z4UdQsXwJHL/wjNRisBMbEXEcc1zRALNmtX4jxN5RRGRRhGRRhGRRhGRRhGRRhGRRhGRRhGRxiYRhRcXgCbyvr0hcu63bBLBJFIZllENQ+R95YhUAAath8iS9H7CXnCLSGNEdDuJeAgRCYv6Hr+IMg1EZLpoTqqAoTU1aZN6K3Lm93hFagA3qNpGHj3q2nSb7tABKmuaBHwiDxATAxVrZ9xoDUAPoN414FZ8ItBcV2zb6unKuGkZHjvm24xHBB23AP992mjpy0ED3d4ht/BepIJpopF1B00SQMvxy33mw5D4Gh4RAi0ivGqgMfzDDMee7CsOEOnnapBziBMwcGeHcnTVRwCRIjarPQ6VK3K3jg31yX11TEUYeEwG/32saudqJUfgHB6+XYsAnkXc33RhzhZWORZJFny7FngRca7YYmjKUZA8+FpL2TPR5h6cC5SrsTLxQmRcgDco3NCYBqudN99aLSc7tn7g":"查询成功"}
original_image {"code":0,"data":{"key_id":"2aaf4d3a-6e7d-43d3-a271-9325b2944787","src":"data:image/png;base64,iVBORw0KGgoAAAANSUhEUgAAAMgAAABQCAMAAABrs8qCAAAAP1BMVEUAAACvrneFhE1OTRZ6eUJMSxRTUhtRUBl0czxlZC2Hhk+GhU5+fUZwbziEg0xBQAk9PAVXVh92dT5xcDlPThcFmkHwAAAAAXRSTlMAQObYZgAAAitJREFUeJzsmfFyoyAQh4/kdMpODD+8vP+z3kDUCIg1GtIl3e+fijJ0P3cBo38Ud6y1W7o9J0IJe8PbjrWbVPhnZGNKKhDZhohwQ0S4ISI/zckxaxcVuRUc+50itzmvHtyLdKeByjPSzah2jjgT1T2a9YpEPCXSR5QL63l+V0ZOCeUDe5bflZEaKCLyph/BASKyRvqOovyrCsnIKu/3kFWLHSLCDRHhhohwQ0QU6b/ZS7sH3c9+EYPw+4t+tNEcC2oP+0U0wpcP3dTugYNR7WBVhPqelMp8+CKEFXQd2mQQXtHHo9zAiggBMECP6/L1SITuaWgRlZwd+l1z47yGvAjBKF9ByMzdKDAnQg0QJ4DuhUaF6y0v0gz/2GZFvoImoQWwsJJ5heG2lCMvMsYZz4VHh/DeE5DLnkWXJOrVrIhQfBB3aIOmnyPtUkpcdrsDMW5ii0g4F6ZIcQ7O3ye7xsIip5Gtz5exIjLeRPybn+7GOUtJaflYL2l5nUGqK22SF2mHiE145+0oYpJ95OL/Urxw3Tu2KPvyfmUfATSRq4pg3ST4gCl+QlGX8YTbf/R8GBr+Fn1wWdvZvwzQkNLRxmfhMMn29oi+ma0DtylxjSlTXa3jEx7jD4jw+jziRdo9Iul3n+8pIDDyMRlxJvJTlx8iwg0R4YaIcENEeKAd/khEeKAnPkHEH1UuokaP6kUmRIQbIsINEeGGiHBDRLghIj/HeZEKRZb5HwAA//+KjVJepRDv5QAAAABJRU5ErkJggg=="},"msg":"查询成功"}
digit_image {"code":0,"data":{"key_id":"b1bd66a8-20d2-4c85-8cad-24325c763d56","src":"data:image/png;base64,iVBORw0KGgoAAAANSUhEUgAAAMgAAABQCAMAAABrs8qCAAAAP1BMVEUAAABwNMg/A5c+ApZmKr5LD6NiJrpKDqJdIbVaHrI/A5dzN8tZHbFZHbFmKr5KDqJlKb1xNcltMcVhJbliJrpxVmC8AAAAAXRSTlMAQObYZgAAAYZJREFUeJzsme1ygkAMRVuClgr9YOn7P2tnOiNFJLjJZtcL5PxyLkZy3CwqvrQ7Yacinw+pZ8wOP81jryuyYVwEDV4kFO0jGVYkjJRtSMv+V2RruAgaLoKGi6DhImi4CBo4In+/nPXlLmLO/LbGP1HlYCL6chyRNskDSSQNF0HDRdBwETSOIVK/LoRLmSY3ZlWE6L6LQCNJuTVrIiei/i7siOjUDKEjOi/kTWxuDStSr76JIY5JxUCdefNTWJGeqn5lGsLsWHV9wNVknqxHe4Q9VN1OXRWLYeu3KEXCfFAqJueeb45SRDpAwsF6i2Ys0YkM9G6SG6ITSVuQDwGRGkqRMy1vWmk+8iPBUITo2yRvv0SstaoS4T5fpLkpGpGGmRRpbkrpr/GXXC+cR4S94FzkRJ6ysMgmV0RM0ilLr0g2Mm324h4HufmwJQ4oMvkjpsnXjxqNSJMAioiabAJXVKOFiGCzQ3sc8aoFjoug4SJouAgaLoKGi6DhImi4CBougsZvAAAA//9YWKoAJAifvwAAAABJRU5ErkJggg=="},"msg":"查询成功"}
"""
```

#### 验证验证码
```python
import requests

data = {
    "key_id": "",
    "captcha": "",
}


def test_verify_captcha(key_id, captcha):
    data["key_id"] = key_id
    data["captcha"] = captcha
    url = f"http://localhost:8081/captcha"
    rsp = requests.post(url, json=data)
    print(rsp.text)


if __name__ == "__main__":
    test_verify_captcha("b1bd66a8-20d2-4c85-8cad-24325c763d57", "1777") # 错误key_id, 正确验证码:{"code":7,"data":{},"msg":"redis: nil"}
    test_verify_captcha("b1bd66a8-20d2-4c85-8cad-24325c763d56", "1776") # 正确key_id, 错误验证码:{"code":7,"data":{},"msg":"验证码错误请重试"}
    test_verify_captcha("b1bd66a8-20d2-4c85-8cad-24325c763d56", "1777") # 正确key_id, 正确验证码:{"code":0,"data":"ok","msg":"操作成功"}
    test_verify_captcha("b1bd66a8-20d2-4c85-8cad-24325c763d56", "1777") # 正确key_id, 正确验证码:{"code":7,"data":{},"msg":"redis: nil"}

```

#### 提交自定义验证码至语料库
```python
import requests

data = {
    "captcha_list": []
}

def test_post_captcha(*args):
    data["captcha_list"] = [i for i in args if isinstance(i, str)]
    url = f"http://localhost:8081/captcha/library"
    rsp = requests.post(url, json=data)
    print(rsp.text)

if __name__ == "__main__":
    test_post_captcha("我就试一试", "看看这个能不能够使用", "66666666", "自动截取前6位") # Redis items: ["我就试一试", "看看这个能不", "666666", "自动截取前6"]：{"code":0,"data":"ok","msg":"操作成功"}

```


## 设计思路
每次编写web服务时，如涉及账号操作，经常用到验证码，但是对编写或者迁移验证码生成的接口的时候，有些烦躁，于是产生了编写插件的想法，经过2天的编写，完成了初期验证码的内容


## 适用范围
- web服务 登录、重置密码等操作过程中的验证码图片，使用base64进行渲染
- 验证码email 模板中验证码字符串 or 字符串验证码的的渲染
- 短信验证码的生成
- 微信公众号 对接web服务的口令获取


## 可优化项
- [ ] 修正项，任何模式操作的验证码内容均保持在长度6之内
- [ ] 随机四则运算的验证码生成，如 43 + 27 = 
- [ ] 位置坐标验证码的生成，如 点击所需汉字的位置
- [ ] 物体信息验证码的生成，如 根据图片信息输入物体名称