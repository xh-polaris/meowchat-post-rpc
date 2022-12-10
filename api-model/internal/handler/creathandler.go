package handler

import (
	"net/http"

	"github.com/xh-polaris/meowchat-post-rpc/api-model/internal/logic"
	"github.com/xh-polaris/meowchat-post-rpc/api-model/internal/svc"
	"github.com/xh-polaris/meowchat-post-rpc/api-model/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func CreatHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.CreatReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := logic.NewCreatLogic(r.Context(), svcCtx)
		resp, err := l.Creat(&req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
