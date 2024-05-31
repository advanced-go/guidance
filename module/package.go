package module

const (
	Authority = "github/advanced-go/guidance"
	Name      = "guidance"
	Version   = "1.1.1"
	Ver1      = "v1"
	Ver2      = "v2"

	ResiliencyResource = "resiliency"
)

const (
	DocumentsAuthority = "github/advanced-go/documents"
	DocumentsResource  = "resiliency"
	DocumentsPath      = "/github/advanced-go/documents:%sresiliency"

	DocumentsV1 = "v1"
	DocumentsV2 = "v2"
)

// Configuration keys used on startup for map values
const (
	PackageNameUserKey     = "user"    // type:text, empty:false
	PackageNamePasswordKey = "pswd"    // type:text, empty:false
	PackageNameRetriesKey  = "retries" // type:int, default:-1, range:0-100
)
