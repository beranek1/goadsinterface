package goadsinterface

type AdsDeviceInfo struct {
	Name    string
	Version AdsVersion
}

type AdsError struct {
	Error    string
	ErrorNum int32
}

type AdsState struct {
	Ads    uint16
	Device uint16
}

type AdsSymbol struct {
	Name        string
	IndexGroup  uint32
	IndexOffset uint32
	Size        uint32
	Type        string
	Comment     string
}

type AdsSymbolInfo map[string]AdsSymbol

type AdsSymbolList []string

type AdsData struct {
	Data any
}

type AdsVersion struct {
	Version  uint8
	Revision uint8
	Build    uint16
}

type AdsLibrary interface {
	GetVersion() (AdsVersion, error)
	GetState() (AdsState, error)
	GetDeviceInfo() (AdsDeviceInfo, error)
	GetSymbol(name string) (AdsSymbol, error)
	GetSymbolInfo() (AdsSymbolInfo, error)
	GetSymbolValue(name string) (AdsData, error)
	GetSymbolList() (AdsSymbolList, error)
	SetState(AdsState) (AdsState, error)
	SetSymbolValue(name string, value AdsData) (AdsData, error)
}
