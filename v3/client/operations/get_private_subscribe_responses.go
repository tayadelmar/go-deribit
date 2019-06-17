// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/adampointer/go-deribit/v3/models"
)

// GetPrivateSubscribeReader is a Reader for the GetPrivateSubscribe structure.
type GetPrivateSubscribeReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetPrivateSubscribeReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetPrivateSubscribeOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 401:
		result := NewGetPrivateSubscribeUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetPrivateSubscribeOK creates a GetPrivateSubscribeOK with default headers values
func NewGetPrivateSubscribeOK() *GetPrivateSubscribeOK {
	return &GetPrivateSubscribeOK{}
}

/*GetPrivateSubscribeOK handles this case with default header values.

ok response
*/
type GetPrivateSubscribeOK struct {
	Payload *models.PrivateSubscribeResponse
}

func (o *GetPrivateSubscribeOK) Error() string {
	return fmt.Sprintf("[GET /private/subscribe][%d] getPrivateSubscribeOK  %+v", 200, o.Payload)
}

func (o *GetPrivateSubscribeOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PrivateSubscribeResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetPrivateSubscribeUnauthorized creates a GetPrivateSubscribeUnauthorized with default headers values
func NewGetPrivateSubscribeUnauthorized() *GetPrivateSubscribeUnauthorized {
	return &GetPrivateSubscribeUnauthorized{}
}

/*GetPrivateSubscribeUnauthorized handles this case with default header values.

not authorised
*/
type GetPrivateSubscribeUnauthorized struct {
	Payload *models.ErrorMessage
}

func (o *GetPrivateSubscribeUnauthorized) Error() string {
	return fmt.Sprintf("[GET /private/subscribe][%d] getPrivateSubscribeUnauthorized  %+v", 401, o.Payload)
}

func (o *GetPrivateSubscribeUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorMessage)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}