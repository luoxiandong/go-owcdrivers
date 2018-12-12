package ontologyTransaction

const AddressPrefix = byte(0x17)

const (
	AssetONT         = 0
	AssetONG         = 1
	AssetONGWithdraw = 2
)

const (
	DefaultGasPrice = uint64(500)
	DefaultGasLimit = uint64(20000)
)

const (
	ONTContractVersion = byte(0x00)
	ONGContractVersion = byte(0x00)
	ONTContractAddress = "0100000000000000000000000000000000000000"
	ONGContractAddress = "0200000000000000000000000000000000000000"
	TxTypeInvoke       = byte(0xD1)
	DefaultAttribute   = byte(0)
)

const (
	MethodTransfer     = "transfer"
	MethodTransferFrom = "transferFrom"
	NativeInvokeName   = "Ontology.Native.Invoke"
)

const (
	OpCodeNewStruct       = byte(0xC6)
	OpCodeToALTStack      = byte(0x6B)
	OpCodeDupFromALTStack = byte(0x6A)
	OpCodeAppend          = byte(0xC8)
	OpCodePush0           = byte(0x00)
	OpCodePush1           = byte(0x51)
	OpCodeFromALTStack    = byte(0x6C)
	OpCodePack            = byte(0xC1)
	OpCodeSysCall         = byte(0x68)
	OpCodeCheckSig        = byte(0xAC)
	OpCodeCheckMultiSig   = byte(0xAE)
)

var (
	CurveOrder     = []byte{0xFF, 0xFF, 0xFF, 0xFF, 0x00, 0x00, 0x00, 0x00, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xBC, 0xE6, 0xFA, 0xAD, 0xA7, 0x17, 0x9E, 0x84, 0xF3, 0xB9, 0xCA, 0xC2, 0xFC, 0x63, 0x25, 0x51}
	HalfCurveOrder = []byte{0x7F, 0xFF, 0xFF, 0xFF, 0x80, 0x00, 0x00, 0x00, 0x7F, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xDE, 0x73, 0x7D, 0x56, 0xD3, 0x8B, 0xCF, 0x42, 0x79, 0xDC, 0xE5, 0x61, 0x7E, 0x31, 0x92, 0xA8}
)
