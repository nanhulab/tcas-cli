package csv

import (
	"bytes"
	"crypto/rand"
	"encoding/base64"
	"encoding/binary"
	"encoding/hex"
	"fmt"
	"github.com/tjfoc/gmsm/sm3"
	"github.com/vtolstov/go-ioctl"
	"golang.org/x/sys/unix"
	"os"
	"unsafe"
)

// https://gitee.com/anolis/hygon-devkit/blob/master/csv/attestation/csv_status.h
const GUEST_ATTESTATION_NONCE_SIZE = 16
const GUEST_ATTESTATION_DATA_SIZE = 64
const VM_ID_SIZE = 16
const VM_VERSION_SIZE = 16
const SN_LEN = 64
const USER_DATA_SIZE = 64
const HASH_BLOCK_LEN = 32
const CSV_CERT_RSVD3_SIZE = 624
const CSV_CERT_RSVD4_SIZE = 368
const CSV_CERT_RSVD5_SIZE = 368
const HIGON_USER_ID_SIZE = 256
const SIZE_INT32 = 4
const ECC_POINT_SIZE = 72
const PAGE_SIZE = 1 << 12

type HigonCsvCert struct {
	Version     uint32
	APIMajor    byte
	APIMinor    byte
	Reserved1   byte
	Reserved2   byte
	PubkeyUsage uint32
	PubkeyAlgo  uint32
	Pubkey      [SIZE_INT32 + ECC_POINT_SIZE*2 + HIGON_USER_ID_SIZE]byte
	Reserved3   [CSV_CERT_RSVD3_SIZE]byte
	Sig1Usage   uint32
	Sig1Algo    uint32
	Sig1        [ECC_POINT_SIZE * 2]byte
	Reserved4   [CSV_CERT_RSVD4_SIZE]byte
	Sig2Usage   uint32
	Sig2Algo    uint32
	Sig2        [ECC_POINT_SIZE * 2]byte
	Reserved5   [CSV_CERT_RSVD5_SIZE]byte
}

type CSV_CERT_t = HigonCsvCert
type HashBlockT [HASH_BLOCK_LEN]byte
type HashBlockU [HASH_BLOCK_LEN]byte
type ecc_signature_t [ECC_POINT_SIZE]byte

type CSVAttestationReport struct {
	UserPubkeyDigest HashBlockT
	VmId             [VM_ID_SIZE]byte
	VmVersion        [VM_VERSION_SIZE]byte
	UserData         [USER_DATA_SIZE]byte
	Mnonce           [GUEST_ATTESTATION_NONCE_SIZE]byte
	Measure          HashBlockT
	Policy           uint32
	SigUsage         uint32
	SigAlgo          uint32
	Anonce           uint32
	Sig1             [ECC_POINT_SIZE * 2 / SIZE_INT32]uint32
	PekCert          CSV_CERT_t
	Sn               [SN_LEN]byte
	Reserved2        [32]byte
	Mac              HashBlockU
}

type ReportDetailInfo struct {
	UserData   string `json:"userData"`
	Monce      string `json:"monce"`
	Measure    string `json:"measure"`
	VMId       string `json:"vmId"`
	VMVersion  string `json:"vmVersion"`
	ChipId     string `json:"chipId"`
	FullReport string `json:"fullReport"`
}

type CSVAttestationUserData struct {
	Data   [GUEST_ATTESTATION_DATA_SIZE]byte
	Mnonce [GUEST_ATTESTATION_NONCE_SIZE]byte
	Hash   HashBlockU
}

type CsvGuestMem struct {
	Va   uintptr
	Size int32
}

func genRandomBytes(buf []byte, len uint32) {
	rand.Read(buf[:len])
}

func computeSM3Hash(data []byte) [HASH_BLOCK_LEN]byte {
	hash := sm3.New()
	hash.Write(data)
	var result [HASH_BLOCK_LEN]byte
	copy(result[:], hash.Sum(nil))
	return result
}

func GetReportInByte(userData []byte) (report []byte, err error) {
	attestationReport, err := GetCSVAttestationReport(userData)
	if err != nil {
		return nil, fmt.Errorf("get csv attestation report failed, error: %s", err)
	}

	attestationReport.Reserved2 = [32]byte{}

	report, err = MarshalCsvAttestationReport(attestationReport)
	if err != nil {
		return nil, fmt.Errorf("marshal csv attestation report failed, error: %s", err)
	}

	return report, err
}

func GetCSVAttestationReport(userData []byte) (*CSVAttestationReport, error) {
	if len(userData) > GUEST_ATTESTATION_DATA_SIZE {
		return nil, fmt.Errorf("user data size is too large, limit to %d\n", GUEST_ATTESTATION_DATA_SIZE)
	}

	user_data_len := PAGE_SIZE
	var reportData [PAGE_SIZE]byte
	copy(reportData[:GUEST_ATTESTATION_DATA_SIZE], userData)
	genRandomBytes(reportData[GUEST_ATTESTATION_DATA_SIZE:GUEST_ATTESTATION_DATA_SIZE+GUEST_ATTESTATION_NONCE_SIZE], GUEST_ATTESTATION_NONCE_SIZE)

	hashData := append(reportData[:GUEST_ATTESTATION_DATA_SIZE], reportData[GUEST_ATTESTATION_DATA_SIZE:GUEST_ATTESTATION_DATA_SIZE+GUEST_ATTESTATION_NONCE_SIZE]...)
	sm3Hash := computeSM3Hash(hashData)
	copy(reportData[GUEST_ATTESTATION_DATA_SIZE+GUEST_ATTESTATION_NONCE_SIZE:], sm3Hash[:])

	fd, err := os.OpenFile("/dev/csv-guest", os.O_RDWR, 0)
	if err != nil {
		return nil, fmt.Errorf("open /dev/csv-guest failed: %s", err)
	}
	defer fd.Close()

	mem := CsvGuestMem{
		Va:   uintptr(unsafe.Pointer(&reportData)),
		Size: int32(user_data_len),
	}

	//https://gitee.com/anolis/hygon-devkit/blob/master/csv/attestation/csv_sdk/ioctl_get_attestation_report.c#L20
	GET_ATTESTATION_REPORT := ioctl.IOWR('D', 1, 16)
	_, _, errno := unix.Syscall(unix.SYS_IOCTL, fd.Fd(), GET_ATTESTATION_REPORT, uintptr(unsafe.Pointer(&mem)))
	if errno != 0 {
		return nil, fmt.Errorf("ioctl GET_ATTESTATION_REPORT failed: %d\n", errno)
	}
	report, err := UnmarshalCsvAttestationReport(reportData[:])
	if err != nil {
		return nil, fmt.Errorf("unmarshal csv attestation report failed, error: %s", err)
	}

	return report, nil
}

func GetReportDetailInfo(report []byte) (*ReportDetailInfo, error) {
	rdi := new(ReportDetailInfo)

	rdi.FullReport = base64.StdEncoding.EncodeToString(report)
	d, err := UnmarshalCsvAttestationReport(report)
	if err != nil {
		return nil, fmt.Errorf("unmarshal csv attetation report failed, error: %s", err)
	}
	var UserData [64]uint8
	j := unsafe.Sizeof(d.UserData) / unsafe.Sizeof(uint32(0))
	for i := 0; i < int(j); i++ {
		tmp := (*uint32)(unsafe.Pointer(&d.UserData[i*4]))
		*tmp ^= d.Anonce
		copy(UserData[i*4:], (*[4]uint8)(unsafe.Pointer(tmp))[:])
	}
	rdi.UserData = string(bytes.TrimRight(UserData[:], "\x00"))

	var measure [32]uint8
	j = unsafe.Sizeof(d.Measure) / unsafe.Sizeof(uint32(0))
	for i := 0; i < int(j); i++ {
		tmp := (*uint32)(unsafe.Pointer(&d.Measure[i*4]))
		*tmp ^= d.Anonce
		copy(measure[i*4:], (*[4]uint8)(unsafe.Pointer(tmp))[:])
	}
	rdi.Measure = hex.EncodeToString(measure[:])

	var mnonce [16]uint8
	j = unsafe.Sizeof(d.Mnonce) / unsafe.Sizeof(uint32(0))
	for i := 0; i < int(j); i++ {
		tmp := (*uint32)(unsafe.Pointer(&d.Mnonce[i*4]))
		*tmp ^= d.Anonce
		copy(mnonce[i*4:], (*[4]uint8)(unsafe.Pointer(tmp))[:])
	}
	rdi.Monce = hex.EncodeToString(mnonce[:])

	var vmid [16]uint8
	j = unsafe.Sizeof(d.VmId) / unsafe.Sizeof(uint32(0))
	for i := 0; i < int(j); i++ {
		tmp := (*uint32)(unsafe.Pointer(&d.VmId[i*4]))
		*tmp ^= d.Anonce
		copy(vmid[i*4:], (*[4]uint8)(unsafe.Pointer(tmp))[:])
	}
	rdi.VMId = hex.EncodeToString(vmid[:])

	var vmversion [16]uint8
	j = unsafe.Sizeof(d.VmVersion) / unsafe.Sizeof(uint32(0))
	for i := 0; i < int(j); i++ {
		tmp := (*uint32)(unsafe.Pointer(&d.VmVersion[i*4]))
		*tmp ^= d.Anonce
		copy(vmversion[i*4:], (*[4]uint8)(unsafe.Pointer(tmp))[:])
	}
	rdi.VMVersion = hex.EncodeToString(vmversion[:])

	var chipID [64]uint8
	j = (uintptr(unsafe.Pointer(&d.Reserved2)) - uintptr(unsafe.Pointer(&d.Sn))) / uintptr(unsafe.Sizeof(uint32(0)))
	for i := 0; i < int(j); i++ {
		chipID32 := (*uint32)(unsafe.Pointer(&d.Sn[i*4]))
		*chipID32 ^= d.Anonce
		copy(chipID[i*4:], (*[4]uint8)(unsafe.Pointer(chipID32))[:])
	}
	rdi.ChipId = string(bytes.TrimRight(chipID[:], "\x00"))

	return rdi, nil
}

func GetSealingKey() (sealingkey string, err error) {
	report, err := GetCSVAttestationReport([]byte("get-sealing-key"))
	if err != nil {
		return "", fmt.Errorf("get sealing key failed, error: %s", err)
	}

	return hex.EncodeToString(report.Reserved2[:]), nil
}

// 将结构体编码为二进制数据
func MarshalCsvAttestationReport(d *CSVAttestationReport) ([]byte, error) {
	buf := new(bytes.Buffer)
	err := binary.Write(buf, binary.LittleEndian, d)
	if err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

func UnmarshalCsvAttestationReport(report []byte) (*CSVAttestationReport, error) {
	buf := bytes.NewReader(report)
	d := new(CSVAttestationReport)
	err := binary.Read(buf, binary.LittleEndian, d)
	if err != nil {
		return nil, err
	}
	return d, nil
}
