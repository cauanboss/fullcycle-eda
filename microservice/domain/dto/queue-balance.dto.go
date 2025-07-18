package dto

type QueueBalanceDTO struct {
	AccountIDFrom        string `json:"account_id_from"`
	AccountIDTo          string `json:"account_id_to"`
	BalanceAccountIDFrom int    `json:"balance_account_id_from"`
	BalanceAccountIDTo   int    `json:"balance_account_id_to"`
}

// Estrutura para receber o evento completo
type BalanceEvent struct {
	Name    string          `json:"Name"`
	Payload QueueBalanceDTO `json:"Payload"`
}
