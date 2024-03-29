package enum

type InvoiceStatus string

const (
	InvoiceStatusNone  InvoiceStatus = ""
	InvoiceStatusDraft InvoiceStatus = "DRAFT"
	InvoiceStatusDue   InvoiceStatus = "DUE"
	InvoiceStatusPaid  InvoiceStatus = "PAID"
	InvoiceStatusVoid  InvoiceStatus = "VOID"
)

var AllInvoiceStatuses = []InvoiceStatus{
	InvoiceStatusNone,
	InvoiceStatusDraft,
	InvoiceStatusDue,
	InvoiceStatusPaid,
	InvoiceStatusVoid,
}

func DecodeInvoiceStatus(s string) InvoiceStatus {
	if IsValidInvoiceStatus(s) {
		return InvoiceStatus(s)
	}
	return InvoiceStatusNone
}

func IsValidInvoiceStatus(s string) bool {
	for _, ms := range AllInvoiceStatuses {
		if ms == InvoiceStatus(s) {
			return true
		}
	}
	return false
}

func (e InvoiceStatus) String() string {
	return string(e)
}
