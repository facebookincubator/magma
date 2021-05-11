// Code generated by go-swagger; DO NOT EDIT.

package policies

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"

	models "magma/orc8r/cloud/go/obsidian/swagger/v1/models"
)

// NewPutNetworksNetworkIDPoliciesRulesRuleIDParams creates a new PutNetworksNetworkIDPoliciesRulesRuleIDParams object
// with the default values initialized.
func NewPutNetworksNetworkIDPoliciesRulesRuleIDParams() *PutNetworksNetworkIDPoliciesRulesRuleIDParams {
	var ()
	return &PutNetworksNetworkIDPoliciesRulesRuleIDParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPutNetworksNetworkIDPoliciesRulesRuleIDParamsWithTimeout creates a new PutNetworksNetworkIDPoliciesRulesRuleIDParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPutNetworksNetworkIDPoliciesRulesRuleIDParamsWithTimeout(timeout time.Duration) *PutNetworksNetworkIDPoliciesRulesRuleIDParams {
	var ()
	return &PutNetworksNetworkIDPoliciesRulesRuleIDParams{

		timeout: timeout,
	}
}

// NewPutNetworksNetworkIDPoliciesRulesRuleIDParamsWithContext creates a new PutNetworksNetworkIDPoliciesRulesRuleIDParams object
// with the default values initialized, and the ability to set a context for a request
func NewPutNetworksNetworkIDPoliciesRulesRuleIDParamsWithContext(ctx context.Context) *PutNetworksNetworkIDPoliciesRulesRuleIDParams {
	var ()
	return &PutNetworksNetworkIDPoliciesRulesRuleIDParams{

		Context: ctx,
	}
}

// NewPutNetworksNetworkIDPoliciesRulesRuleIDParamsWithHTTPClient creates a new PutNetworksNetworkIDPoliciesRulesRuleIDParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPutNetworksNetworkIDPoliciesRulesRuleIDParamsWithHTTPClient(client *http.Client) *PutNetworksNetworkIDPoliciesRulesRuleIDParams {
	var ()
	return &PutNetworksNetworkIDPoliciesRulesRuleIDParams{
		HTTPClient: client,
	}
}

/*PutNetworksNetworkIDPoliciesRulesRuleIDParams contains all the parameters to send to the API endpoint
for the put networks network ID policies rules rule ID operation typically these are written to a http.Request
*/
type PutNetworksNetworkIDPoliciesRulesRuleIDParams struct {

	/*PolicyRule
	  Policy rule

	*/
	PolicyRule *models.PolicyRule
	/*NetworkID
	  Network ID

	*/
	NetworkID string
	/*RuleID
	  Rule Id

	*/
	RuleID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the put networks network ID policies rules rule ID params
func (o *PutNetworksNetworkIDPoliciesRulesRuleIDParams) WithTimeout(timeout time.Duration) *PutNetworksNetworkIDPoliciesRulesRuleIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the put networks network ID policies rules rule ID params
func (o *PutNetworksNetworkIDPoliciesRulesRuleIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the put networks network ID policies rules rule ID params
func (o *PutNetworksNetworkIDPoliciesRulesRuleIDParams) WithContext(ctx context.Context) *PutNetworksNetworkIDPoliciesRulesRuleIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the put networks network ID policies rules rule ID params
func (o *PutNetworksNetworkIDPoliciesRulesRuleIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the put networks network ID policies rules rule ID params
func (o *PutNetworksNetworkIDPoliciesRulesRuleIDParams) WithHTTPClient(client *http.Client) *PutNetworksNetworkIDPoliciesRulesRuleIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the put networks network ID policies rules rule ID params
func (o *PutNetworksNetworkIDPoliciesRulesRuleIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithPolicyRule adds the policyRule to the put networks network ID policies rules rule ID params
func (o *PutNetworksNetworkIDPoliciesRulesRuleIDParams) WithPolicyRule(policyRule *models.PolicyRule) *PutNetworksNetworkIDPoliciesRulesRuleIDParams {
	o.SetPolicyRule(policyRule)
	return o
}

// SetPolicyRule adds the policyRule to the put networks network ID policies rules rule ID params
func (o *PutNetworksNetworkIDPoliciesRulesRuleIDParams) SetPolicyRule(policyRule *models.PolicyRule) {
	o.PolicyRule = policyRule
}

// WithNetworkID adds the networkID to the put networks network ID policies rules rule ID params
func (o *PutNetworksNetworkIDPoliciesRulesRuleIDParams) WithNetworkID(networkID string) *PutNetworksNetworkIDPoliciesRulesRuleIDParams {
	o.SetNetworkID(networkID)
	return o
}

// SetNetworkID adds the networkId to the put networks network ID policies rules rule ID params
func (o *PutNetworksNetworkIDPoliciesRulesRuleIDParams) SetNetworkID(networkID string) {
	o.NetworkID = networkID
}

// WithRuleID adds the ruleID to the put networks network ID policies rules rule ID params
func (o *PutNetworksNetworkIDPoliciesRulesRuleIDParams) WithRuleID(ruleID string) *PutNetworksNetworkIDPoliciesRulesRuleIDParams {
	o.SetRuleID(ruleID)
	return o
}

// SetRuleID adds the ruleId to the put networks network ID policies rules rule ID params
func (o *PutNetworksNetworkIDPoliciesRulesRuleIDParams) SetRuleID(ruleID string) {
	o.RuleID = ruleID
}

// WriteToRequest writes these params to a swagger request
func (o *PutNetworksNetworkIDPoliciesRulesRuleIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.PolicyRule != nil {
		if err := r.SetBodyParam(o.PolicyRule); err != nil {
			return err
		}
	}

	// path param network_id
	if err := r.SetPathParam("network_id", o.NetworkID); err != nil {
		return err
	}

	// path param rule_id
	if err := r.SetPathParam("rule_id", o.RuleID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
