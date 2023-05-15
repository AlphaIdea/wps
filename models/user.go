package models

type User struct {
	// 用户 ID，字符串长度小于 32
	Id string `json:"id,omitempty"`
	// 用户名
	Name string `json:"name,omitempty"`
	// 用户头像
	AvatarUrl string `json:"avatar_url,omitempty"`
	// 用户操作权限，write：可编辑，read：预览 (必填)
	Permission string `json:"permission,omitempty"`
}

type UserACL struct {
	// 导出PDF、下载权限，默认为 1
	Export int32 `json:"export,omitempty"`
	// 打印文档权限，默认为 1
	Print int32 `json:"print,omitempty"`
	// 只读可评论权限，默认为 0
	Comment int32 `json:"comment,omitempty"`
	// 复制权限，默认为 1
	Copy     int32 `json:"copy,omitempty"`
	Read     int32 `json:"read,omitempty"`
	Update   int32 `json:"update,omitempty"`
	Download int32 `json:"download,omitempty"`
	Share    int32 `json:"share,omitempty"`
	// 重命名权限，默认为 0
	Rename int32 `json:"rename,omitempty"`
	// 历史版本权限，默认为 1
	History int32 `json:"history,omitempty"`
}

type GetUserInfosInput struct {
	Ids []string `json:"ids,omitempty"`
}

type GetUserInfos struct {
	Users []*User `json:"users,omitempty"`
}
