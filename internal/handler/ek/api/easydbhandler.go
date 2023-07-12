package api

import (
	"git.code.tencent.com/JXKJ-go/go-zen/core/eerrors"
	"git.code.tencent.com/JXKJ-go/go-zen/pkg/baseutil"
	"net/http"

	"easyStart/internal/logic/ek/api"
	"easyStart/internal/svc"
	"easyStart/pkg/response"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func EasyDbHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := api.NewEasyDbLogic(r.Context(), svcCtx)
		logx.Info(baseutil.GetRequestParam(r))
		resp, err := l.EasyDb()
		if err != nil {
			param, e := baseutil.GetRequestParam(r)
			logx.Infow("EasyDbHandler", logx.Field("err", err),
				//logx.Field("session", svcCtx.GetSession().GetAllSession()),
				logx.Field("请求参数", param),
				logx.Field("请求参数err", e),
			)
			if eerr, ok := err.(*eerrors.GoError); ok {
				httpx.OkJsonCtx(r.Context(), w, response.Resp{ReturnCode: eerr.GetReason(), ReturnMsg: eerr.GetMessage()})
				return
			}
			httpx.OkJsonCtx(r.Context(), w, response.Resp{ReturnCode: "FAIL", ReturnMsg: err.Error()})
		} else {
			httpx.OkJsonCtx(r.Context(), w, response.Resp{ReturnCode: "SUCCESS", ReturnMsg: resp})
		}
	}
}
