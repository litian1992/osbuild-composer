// Package v2 provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/deepmap/oapi-codegen version v1.8.2 DO NOT EDIT.
package v2

import (
	"bytes"
	"compress/gzip"
	"encoding/base64"
	"fmt"
	"net/http"
	"net/url"
	"path"
	"strings"

	"github.com/deepmap/oapi-codegen/pkg/runtime"
	"github.com/getkin/kin-openapi/openapi3"
	"github.com/labstack/echo/v4"
)

const (
	BearerScopes = "Bearer.Scopes"
)

// Defines values for ComposeStatusValue.
const (
	ComposeStatusValueFailure ComposeStatusValue = "failure"

	ComposeStatusValuePending ComposeStatusValue = "pending"

	ComposeStatusValueSuccess ComposeStatusValue = "success"
)

// Defines values for ImageStatusValue.
const (
	ImageStatusValueBuilding ImageStatusValue = "building"

	ImageStatusValueFailure ImageStatusValue = "failure"

	ImageStatusValuePending ImageStatusValue = "pending"

	ImageStatusValueRegistering ImageStatusValue = "registering"

	ImageStatusValueSuccess ImageStatusValue = "success"

	ImageStatusValueUploading ImageStatusValue = "uploading"
)

// Defines values for ImageTypes.
const (
	ImageTypesAws ImageTypes = "aws"

	ImageTypesAwsHaRhui ImageTypes = "aws-ha-rhui"

	ImageTypesAwsRhui ImageTypes = "aws-rhui"

	ImageTypesAwsSapRhui ImageTypes = "aws-sap-rhui"

	ImageTypesAzure ImageTypes = "azure"

	ImageTypesAzureRhui ImageTypes = "azure-rhui"

	ImageTypesEdgeCommit ImageTypes = "edge-commit"

	ImageTypesEdgeContainer ImageTypes = "edge-container"

	ImageTypesEdgeInstaller ImageTypes = "edge-installer"

	ImageTypesGcp ImageTypes = "gcp"

	ImageTypesGuestImage ImageTypes = "guest-image"

	ImageTypesImageInstaller ImageTypes = "image-installer"

	ImageTypesVsphere ImageTypes = "vsphere"
)

// Defines values for UploadStatusValue.
const (
	UploadStatusValueFailure UploadStatusValue = "failure"

	UploadStatusValuePending UploadStatusValue = "pending"

	UploadStatusValueRunning UploadStatusValue = "running"

	UploadStatusValueSuccess UploadStatusValue = "success"
)

// Defines values for UploadTypes.
const (
	UploadTypesAws UploadTypes = "aws"

	UploadTypesAwsS3 UploadTypes = "aws.s3"

	UploadTypesAzure UploadTypes = "azure"

	UploadTypesGcp UploadTypes = "gcp"
)

// AWSEC2UploadOptions defines model for AWSEC2UploadOptions.
type AWSEC2UploadOptions struct {
	Region            string   `json:"region"`
	ShareWithAccounts []string `json:"share_with_accounts"`
	SnapshotName      *string  `json:"snapshot_name,omitempty"`
}

// AWSEC2UploadStatus defines model for AWSEC2UploadStatus.
type AWSEC2UploadStatus struct {
	Ami    string `json:"ami"`
	Region string `json:"region"`
}

// AWSS3UploadOptions defines model for AWSS3UploadOptions.
type AWSS3UploadOptions struct {
	Region string `json:"region"`
}

// AWSS3UploadStatus defines model for AWSS3UploadStatus.
type AWSS3UploadStatus struct {
	Url string `json:"url"`
}

// AzureUploadOptions defines model for AzureUploadOptions.
type AzureUploadOptions struct {
	// Name of the uploaded image. It must be unique in the given resource group.
	// If name is omitted from the request, a random one based on a UUID is
	// generated.
	ImageName *string `json:"image_name,omitempty"`

	// Location where the image should be uploaded and registered.
	// How to list all locations:
	// https://docs.microsoft.com/en-us/cli/azure/account?view=azure-cli-latest#az_account_list_locations'
	Location string `json:"location"`

	// Name of the resource group where the image should be uploaded.
	ResourceGroup string `json:"resource_group"`

	// ID of subscription where the image should be uploaded.
	SubscriptionId string `json:"subscription_id"`

	// ID of the tenant where the image should be uploaded.
	// How to find it in the Azure Portal:
	// https://docs.microsoft.com/en-us/azure/active-directory/fundamentals/active-directory-how-to-find-tenant
	TenantId string `json:"tenant_id"`
}

// AzureUploadStatus defines model for AzureUploadStatus.
type AzureUploadStatus struct {
	ImageName string `json:"image_name"`
}

// ComposeId defines model for ComposeId.
type ComposeId struct {
	// Embedded struct due to allOf(#/components/schemas/ObjectReference)
	ObjectReference `yaml:",inline"`
	// Embedded fields due to inline allOf schema
	Id string `json:"id"`
}

// ComposeLogs defines model for ComposeLogs.
type ComposeLogs struct {
	// Embedded struct due to allOf(#/components/schemas/ObjectReference)
	ObjectReference `yaml:",inline"`
	// Embedded fields due to inline allOf schema
	ImageBuilds []interface{} `json:"image_builds"`
	Koji        *KojiLogs     `json:"koji,omitempty"`
}

// ComposeManifests defines model for ComposeManifests.
type ComposeManifests struct {
	// Embedded struct due to allOf(#/components/schemas/ObjectReference)
	ObjectReference `yaml:",inline"`
	// Embedded fields due to inline allOf schema
	Manifests []interface{} `json:"manifests"`
}

// ComposeMetadata defines model for ComposeMetadata.
type ComposeMetadata struct {
	// Embedded struct due to allOf(#/components/schemas/ObjectReference)
	ObjectReference `yaml:",inline"`
	// Embedded fields due to inline allOf schema
	// ID (hash) of the built commit
	OstreeCommit *string `json:"ostree_commit,omitempty"`

	// Package list including NEVRA
	Packages *[]PackageMetadata `json:"packages,omitempty"`
}

// ComposeRequest defines model for ComposeRequest.
type ComposeRequest struct {
	Customizations *Customizations `json:"customizations,omitempty"`
	Distribution   string          `json:"distribution"`
	ImageRequest   *ImageRequest   `json:"image_request,omitempty"`
	ImageRequests  *[]ImageRequest `json:"image_requests,omitempty"`
	Koji           *Koji           `json:"koji,omitempty"`
}

// ComposeStatus defines model for ComposeStatus.
type ComposeStatus struct {
	// Embedded struct due to allOf(#/components/schemas/ObjectReference)
	ObjectReference `yaml:",inline"`
	// Embedded fields due to inline allOf schema
	ImageStatus   ImageStatus        `json:"image_status"`
	ImageStatuses *[]ImageStatus     `json:"image_statuses,omitempty"`
	KojiStatus    *KojiStatus        `json:"koji_status,omitempty"`
	Status        ComposeStatusValue `json:"status"`
}

// ComposeStatusError defines model for ComposeStatusError.
type ComposeStatusError struct {
	Details *interface{} `json:"details,omitempty"`
	Id      int          `json:"id"`
	Reason  string       `json:"reason"`
}

// ComposeStatusValue defines model for ComposeStatusValue.
type ComposeStatusValue string

// Customizations defines model for Customizations.
type Customizations struct {
	Filesystem *[]Filesystem `json:"filesystem,omitempty"`
	Packages   *[]string     `json:"packages,omitempty"`

	// Extra repositories for packages specified in customizations. These
	// repositories will only be used to depsolve and retrieve packages
	// for the OS itself (they will not be available for the build root or
	// any other part of the build process). The package_sets field for these
	// repositories is ignored.
	PayloadRepositories *[]Repository `json:"payload_repositories,omitempty"`
	Subscription        *Subscription `json:"subscription,omitempty"`
	Users               *[]User       `json:"users,omitempty"`
}

// Error defines model for Error.
type Error struct {
	// Embedded struct due to allOf(#/components/schemas/ObjectReference)
	ObjectReference `yaml:",inline"`
	// Embedded fields due to inline allOf schema
	Code        string       `json:"code"`
	Details     *interface{} `json:"details,omitempty"`
	OperationId string       `json:"operation_id"`
	Reason      string       `json:"reason"`
}

// ErrorList defines model for ErrorList.
type ErrorList struct {
	// Embedded struct due to allOf(#/components/schemas/List)
	List `yaml:",inline"`
	// Embedded fields due to inline allOf schema
	Items []Error `json:"items"`
}

// Filesystem defines model for Filesystem.
type Filesystem struct {
	MinSize    uint64 `json:"min_size"`
	Mountpoint string `json:"mountpoint"`
}

// GCPUploadOptions defines model for GCPUploadOptions.
type GCPUploadOptions struct {
	// Name of an existing STANDARD Storage class Bucket.
	Bucket string `json:"bucket"`

	// The name to use for the imported and shared Compute Engine image.
	// The image name must be unique within the GCP project, which is used
	// for the OS image upload and import. If not specified a random
	// 'composer-api-<uuid>' string is used as the image name.
	ImageName *string `json:"image_name,omitempty"`

	// The GCP region where the OS image will be imported to and shared from.
	// The value must be a valid GCP location. See https://cloud.google.com/storage/docs/locations.
	// If not specified, the multi-region location closest to the source
	// (source Storage Bucket location) is chosen automatically.
	Region string `json:"region"`

	// List of valid Google accounts to share the imported Compute Engine image with.
	// Each string must contain a specifier of the account type. Valid formats are:
	//   - 'user:{emailid}': An email address that represents a specific
	//     Google account. For example, 'alice@example.com'.
	//   - 'serviceAccount:{emailid}': An email address that represents a
	//     service account. For example, 'my-other-app@appspot.gserviceaccount.com'.
	//   - 'group:{emailid}': An email address that represents a Google group.
	//     For example, 'admins@example.com'.
	//   - 'domain:{domain}': The G Suite domain (primary) that represents all
	//     the users of that domain. For example, 'google.com' or 'example.com'.
	// If not specified, the imported Compute Engine image is not shared with any
	// account.
	ShareWithAccounts *[]string `json:"share_with_accounts,omitempty"`
}

// GCPUploadStatus defines model for GCPUploadStatus.
type GCPUploadStatus struct {
	ImageName string `json:"image_name"`
	ProjectId string `json:"project_id"`
}

// ImageRequest defines model for ImageRequest.
type ImageRequest struct {
	Architecture string       `json:"architecture"`
	ImageType    ImageTypes   `json:"image_type"`
	Ostree       *OSTree      `json:"ostree,omitempty"`
	Repositories []Repository `json:"repositories"`

	// This should really be oneOf but AWSS3UploadOptions is a subset of
	// AWSEC2UploadOptions. This means that all AWSEC2UploadOptions objects
	// are also valid AWSS3UploadOptionas objects which violates the oneOf
	// rules. Therefore, we have to use anyOf here but be aware that it isn't
	// possible to mix and match more schemas together.
	UploadOptions *UploadOptions `json:"upload_options,omitempty"`
}

// ImageStatus defines model for ImageStatus.
type ImageStatus struct {
	Error        *ComposeStatusError `json:"error,omitempty"`
	Status       ImageStatusValue    `json:"status"`
	UploadStatus *UploadStatus       `json:"upload_status,omitempty"`
}

// ImageStatusValue defines model for ImageStatusValue.
type ImageStatusValue string

// ImageTypes defines model for ImageTypes.
type ImageTypes string

// Koji defines model for Koji.
type Koji struct {
	Name    string `json:"name"`
	Release string `json:"release"`
	Server  string `json:"server"`
	TaskId  int    `json:"task_id"`
	Version string `json:"version"`
}

// KojiLogs defines model for KojiLogs.
type KojiLogs struct {
	Import interface{} `json:"import"`
	Init   interface{} `json:"init"`
}

// KojiStatus defines model for KojiStatus.
type KojiStatus struct {
	BuildId *int `json:"build_id,omitempty"`
}

// List defines model for List.
type List struct {
	Kind  string `json:"kind"`
	Page  int    `json:"page"`
	Size  int    `json:"size"`
	Total int    `json:"total"`
}

// OSTree defines model for OSTree.
type OSTree struct {
	// Can be either a commit (example: 02604b2da6e954bd34b8b82a835e5a77d2b60ffa), or a branch-like reference (example: rhel/8/x86_64/edge)
	Parent *string `json:"parent,omitempty"`
	Ref    *string `json:"ref,omitempty"`
	Url    *string `json:"url,omitempty"`
}

// ObjectReference defines model for ObjectReference.
type ObjectReference struct {
	Href string `json:"href"`
	Id   string `json:"id"`
	Kind string `json:"kind"`
}

// PackageMetadata defines model for PackageMetadata.
type PackageMetadata struct {
	Arch      string  `json:"arch"`
	Epoch     *string `json:"epoch,omitempty"`
	Name      string  `json:"name"`
	Release   string  `json:"release"`
	Sigmd5    string  `json:"sigmd5"`
	Signature *string `json:"signature,omitempty"`
	Type      string  `json:"type"`
	Version   string  `json:"version"`
}

// Repository defines model for Repository.
type Repository struct {
	Baseurl  *string `json:"baseurl,omitempty"`
	CheckGpg *bool   `json:"check_gpg,omitempty"`

	// GPG key used to sign packages in this repository.
	Gpgkey     *string `json:"gpgkey,omitempty"`
	IgnoreSsl  *bool   `json:"ignore_ssl,omitempty"`
	Metalink   *string `json:"metalink,omitempty"`
	Mirrorlist *string `json:"mirrorlist,omitempty"`

	// Naming package sets for a repository assigns it to a specific part
	// (pipeline) of the build process.
	PackageSets *[]string `json:"package_sets,omitempty"`

	// Determines whether a valid subscription is required to access this repository.
	Rhsm *bool `json:"rhsm,omitempty"`
}

// Subscription defines model for Subscription.
type Subscription struct {
	ActivationKey string `json:"activation_key"`
	BaseUrl       string `json:"base_url"`
	Insights      bool   `json:"insights"`
	Organization  string `json:"organization"`
	ServerUrl     string `json:"server_url"`
}

// This should really be oneOf but AWSS3UploadOptions is a subset of
// AWSEC2UploadOptions. This means that all AWSEC2UploadOptions objects
// are also valid AWSS3UploadOptionas objects which violates the oneOf
// rules. Therefore, we have to use anyOf here but be aware that it isn't
// possible to mix and match more schemas together.
type UploadOptions interface{}

// UploadStatus defines model for UploadStatus.
type UploadStatus struct {
	Options interface{}       `json:"options"`
	Status  UploadStatusValue `json:"status"`
	Type    UploadTypes       `json:"type"`
}

// UploadStatusValue defines model for UploadStatusValue.
type UploadStatusValue string

// UploadTypes defines model for UploadTypes.
type UploadTypes string

// User defines model for User.
type User struct {
	Groups *[]string `json:"groups,omitempty"`
	Key    *string   `json:"key,omitempty"`
	Name   string    `json:"name"`
}

// Page defines model for page.
type Page string

// Size defines model for size.
type Size string

// PostComposeJSONBody defines parameters for PostCompose.
type PostComposeJSONBody ComposeRequest

// GetErrorListParams defines parameters for GetErrorList.
type GetErrorListParams struct {
	// Page index
	Page *Page `json:"page,omitempty"`

	// Number of items in each page
	Size *Size `json:"size,omitempty"`
}

// PostComposeJSONRequestBody defines body for PostCompose for application/json ContentType.
type PostComposeJSONRequestBody PostComposeJSONBody

// ServerInterface represents all server handlers.
type ServerInterface interface {
	// Create compose
	// (POST /compose)
	PostCompose(ctx echo.Context) error
	// The status of a compose
	// (GET /composes/{id})
	GetComposeStatus(ctx echo.Context, id string) error
	// Get logs for a compose.
	// (GET /composes/{id}/logs)
	GetComposeLogs(ctx echo.Context, id string) error
	// Get the manifests for a compose.
	// (GET /composes/{id}/manifests)
	GetComposeManifests(ctx echo.Context, id string) error
	// Get the metadata for a compose.
	// (GET /composes/{id}/metadata)
	GetComposeMetadata(ctx echo.Context, id string) error
	// Get a list of all possible errors
	// (GET /errors)
	GetErrorList(ctx echo.Context, params GetErrorListParams) error
	// Get error description
	// (GET /errors/{id})
	GetError(ctx echo.Context, id string) error
	// Get the openapi spec in json format
	// (GET /openapi)
	GetOpenapi(ctx echo.Context) error
}

// ServerInterfaceWrapper converts echo contexts to parameters.
type ServerInterfaceWrapper struct {
	Handler ServerInterface
}

// PostCompose converts echo context to params.
func (w *ServerInterfaceWrapper) PostCompose(ctx echo.Context) error {
	var err error

	ctx.Set(BearerScopes, []string{""})

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.PostCompose(ctx)
	return err
}

// GetComposeStatus converts echo context to params.
func (w *ServerInterfaceWrapper) GetComposeStatus(ctx echo.Context) error {
	var err error
	// ------------- Path parameter "id" -------------
	var id string

	err = runtime.BindStyledParameterWithLocation("simple", false, "id", runtime.ParamLocationPath, ctx.Param("id"), &id)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter id: %s", err))
	}

	ctx.Set(BearerScopes, []string{""})

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.GetComposeStatus(ctx, id)
	return err
}

// GetComposeLogs converts echo context to params.
func (w *ServerInterfaceWrapper) GetComposeLogs(ctx echo.Context) error {
	var err error
	// ------------- Path parameter "id" -------------
	var id string

	err = runtime.BindStyledParameterWithLocation("simple", false, "id", runtime.ParamLocationPath, ctx.Param("id"), &id)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter id: %s", err))
	}

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.GetComposeLogs(ctx, id)
	return err
}

// GetComposeManifests converts echo context to params.
func (w *ServerInterfaceWrapper) GetComposeManifests(ctx echo.Context) error {
	var err error
	// ------------- Path parameter "id" -------------
	var id string

	err = runtime.BindStyledParameterWithLocation("simple", false, "id", runtime.ParamLocationPath, ctx.Param("id"), &id)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter id: %s", err))
	}

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.GetComposeManifests(ctx, id)
	return err
}

// GetComposeMetadata converts echo context to params.
func (w *ServerInterfaceWrapper) GetComposeMetadata(ctx echo.Context) error {
	var err error
	// ------------- Path parameter "id" -------------
	var id string

	err = runtime.BindStyledParameterWithLocation("simple", false, "id", runtime.ParamLocationPath, ctx.Param("id"), &id)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter id: %s", err))
	}

	ctx.Set(BearerScopes, []string{""})

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.GetComposeMetadata(ctx, id)
	return err
}

// GetErrorList converts echo context to params.
func (w *ServerInterfaceWrapper) GetErrorList(ctx echo.Context) error {
	var err error

	ctx.Set(BearerScopes, []string{""})

	// Parameter object where we will unmarshal all parameters from the context
	var params GetErrorListParams
	// ------------- Optional query parameter "page" -------------

	err = runtime.BindQueryParameter("form", true, false, "page", ctx.QueryParams(), &params.Page)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter page: %s", err))
	}

	// ------------- Optional query parameter "size" -------------

	err = runtime.BindQueryParameter("form", true, false, "size", ctx.QueryParams(), &params.Size)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter size: %s", err))
	}

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.GetErrorList(ctx, params)
	return err
}

// GetError converts echo context to params.
func (w *ServerInterfaceWrapper) GetError(ctx echo.Context) error {
	var err error
	// ------------- Path parameter "id" -------------
	var id string

	err = runtime.BindStyledParameterWithLocation("simple", false, "id", runtime.ParamLocationPath, ctx.Param("id"), &id)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter id: %s", err))
	}

	ctx.Set(BearerScopes, []string{""})

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.GetError(ctx, id)
	return err
}

// GetOpenapi converts echo context to params.
func (w *ServerInterfaceWrapper) GetOpenapi(ctx echo.Context) error {
	var err error

	ctx.Set(BearerScopes, []string{""})

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.GetOpenapi(ctx)
	return err
}

// This is a simple interface which specifies echo.Route addition functions which
// are present on both echo.Echo and echo.Group, since we want to allow using
// either of them for path registration
type EchoRouter interface {
	CONNECT(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	DELETE(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	GET(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	HEAD(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	OPTIONS(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	PATCH(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	POST(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	PUT(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	TRACE(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
}

// RegisterHandlers adds each server route to the EchoRouter.
func RegisterHandlers(router EchoRouter, si ServerInterface) {
	RegisterHandlersWithBaseURL(router, si, "")
}

// Registers handlers, and prepends BaseURL to the paths, so that the paths
// can be served under a prefix.
func RegisterHandlersWithBaseURL(router EchoRouter, si ServerInterface, baseURL string) {

	wrapper := ServerInterfaceWrapper{
		Handler: si,
	}

	router.POST(baseURL+"/compose", wrapper.PostCompose)
	router.GET(baseURL+"/composes/:id", wrapper.GetComposeStatus)
	router.GET(baseURL+"/composes/:id/logs", wrapper.GetComposeLogs)
	router.GET(baseURL+"/composes/:id/manifests", wrapper.GetComposeManifests)
	router.GET(baseURL+"/composes/:id/metadata", wrapper.GetComposeMetadata)
	router.GET(baseURL+"/errors", wrapper.GetErrorList)
	router.GET(baseURL+"/errors/:id", wrapper.GetError)
	router.GET(baseURL+"/openapi", wrapper.GetOpenapi)

}

// Base64 encoded, gzipped, json marshaled Swagger object
var swaggerSpec = []string{

	"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",
	"AP//e2tK5q1ZAAA=",
}

// GetSwagger returns the content of the embedded swagger specification file
// or error if failed to decode
func decodeSpec() ([]byte, error) {
	zipped, err := base64.StdEncoding.DecodeString(strings.Join(swaggerSpec, ""))
	if err != nil {
		return nil, fmt.Errorf("error base64 decoding spec: %s", err)
	}
	zr, err := gzip.NewReader(bytes.NewReader(zipped))
	if err != nil {
		return nil, fmt.Errorf("error decompressing spec: %s", err)
	}
	var buf bytes.Buffer
	_, err = buf.ReadFrom(zr)
	if err != nil {
		return nil, fmt.Errorf("error decompressing spec: %s", err)
	}

	return buf.Bytes(), nil
}

var rawSpec = decodeSpecCached()

// a naive cached of a decoded swagger spec
func decodeSpecCached() func() ([]byte, error) {
	data, err := decodeSpec()
	return func() ([]byte, error) {
		return data, err
	}
}

// Constructs a synthetic filesystem for resolving external references when loading openapi specifications.
func PathToRawSpec(pathToFile string) map[string]func() ([]byte, error) {
	var res = make(map[string]func() ([]byte, error))
	if len(pathToFile) > 0 {
		res[pathToFile] = rawSpec
	}

	return res
}

// GetSwagger returns the Swagger specification corresponding to the generated code
// in this file. The external references of Swagger specification are resolved.
// The logic of resolving external references is tightly connected to "import-mapping" feature.
// Externally referenced files must be embedded in the corresponding golang packages.
// Urls can be supported but this task was out of the scope.
func GetSwagger() (swagger *openapi3.T, err error) {
	var resolvePath = PathToRawSpec("")

	loader := openapi3.NewLoader()
	loader.IsExternalRefsAllowed = true
	loader.ReadFromURIFunc = func(loader *openapi3.Loader, url *url.URL) ([]byte, error) {
		var pathToFile = url.String()
		pathToFile = path.Clean(pathToFile)
		getSpec, ok := resolvePath[pathToFile]
		if !ok {
			err1 := fmt.Errorf("path not found: %s", pathToFile)
			return nil, err1
		}
		return getSpec()
	}
	var specData []byte
	specData, err = rawSpec()
	if err != nil {
		return
	}
	swagger, err = loader.LoadFromData(specData)
	if err != nil {
		return
	}
	return
}
