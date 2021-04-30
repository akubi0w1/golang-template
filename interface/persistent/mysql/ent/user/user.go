// Code generated by entc, DO NOT EDIT.

package user

const (
	// Label holds the string label denoting the user type in the database.
	Label = "user"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldAccountID holds the string denoting the account_id field in the database.
	FieldAccountID = "account_id"
	// FieldEmail holds the string denoting the email field in the database.
	FieldEmail = "email"
	// FieldPassword holds the string denoting the password field in the database.
	FieldPassword = "password"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// FieldUpdatedAt holds the string denoting the updated_at field in the database.
	FieldUpdatedAt = "updated_at"
	// FieldDeletedAt holds the string denoting the deleted_at field in the database.
	FieldDeletedAt = "deleted_at"
	// EdgeProfile holds the string denoting the profile edge name in mutations.
	EdgeProfile = "profile"
	// EdgePosts holds the string denoting the posts edge name in mutations.
	EdgePosts = "posts"
	// EdgeUpload holds the string denoting the upload edge name in mutations.
	EdgeUpload = "upload"
	// Table holds the table name of the user in the database.
	Table = "users"
	// ProfileTable is the table the holds the profile relation/edge.
	ProfileTable = "profiles"
	// ProfileInverseTable is the table name for the Profile entity.
	// It exists in this package in order to avoid circular dependency with the "profile" package.
	ProfileInverseTable = "profiles"
	// ProfileColumn is the table column denoting the profile relation/edge.
	ProfileColumn = "user_profile"
	// PostsTable is the table the holds the posts relation/edge.
	PostsTable = "posts"
	// PostsInverseTable is the table name for the Post entity.
	// It exists in this package in order to avoid circular dependency with the "post" package.
	PostsInverseTable = "posts"
	// PostsColumn is the table column denoting the posts relation/edge.
	PostsColumn = "user_posts"
	// UploadTable is the table the holds the upload relation/edge.
	UploadTable = "images"
	// UploadInverseTable is the table name for the Image entity.
	// It exists in this package in order to avoid circular dependency with the "image" package.
	UploadInverseTable = "images"
	// UploadColumn is the table column denoting the upload relation/edge.
	UploadColumn = "user_upload"
)

// Columns holds all SQL columns for user fields.
var Columns = []string{
	FieldID,
	FieldAccountID,
	FieldEmail,
	FieldPassword,
	FieldCreatedAt,
	FieldUpdatedAt,
	FieldDeletedAt,
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	return false
}
