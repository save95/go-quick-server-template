package lang

import (
	"context"
	"fmt"

	"server-api/global"
	"server-api/repository/platform/dao"

	"github.com/eko/gocache/v2/store"
	"github.com/go-redis/redis/v8"
	"github.com/save95/go-pkg/model/pager"
	"github.com/save95/xerror"
	"github.com/save95/xerror/xcode"
	"golang.org/x/sync/singleflight"
)

var single singleflight.Group

const (
	singleKey    = "lang"
	msgKeyFormat = "lang:%s:%d"
)

func getMsgKey(language global.Language, code int) string {
	return fmt.Sprintf(msgKeyFormat, language, code)
}

func Init() error {
	ctx := context.Background()

	// 先从 db 中缓存，如果 db 无数据，则使用本地缓存
	if err := initForMysql(ctx); nil != err {
		if err := initForLocal(ctx); nil != err {
			return err
		}
	}

	return nil
}

func initForMysql(ctx context.Context) error {
	start := 0
	limit := 100
	hasMore := true

	langs := []global.Language{
		global.LanguageZhCN,
		global.LanguageZhHK,
		global.LanguageEnUS,
	}

	for hasMore {
		records, _, err := dao.NewLang().Paginate(pager.Option{
			Start: start,
			Limit: limit,
		})
		if nil != err {
			return err
		}

		start += limit
		hasMore = len(records) >= limit
		for _, record := range records {
			for _, language := range langs {
				msg := record.ZhCN
				switch language {
				case global.LanguageZhCN:
					msg = record.ZhCN
				case global.LanguageZhHK:
					msg = record.ZhHK
				case global.LanguageEnUS:
					msg = record.EnUS
				default:
					return xerror.New("undefined lang")
				}

				// 如果未翻译，则跳过
				if len(msg) == 0 {
					continue
				}

				key := getMsgKey(language, record.Code)
				err := global.CacheManager.Set(ctx, key, msg, &store.Options{
					Tags: []string{"lang", string(language)},
				})
				if err != nil {
					//global.Log.Errorf("lang set failed: lang=%s, code=%d, msg=%s", language, k, msg)
					return xerror.Wrapf(err, "lang set failed: lang=%s, code=%d, msg=%s", language, record.Code, msg)
				}
			}

		}
	}

	return nil
}

func initForLocal(ctx context.Context) error {
	languages := map[global.Language]map[int]string{
		//global.LanguageZhCN: zh_CN,
		global.LanguageZhHK: zh_HK,
		global.LanguageEnUS: en_US,
	}

	for language, m := range languages {
		for k, msg := range m {
			key := getMsgKey(language, k)
			err := global.CacheManager.Set(ctx, key, msg, &store.Options{
				Tags: []string{"lang", string(language)},
			})
			if err != nil {
				//global.Log.Errorf("lang set failed: lang=%s, code=%d, msg=%s", language, k, msg)
				return xerror.Wrapf(err, "lang set failed: lang=%s, code=%d, msg=%s", language, k, msg)
			}
		}
	}

	return nil
}

// Handle 语言包处理器
func Handle() func(code int, language string) string {
	ctx := context.Background()

	return func(code int, language string) string {
		v, err, _ := single.Do(singleKey, func() (interface{}, error) {
			lang := global.Language(language)
			key := getMsgKey(lang, code)
			return global.CacheManager.Get(ctx, key)
		})

		if nil != err {
			if err != redis.Nil {
				global.Log.Errorf("get lang failed: lang=%s, err=%+v", language, err)
			}
			return ""
		}

		return v.(string)
	}
}

// GetContent 获得语言包对应文字
// 请求头中必须包含 `X-Use-Language` 才可以
func GetContent(ctx context.Context, code xcode.XCode) string {
	header, err := global.MustParseAPPHeader(ctx)
	if nil != err {
		global.Log.Errorf("parse header failed in lang.GetContent: %+v", err)
		return ""
	}

	msg := Handle()(code.Code(), header.UseLanguage)
	if len(msg) == 0 {
		return code.String()
	}
	return msg
}
