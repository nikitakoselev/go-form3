// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"log"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/validate"
)

// ResourceType resource type
// swagger:model ResourceType
type ResourceType string

const (

	// ResourceTypePaymentDefaults captures enum value "payment_defaults"
	ResourceTypePaymentDefaults ResourceType = "payment_defaults"

	// ResourceTypePartyProducts captures enum value "party_products"
	ResourceTypePartyProducts ResourceType = "party_products"

	// ResourceTypeSchemeMessages captures enum value "scheme_messages"
	ResourceTypeSchemeMessages ResourceType = "scheme_messages"

	// ResourceTypeReversalSubmissionValidations captures enum value "reversal_submission_validations"
	ResourceTypeReversalSubmissionValidations ResourceType = "reversal_submission_validations"

	// ResourceTypePaymentAdmissions captures enum value "payment_admissions"
	ResourceTypePaymentAdmissions ResourceType = "payment_admissions"

	// ResourceTypeReturnAdmissions captures enum value "return_admissions"
	ResourceTypeReturnAdmissions ResourceType = "return_admissions"

	// ResourceTypeFxDeals captures enum value "fx_deals"
	ResourceTypeFxDeals ResourceType = "fx_deals"

	// ResourceTypePaymentBatches captures enum value "payment_batches"
	ResourceTypePaymentBatches ResourceType = "payment_batches"

	// ResourceTypePaymentSubmissionValidations captures enum value "payment_submission_validations"
	ResourceTypePaymentSubmissionValidations ResourceType = "payment_submission_validations"

	// ResourceTypeRecallSubmissionValidations captures enum value "recall_submission_validations"
	ResourceTypeRecallSubmissionValidations ResourceType = "recall_submission_validations"

	// ResourceTypeReturnReversalAdmissions captures enum value "return_reversal_admissions"
	ResourceTypeReturnReversalAdmissions ResourceType = "return_reversal_admissions"

	// ResourceTypeReturnSubmissions captures enum value "return_submissions"
	ResourceTypeReturnSubmissions ResourceType = "return_submissions"

	// ResourceTypeRecallDecisionSubmissionValidations captures enum value "recall_decision_submission_validations"
	ResourceTypeRecallDecisionSubmissionValidations ResourceType = "recall_decision_submission_validations"

	// ResourceTypeReturnReversals captures enum value "return_reversals"
	ResourceTypeReturnReversals ResourceType = "return_reversals"

	// ResourceTypeRecallReversalAdmissions captures enum value "recall_reversal_admissions"
	ResourceTypeRecallReversalAdmissions ResourceType = "recall_reversal_admissions"

	// ResourceTypeAccounts captures enum value "accounts"
	ResourceTypeAccounts ResourceType = "accounts"

	// ResourceTypeReversals captures enum value "reversals"
	ResourceTypeReversals ResourceType = "reversals"

	// ResourceTypeAccountIndirects captures enum value "account_indirects"
	ResourceTypeAccountIndirects ResourceType = "account_indirects"

	// ResourceTypeRecallDecisionAdmissions captures enum value "recall_decision_admissions"
	ResourceTypeRecallDecisionAdmissions ResourceType = "recall_decision_admissions"

	// ResourceTypeRecallDecisionSubmissions captures enum value "recall_decision_submissions"
	ResourceTypeRecallDecisionSubmissions ResourceType = "recall_decision_submissions"

	// ResourceTypeContactAccounts captures enum value "contact_accounts"
	ResourceTypeContactAccounts ResourceType = "contact_accounts"

	// ResourceTypeLimits captures enum value "limits"
	ResourceTypeLimits ResourceType = "limits"

	// ResourceTypeDirectAccount captures enum value "direct_account"
	ResourceTypeDirectAccount ResourceType = "direct_account"

	// ResourceTypePaymentAutomaticReturns captures enum value "payment_automatic_returns"
	ResourceTypePaymentAutomaticReturns ResourceType = "payment_automatic_returns"

	// ResourceTypePaymentAdviceSubmissionValidations captures enum value "payment_advice_submission_validations"
	ResourceTypePaymentAdviceSubmissionValidations ResourceType = "payment_advice_submission_validations"

	// ResourceTypeBics captures enum value "bics"
	ResourceTypeBics ResourceType = "bics"

	// ResourceTypePartyProductEvents captures enum value "party_product_events"
	ResourceTypePartyProductEvents ResourceType = "party_product_events"

	// ResourceTypeAccountConfigurations captures enum value "account_configurations"
	ResourceTypeAccountConfigurations ResourceType = "account_configurations"

	// ResourceTypeAccountRoutings captures enum value "account_routings"
	ResourceTypeAccountRoutings ResourceType = "account_routings"

	// ResourceTypeAccountEvents captures enum value "account_events"
	ResourceTypeAccountEvents ResourceType = "account_events"

	// ResourceTypePositions captures enum value "positions"
	ResourceTypePositions ResourceType = "positions"

	// ResourceTypeBankIds captures enum value "bank_ids"
	ResourceTypeBankIds ResourceType = "bank_ids"

	// ResourceTypePartyActors captures enum value "party_actors"
	ResourceTypePartyActors ResourceType = "party_actors"

	// ResourceTypeRecallAdmissions captures enum value "recall_admissions"
	ResourceTypeRecallAdmissions ResourceType = "recall_admissions"

	// ResourceTypeProductAssociations captures enum value "product_associations"
	ResourceTypeProductAssociations ResourceType = "product_associations"

	// ResourceTypeReturns captures enum value "returns"
	ResourceTypeReturns ResourceType = "returns"

	// ResourceTypePaymentAdviceSubmissions captures enum value "payment_advice_submissions"
	ResourceTypePaymentAdviceSubmissions ResourceType = "payment_advice_submissions"

	// ResourceTypeReversalSubmissions captures enum value "reversal_submissions"
	ResourceTypeReversalSubmissions ResourceType = "reversal_submissions"

	// ResourceTypeRecallDecisions captures enum value "recall_decisions"
	ResourceTypeRecallDecisions ResourceType = "recall_decisions"

	// ResourceTypeParties captures enum value "parties"
	ResourceTypeParties ResourceType = "parties"

	// ResourceTypePayments captures enum value "payments"
	ResourceTypePayments ResourceType = "payments"

	// ResourceTypeReversalAdmissions captures enum value "reversal_admissions"
	ResourceTypeReversalAdmissions ResourceType = "reversal_admissions"

	// ResourceTypeDirectDebit captures enum value "direct_debit"
	ResourceTypeDirectDebit ResourceType = "direct_debit"

	// ResourceTypeSchemeMessageAdmissions captures enum value "scheme_message_admissions"
	ResourceTypeSchemeMessageAdmissions ResourceType = "scheme_message_admissions"

	// ResourceTypePartyProductUpdatedEvents captures enum value "party_product_updated_events"
	ResourceTypePartyProductUpdatedEvents ResourceType = "party_product_updated_events"

	// ResourceTypeRecallReversals captures enum value "recall_reversals"
	ResourceTypeRecallReversals ResourceType = "recall_reversals"

	// ResourceTypePartyAccounts captures enum value "party_accounts"
	ResourceTypePartyAccounts ResourceType = "party_accounts"

	// ResourceTypePaymentSubmissions captures enum value "payment_submissions"
	ResourceTypePaymentSubmissions ResourceType = "payment_submissions"

	// ResourceTypeContacts captures enum value "contacts"
	ResourceTypeContacts ResourceType = "contacts"

	// ResourceTypeReturnSubmissionValidations captures enum value "return_submission_validations"
	ResourceTypeReturnSubmissionValidations ResourceType = "return_submission_validations"

	// ResourceTypeRecallSubmissions captures enum value "recall_submissions"
	ResourceTypeRecallSubmissions ResourceType = "recall_submissions"

	// ResourceTypePaymentAdvices captures enum value "payment_advices"
	ResourceTypePaymentAdvices ResourceType = "payment_advices"

	// ResourceTypeRecalls captures enum value "recalls"
	ResourceTypeRecalls ResourceType = "recalls"
)

// for schema
var resourceTypeEnum []interface{}

func init() {
	var res []ResourceType
	if err := json.Unmarshal([]byte(`["payment_defaults","party_products","scheme_messages","reversal_submission_validations","payment_admissions","return_admissions","fx_deals","payment_batches","payment_submission_validations","recall_submission_validations","return_reversal_admissions","return_submissions","recall_decision_submission_validations","return_reversals","recall_reversal_admissions","accounts","reversals","account_indirects","recall_decision_admissions","recall_decision_submissions","contact_accounts","limits","direct_account","payment_automatic_returns","payment_advice_submission_validations","bics","party_product_events","account_configurations","account_routings","account_events","positions","bank_ids","party_actors","recall_admissions","product_associations","returns","payment_advice_submissions","reversal_submissions","recall_decisions","parties","payments","reversal_admissions","direct_debit","scheme_message_admissions","party_product_updated_events","recall_reversals","party_accounts","payment_submissions","contacts","return_submission_validations","recall_submissions","payment_advices","recalls"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		resourceTypeEnum = append(resourceTypeEnum, v)
	}
}

func (m ResourceType) validateResourceTypeEnum(path, location string, value ResourceType) error {
	if err := validate.Enum(path, location, value, resourceTypeEnum); err != nil {
		return err
	}
	return nil
}

// Validate validates this resource type
func (m ResourceType) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateResourceTypeEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ResourceType) Json() string {
	json, err := json.MarshalIndent(m, "  ", "  ")
	if err != nil {
		log.Fatal(err)
	}
	return string(json)
}
