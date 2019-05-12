package lorem_restful

import (
	golorem "github.com/drhodes/golorem"
)

// 定义业务接口
type Service interface {
	// 生成一个最少min个, 最多max个字母的单词
	Word(min, max int) string

	// 生成一个最少min个, 最多max个单词的句子
	Sentence(min, max int) string

	// 生成一个最小min个, 最多max个句子的段落
	Paragraph(min, max int) string
}

type LoremService struct {
}

// 实现业务逻辑
func (LoremService) Word(min, max int) string {
	return golorem.Word(min, max)
}

func (LoremService) Sentence(min, max int) string {
	return golorem.Sentence(min, max)
}

func (LoremService) Paragraph(min, max int) string {
	return golorem.Paragraph(min, max)
}
