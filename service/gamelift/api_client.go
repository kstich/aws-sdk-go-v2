// Code generated by smithy-go-codegen DO NOT EDIT.

package gamelift

import (
	"context"
	"github.com/aws/aws-sdk-go-v2/aws"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	awshttp "github.com/aws/aws-sdk-go-v2/aws/transport/http"
	"github.com/awslabs/smithy-go/middleware"
	"net/http"
	"time"
)

const ServiceID = "GameLift"

// Amazon GameLift Service Amazon GameLift provides a range of multiplayer game
// hosting solutions. As a fully managed service, GameLift helps you:
//
//     * Set up
// EC2-based computing resources and use GameLift FleetIQ to and deploy your game
// servers on low-cost, reliable Spot instances.
//
//     * Track game server
// availability and route players into game sessions to minimize latency.
//
//     *
// Automatically scale your resources to meet player demand and manage costs
//
//     *
// Optionally add FlexMatch matchmaking.
//
//     <p>With GameLift as a managed
// service, you have the option to deploy your custom game server or use Amazon
// GameLift Realtime Servers to quickly stand up lightweight game servers for your
// game. Realtime Servers provides an efficient game server framework with core
// Amazon GameLift infrastructure already built in.</p> <p> <b>Now in Public
// Preview:</b> </p> <p>Use GameLift FleetIQ as a standalone feature with EC2
// instances and Auto Scaling groups. GameLift FleetIQ provides optimizations that
// make low-cost Spot instances viable for game hosting. This extension of GameLift
// FleetIQ gives you access to these optimizations while managing your EC2
// instances and Auto Scaling groups within your own AWS account.</p> <p> <b>Get
// Amazon GameLift Tools and Resources</b> </p> <p>This reference guide describes
// the low-level service API for Amazon GameLift and provides links to
// language-specific SDK reference topics. See also <a
// href="https://docs.aws.amazon.com/gamelift/latest/developerguide/gamelift-components.html">
// Amazon GameLift Tools and Resources</a>.</p> <p> <b>API Summary</b> </p> <p>The
// Amazon GameLift service API includes two key sets of actions:</p> <ul> <li>
// <p>Manage game sessions and player access -- Integrate this functionality into
// game client services in order to create new game sessions, retrieve information
// on existing game sessions; reserve a player slot in a game session, request
// matchmaking, etc.</p> </li> <li> <p>Configure and manage game server resources
// -- Manage your Amazon GameLift hosting resources, including builds, scripts,
// fleets, queues, and aliases. Set up matchmakers, configure auto-scaling,
// retrieve game logs, and get hosting and game metrics.</p> </li> </ul> <p> <b> <a
// href="https://docs.aws.amazon.com/gamelift/latest/developerguide/reference-awssdk.html">
// Task-based list of API actions</a> </b> </p>
type Client struct {
	options Options
}

// New returns an initialized Client based on the functional options. Provide
// additional functional options to further configure the behavior of the client,
// such as changing the client's endpoint or adding custom middleware behavior.
func New(options Options, optFns ...func(*Options)) *Client {
	options = options.Copy()

	resolveRetryer(&options)

	resolveHTTPClient(&options)

	resolveHTTPSignerV4(&options)

	resolveDefaultEndpointConfiguration(&options)

	for _, fn := range optFns {
		fn(&options)
	}

	client := &Client{
		options: options,
	}

	return client
}

type Options struct {
	// Set of options to modify how an operation is invoked. These apply to all
	// operations invoked for this client. Use functional options on operation call to
	// modify this list for per operation behavior.
	APIOptions []APIOptionFunc

	// The credentials object to use when signing requests.
	Credentials aws.CredentialsProvider

	// The endpoint options to be used when attempting to resolve an endpoint.
	EndpointOptions ResolverOptions

	// The service endpoint resolver.
	EndpointResolver EndpointResolver

	// Signature Version 4 (SigV4) Signer
	HTTPSignerV4 HTTPSignerV4

	// An integer value representing the logging level.
	LogLevel aws.LogLevel

	// The logger writer interface to write logging messages to.
	Logger aws.Logger

	// The region to send requests to. (Required)
	Region string

	// Retryer guides how HTTP requests should be retried in case of recoverable
	// failures. When nil the API client will use a default retryer.
	Retryer retry.Retryer

	// The HTTP client to invoke API calls with. Defaults to client's default HTTP
	// implementation if nil.
	HTTPClient HTTPClient
}

func (o Options) GetCredentials() aws.CredentialsProvider {
	return o.Credentials
}

func (o Options) GetEndpointOptions() ResolverOptions {
	return o.EndpointOptions
}

func (o Options) GetEndpointResolver() EndpointResolver {
	return o.EndpointResolver
}

func (o Options) GetHTTPSignerV4() HTTPSignerV4 {
	return o.HTTPSignerV4
}

func (o Options) GetLogLevel() aws.LogLevel {
	return o.LogLevel
}

func (o Options) GetLogger() aws.Logger {
	return o.Logger
}

func (o Options) GetRegion() string {
	return o.Region
}

func (o Options) GetRetryer() retry.Retryer {
	return o.Retryer
}

type HTTPClient interface {
	Do(*http.Request) (*http.Response, error)
}

// Copy creates a clone where the APIOptions list is deep copied.
func (o Options) Copy() Options {
	to := o
	to.APIOptions = make([]APIOptionFunc, len(o.APIOptions))
	copy(to.APIOptions, o.APIOptions)
	return to
}

type APIOptionFunc func(*middleware.Stack) error

// NewFromConfig returns a new client from the provided config.
func NewFromConfig(cfg aws.Config, optFns ...func(*Options)) *Client {
	opts := Options{
		Region:      cfg.Region,
		Retryer:     cfg.Retryer,
		LogLevel:    cfg.LogLevel,
		Logger:      cfg.Logger,
		HTTPClient:  cfg.HTTPClient,
		Credentials: cfg.Credentials,
	}
	return New(opts, optFns...)
}

func resolveHTTPClient(o *Options) {
	if o.HTTPClient != nil {
		return
	}
	o.HTTPClient = aws.NewBuildableHTTPClient()
}

func resolveRetryer(o *Options) {
	if o.Retryer != nil {
		return
	}
	o.Retryer = retry.NewStandard()
}

func addClientUserAgent(stack *middleware.Stack) {
	awsmiddleware.AddUserAgentKey("gamelift")(stack)
}

func addHTTPSignerV4Middleware(stack *middleware.Stack, o Options) {
	stack.Finalize.Add(v4.NewSignHTTPRequestMiddleware(o.Credentials, o.HTTPSignerV4), middleware.After)
}

type HTTPSignerV4 interface {
	SignHTTP(ctx context.Context, credentials aws.Credentials, r *http.Request, payloadHash string, service string, region string, signingTime time.Time) error
}

func resolveHTTPSignerV4(o *Options) {
	if o.HTTPSignerV4 != nil {
		return
	}
	o.HTTPSignerV4 = v4.NewSigner()
}

func addResponseErrorWrapper(stack *middleware.Stack) {
	awshttp.AddResponseErrorWrapper(stack)
}
