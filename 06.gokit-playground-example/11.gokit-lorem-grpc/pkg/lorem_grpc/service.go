package lorem_grpc

import (
	"context"
	"errors"
	"strings"

	// 这里需要重申导入名称, 应该是在golorem内部的package名称为lorem而不是golorem.
	golorem "github.com/drhodes/golorem"
)

var (
	ErrRequestTypeNotFound = errors.New("Request type only valid for word, sentence and paragraph")
)

// Service 定义业务接口
type Service interface {
	// 其实就是将之前项目中的word, sentence, paragraph3个方法合并了.
	Lorem(ctx context.Context, requestType string, min, max int) (string, error)
}

// LoremService 实现业务逻辑
type LoremService struct {
}

func (LoremService) Lorem(_ context.Context, requestType string, min, max int) (result string, err error) {
	if strings.EqualFold(requestType, "Word") {
		result = golorem.Word(min, max)
	} else if strings.EqualFold(requestType, "Sentence") {
		result = golorem.Sentence(min, max)
	} else if strings.EqualFold(requestType, "Paragraph") {
		result = golorem.Paragraph(min, max)
	} else {
		err = ErrRequestTypeNotFound
	}
	return 
}
