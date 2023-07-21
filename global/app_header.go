package global

import (
	"net/url"
)

// appHeaderI18N 国际化相关参数
// 其值为标准的国际化标识，详细参考：[附：国际化语言标识](../docs/i18n.md)
type appHeaderI18N struct {
	UseLanguage string `header:"X-Use-Language" form:"ul"` // 使用语言，默认是英文
}

// appHeaderDevice 设备相关参数
type appHeaderDevice struct {
	DeviceModel string `header:"X-Device-Model" form:"db"` // 设备型号。如：iPhone 14 Pro
	DeviceBand  string `header:"X-Device-Band" form:"db"`  // 设备厂商品牌。如：xiaomi, huawei
	OSVersion   string `header:"X-Os-Version" form:"osv"`  // 操作系统版本。如：ios 16.5, android 10.8
	AppVersion  string `header:"X-App-Version" form:"arv"` // APP 版本号，格式：app版本/RN版本，如：1.0/0.1.202039271
}

// appHeaderCustomArgs 自定义参数
type appHeaderCustomArgs struct {
	ChannelFlag string `header:"X-Channel-Flag" form:"cf"` // 来源渠道
}

// appHeaderUTM 广告统计分析相关参数
// @see https://www.shangyexinzhi.com/article/4258360.html
// @see https://zhuanlan.zhihu.com/p/378091279
type appHeaderUTM struct {
	UTMSourceEncode   string `header:"X-Utm-Source" form:"utm_source"`     // 广告投放来源
	UTMMediumEncode   string `header:"X-Utm-Medium" form:"utm_medium"`     // 广告投放媒介：cpc
	UTMCampaignEncode string `header:"X-Utm-Campaign" form:"utm_campaign"` // 广告投放名称
	UTMTermEncode     string `header:"X-Utm-Term" form:"utm_term"`         // 广告投放字词关键字
	UTMContentEncode  string `header:"X-Utm-Content" form:"utm_content"`   // 广告投放内容
}

// APPHeader  通用请求 header
type APPHeader struct {
	appHeaderI18N
	appHeaderDevice
	appHeaderCustomArgs
	appHeaderUTM
}

func (c APPHeader) Language() Language {
	if len(c.UseLanguage) == 0 {
		return LanguageZhCN
	}

	return Language(c.UseLanguage)
}

func (c APPHeader) ChannelID() uint {
	// todo 解析来源渠道
	return 0
}

func (c APPHeader) UTMSource() string {
	str, _ := url.QueryUnescape(c.UTMSourceEncode)
	return str
}

func (c APPHeader) UTMMedium() string {
	str, _ := url.QueryUnescape(c.UTMMediumEncode)
	return str
}

func (c APPHeader) UTMCampaign() string {
	str, _ := url.QueryUnescape(c.UTMCampaignEncode)
	return str
}

func (c APPHeader) UTMTerm() string {
	str, _ := url.QueryUnescape(c.UTMTermEncode)
	return str
}

func (c APPHeader) UTMContent() string {
	str, _ := url.QueryUnescape(c.UTMContentEncode)
	return str
}
