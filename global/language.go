package global

import "github.com/save95/xerror"

type Language string

// 支持的国际化语言标识
// @see [附：国际化语言标识](../docs/i18n.md)
const (
	LanguageEnUS Language = "en_US" // 英语(美国)
	LanguageZhCN Language = "zh_CN" // 简体中文(中国)
	LanguageZhHK Language = "zh_HK" // 繁体中文(香港)
)

var supportedLanguages = []Language{
	LanguageEnUS,
	LanguageZhCN,
	LanguageZhHK,
}

func NewLanguage(s string) (Language, error) {
	for _, language := range supportedLanguages {
		if language == Language(s) {
			return language, nil
		}
	}

	return "", xerror.New("不支持的语言")
}
