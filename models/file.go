package models

// Watermark 水印
type Watermark struct {
	// 水印类型， 0 为无水印； 1 为文字水印
	Type int32 `json:"type,omitempty"`
	// 文字水印的文字，当 type 为 1 时，此字段必填
	Value string `json:"value,omitempty"`
	// 水印的透明度，非必填，有默认值
	FillStyle string `json:"fillstyle,omitempty"`
	// 水印的字体，非必填，有默认值
	Font string `json:"font,omitempty"`
	// 水印的旋转度，非必填，有默认值
	Rotate float32 `json:"rotate,omitempty"`
	// 水印水平间距，非必填，有默认值
	Horizontal int32 `json:"horizontal,omitempty"`
	// 水印垂直间距，非必填，有默认值
	Vertical int32 `json:"vertical,omitempty"`
}

// FileMetadata 文件元数据
type FileMetadata struct {
	// 文件id，字符串长度小于 40，和 URL 中的 fileid 必须一致
	Id string `binding:"omitempty,max=40" json:"id,omitempty"`
	// 文件名(必须带文件后缀) (必填)
	Name string `json:"name,omitempty"`
	// 当前版本号，必须大于 0，同时位数小于 11 (必填)
	Version int32 `json:"version,omitempty"`
	// 文件大小，单位为B(文件真实大小，否则会出现异常) (必填)
	Size int64  `json:"size,omitempty"`
	Type string `json:"type,omitempty"`
	// 文档下载地址 (必填)
	DownloadUrl string `json:"download_url,omitempty"`
	// 创建时间，时间戳，单位为秒 (必填)
	CreateTime int64 `json:"create_time,omitempty"`
	// 修改时间，时间戳，单位为秒 (必填)
	ModifyTime int64      `json:"modify_time,omitempty"`
	Creator    *User      `json:"creator,omitempty"`
	Modifier   *User      `json:"modifier,omitempty"`
	UniqueId   string     `json:"unique_id,omitempty"`
	UserAcl    *UserACL   `json:"user_acl,omitempty"`
	VerType    string     `json:"ver_type,omitempty"`
	Watermark  *Watermark `json:"watermark,omitempty"`
}

type File struct {
	// 文件 id，字符串长度小于 40
	Id string `json:"id,omitempty"`
	// 文件名
	Name string `json:"name,omitempty"`
	// 当前版本号，位数小于 11
	Version int32 `json:"version,omitempty"`
	// 文件大小，单位是 B
	Size int64  `json:"size,omitempty"`
	Type string `json:"type,omitempty"`
	// 文件下载地址
	DownloadUrl string `json:"download_url,omitempty"`
	// 创建时间，时间戳，单位为秒
	CreateTime int64 `json:"create_time,omitempty"`
	// 修改时间，时间戳，单位为秒
	ModifyTime int64 `json:"modify_time,omitempty"`
	// 创建者 id，字符串长度小于 40
	CreatorId string `json:"creator,omitempty"`
	// 修改者 id，字符串长度小于 40
	ModifierId string     `json:"modifier,omitempty"`
	UniqueId   string     `json:"unique_id,omitempty"`
	LinkId     string     `json:"link_id,omitempty"`
	UserAcl    *UserACL   `json:"user_acl,omitempty"`
	Watermark  *Watermark `json:"watermark,omitempty"`
}

type FileVersion struct {
	File *File `json:"file,omitempty"`
}

type PutFileInput struct {
	Name string `json:"name,omitempty"`
}

type FileEdit struct {
	File *File `json:"file,omitempty"`
	User *User `json:"user,omitempty"`
}

type PostFile struct {
	File *File `json:"file,omitempty"`
}

type FileHistoryVersions struct {
	Histories []*FileMetadata `json:"histories,omitempty"`
}

type FileHistoryVersionsInput struct {
	Id     string `json:"id,omitempty"`
	Offset int32  `json:"offset,omitempty"`
	Count  int32  `json:"count,omitempty"`
}

type GetTemplateInfo struct {
	RedirectUrl string `json:"redirect_url"`
	UserId      string `json:"user_id"`
}
