package request

type Role struct {
	ParentID        int    `json:"parent_role" binding:"required"`
	AuthRoleID      int    `json:"role_id" binding:"required"`
	AuthRoleName    string `json:"role_name" binding:"required"`
	AuthRoleRemarks string `json:"role_remarks"`
}
