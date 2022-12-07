package model

type Post struct {
	ID              int64
	AttachmentPrice int64
	UserID          int64
}

type PostAttachmentBill struct {
	PostID     int64
	UserID     int64
	PaidAmount int64
}
