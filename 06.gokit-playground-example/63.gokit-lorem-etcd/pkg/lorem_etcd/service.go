package lorem_etcd

import (
	"errors"
	"fmt"
	"strings"

	golorem "github.com/drhodes/golorem"
)

var (
	ErrRequestTypeNotFound = errors.New("Request type only valid for word, sentence and paragraph")
)

// Service 定义业务接口
type Service interface {
	Lorem(requestType string, min, max int) (string, error)
}

// LoremService 实现业务逻辑
type LoremService struct {
}

// Lorem 随机文本服务,
// @param: requestType 生成类型, 可选值有Word, Sentence, Paragraph.
// Word: 生成一个最少min个, 最多max个字母的单词
// Sentence: 生成一个最少min个, 最多max个单词的句子
// Paragraph: 生成一个最小min个, 最多max个句子的段落
func (LoremService) Lorem(requestType string, min, max int) (result string, err error) {
	if strings.EqualFold(requestType, "Word") {
		result = golorem.Word(min, max)
	} else if strings.EqualFold(requestType, "Sentence") {
		result = golorem.Sentence(min, max)
	} else if strings.EqualFold(requestType, "Paragraph") {
		result = golorem.Paragraph(min, max)
	} else {
		err = ErrRequestTypeNotFound
	}
	fmt.Println("=====================")
	return
}
