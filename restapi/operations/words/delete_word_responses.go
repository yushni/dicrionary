// Code generated by go-swagger; DO NOT EDIT.

package words

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"dictionary/models"
)

// DeleteWordBadRequestCode is the HTTP code returned for type DeleteWordBadRequest
const DeleteWordBadRequestCode int = 400

/*DeleteWordBadRequest Invalid Word ID

swagger:response deleteWordBadRequest
*/
type DeleteWordBadRequest struct {

	/*
	  In: Body
	*/
	Payload *models.Error400 `json:"body,omitempty"`
}

// NewDeleteWordBadRequest creates DeleteWordBadRequest with default headers values
func NewDeleteWordBadRequest() *DeleteWordBadRequest {

	return &DeleteWordBadRequest{}
}

// WithPayload adds the payload to the delete word bad request response
func (o *DeleteWordBadRequest) WithPayload(payload *models.Error400) *DeleteWordBadRequest {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the delete word bad request response
func (o *DeleteWordBadRequest) SetPayload(payload *models.Error400) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *DeleteWordBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(400)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// DeleteWordNotFoundCode is the HTTP code returned for type DeleteWordNotFound
const DeleteWordNotFoundCode int = 404

/*DeleteWordNotFound Word not found

swagger:response deleteWordNotFound
*/
type DeleteWordNotFound struct {

	/*
	  In: Body
	*/
	Payload *models.Error404 `json:"body,omitempty"`
}

// NewDeleteWordNotFound creates DeleteWordNotFound with default headers values
func NewDeleteWordNotFound() *DeleteWordNotFound {

	return &DeleteWordNotFound{}
}

// WithPayload adds the payload to the delete word not found response
func (o *DeleteWordNotFound) WithPayload(payload *models.Error404) *DeleteWordNotFound {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the delete word not found response
func (o *DeleteWordNotFound) SetPayload(payload *models.Error404) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *DeleteWordNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(404)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// DeleteWordInternalServerErrorCode is the HTTP code returned for type DeleteWordInternalServerError
const DeleteWordInternalServerErrorCode int = 500

/*DeleteWordInternalServerError Internal Server Error

swagger:response deleteWordInternalServerError
*/
type DeleteWordInternalServerError struct {

	/*
	  In: Body
	*/
	Payload *models.Error500 `json:"body,omitempty"`
}

// NewDeleteWordInternalServerError creates DeleteWordInternalServerError with default headers values
func NewDeleteWordInternalServerError() *DeleteWordInternalServerError {

	return &DeleteWordInternalServerError{}
}

// WithPayload adds the payload to the delete word internal server error response
func (o *DeleteWordInternalServerError) WithPayload(payload *models.Error500) *DeleteWordInternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the delete word internal server error response
func (o *DeleteWordInternalServerError) SetPayload(payload *models.Error500) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *DeleteWordInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(500)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
