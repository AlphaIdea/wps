package wps

const (
	// UAField 用户代理
	UAField = "weboffice-httpclient"

	// XFileIdField 文件 id
	XFileIdField = "x-weboffice-file-id"

	// XUserTokenField 校验身份的 token，值根据对接的企业而定
	XUserTokenField = "x-wps-weboffice-token"
)

const (
	AppIdParam       = "_w_appid"
	FileTypeParam    = "_w_filetype"
	SignatureParam   = "_w_signature"
	TokenTypeParam   = "_w_tokentype"
	AccessTokenParam = "_w_access_token"
)

// 用户操作权限
const (
	Read  = "read"
	Write = "write"
)

// 文件类型
const (
	Excel = "s"
	Word  = "w"
	PPT   = "p"
	PDF   = "f"
)

var (
	Excels = []string{"xls", "xlt", "et", "xlsx", "xltx", "csv", "xlsm", "xltm", "ett"}
	Words  = []string{"doc", "dot", "wps", "wpt", "docx", "dotx", "docm", "dotm", "rtf", "txt", "xml", "mhtml", "mht", "html", "htm", "uof"}
	PPTs   = []string{"ppt", "pptx", "pptm", "ppsx", "ppsm", "pps", "potx", "potm", "dpt", "dps", "pot"}
	PDFs   = []string{"pdf"}
)

// 通用权限
const (
	Off = iota
	On
)
