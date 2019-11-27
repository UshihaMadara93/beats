// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package iam

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/iam-2010-05-08/CreateOpenIDConnectProviderRequest
type CreateOpenIDConnectProviderInput struct {
	_ struct{} `type:"structure"`

	// A list of client IDs (also known as audiences). When a mobile or web app
	// registers with an OpenID Connect provider, they establish a value that identifies
	// the application. (This is the value that's sent as the client_id parameter
	// on OAuth requests.)
	//
	// You can register multiple client IDs with the same provider. For example,
	// you might have multiple applications that use the same OIDC provider. You
	// cannot register more than 100 client IDs with a single IAM OIDC provider.
	//
	// There is no defined format for a client ID. The CreateOpenIDConnectProviderRequest
	// operation accepts client IDs up to 255 characters long.
	ClientIDList []string `type:"list"`

	// A list of server certificate thumbprints for the OpenID Connect (OIDC) identity
	// provider's server certificates. Typically this list includes only one entry.
	// However, IAM lets you have up to five thumbprints for an OIDC provider. This
	// lets you maintain multiple thumbprints if the identity provider is rotating
	// certificates.
	//
	// The server certificate thumbprint is the hex-encoded SHA-1 hash value of
	// the X.509 certificate used by the domain where the OpenID Connect provider
	// makes its keys available. It is always a 40-character string.
	//
	// You must provide at least one thumbprint when creating an IAM OIDC provider.
	// For example, assume that the OIDC provider is server.example.com and the
	// provider stores its keys at https://keys.server.example.com/openid-connect.
	// In that case, the thumbprint string would be the hex-encoded SHA-1 hash value
	// of the certificate used by https://keys.server.example.com.
	//
	// For more information about obtaining the OIDC provider's thumbprint, see
	// Obtaining the Thumbprint for an OpenID Connect Provider (https://docs.aws.amazon.com/IAM/latest/UserGuide/identity-providers-oidc-obtain-thumbprint.html)
	// in the IAM User Guide.
	//
	// ThumbprintList is a required field
	ThumbprintList []string `type:"list" required:"true"`

	// The URL of the identity provider. The URL must begin with https:// and should
	// correspond to the iss claim in the provider's OpenID Connect ID tokens. Per
	// the OIDC standard, path components are allowed but query parameters are not.
	// Typically the URL consists of only a hostname, like https://server.example.org
	// or https://example.com.
	//
	// You cannot register the same provider multiple times in a single AWS account.
	// If you try to submit a URL that has already been used for an OpenID Connect
	// provider in the AWS account, you will get an error.
	//
	// Url is a required field
	Url *string `min:"1" type:"string" required:"true"`
}

// String returns the string representation
func (s CreateOpenIDConnectProviderInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *CreateOpenIDConnectProviderInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "CreateOpenIDConnectProviderInput"}

	if s.ThumbprintList == nil {
		invalidParams.Add(aws.NewErrParamRequired("ThumbprintList"))
	}

	if s.Url == nil {
		invalidParams.Add(aws.NewErrParamRequired("Url"))
	}
	if s.Url != nil && len(*s.Url) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("Url", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Contains the response to a successful CreateOpenIDConnectProvider request.
// Please also see https://docs.aws.amazon.com/goto/WebAPI/iam-2010-05-08/CreateOpenIDConnectProviderResponse
type CreateOpenIDConnectProviderOutput struct {
	_ struct{} `type:"structure"`

	// The Amazon Resource Name (ARN) of the new IAM OpenID Connect provider that
	// is created. For more information, see OpenIDConnectProviderListEntry.
	OpenIDConnectProviderArn *string `min:"20" type:"string"`
}

// String returns the string representation
func (s CreateOpenIDConnectProviderOutput) String() string {
	return awsutil.Prettify(s)
}

const opCreateOpenIDConnectProvider = "CreateOpenIDConnectProvider"

// CreateOpenIDConnectProviderRequest returns a request value for making API operation for
// AWS Identity and Access Management.
//
// Creates an IAM entity to describe an identity provider (IdP) that supports
// OpenID Connect (OIDC) (http://openid.net/connect/).
//
// The OIDC provider that you create with this operation can be used as a principal
// in a role's trust policy. Such a policy establishes a trust relationship
// between AWS and the OIDC provider.
//
// When you create the IAM OIDC provider, you specify the following:
//
//    * The URL of the OIDC identity provider (IdP) to trust
//
//    * A list of client IDs (also known as audiences) that identify the application
//    or applications that are allowed to authenticate using the OIDC provider
//
//    * A list of thumbprints of the server certificate(s) that the IdP uses
//
// You get all of this information from the OIDC IdP that you want to use to
// access AWS.
//
// The trust for the OIDC provider is derived from the IAM provider that this
// operation creates. Therefore, it is best to limit access to the CreateOpenIDConnectProvider
// operation to highly privileged users.
//
//    // Example sending a request using CreateOpenIDConnectProviderRequest.
//    req := client.CreateOpenIDConnectProviderRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/iam-2010-05-08/CreateOpenIDConnectProvider
func (c *Client) CreateOpenIDConnectProviderRequest(input *CreateOpenIDConnectProviderInput) CreateOpenIDConnectProviderRequest {
	op := &aws.Operation{
		Name:       opCreateOpenIDConnectProvider,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &CreateOpenIDConnectProviderInput{}
	}

	req := c.newRequest(op, input, &CreateOpenIDConnectProviderOutput{})
	return CreateOpenIDConnectProviderRequest{Request: req, Input: input, Copy: c.CreateOpenIDConnectProviderRequest}
}

// CreateOpenIDConnectProviderRequest is the request type for the
// CreateOpenIDConnectProvider API operation.
type CreateOpenIDConnectProviderRequest struct {
	*aws.Request
	Input *CreateOpenIDConnectProviderInput
	Copy  func(*CreateOpenIDConnectProviderInput) CreateOpenIDConnectProviderRequest
}

// Send marshals and sends the CreateOpenIDConnectProvider API request.
func (r CreateOpenIDConnectProviderRequest) Send(ctx context.Context) (*CreateOpenIDConnectProviderResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &CreateOpenIDConnectProviderResponse{
		CreateOpenIDConnectProviderOutput: r.Request.Data.(*CreateOpenIDConnectProviderOutput),
		response:                          &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// CreateOpenIDConnectProviderResponse is the response type for the
// CreateOpenIDConnectProvider API operation.
type CreateOpenIDConnectProviderResponse struct {
	*CreateOpenIDConnectProviderOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// CreateOpenIDConnectProvider request.
func (r *CreateOpenIDConnectProviderResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}