package global

import (
	"sync"

	"github.com/mojocn/base64Captcha"
)

var (
	captchaConfig     *CaptchaConfig
	captchaConfigOnce sync.Once
)

type CaptchaConfig struct {
	CaptchaType     string
	VerifyKey       string
	ConfigAudio     base64Captcha.ConfigAudio
	ConfigCharacter base64Captcha.ConfigCharacter
	ConfigDigit     base64Captcha.ConfigDigit
}

// GetCaptchaConfig 获取base64验证码基本配置
func GetCaptchaConfig() *CaptchaConfig {
	captchaConfigOnce.Do(func() {
		captchaConfig = &CaptchaConfig{
			CaptchaType: "character",
			VerifyKey:   "",
			ConfigAudio: base64Captcha.ConfigAudio{},
			ConfigCharacter: base64Captcha.ConfigCharacter{
				Height:             60,
				Width:              160,
				Mode:               base64Captcha.CaptchaModeNumberAlphabet,
				IsUseSimpleFont:    true,
				ComplexOfNoiseText: 0,
				ComplexOfNoiseDot:  0,
				IsShowHollowLine:   true,
				IsShowNoiseDot:     true,
				IsShowNoiseText:    true,
				IsShowSlimeLine:    true,
				IsShowSineLine:     false,
				CaptchaLen:         4,
			},
			ConfigDigit: base64Captcha.ConfigDigit{},
		}
	})

	return captchaConfig
}
