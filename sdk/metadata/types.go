package metadata

// Response types for metadata API

// ResultGetServerTime represents the response for GetServerTime
type ResultGetServerTime struct {
	Code         string            `json:"code"`
	Data         *ServerTime       `json:"data"`
	ErrorParam   map[string]string `json:"errorParam"`
	RequestTime  string            `json:"requestTime"`
	ResponseTime string            `json:"responseTime"`
	TraceId      string            `json:"traceId"`
}

// ServerTime represents server time information
type ServerTime struct {
	TimeMillis string `json:"timeMillis"`
}

// ResultMetaData represents the response for GetMetaData
type ResultMetaData struct {
	Code         string            `json:"code"`
	Data         *MetaData         `json:"data"`
	ErrorParam   map[string]string `json:"errorParam"`
	RequestTime  string            `json:"requestTime"`
	ResponseTime string            `json:"responseTime"`
	TraceId      string            `json:"traceId"`
}

// MetaData represents global metadata
type MetaData struct {
	Global       *Global     `json:"global"`
	CoinList     []Coin      `json:"coinList"`
	ContractList []Contract  `json:"contractList"`
	MultiChain   *MultiChain `json:"multiChain"`
}

// Global represents global metadata information
type Global struct {
	AppName                      string `json:"appName"`
	AppEnv                       string `json:"appEnv"`
	AppOnlySignOn                string `json:"appOnlySignOn"`
	FeeAccountId                 string `json:"feeAccountId"`
	FeeAccountL2Key              string `json:"feeAccountL2Key"`
	PoolAccountId                string `json:"poolAccountId"`
	PoolAccountL2Key             string `json:"poolAccountL2Key"`
	FastWithdrawAccountId        string `json:"fastWithdrawAccountId"`
	FastWithdrawAccountL2Key     string `json:"fastWithdrawAccountL2Key"`
	FastWithdrawMaxAmount        string `json:"fastWithdrawMaxAmount"`
	FastWithdrawRegistryAddress  string `json:"fastWithdrawRegistryAddress"`
	StarkExChainId               string `json:"starkExChainId"`
	StarkExContractAddress       string `json:"starkExContractAddress"`
	StarkExCollateralCoin        *Coin  `json:"starkExCollateralCoin"`
	StarkExMaxFundingRate        int32  `json:"starkExMaxFundingRate"`
	StarkExOrdersTreeHeight      int32  `json:"starkExOrdersTreeHeight"`
	StarkExPositionsTreeHeight   int32  `json:"starkExPositionsTreeHeight"`
	StarkExFundingValidityPeriod int32  `json:"starkExFundingValidityPeriod"`
	StarkExPriceValidityPeriod   int32  `json:"starkExPriceValidityPeriod"`
	MaintenanceReason            string `json:"maintenanceReason"`
}

// Coin represents coin metadata
type Coin struct {
	CoinId            string `json:"coinId"`
	CoinName          string `json:"coinName"`
	StepSize          string `json:"stepSize"`
	ShowStepSize      string `json:"showStepSize"`
	IconUrl           string `json:"iconUrl"`
	StarkExAssetId    string `json:"starkExAssetId"`
	StarkExResolution string `json:"starkExResolution"`
}

// Contract represents contract metadata
type Contract struct {
	ContractId                   string `json:"contractId"`
	ContractName                 string `json:"contractName"`
	BaseCoinId                   string `json:"baseCoinId"`
	QuoteCoinId                  string `json:"quoteCoinId"`
	TickSize                     string `json:"tickSize"`
	StepSize                     string `json:"stepSize"`
	MinOrderSize                 string `json:"minOrderSize"`
	MaxOrderSize                 string `json:"maxOrderSize"`
	MaxOrderBuyPriceRatio        string `json:"maxOrderBuyPriceRatio"`
	MaxOrderSellPriceRatio       string `json:"maxOrderSellPriceRatio"`
	MaxLongLeverage              string `json:"maxLongLeverage"`
	MaxShortLeverage             string `json:"maxShortLeverage"`
	InitialMarginRate            string `json:"initialMarginRate"`
	MaintenanceMarginRate        string `json:"maintenanceMarginRate"`
	FundingRateCoefficient       string `json:"fundingRateCoefficient"`
	FundingRateInterval          int32  `json:"fundingRateInterval"`
	IsOpenPosition               bool   `json:"isOpenPosition"`
	IsOpenTpsl                   bool   `json:"isOpenTpsl"`
	IsOpenConditionalTransfer    bool   `json:"isOpenConditionalTransfer"`
	IsOpenDeleverage             bool   `json:"isOpenDeleverage"`
	IsOpenLiquidate              bool   `json:"isOpenLiquidate"`
	IsOpenAutoDeleverage         bool   `json:"isOpenAutoDeleverage"`
	IsOpenAutoLiquidate          bool   `json:"isOpenAutoLiquidate"`
	IsOpenAutoReducePosition     bool   `json:"isOpenAutoReducePosition"`
	IsOpenAutoReduceMargin       bool   `json:"isOpenAutoReduceMargin"`
	IsOpenAutoReduceCollateral   bool   `json:"isOpenAutoReduceCollateral"`
	IsOpenAutoReduceDebt         bool   `json:"isOpenAutoReduceDebt"`
	IsOpenAutoReduceRisk         bool   `json:"isOpenAutoReduceRisk"`
	IsOpenAutoReduceExposure     bool   `json:"isOpenAutoReduceExposure"`
	IsOpenAutoReduceMarginRate   bool   `json:"isOpenAutoReduceMarginRate"`
	IsOpenAutoReduceDebtRate     bool   `json:"isOpenAutoReduceDebtRate"`
	IsOpenAutoReduceRiskRate     bool   `json:"isOpenAutoReduceRiskRate"`
	IsOpenAutoReduceExposureRate bool   `json:"isOpenAutoReduceExposureRate"`
	StarkExResolution            string `json:"starkExResolution"`
	StarkExSyntheticAssetId      string `json:"starkExSyntheticAssetId"`
	DefaultTakerFeeRate          string `json:"defaultTakerFeeRate"`
}

// MultiChain represents multi-chain withdrawal information
type MultiChain struct {
	CoinId      string  `json:"coinId"`
	MaxWithdraw string  `json:"maxWithdraw"`
	MinWithdraw string  `json:"minWithdraw"`
	MinDeposit  string  `json:"minDeposit"`
	ChainList   []Chain `json:"chainList"`
}

// Chain represents blockchain chain information
type Chain struct {
	Chain              string            `json:"chain"`
	ChainId            string            `json:"chainId"`
	ChainIconUrl       string            `json:"chainIconUrl"`
	ContractAddress    string            `json:"contractAddress"`
	DepositGasFeeLess  bool              `json:"depositGasFeeLess"`
	FeeLess            bool              `json:"feeLess"`
	FeeRate            string            `json:"feeRate"`
	GasLess            bool              `json:"gasLess"`
	GasToken           string            `json:"gasToken"`
	MinFee             string            `json:"minFee"`
	RpcUrl             string            `json:"rpcUrl"`
	WebTxUrl           string            `json:"webTxUrl"`
	WithdrawGasFeeLess bool              `json:"withdrawGasFeeLess"`
	TokenList          []MultiChainToken `json:"tokenList"`
	TxConfirm          string            `json:"txConfirm"`
	BlockTime          string            `json:"blockTime"`
	AllowAaDeposit     bool              `json:"allowAaDeposit"`
	AllowAaWithdraw    bool              `json:"allowAaWithdraw"`
	AllowDeposit       bool              `json:"allowDeposit"`
	AllowWithdraw      bool              `json:"allowWithdraw"`
	AppRpcUrl          string            `json:"appRpcUrl"`
}

// MultiChainToken represents multi-chain token information
type MultiChainToken struct {
	TokenAddress   string `json:"tokenAddress"`
	Decimals       string `json:"decimals"`
	IconUrl        string `json:"iconUrl"`
	Token          string `json:"token"`
	PullOff        bool   `json:"pullOff"`
	WithdrawEnable bool   `json:"withdrawEnable"`
	UseFixedRate   bool   `json:"useFixedRate"`
	FixedRate      string `json:"fixedRate"`
}
