package handler

import (
	"net/http"

	"github.com/xh-polaris/meowchat-post-rpc/api-model/internal/logic"
	"github.com/xh-polaris/meowchat-post-rpc/api-model/internal/svc"
	"github.com/xh-polaris/meowchat-post-rpc/api-model/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func DestroyHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.DestroyReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := logic.NewDestroyLogic(r.Context(), svcCtx)
		resp, err := l.Destroy(&req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
