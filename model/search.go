package model

type SearchResult struct {
	AnswerList []*Answer `json:"answerList"`
	QuestionList []*Question `json:"questionList"`
}