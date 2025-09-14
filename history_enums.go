package wigal

// TEXT | FLASH | CALL | RECORD | IVR | TEXT2SPEECH | OTP | SIMACTIVE | KYCVALIDATE
type ServiceType int8

const (
	TEXT ServiceType = iota
	FLASH
	CALL
	RECORD
	IVR
	TEXT2SPEECH
	OTP
	SIMACTIVE
	KYCVALIDATE
)

func (s ServiceType) String() string {
	switch s {
	case TEXT:
		return "TEXT"
	case FLASH:
		return "FLASH"
	case CALL:
		return "CALL"
	case RECORD:
		return "RECORD"
	case IVR:
		return "IVR"
	case TEXT2SPEECH:
		return "TEXT2SPEECH"
	case OTP:
		return "OTP"
	case SIMACTIVE:
		return "SIMACTIVE"
	case KYCVALIDATE:
		return "KYCVALIDATE"
	}

	return "TEXT"
}

type Status int8

const (
	DELETED Status = iota
	UNKNOWN
	EXPIRED
	WAITING
	FAILED
	REJECTD
	UNDELIV
	ACCEPTD
	DELIVRD
	BUSY
	ANSWERED
)

// DELETED | UNKNOWN | EXPIRED | WAITING | FAILED | REJECTD | UNDELIV | ACCEPTD | DELIVRD | BUSY | ANSWERED
func (s Status) String() string {
	switch s {
	case DELETED:
		return "DELETED"
	case UNKNOWN:
		return "UNKNOWN"
	case EXPIRED:
		return "EXPIRED"
	case WAITING:
		return "WAITING"
	case FAILED:
		return "FAILED"
	case REJECTD:
		return "REJECTD"
	case UNDELIV:
		return "UNDELIV"
	case ACCEPTD:
		return "ACCEPTD"
	case DELIVRD:
		return "DELIVRD"
	case BUSY:
		return "BUSY"
	case ANSWERED:
		return "ANSWERED"
	}

	return "UNKNOWN"
}

type Service int8

// SMS | USSD | VOICE | SIMACTIVE | KYCVERIFY
const (
	SMS Service = iota
	USSD
	VOICE
	KYCVERIFY
)

func (s Service) String() string {
	switch s {
	case SMS:
		return "SMS"
	case USSD:
		return "USSD"
	case VOICE:
		return "VOICE"
	case Service(SIMACTIVE):
		return "SIMACTIVE"
	case KYCVERIFY:
		return "KYCVERIFY"
	}

	return "SMS"
}
