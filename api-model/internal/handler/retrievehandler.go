package handler

import (
	"net/http"

	"github.com/xh-polaris/meowchat-post-rpc/api-model/internal/logic"
	"github.com/xh-polaris/meowchat-post-rpc/api-model/internal/svc"
	"github.com/xh-polaris/meowchat-post-rpc/api-model/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func RetrieveHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.RetrieveReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := logic.NewRetrieveLogic(r.Context(), svcCtx)
		resp, err := l.Retrieve(&req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
