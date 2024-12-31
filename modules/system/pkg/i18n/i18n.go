// Package i18n
// @Link  https://github.com/huagelong/devinggo
// @Copyright  Copyright (c) 2024 devinggo
// @Author  Kai <hpuwang@gmail.com>
// @License  https://github.com/huagelong/devinggo/blob/master/LICENSE

package i18n

import (
	"devinggo/modules/system/pkg/utils/request"
	"context"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/text/gstr"
)

func InitI18n(ctx context.Context) {
	r := request.GetHttpRequest(ctx)
	lang := r.Header.Get("Accept-Language")
	if g.IsEmpty(lang) {
		lang = "zh-CN"
	} else {
		langarr := gstr.Split(lang, ";")
		lang = langarr[0]
		if gstr.Contains(lang, "en") {
			lang = "en"
		}
	}
	g.I18n().SetPath("resource/i18n")
	g.I18n().SetLanguage(lang)
}

func Newf(ctx context.Context, msg string, params ...string) string {
	v := make([]interface{}, len(params))
	for i, p := range params {
		v[i] = p
	}
	//InitI18n(ctx)
	return g.I18n().Tf(ctx, msg, v...)
}
