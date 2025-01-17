// Code generated by go-swagger; DO NOT EDIT.

package metrics_api

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/form3tech-oss/go-form3/v4/pkg/generated/models"
)

// PostMetricsPrometheusAPIV1QueryReader is a Reader for the PostMetricsPrometheusAPIV1Query structure.
type PostMetricsPrometheusAPIV1QueryReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostMetricsPrometheusAPIV1QueryReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewPostMetricsPrometheusAPIV1QueryOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewPostMetricsPrometheusAPIV1QueryBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 403:
		result := NewPostMetricsPrometheusAPIV1QueryForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 500:
		result := NewPostMetricsPrometheusAPIV1QueryInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewPostMetricsPrometheusAPIV1QueryOK creates a PostMetricsPrometheusAPIV1QueryOK with default headers values
func NewPostMetricsPrometheusAPIV1QueryOK() *PostMetricsPrometheusAPIV1QueryOK {
	return &PostMetricsPrometheusAPIV1QueryOK{}
}

/*PostMetricsPrometheusAPIV1QueryOK handles this case with default header values.

See Prometheus api https://prometheus.io/docs/prometheus/latest/querying/api/#instant-queries
*/
type PostMetricsPrometheusAPIV1QueryOK struct {

	//Payload

	// isStream: false
	*models.MetricsQueryResponse
}

func (o *PostMetricsPrometheusAPIV1QueryOK) Error() string {
	return fmt.Sprintf("[POST /metrics/prometheus/api/v1/query][%d] postMetricsPrometheusApiV1QueryOK", 200)
}

func (o *PostMetricsPrometheusAPIV1QueryOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.MetricsQueryResponse = new(models.MetricsQueryResponse)

	// response payload

	if err := consumer.Consume(response.Body(), o.MetricsQueryResponse); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostMetricsPrometheusAPIV1QueryBadRequest creates a PostMetricsPrometheusAPIV1QueryBadRequest with default headers values
func NewPostMetricsPrometheusAPIV1QueryBadRequest() *PostMetricsPrometheusAPIV1QueryBadRequest {
	return &PostMetricsPrometheusAPIV1QueryBadRequest{}
}

/*PostMetricsPrometheusAPIV1QueryBadRequest handles this case with default header values.

Bad Request
*/
type PostMetricsPrometheusAPIV1QueryBadRequest struct {

	//Payload

	// isStream: false
	*models.MetricsQueryResponse
}

func (o *PostMetricsPrometheusAPIV1QueryBadRequest) Error() string {
	return fmt.Sprintf("[POST /metrics/prometheus/api/v1/query][%d] postMetricsPrometheusApiV1QueryBadRequest", 400)
}

func (o *PostMetricsPrometheusAPIV1QueryBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.MetricsQueryResponse = new(models.MetricsQueryResponse)

	// response payload

	if err := consumer.Consume(response.Body(), o.MetricsQueryResponse); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostMetricsPrometheusAPIV1QueryForbidden creates a PostMetricsPrometheusAPIV1QueryForbidden with default headers values
func NewPostMetricsPrometheusAPIV1QueryForbidden() *PostMetricsPrometheusAPIV1QueryForbidden {
	return &PostMetricsPrometheusAPIV1QueryForbidden{}
}

/*PostMetricsPrometheusAPIV1QueryForbidden handles this case with default header values.

Forbidden
*/
type PostMetricsPrometheusAPIV1QueryForbidden struct {
}

func (o *PostMetricsPrometheusAPIV1QueryForbidden) Error() string {
	return fmt.Sprintf("[POST /metrics/prometheus/api/v1/query][%d] postMetricsPrometheusApiV1QueryForbidden", 403)
}

func (o *PostMetricsPrometheusAPIV1QueryForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPostMetricsPrometheusAPIV1QueryInternalServerError creates a PostMetricsPrometheusAPIV1QueryInternalServerError with default headers values
func NewPostMetricsPrometheusAPIV1QueryInternalServerError() *PostMetricsPrometheusAPIV1QueryInternalServerError {
	return &PostMetricsPrometheusAPIV1QueryInternalServerError{}
}

/*PostMetricsPrometheusAPIV1QueryInternalServerError handles this case with default header values.

Internal Server Error
*/
type PostMetricsPrometheusAPIV1QueryInternalServerError struct {
}

func (o *PostMetricsPrometheusAPIV1QueryInternalServerError) Error() string {
	return fmt.Sprintf("[POST /metrics/prometheus/api/v1/query][%d] postMetricsPrometheusApiV1QueryInternalServerError", 500)
}

func (o *PostMetricsPrometheusAPIV1QueryInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
