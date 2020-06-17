// Code generated by go-swagger; DO NOT EDIT.

package words

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"dictionary/models"
)

// GetWordsOKCode is the HTTP code returned for type GetWordsOK
const GetWordsOKCode int = 200

/*GetWordsOK Successful operation

swagger:response getWordsOK
*/
type GetWordsOK struct {

	/*
	  In: Body
	*/
	Payload []*models.Word `json:"body,omitempty"`
}

// NewGetWordsOK creates GetWordsOK with default headers values
func NewGetWordsOK() *GetWordsOK {

	return &GetWordsOK{}
}

// WithPayload adds the payload to the get words o k response
func (o *GetWordsOK) WithPayload(payload []*models.Word) *GetWordsOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get words o k response
func (o *GetWordsOK) SetPayload(payload []*models.Word) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetWordsOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	payload := o.Payload
	if payload == nil {
		// return empty array
		payload = make([]*models.Word, 0, 50)
	}

	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}

// GetWordsBadRequestCode is the HTTP code returned for type GetWordsBadRequest
const GetWordsBadRequestCode int = 400

/*GetWordsBadRequest Invalid input

swagger:response getWordsBadRequest
*/
type GetWordsBadRequest struct {

	/*
	  In: Body
	*/
	Payload *models.Error400 `json:"body,omitempty"`
}

// NewGetWordsBadRequest creates GetWordsBadRequest with default headers values
func NewGetWordsBadRequest() *GetWordsBadRequest {

	return &GetWordsBadRequest{}
}

// WithPayload adds the payload to the get words bad request response
func (o *GetWordsBadRequest) WithPayload(payload *models.Error400) *GetWordsBadRequest {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get words bad request response
func (o *GetWordsBadRequest) SetPayload(payload *models.Error400) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetWordsBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(400)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetWordsInternalServerErrorCode is the HTTP code returned for type GetWordsInternalServerError
const GetWordsInternalServerErrorCode int = 500

/*GetWordsInternalServerError Internal Server Error

swagger:response getWordsInternalServerError
*/
type GetWordsInternalServerError struct {

	/*
	  In: Body
	*/
	Payload *models.Error500 `json:"body,omitempty"`
}

// NewGetWordsInternalServerError creates GetWordsInternalServerError with default headers values
func NewGetWordsInternalServerError() *GetWordsInternalServerError {

	return &GetWordsInternalServerError{}
}

// WithPayload adds the payload to the get words internal server error response
func (o *GetWordsInternalServerError) WithPayload(payload *models.Error500) *GetWordsInternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get words internal server error response
func (o *GetWordsInternalServerError) SetPayload(payload *models.Error500) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetWordsInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(500)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
