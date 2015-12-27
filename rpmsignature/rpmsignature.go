package rpmsignature

import (
	"reflect"
	"github.com/oneumyvakin/rpmfile-go/common"
)

type Rpm_signature struct {
	header_start          int64
	header_end            int64
	number_of_entries     int32
	header_structure_size int32
	entries               []common.Entry_header
	store                 []byte

	Headersignatures []byte
	Badsha1_1        string
	Badsha1_2        string
	Dsa              string
	Rsa              string
	Sha1             string
	Size             int32
	Lemd5_1          string
	Pgp              []byte
	Lemd5_2          string
	Md5              []byte
	Gpg              []byte
	Pgp5             []byte
	Payloadsize      int32
}

func (rpm_signature *Rpm_signature) Get_tags() map[int]string {
	return map[int]string{
		62:   "Headersignatures",
		264:  "Badsha1_1",
		265:  "Badsha1_2",
		267:  "Dsa",
		268:  "Rsa",
		269:  "Sha1",
		1000: "Size",
		1001: "Lemd5_1",
		1002: "Pgp",
		1003: "Lemd5_2",
		1004: "Md5",
		1005: "Gpg",
		1006: "Pgp5",
		1007: "Payloadsize",
	}
}

func (rpm_signature *Rpm_signature) Get_field(name string) (field reflect.Value) {
	reflect_of_rpm_header := reflect.ValueOf(rpm_signature).Elem()
	field = reflect_of_rpm_header.FieldByName(name)

	return
}

func (r Rpm_signature) Get_number_of_entries() int32 {
	return r.number_of_entries
}
func (r Rpm_signature) Get_entries() []common.Entry_header {
	return r.entries
}
func (r Rpm_signature) Get_header_structure_size() int32 {
	return r.header_structure_size
}
func (r Rpm_signature) Get_header_start() int64 {
	return r.header_start
}
func (r Rpm_signature) Get_header_end() int64 {
	return r.header_end
}
func (r Rpm_signature) Get_store() []byte {
	return r.store
}

func (r *Rpm_signature) Set_number_of_entries(number_of_entries int32) {
	r.number_of_entries = number_of_entries
}
func (r *Rpm_signature) Set_entries(entries []common.Entry_header) {
	r.entries = entries
}
func (r *Rpm_signature) Set_header_structure_size(header_structure_size int32) {
	r.header_structure_size = header_structure_size
}
func (r *Rpm_signature) Set_header_start(header_start int64) {
	r.header_start = header_start
}
func (r *Rpm_signature) Set_header_end(header_end int64) {
	r.header_end = header_end
}
func (r *Rpm_signature) Set_store(store []byte) {
	r.store = store
}
