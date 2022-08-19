package controller

import (
	"github.com/gin-gonic/gin"
)

type Authorize struct {
}

// 生成全局唯一访问凭证 ACCESS-TOKEN
func (s Authorize) Token(ctx *gin.Context) {

}
