package db

// Code generated by xo. DO NOT EDIT.

type ExtraSchemaTableEntity string

const (
	ExtraSchemaTableEntityExtraSchemaBook                    ExtraSchemaTableEntity = "extra_schema.books"
	ExtraSchemaTableEntityExtraSchemaBookAuthor              ExtraSchemaTableEntity = "extra_schema.book_authors"
	ExtraSchemaTableEntityExtraSchemaBookAuthorsSurrogateKey ExtraSchemaTableEntity = "extra_schema.book_authors_surrogate_key"
	ExtraSchemaTableEntityExtraSchemaBookReview              ExtraSchemaTableEntity = "extra_schema.book_reviews"
	ExtraSchemaTableEntityExtraSchemaBookSeller              ExtraSchemaTableEntity = "extra_schema.book_sellers"
	ExtraSchemaTableEntityExtraSchemaDemoWorkItem            ExtraSchemaTableEntity = "extra_schema.demo_work_items"
	ExtraSchemaTableEntityExtraSchemaDummyJoin               ExtraSchemaTableEntity = "extra_schema.dummy_join"
	ExtraSchemaTableEntityExtraSchemaNotification            ExtraSchemaTableEntity = "extra_schema.notifications"
	ExtraSchemaTableEntityExtraSchemaPagElement              ExtraSchemaTableEntity = "extra_schema.pag_element"
	ExtraSchemaTableEntityExtraSchemaUser                    ExtraSchemaTableEntity = "extra_schema.users"
	ExtraSchemaTableEntityExtraSchemaUserAPIKey              ExtraSchemaTableEntity = "extra_schema.user_api_keys"
	ExtraSchemaTableEntityExtraSchemaWorkItem                ExtraSchemaTableEntity = "extra_schema.work_items"
	ExtraSchemaTableEntityExtraSchemaWorkItemAdmin           ExtraSchemaTableEntity = "extra_schema.work_item_admin"
	ExtraSchemaTableEntityExtraSchemaWorkItemAssignee        ExtraSchemaTableEntity = "extra_schema.work_item_assignee"
)

var ExtraSchemaEntityFields = map[ExtraSchemaTableEntity]map[string]DbField{
	ExtraSchemaTableEntityExtraSchemaBook: {
		"bookID": DbField{Type: ColumnSimpleTypeInteger, Db: "book_id", Nullable: false, Public: true},
		"name":   DbField{Type: ColumnSimpleTypeString, Db: "name", Nullable: false, Public: true},
	},
	ExtraSchemaTableEntityExtraSchemaBookAuthor: {
		"authorID":  DbField{Type: ColumnSimpleTypeString, Db: "author_id", Nullable: false, Public: true},
		"bookID":    DbField{Type: ColumnSimpleTypeInteger, Db: "book_id", Nullable: false, Public: true},
		"pseudonym": DbField{Type: ColumnSimpleTypeString, Db: "pseudonym", Nullable: true, Public: true},
	},
	ExtraSchemaTableEntityExtraSchemaBookAuthorsSurrogateKey: {
		"authorID":                  DbField{Type: ColumnSimpleTypeString, Db: "author_id", Nullable: false, Public: true},
		"bookAuthorsSurrogateKeyID": DbField{Type: ColumnSimpleTypeInteger, Db: "book_authors_surrogate_key_id", Nullable: false, Public: true},
		"bookID":                    DbField{Type: ColumnSimpleTypeInteger, Db: "book_id", Nullable: false, Public: true},
		"pseudonym":                 DbField{Type: ColumnSimpleTypeString, Db: "pseudonym", Nullable: true, Public: true},
	},
	ExtraSchemaTableEntityExtraSchemaBookReview: {
		"bookID":       DbField{Type: ColumnSimpleTypeInteger, Db: "book_id", Nullable: false, Public: true},
		"bookReviewID": DbField{Type: ColumnSimpleTypeInteger, Db: "book_review_id", Nullable: false, Public: true},
		"reviewer":     DbField{Type: ColumnSimpleTypeString, Db: "reviewer", Nullable: false, Public: true},
	},
	ExtraSchemaTableEntityExtraSchemaBookSeller: {
		"bookID": DbField{Type: ColumnSimpleTypeInteger, Db: "book_id", Nullable: false, Public: true},
		"seller": DbField{Type: ColumnSimpleTypeString, Db: "seller", Nullable: false, Public: true},
	},
	ExtraSchemaTableEntityExtraSchemaDemoWorkItem: {
		"checked":    DbField{Type: ColumnSimpleTypeBoolean, Db: "checked", Nullable: false, Public: true},
		"workItemID": DbField{Type: ColumnSimpleTypeInteger, Db: "work_item_id", Nullable: false, Public: true},
	},
	ExtraSchemaTableEntityExtraSchemaDummyJoin: {
		"dummyJoinID": DbField{Type: ColumnSimpleTypeInteger, Db: "dummy_join_id", Nullable: false, Public: true},
		"name":        DbField{Type: ColumnSimpleTypeString, Db: "name", Nullable: true, Public: true},
	},
	ExtraSchemaTableEntityExtraSchemaNotification: {
		"body":           DbField{Type: ColumnSimpleTypeString, Db: "body", Nullable: false, Public: true},
		"notificationID": DbField{Type: ColumnSimpleTypeInteger, Db: "notification_id", Nullable: false, Public: true},
		"receiver":       DbField{Type: ColumnSimpleTypeString, Db: "receiver", Nullable: true, Public: true},
		"sender":         DbField{Type: ColumnSimpleTypeString, Db: "sender", Nullable: false, Public: true},
	},
	ExtraSchemaTableEntityExtraSchemaPagElement: {
		"createdAt":          DbField{Type: ColumnSimpleTypeDateTime, Db: "created_at", Nullable: false, Public: true},
		"dummy":              DbField{Type: ColumnSimpleTypeInteger, Db: "dummy", Nullable: true, Public: true},
		"name":               DbField{Type: ColumnSimpleTypeString, Db: "name", Nullable: false, Public: true},
		"paginatedElementID": DbField{Type: ColumnSimpleTypeString, Db: "paginated_element_id", Nullable: false, Public: true},
	},
	ExtraSchemaTableEntityExtraSchemaUser: {
		"apiKeyID":  DbField{Type: ColumnSimpleTypeInteger, Db: "api_key_id", Nullable: true, Public: true},
		"createdAt": DbField{Type: ColumnSimpleTypeDateTime, Db: "created_at", Nullable: false, Public: true},
		"deletedAt": DbField{Type: ColumnSimpleTypeDateTime, Db: "deleted_at", Nullable: true, Public: true},
		"name":      DbField{Type: ColumnSimpleTypeString, Db: "name", Nullable: false, Public: true},
		"userID":    DbField{Type: ColumnSimpleTypeString, Db: "user_id", Nullable: false, Public: true},
	},
	ExtraSchemaTableEntityExtraSchemaUserAPIKey: {
		"apiKey":       DbField{Type: ColumnSimpleTypeString, Db: "api_key", Nullable: false, Public: true},
		"expiresOn":    DbField{Type: ColumnSimpleTypeDateTime, Db: "expires_on", Nullable: false, Public: true},
		"userAPIKeyID": DbField{Type: ColumnSimpleTypeInteger, Db: "user_api_key_id", Nullable: false, Public: true},
		"userID":       DbField{Type: ColumnSimpleTypeString, Db: "user_id", Nullable: false, Public: true},
	},
	ExtraSchemaTableEntityExtraSchemaWorkItem: {
		"description": DbField{Type: ColumnSimpleTypeString, Db: "description", Nullable: true, Public: true},
		"title":       DbField{Type: ColumnSimpleTypeString, Db: "title", Nullable: true, Public: true},
		"workItemID":  DbField{Type: ColumnSimpleTypeInteger, Db: "work_item_id", Nullable: false, Public: true},
	},
	ExtraSchemaTableEntityExtraSchemaWorkItemAdmin: {
		"admin":      DbField{Type: ColumnSimpleTypeString, Db: "admin", Nullable: false, Public: true},
		"workItemID": DbField{Type: ColumnSimpleTypeInteger, Db: "work_item_id", Nullable: false, Public: true},
	},
	ExtraSchemaTableEntityExtraSchemaWorkItemAssignee: {
		"assignee":   DbField{Type: ColumnSimpleTypeString, Db: "assignee", Nullable: false, Public: true},
		"workItemID": DbField{Type: ColumnSimpleTypeInteger, Db: "work_item_id", Nullable: false, Public: true},
	},
}
