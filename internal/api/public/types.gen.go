// Package public provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/deepmap/oapi-codegen version v1.12.4 DO NOT EDIT.
package public

const (
	X_rh_identityScopes = "x_rh_identity.Scopes"
)

// Defines values for CreateDomainDomainType.
const (
	CreateDomainDomainTypeIpa CreateDomainDomainType = "ipa"
)

// Defines values for CreateDomainResponseDomainType.
const (
	CreateDomainResponseDomainTypeIpa CreateDomainResponseDomainType = "ipa"
)

// CheckHosts Define the input data for the /check_host action.
//
// This action is launched from the ipa server to check the host that is being
// auto-enrolled.
type CheckHosts struct {
	// DomainName The domain name where to enroll the host vm.
	DomainName string `json:"domain_name"`

	// DomainType Indicate the type of domain. Actually only ipa is supported.
	DomainType string `json:"domain_type"`

	// InventoryId The id of the host vm into the insight host inventory.
	InventoryId string `json:"inventory_id"`

	// SubscriptionManagerId The subscription manager id for auto-enroll the host vm.
	SubscriptionManagerId string `json:"subscription_manager_id"`
}

// CreateDomain A domain resource
type CreateDomain struct {
	// AutoEnrollmentEnabled Enable or disable host vm auto-enrollment for this domain
	AutoEnrollmentEnabled bool `json:"auto_enrollment_enabled"`

	// DomainName Domain name
	DomainName string `json:"domain_name"`

	// DomainType Type of this domain. Currently only ipa is supported.
	DomainType CreateDomainDomainType `json:"domain_type"`

	// Ipa Options for ipa domains
	Ipa CreateDomainIpa `json:"ipa"`
}

// CreateDomainDomainType Type of this domain. Currently only ipa is supported.
type CreateDomainDomainType string

// CreateDomainIpa Options for ipa domains
type CreateDomainIpa struct {
	// CaList A base64 representation of all the list of chain of certificates, including the server ca.
	CaList string `json:"ca_list"`

	// RealmName The realm name for the domain
	RealmName *string `json:"realm_name,omitempty"`

	// ServerList List of auto-enrollment enabled servers for this domain.
	ServerList *[]string `json:"server_list,omitempty"`
}

// CreateDomainResponse A domain resource
type CreateDomainResponse struct {
	// AutoEnrollmentEnabled Enable or disable host vm auto-enrollment for this domain
	AutoEnrollmentEnabled bool `json:"auto_enrollment_enabled"`

	// DomainName Domain name
	DomainName string `json:"domain_name"`

	// DomainType Type of this domain. Currently only ipa is supported.
	DomainType CreateDomainResponseDomainType `json:"domain_type"`

	// DomainUuid Internal id for this domain
	DomainUuid string `json:"domain_uuid"`

	// Ipa Options for ipa domains
	Ipa CreateDomainIpa `json:"ipa"`
}

// CreateDomainResponseDomainType Type of this domain. Currently only ipa is supported.
type CreateDomainResponseDomainType string

// Error General error schema
type Error struct {
	// Detail A human-readable explanation specific to this occurrence of the problem. This field’s value can be localized.
	Detail string `json:"detail"`

	// Id A unique identifier for this particular occurrence of the problem.
	Id string `json:"id"`

	// Status The HTTP status code applicable to this problem, expressed as a string value. This SHOULD be provided.
	Status *string `json:"status,omitempty"`
}

// ErrorResponse General error response returned by the hmsidm API
type ErrorResponse struct {
	// Errors Error objects provide additional information about problems encountered while performing an operation.
	Errors *[]Error `json:"errors,omitempty"`
}

// HostConf Represent the request payload for the /hostconf/:fqdn endpoint.
type HostConf struct {
	// Fqdn The full qualified domain name of the host vm that is being enrolled.
	Fqdn *string `json:"fqdn,omitempty"`

	// SubscriptionManagerId The id for the subscription manager.
	SubscriptionManagerId *string `json:"subscription_manager_id,omitempty"`
}

// HostConfResponse The response for the action to retrieve the host vm information when
// it is being enrolled. This action is taken from the host vm.
type HostConfResponse struct {
	DomainName *string `json:"domain_name,omitempty"`
	DomainType *string `json:"domain_type,omitempty"`
	Ipa        *struct {
		CaCert     *string   `json:"ca_cert,omitempty"`
		RealmName  *string   `json:"realm_name,omitempty"`
		ServerList *[]string `json:"server_list,omitempty"`
	} `json:"ipa,omitempty"`
}

// ListDomainsData The data listed for the domains.
type ListDomainsData struct {
	AutoEnrollmentEnabled *bool   `json:"auto_enrollment_enabled,omitempty"`
	DomainName            *string `json:"domain_name,omitempty"`
	DomainType            *string `json:"domain_type,omitempty"`
	DomainUuid            *string `json:"domain_uuid,omitempty"`
}

// ListDomainsResponse Represent a paginated result for a list of domains
type ListDomainsResponse struct {
	// Data The content for this page.
	Data []ListDomainsData `json:"data"`

	// Links Represent the navigation links for the data paginated.
	Links PaginationLinks `json:"links"`

	// Meta Metadata for the paginated responses.
	Meta PaginationMeta `json:"meta"`
}

// PaginationLinks Represent the navigation links for the data paginated.
type PaginationLinks struct {
	// First Reference to the first page of the request.
	First *string `json:"first,omitempty"`

	// Last Reference to the last page of the request.
	Last *string `json:"last,omitempty"`

	// Next Reference to the next page of the request.
	Next *string `json:"next,omitempty"`

	// Previous Reference to the previous page of the request.
	Previous *string `json:"previous,omitempty"`
}

// PaginationMeta Metadata for the paginated responses.
type PaginationMeta struct {
	Count *int32 `json:"count,omitempty"`
}

// ReadDomainIpa Options for ipa domains
type ReadDomainIpa struct {
	// CaList A base64 representation of all the list of chain of certificates, including the server ca.
	CaList string `json:"ca_list"`

	// RealmName The realm name for the domain
	RealmName *string `json:"realm_name,omitempty"`

	// ServerList List of auto-enrollment enabled servers for this domain.
	ServerList *[]string `json:"server_list,omitempty"`
}

// ReadDomainResponse Represent the information read for a specific domain.
type ReadDomainResponse struct {
	// AutoEnrollmentEnabled Inidicate if the auto-enrollment is enabled or disabled for this domain.
	AutoEnrollmentEnabled *bool `json:"auto_enrollment_enabled,omitempty"`

	// DomainName The name of the domain.
	DomainName *string `json:"domain_name,omitempty"`

	// DomainType Actually only "ipa" is supported.
	DomainType *string `json:"domain_type,omitempty"`

	// DomainUuid Internal identifier for the domain, it is unique into system.
	DomainUuid *string `json:"domain_uuid,omitempty"`

	// Ipa Options for ipa domains
	Ipa *ReadDomainIpa `json:"ipa,omitempty"`
}

// CheckHostParams defines parameters for CheckHost.
type CheckHostParams struct {
	// XRhIdentity Identity header.
	XRhIdentity []byte `json:"X-Rh-Identity"`

	// XRhInsightsRequestId Request id for distributed tracing.
	XRhInsightsRequestId string `json:"X-Rh-Insights-Request-Id"`
}

// ListDomainsParams defines parameters for ListDomains.
type ListDomainsParams struct {
	// Offset pagination offset
	Offset *int `form:"offset,omitempty" json:"offset,omitempty"`

	// Limit Number of items per page
	Limit *int `form:"limit,omitempty" json:"limit,omitempty"`

	// XRhIdentity Identity header for the request
	XRhIdentity []byte `json:"X-Rh-Identity"`

	// XRhInsightsRequestId Request id for distributed tracing.
	XRhInsightsRequestId string `json:"X-Rh-Insights-Request-Id"`
}

// CreateDomainParams defines parameters for CreateDomain.
type CreateDomainParams struct {
	// XRhIdentity Identity header for the request
	XRhIdentity []byte `json:"X-Rh-Identity"`

	// XRhInsightsRequestId Request id for distributed tracing.
	XRhInsightsRequestId string `json:"X-Rh-Insights-Request-Id"`
}

// DeleteDomainParams defines parameters for DeleteDomain.
type DeleteDomainParams struct {
	// XRhIdentity Identity header for the request
	XRhIdentity []byte `json:"X-Rh-Identity"`

	// XRhInsightsRequestId Request id for distributed tracing.
	XRhInsightsRequestId string `json:"X-Rh-Insights-Request-Id"`
}

// ReadDomainParams defines parameters for ReadDomain.
type ReadDomainParams struct {
	// XRhIdentity Identity header for the request
	XRhIdentity []byte `json:"X-Rh-Identity"`

	// XRhInsightsRequestId Request id for distributed tracing.
	XRhInsightsRequestId string `json:"X-Rh-Insights-Request-Id"`
}

// HostConfParams defines parameters for HostConf.
type HostConfParams struct {
	// XRhIdentity The identity header of the request.
	XRhIdentity []byte `json:"X-Rh-Identity"`

	// XRhInsightsRequestId Unique request id for distributing tracing.
	XRhInsightsRequestId string `json:"X-Rh-Insights-Request-Id"`
}

// CheckHostJSONRequestBody defines body for CheckHost for application/vnd.api+json ContentType.
type CheckHostJSONRequestBody = CheckHosts

// CreateDomainJSONRequestBody defines body for CreateDomain for application/vnd.api+json ContentType.
type CreateDomainJSONRequestBody = CreateDomain

// HostConfJSONRequestBody defines body for HostConf for application/vnd.api+json ContentType.
type HostConfJSONRequestBody = HostConf
