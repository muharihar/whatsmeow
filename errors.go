package whatsapp

import (
	"encoding/json"
	"errors"
	"fmt"
)

var (
	ErrAlreadyConnected           = errors.New("already connected")
	ErrAlreadyLoggedIn            = errors.New("already logged in")
	ErrInvalidSession             = errors.New("invalid session")
	ErrLoginInProgress            = errors.New("login or restore already running")
	ErrNotConnected               = errors.New("not connected")
	ErrInvalidWsData              = errors.New("received invalid data")
	ErrInvalidWsState             = errors.New("can't handle binary data when not logged in")
	ErrConnectionTimeout          = errors.New("connection timed out")
	ErrMissingMessageTag          = errors.New("no messageTag specified or to short")
	ErrInvalidHmac                = errors.New("invalid hmac")
	ErrInvalidServerResponse      = errors.New("invalid response received from server")
	ErrServerRespondedWith404     = errors.New("server responded with status 404")
	ErrMediaDownloadFailedWith404 = errors.New("download failed with status code 404")
	ErrMediaDownloadFailedWith410 = errors.New("download failed with status code 410")
	ErrLoginTimedOut              = errors.New("login timed out")

	ErrBadRequest   = errors.New("400 (bad request)")
	ErrUnpaired     = errors.New("401 (unpaired from phone)")
	ErrAccessDenied = errors.New("403 (access denied)")
	ErrLoggedIn     = errors.New("405 (already logged in)")
	ErrReplaced     = errors.New("409 (logged in from another location)")

	ErrNoURLPresent       = errors.New("no url present")
	ErrFileLengthMismatch = errors.New("file length does not match")
	ErrInvalidHashLength  = errors.New("hash too short")
	ErrTooShortFile       = errors.New("file too short")
	ErrInvalidMediaHMAC   = errors.New("invalid media hmac")

	ErrCantGetInviteLink = errors.New("you don't have the permission to view the invite link")
	ErrJoinUnauthorized  = errors.New("you're not allowed to join that group")

	ErrInvalidWebsocket = errors.New("invalid websocket")
)

type ErrConnectionFailed struct {
	Err error
}

func (e *ErrConnectionFailed) Error() string {
	return fmt.Sprintf("connection to WhatsApp servers failed: %v", e.Err)
}

type ErrConnectionClosed struct {
	Code int
	Text string
}

func (e *ErrConnectionClosed) Error() string {
	return fmt.Sprintf("server closed connection,code: %d,text: %s", e.Code, e.Text)
}

type StatusResponse struct {
	Status      int    `json:"status"`
	RequestType string `json:"-"`

	Extra map[string]interface{} `json:"-"`
}

func (sr *StatusResponse) UnmarshalJSON(data []byte) error {
	err := json.Unmarshal(data, &sr.Extra)
	if err != nil {
		return err
	}
	sr.Status = int(sr.Extra["status"].(float64))
	return nil
}

func (sr StatusResponse) Error() string {
	return fmt.Sprintf("%s responded with %d", sr.RequestType, sr.Status)
}