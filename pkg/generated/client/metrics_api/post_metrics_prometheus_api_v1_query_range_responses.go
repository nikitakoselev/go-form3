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

// PostMetricsPrometheusAPIV1QueryRangeReader is a Reader for the PostMetricsPrometheusAPIV1QueryRange structure.
type PostMetricsPrometheusAPIV1QueryRangeReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostMetricsPrometheusAPIV1QueryRangeReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewPostMetricsPrometheusAPIV1QueryRangeOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewPostMetricsPrometheusAPIV1QueryRangeBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 403:
		result := NewPostMetricsPrometheusAPIV1QueryRangeForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 500:
		result := NewPostMetricsPrometheusAPIV1QueryRangeInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewPostMetricsPrometheusAPIV1QueryRangeOK creates a PostMetricsPrometheusAPIV1QueryRangeOK with default headers values
func NewPostMetricsPrometheusAPIV1QueryRangeOK() *PostMetricsPrometheusAPIV1QueryRangeOK {
	return &PostMetricsPrometheusAPIV1QueryRangeOK{}
}

/*PostMetricsPrometheusAPIV1QueryRangeOK handles this case with default header values.

See Prometheus api https://prometheus.io/docs/prometheus/latest/querying/api/#instant-queries
*/
type PostMetricsPrometheusAPIV1QueryRangeOK struct {

	//Payload

	// isStream: false
	*models.MetricsQueryResponse
}

func (o *PostMetricsPrometheusAPIV1QueryRangeOK) Error() string {
	return fmt.Sprintf("[POST /metrics/prometheus/api/v1/query_range][%d] postMetricsPrometheusApiV1QueryRangeOK", 200)
}

func (o *PostMetricsPrometheusAPIV1QueryRangeOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.MetricsQueryResponse = new(models.MetricsQueryResponse)

	// response payload

	if err := consumer.Consume(response.Body(), o.MetricsQueryResponse); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostMetricsPrometheusAPIV1QueryRangeBadRequest creates a PostMetricsPrometheusAPIV1QueryRangeBadRequest with default headers values
func NewPostMetricsPrometheusAPIV1QueryRangeBadRequest() *PostMetricsPrometheusAPIV1QueryRangeBadRequest {
	return &PostMetricsPrometheusAPIV1QueryRangeBadRequest{}
}

/*PostMetricsPrometheusAPIV1QueryRangeBadRequest handles this case with default header values.

Bad Request
*/
type PostMetricsPrometheusAPIV1QueryRangeBadRequest struct {

	//Payload

	// isStream: false
	*models.MetricsQueryResponse
}

func (o *PostMetricsPrometheusAPIV1QueryRangeBadRequest) Error() string {
	return fmt.Sprintf("[POST /metrics/prometheus/api/v1/query_range][%d] postMetricsPrometheusApiV1QueryRangeBadRequest", 400)
}

func (o *PostMetricsPrometheusAPIV1QueryRangeBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.MetricsQueryResponse = new(models.MetricsQueryResponse)

	// response payload

	if err := consumer.Consume(response.Body(), o.MetricsQueryResponse); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostMetricsPrometheusAPIV1QueryRangeForbidden creates a PostMetricsPrometheusAPIV1QueryRangeForbidden with default headers values
func NewPostMetricsPrometheusAPIV1QueryRangeForbidden() *PostMetricsPrometheusAPIV1QueryRangeForbidden {
	return &PostMetricsPrometheusAPIV1QueryRangeForbidden{}
}

/*PostMetricsPrometheusAPIV1QueryRangeForbidden handles this case with default header values.

Forbidden
*/
type PostMetricsPrometheusAPIV1QueryRangeForbidden struct {
}

func (o *PostMetricsPrometheusAPIV1QueryRangeForbidden) Error() string {
	return fmt.Sprintf("[POST /metrics/prometheus/api/v1/query_range][%d] postMetricsPrometheusApiV1QueryRangeForbidden", 403)
}

func (o *PostMetricsPrometheusAPIV1QueryRangeForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPostMetricsPrometheusAPIV1QueryRangeInternalServerError creates a PostMetricsPrometheusAPIV1QueryRangeInternalServerError with default headers values
func NewPostMetricsPrometheusAPIV1QueryRangeInternalServerError() *PostMetricsPrometheusAPIV1QueryRangeInternalServerError {
	return &PostMetricsPrometheusAPIV1QueryRangeInternalServerError{}
}

/*PostMetricsPrometheusAPIV1QueryRangeInternalServerError handles this case with default header values.

Internal Server Error
*/
type PostMetricsPrometheusAPIV1QueryRangeInternalServerError struct {
}

func (o *PostMetricsPrometheusAPIV1QueryRangeInternalServerError) Error() string {
	return fmt.Sprintf("[POST /metrics/prometheus/api/v1/query_range][%d] postMetricsPrometheusApiV1QueryRangeInternalServerError", 500)
}

func (o *PostMetricsPrometheusAPIV1QueryRangeInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
