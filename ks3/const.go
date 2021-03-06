package ks3

import "os"

// ACLType bucket/object ACL
type ACLType string

const (
	// ACLPrivate definition : private read and write
	ACLPrivate ACLType = "private"

	// ACLPublicRead definition : public read and private write
	ACLPublicRead ACLType = "public-read"

	// ACLPublicReadWrite definition : public read and public write
	ACLPublicReadWrite ACLType = "public-read-write"

	// ACLDefault Object. It's only applicable for object.
	ACLDefault ACLType = "default"
)

// bucket versioning status
type VersioningStatus string

const (
	// Versioning Status definition: Enabled
	VersionEnabled VersioningStatus = "Enabled"

	// Versioning Status definition: Suspended
	VersionSuspended VersioningStatus = "Suspended"
)

// MetadataDirectiveType specifying whether use the metadata of source object when copying object.
type MetadataDirectiveType string

const (
	// MetaCopy the target object's metadata is copied from the source one
	MetaCopy MetadataDirectiveType = "COPY"

	// MetaReplace the target object's metadata is created as part of the copy request (not same as the source one)
	MetaReplace MetadataDirectiveType = "REPLACE"
)

// TaggingDirectiveType specifying whether use the tagging of source object when copying object.
type TaggingDirectiveType string

const (
	// TaggingCopy the target object's tagging is copied from the source one
	TaggingCopy TaggingDirectiveType = "COPY"

	// TaggingReplace the target object's tagging is created as part of the copy request (not same as the source one)
	TaggingReplace TaggingDirectiveType = "REPLACE"
)

// AlgorithmType specifying the server side encryption algorithm name
type AlgorithmType string

const (
	KMSAlgorithm AlgorithmType = "KMS"
	AESAlgorithm AlgorithmType = "AES256"
	SM4Algorithm AlgorithmType = "SM4"
)

// StorageClassType bucket storage type
type StorageClassType string

const (
	// StorageStandard standard
	StorageStandard StorageClassType = "Standard"

	// StorageIA infrequent access
	StorageIA StorageClassType = "IA"

	// StorageArchive archive
	StorageArchive StorageClassType = "Archive"

	// StorageColdArchive cold archive
	StorageColdArchive StorageClassType = "ColdArchive"
)

//RedundancyType bucket data Redundancy type
type DataRedundancyType string

const (
	// RedundancyLRS Local redundancy, default value
	RedundancyLRS DataRedundancyType = "LRS"

	// RedundancyZRS Same city redundancy
	RedundancyZRS DataRedundancyType = "ZRS"
)

//ObjecthashFuncType
type ObjecthashFuncType string

const (
	HashFuncSha1   ObjecthashFuncType = "SHA-1"
	HashFuncSha256 ObjecthashFuncType = "SHA-256"
)

// PayerType the type of request payer
type PayerType string

const (
	// Requester the requester who send the request
	Requester PayerType = "Requester"

	// BucketOwner the requester who send the request
	BucketOwner PayerType = "BucketOwner"
)

//RestoreMode the restore mode for coldArchive object
type RestoreMode string

const (
	//RestoreExpedited object will be restored in 1 hour
	RestoreExpedited RestoreMode = "Expedited"

	//RestoreStandard object will be restored in 2-5 hours
	RestoreStandard RestoreMode = "Standard"

	//RestoreBulk object will be restored in 5-10 hours
	RestoreBulk RestoreMode = "Bulk"
)

// HTTPMethod HTTP request method
type HTTPMethod string

const (
	// HTTPGet HTTP GET
	HTTPGet HTTPMethod = "GET"

	// HTTPPut HTTP PUT
	HTTPPut HTTPMethod = "PUT"

	// HTTPHead HTTP HEAD
	HTTPHead HTTPMethod = "HEAD"

	// HTTPPost HTTP POST
	HTTPPost HTTPMethod = "POST"

	// HTTPDelete HTTP DELETE
	HTTPDelete HTTPMethod = "DELETE"
)

// HTTP headers
const (
	HTTPHeaderAcceptEncoding     string = "Accept-Encoding"
	HTTPHeaderAuthorization             = "Authorization"
	HTTPHeaderCacheControl              = "Cache-Control"
	HTTPHeaderContentDisposition        = "Content-Disposition"
	HTTPHeaderContentEncoding           = "Content-Encoding"
	HTTPHeaderContentLength             = "Content-Length"
	HTTPHeaderContentMD5                = "Content-MD5"
	HTTPHeaderContentType               = "Content-Type"
	HTTPHeaderContentLanguage           = "Content-Language"
	HTTPHeaderDate                      = "Date"
	HTTPHeaderEtag                      = "ETag"
	HTTPHeaderExpires                   = "Expires"
	HTTPHeaderHost                      = "Host"
	HTTPHeaderLastModified              = "Last-Modified"
	HTTPHeaderRange                     = "Range"
	HTTPHeaderLocation                  = "Location"
	HTTPHeaderOrigin                    = "Origin"
	HTTPHeaderServer                    = "Server"
	HTTPHeaderUserAgent                 = "User-Agent"
	HTTPHeaderIfModifiedSince           = "If-Modified-Since"
	HTTPHeaderIfUnmodifiedSince         = "If-Unmodified-Since"
	HTTPHeaderIfMatch                   = "If-Match"
	HTTPHeaderIfNoneMatch               = "If-None-Match"
	HTTPHeaderACReqMethod               = "Access-Control-Request-Method"
	HTTPHeaderACReqHeaders              = "Access-Control-Request-Headers"

	HTTPHeaderKs3ACL                         = "X-Kss-Acl"
	HTTPHeaderKs3MetaPrefix                  = "X-Kss-Meta-"
	HTTPHeaderKs3ObjectACL                   = "X-Kss-Acl"
	HTTPHeaderKs3SecurityToken               = "X-Kss-Security-Token"
	HTTPHeaderKs3ServerSideEncryption        = "X-Kss-Server-Side-Encryption"
	HTTPHeaderKs3ServerSideEncryptionKeyID   = "X-Kss-Server-Side-Encryption-Key-Id"
	HTTPHeaderKs3ServerSideDataEncryption    = "X-Kss-Server-Side-Data-Encryption"
	HTTPHeaderSSECAlgorithm                  = "X-Kss-Server-Side-Encryption-Customer-Algorithm"
	HTTPHeaderSSECKey                        = "X-Kss-Server-Side-Encryption-Customer-Key"
	HTTPHeaderSSECKeyMd5                     = "X-Kss-Server-Side-Encryption-Customer-Key-MD5"
	HTTPHeaderKs3CopySource                  = "X-Kss-Copy-Source"
	HTTPHeaderKs3CopySourceRange             = "X-Kss-Copy-Source-Range"
	HTTPHeaderKs3CopySourceIfMatch           = "X-Kss-Copy-Source-If-Match"
	HTTPHeaderKs3CopySourceIfNoneMatch       = "X-Kss-Copy-Source-If-None-Match"
	HTTPHeaderKs3CopySourceIfModifiedSince   = "X-Kss-Copy-Source-If-Modified-Since"
	HTTPHeaderKs3CopySourceIfUnmodifiedSince = "X-Kss-Copy-Source-If-Unmodified-Since"
	HTTPHeaderKs3MetadataDirective           = "X-Kss-Metadata-Directive"
	HTTPHeaderKs3NextAppendPosition          = "X-Kss-Next-Append-Position"
	HTTPHeaderKs3RequestID                   = "X-Kss-Request-Id"
	HTTPHeaderKs3CRC64                       = "X-Kss-Hash-Crc64ecma"
	HTTPHeaderKs3SymlinkTarget               = "X-Kss-Symlink-Target"
	HTTPHeaderKs3StorageClass                = "X-Kss-Storage-Class"
	HTTPHeaderKs3Callback                    = "X-Kss-Callback"
	HTTPHeaderKs3CallbackVar                 = "X-Kss-Callback-Var"
	HTTPHeaderKs3Requester                   = "X-Kss-Request-Payer"
	HTTPHeaderKs3Tagging                     = "X-Kss-Tagging"
	HTTPHeaderKs3TaggingDirective            = "X-Kss-Tagging-Directive"
	HTTPHeaderKs3TrafficLimit                = "X-Kss-Traffic-Limit"
	HTTPHeaderKs3ForbidOverWrite             = "X-Kss-Forbid-Overwrite"
	HTTPHeaderKs3RangeBehavior               = "X-Kss-Range-Behavior"
	HTTPHeaderKs3TaskID                      = "X-Kss-Task-Id"
	HTTPHeaderKs3HashCtx                     = "X-Kss-Hash-Ctx"
	HTTPHeaderKs3Md5Ctx                      = "X-Kss-Md5-Ctx"
	HTTPHeaderAllowSameActionOverLap         = "X-Kss-Allow-Same-Action-Overlap"
)

// HTTP Param
const (
	HTTPParamExpires       = "Expires"
	HTTPParamAccessKeyID   = "KS3AccessKeyId"
	HTTPParamSignature     = "Signature"
	HTTPParamSecurityToken = "security-token"
	HTTPParamPlaylistName  = "playlistName"

	HTTPParamSignatureVersion    = "X-Kss-signature-version"
	HTTPParamExpiresV2           = "X-Kss-expires"
	HTTPParamAccessKeyIDV2       = "X-Kss-access-key-id"
	HTTPParamSignatureV2         = "X-Kss-signature"
	HTTPParamAdditionalHeadersV2 = "X-Kss-additional-headers"
)

// Other constants
const (
	MaxPartSize = 5 * 1024 * 1024 * 1024 // Max part size, 5GB
	MinPartSize = 100 * 1024             // Min part size, 100KB

	FilePermMode = os.FileMode(0664) // Default file permission

	TempFilePrefix = "ks3-go-temp-" // Temp file prefix
	TempFileSuffix = ".temp"        // Temp file suffix

	CheckpointFileSuffix = ".cp" // Checkpoint file suffix

	NullVersion = "null"

	Version = "v0.0.1" // Go SDK version
)

// FrameType
const (
	DataFrameType        = 8388609
	ContinuousFrameType  = 8388612
	EndFrameType         = 8388613
	MetaEndFrameCSVType  = 8388614
	MetaEndFrameJSONType = 8388615
)

// AuthVersion the version of auth
type AuthVersionType string

const (
	// AuthV1 v1
	AuthV1 AuthVersionType = "v1"
	// AuthV2 v2
	AuthV2 AuthVersionType = "v2"
)
