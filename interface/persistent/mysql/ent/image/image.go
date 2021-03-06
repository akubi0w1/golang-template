// Code generated by entc, DO NOT EDIT.

package image

const (
	// Label holds the string label denoting the image type in the database.
	Label = "image"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldURL holds the string denoting the url field in the database.
	FieldURL = "url"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// EdgePosts holds the string denoting the posts edge name in mutations.
	EdgePosts = "posts"
	// EdgeUploadedBy holds the string denoting the uploadedby edge name in mutations.
	EdgeUploadedBy = "uploadedBy"
	// Table holds the table name of the image in the database.
	Table = "images"
	// PostsTable is the table the holds the posts relation/edge.
	PostsTable = "images"
	// PostsInverseTable is the table name for the Post entity.
	// It exists in this package in order to avoid circular dependency with the "post" package.
	PostsInverseTable = "posts"
	// PostsColumn is the table column denoting the posts relation/edge.
	PostsColumn = "post_images"
	// UploadedByTable is the table the holds the uploadedBy relation/edge.
	UploadedByTable = "images"
	// UploadedByInverseTable is the table name for the User entity.
	// It exists in this package in order to avoid circular dependency with the "user" package.
	UploadedByInverseTable = "users"
	// UploadedByColumn is the table column denoting the uploadedBy relation/edge.
	UploadedByColumn = "user_upload"
)

// Columns holds all SQL columns for image fields.
var Columns = []string{
	FieldID,
	FieldURL,
	FieldCreatedAt,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the "images"
// table and are not defined as standalone fields in the schema.
var ForeignKeys = []string{
	"post_images",
	"user_upload",
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	for i := range ForeignKeys {
		if column == ForeignKeys[i] {
			return true
		}
	}
	return false
}
