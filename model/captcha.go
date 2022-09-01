package model

type Captcha struct {
	KeyId string `json:"key_id"`
	Src   string `json:"src"`
}

type CaptchaReq struct {
	KeyId   string `json:"key_id"`
	Captcha string `json:"captcha"`
}

type CaptchaListReq struct {
	CaptchaList []string `json:"captcha_list"`
}
