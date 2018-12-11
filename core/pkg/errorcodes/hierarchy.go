package errorcodes

import (
	"google.golang.org/grpc/codes"
)

var errorHierarchy = map[Code][]Code{
	ErrEnvelopeUntrusted:         {ErrEnvelope},
	ErrEnvelopeNoDeviceFound:     {ErrEnvelope},
	ErrEntityInput:               {ErrEntity},
	ErrEntityExists:              {ErrEntity},
	ErrNetAnotherClientConnected: {ErrNet},
	ErrNetStream:                 {ErrNet},
	ErrNetQueue:                  {ErrNet},
	ErrEventSender:               {ErrEvent},
	ErrEventData:                 {ErrEvent},
	ErrValidationInput:           {ErrValidation},
	ErrValidationMyself:          {ErrValidation},
	ErrDbCreate:                  {ErrDb},
	ErrDbUpdate:                  {ErrDb},
	ErrDbDelete:                  {ErrDb},
	ErrDbNothingFound:            {ErrDb},
	ErrCryptoSig:                 {ErrCrypto},
	ErrCryptoEncrypt:             {ErrCrypto},
	ErrCryptoDecrypt:             {ErrCrypto},
	ErrCryptoKey:                 {ErrCrypto},
	ErrCryptoKeyGen:              {ErrCryptoKey},
	ErrCryptoKeyDecode:           {ErrCryptoKey},
	ErrCryptoSigchain:            {ErrCrypto},
	ErrContactReqKey:             {ErrContactReq, ErrCryptoKeyDecode},
	ErrContactReqKeyMissing:      {ErrContactReq},
	ErrContactReqExisting:        {ErrContactReq},
	ErrContactReqMyself:          {ErrContactReq, ErrValidationMyself},
}

var grpcStatuses = map[Code]codes.Code{
	ErrEnvelope:                  codes.DataLoss,
	ErrPanic:                     codes.Internal,
	ErrEntityInput:               codes.InvalidArgument,
	ErrEntityExists:              codes.Internal,
	ErrNetAnotherClientConnected: codes.Internal,
	ErrNet:                       codes.FailedPrecondition,
	ErrEventSender:               codes.DataLoss,
	ErrEventData:                 codes.DataLoss,
	ErrValidation:                codes.InvalidArgument,
	ErrDb:                        codes.Internal,
	ErrDbNothingFound:            codes.NotFound,
	ErrCryptoSig:                 codes.InvalidArgument,
	ErrCryptoEncrypt:             codes.InvalidArgument,
	ErrCryptoDecrypt:             codes.InvalidArgument,
	ErrCryptoKey:                 codes.InvalidArgument,
	ErrCryptoKeyGen:              codes.Internal,
	ErrCryptoKeyDecode:           codes.InvalidArgument,
	ErrCryptoSigchain:            codes.FailedPrecondition,
	ErrContactReqKey:             codes.InvalidArgument,
	ErrContactReqKeyMissing:      codes.InvalidArgument,
	ErrContactReqExisting:        codes.AlreadyExists,
	ErrContactReqMyself:          codes.InvalidArgument,
}
