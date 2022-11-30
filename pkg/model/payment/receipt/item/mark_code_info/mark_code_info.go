package mark_code_info //nolint:revive // better package readability

type MarkCodeInfo struct {
	MarkCodeRaw string `json:"mark_code_raw"`
	Unknown     string `json:"unknown"`
	Ean8        string `json:"ean_8"`
	Ean13       string `json:"ean_13"`
	Itf14       string `json:"itf_14"`
	Gs10        string `json:"gs_10"`
	Gs1m        string `json:"gs_1m"`
	Short       string `json:"short"`
	Fur         string `json:"fur"`
	Egais20     string `json:"egais_20"`
	Egais30     string `json:"egais_30"`
}

func NewMarkCodeInfo() *MarkCodeInfo {
	return &MarkCodeInfo{}
}

func (i *MarkCodeInfo) WithMarkCodeRaw(markCodeRaw string) *MarkCodeInfo {
	i.MarkCodeRaw = markCodeRaw
	return i
}

func (i *MarkCodeInfo) WithUnknown(unknown string) *MarkCodeInfo {
	i.Unknown = unknown
	return i
}

func (i *MarkCodeInfo) WithEan8(ean8 string) *MarkCodeInfo {
	i.Ean8 = ean8
	return i
}

func (i *MarkCodeInfo) WithEan13(ean13 string) *MarkCodeInfo {
	i.Ean13 = ean13
	return i
}

func (i *MarkCodeInfo) WithItf14(itf14 string) *MarkCodeInfo {
	i.Itf14 = itf14
	return i
}

func (i *MarkCodeInfo) WithGs10(gs10 string) *MarkCodeInfo {
	i.Gs10 = gs10
	return i
}

func (i *MarkCodeInfo) WithGs1m(gs1m string) *MarkCodeInfo {
	i.Gs1m = gs1m
	return i
}

func (i *MarkCodeInfo) WithShort(short string) *MarkCodeInfo {
	i.Short = short
	return i
}

func (i *MarkCodeInfo) WithFur(fur string) *MarkCodeInfo {
	i.Fur = fur
	return i
}

func (i *MarkCodeInfo) WithEgais20(egais20 string) *MarkCodeInfo {
	i.Egais20 = egais20
	return i
}

func (i *MarkCodeInfo) WithEgais30(egais30 string) *MarkCodeInfo {
	i.Egais30 = egais30
	return i
}
