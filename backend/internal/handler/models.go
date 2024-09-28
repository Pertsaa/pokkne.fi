package handler

type SentenceCreate struct {
	Sentence string `json:"sentence"`
}

type SentenceBatchCreate struct {
	Sentences []string `json:"sentences"`
}

type SentenceBatchCreateResponse struct {
	Count int64 `json:"count"`
}

type ChatRequest struct {
	Prompt string `json:"prompt"`
	Mode   string `json:"mode"`
}

type ChatResponse struct {
	Message string `json:"message"`
}
