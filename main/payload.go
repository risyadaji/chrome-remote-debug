package main

// Payload ..
type Payload struct {
	BankName    string `json:"bankName"`
	AccountName string `json:"accountName"`
	FileName    string `json:"fileName"`
	UploadLink  string `json:"uploadLink"`
	APIKey      string `json:"apiKey"`
}
