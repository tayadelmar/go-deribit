package private

type GetSubaccountsResponse []struct {
	Email        string `json:"email" mapstructure:"email"`
	ID           int64  `json:"id" mapstructure:"id"`
	IsPassword   bool   `json:"is_password" mapstructure:"is_password"`
	LoginEnabled bool   `json:"login_enabled" mapstructure:"login_enabled"`
	Portfolio    struct {
		Btc struct {
			AvailableFunds           float64 `json:"available_funds" mapstructure:"available_funds"`
			AvailableWithdrawalFunds float64 `json:"available_withdrawal_funds" mapstructure:"available_withdrawal_funds"`
			Balance                  float64 `json:"balance" mapstructure:"balance"`
			Currency                 string  `json:"currency" mapstructure:"currency"`
			Equity                   float64 `json:"equity" mapstructure:"equity"`
			InitialMargin            float64 `json:"initial_margin" mapstructure:"initial_margin"`
			MaintenanceMargin        float64 `json:"maintenance_margin" mapstructure:"maintenance_margin"`
			MarginBalance            float64 `json:"margin_balance" mapstructure:"margin_balance"`
		} `json:"btc" mapstructure:"btc"`
		Eth struct {
			AvailableFunds           int64  `json:"available_funds" mapstructure:"available_funds"`
			AvailableWithdrawalFunds int64  `json:"available_withdrawal_funds" mapstructure:"available_withdrawal_funds"`
			Balance                  int64  `json:"balance" mapstructure:"balance"`
			Currency                 string `json:"currency" mapstructure:"currency"`
			Equity                   int64  `json:"equity" mapstructure:"equity"`
			InitialMargin            int64  `json:"initial_margin" mapstructure:"initial_margin"`
			MaintenanceMargin        int64  `json:"maintenance_margin" mapstructure:"maintenance_margin"`
			MarginBalance            int64  `json:"margin_balance" mapstructure:"margin_balance"`
		} `json:"eth" mapstructure:"eth"`
	} `json:"portfolio" mapstructure:"portfolio"`
	ReceiveNotifications bool   `json:"receive_notifications" mapstructure:"receive_notifications"`
	SystemName           string `json:"system_name" mapstructure:"system_name"`
	TfaEnabled           bool   `json:"tfa_enabled" mapstructure:"tfa_enabled"`
	Type                 string `json:"type" mapstructure:"type"`
	Username             string `json:"username" mapstructure:"username"`
}