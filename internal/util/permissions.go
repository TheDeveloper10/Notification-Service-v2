package util

type PermissionsNumeric uint8
type PermissionString string
type Permissions []PermissionString

const (
	ManageClients PermissionsNumeric = 1 << iota
	WriteTemplates
	ReadTemplates
	SendNotifications
	ReadNotifications
)

var permissionsMap = map[PermissionString]PermissionsNumeric{
	"manage_clients":     ManageClients,
	"write_templates":    WriteTemplates,
	"read_templates":     ReadTemplates,
	"send_notifications": SendNotifications,
	"read_notifications": ReadNotifications,
}

func (pn *PermissionsNumeric) ToArray() Permissions {
	var arr Permissions

	for k, v := range permissionsMap {
		if v&*pn > 0 {
			arr = append(arr, k)
		}
	}

	return arr
}

func (pn *PermissionsNumeric) HasPermission(perm PermissionsNumeric) bool {
	return *pn&perm > 0
}

func (perms *Permissions) ToNumeric() PermissionsNumeric {
	var n PermissionsNumeric = 0

	for _, jp := range *perms {
		n += jp.ToNumeric()
	}

	return n
}

func (ps *PermissionString) ToNumeric() PermissionsNumeric {
	res, ok := permissionsMap[*ps]
	if ok {
		return res
	} else {
		return 0
	}
}
