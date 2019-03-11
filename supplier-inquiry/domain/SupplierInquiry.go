package domain

type BusinessId struct {
	Id string
}

type SupplierInquiry struct {
	BusinessAddressLine string
	MemberUserId        int
	CreationTime        string
	BusinessId          BusinessId `json:"businessId"`
	BusinessName        string
	ProductOrderId      int
	//"rowModified":null,
	TransactionId int
	//"completionTime":null,
	ReferenceNumber string
	CtoKey          int    `json:"ctoaId"`
	Id              string `json:"id"`
	//"rowCreated":null,
	MemberId int
	Status   string
}

type SupplierInquiryList struct {
	TotalCount        int
	SupplierInquiries []SupplierInquiry
}
