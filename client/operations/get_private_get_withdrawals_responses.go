// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/adampointer/go-deribit/models"
)

// GetPrivateGetWithdrawalsReader is a Reader for the GetPrivateGetWithdrawals structure.
type GetPrivateGetWithdrawalsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetPrivateGetWithdrawalsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetPrivateGetWithdrawalsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetPrivateGetWithdrawalsOK creates a GetPrivateGetWithdrawalsOK with default headers values
func NewGetPrivateGetWithdrawalsOK() *GetPrivateGetWithdrawalsOK {
	return &GetPrivateGetWithdrawalsOK{}
}

/*GetPrivateGetWithdrawalsOK handles this case with default header values.

foo
*/
type GetPrivateGetWithdrawalsOK struct {
	Payload *models.PrivateGetWithdrawalsResponse
}

func (o *GetPrivateGetWithdrawalsOK) Error() string {
	return fmt.Sprintf("[GET /private/get_withdrawals][%d] getPrivateGetWithdrawalsOK  %+v", 200, o.Payload)
}

func (o *GetPrivateGetWithdrawalsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PrivateGetWithdrawalsResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}