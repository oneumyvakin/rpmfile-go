package rpmfile

import (
	"bytes"
	"encoding/binary"
	"encoding/hex"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"errors"
	"reflect"

	"rpmfile/common"
	"rpmfile/rpmheader"
	"rpmfile/rpmsignature"
)

type Rpm_file struct {
	Logger *log.Logger

	rpm_file  *os.File
	Signature rpmsignature.Rpm_signature
	Header    rpmheader.Rpm_header
}

type rpm_data_structure interface {
	Get_tags() map[int]string
	Get_field(field string) reflect.Value

	Get_number_of_entries() int32
	Get_entries() []common.Entry_header
	Get_header_structure_size() int32
	Get_header_start() int64
	Get_header_end() int64
	Get_store() []byte

	Set_number_of_entries(int32)
	Set_entries([]common.Entry_header)
	Set_header_structure_size(int32)
	Set_header_start(int64)
	Set_header_end(int64)
	Set_store([]byte)
}

func (rpm *Rpm_file) extract_data(ds rpm_data_structure) {

	tag := ds.Get_tags()

	for _, entry := range ds.Get_entries() {
		var field_name string
		var ok bool

		if field_name, ok = tag[int(entry.Tag)]; !ok {
			rpm.Logger.Printf("entry.Tag: %#v, type: %d, count %d not found in tag list\n", entry.Tag, entry.Ty, entry.Count)
			continue
		}

		data_field := ds.Get_field(field_name)

		if data_field == (reflect.Value{}) {
			rpm.Logger.Printf("Rpm_header field not found id %d, type: %d name: %s\n", entry.Tag, entry.Ty, tag[int(entry.Tag)])
			continue
		}

		if !data_field.CanSet() {
			rpm.Logger.Println("Field is not settable: ", tag[int(entry.Tag)])
			continue
		}

		rpm.Logger.Printf("Extract data from %v to field type %s: ", entry, data_field.Type())

		switch {
		case entry.Ty == 2:
			value, _ := rpm.extract_int8(ds.Get_store(), entry.Offset, entry.Count)

			if data_field.Type() != reflect.TypeOf(value) {
				rpm.Logger.Printf("Type mismatch between field and extracted data %s != %s. Doing nothing.\n", data_field.Type(), reflect.TypeOf(value))
				continue
			}

			data_field.Set(reflect.ValueOf(value))

		case entry.Ty == 3:
			value, _ := rpm.extract_int16(ds.Get_store(), entry.Offset, entry.Count)

			if data_field.Type() != reflect.TypeOf(value) {
				rpm.Logger.Printf("Type mismatch between field and extracted data %s != %s. Doing nothing.\n", data_field.Type(), reflect.TypeOf(value))
				continue
			}

			data_field.Set(reflect.ValueOf(value))

		case entry.Ty == 4:
			value, _ := rpm.extract_int32(ds.Get_store(), entry.Offset, entry.Count)

			if data_field.Type() != reflect.TypeOf(value) {
				rpm.Logger.Printf("Type mismatch between field and extracted data %s != %s. Doing nothing.\n", data_field.Type(), reflect.TypeOf(value))
				continue
			}

			data_field.Set(reflect.ValueOf(value))

		case entry.Ty == 5:
			value, _ := rpm.extract_int64(ds.Get_store(), entry.Offset, entry.Count)

			if data_field.Type() != reflect.TypeOf(value) {
				rpm.Logger.Printf("Type mismatch between field and extracted data %s != %s. Doing nothing.\n", data_field.Type(), reflect.TypeOf(value))
				continue
			}

			data_field.Set(reflect.ValueOf(value))

		case entry.Ty == 6 || entry.Ty == 9:
			value, _ := rpm.extract_string(ds.Get_store(), entry.Offset, entry.Count)

			if data_field.Type() != reflect.TypeOf(value) {
				rpm.Logger.Printf("Type mismatch between field and extracted data %s != %s. Doing nothing.\n", data_field.Type(), reflect.TypeOf(value))
				continue
			}

			data_field.SetString(value)

		case entry.Ty == 7:
			value, _ := rpm.extract_bin(ds.Get_store(), entry.Offset, entry.Count)

			if data_field.Type() != reflect.TypeOf(value) {
				rpm.Logger.Printf("Type mismatch between field and extracted data %s != %s. Doing nothing.\n", data_field.Type(), reflect.TypeOf(value))
				continue
			}

			data_field.SetBytes(value)

		case entry.Ty == 8:
			value, _ := rpm.extract_array(ds.Get_store(), entry.Offset, entry.Count)

			if data_field.Type() != reflect.TypeOf(value) {
				rpm.Logger.Printf("Type mismatch between field and extracted data %s != %s. Doing nothing.\n", data_field.Type(), reflect.TypeOf(value))
				continue
			}

			data_field.Set(reflect.ValueOf(value))

		default:
			rpm.Logger.Println("Data type %s is not supported. Doing nothing.")
		}
		rpm.Logger.Printf("Done\n")
	}

}

func (rpm *Rpm_file) extract_int8(store []byte, offset int32, count int32) (values int16, err error) {
	// for Ty = 2

	buf := bytes.NewReader(store[offset : offset+1*count])
	err = binary.Read(buf, binary.BigEndian, &values)
	if err != nil {
		rpm.Logger.Println("extract_int16: binary.Read failed:", err)
		return
	}

	return
}

func (rpm *Rpm_file) extract_int16(store []byte, offset int32, count int32) (values int16, err error) {
	// for Ty = 3

	buf := bytes.NewReader(store[offset : offset+2*count])
	err = binary.Read(buf, binary.BigEndian, &values)
	if err != nil {
		rpm.Logger.Println("extract_int16: binary.Read failed:", err)
		return
	}

	return
}

func (rpm *Rpm_file) extract_int32(store []byte, offset int32, count int32) (values int32, err error) {
	// for Ty = 4

	buf := bytes.NewReader(store[offset : offset+4*count])
	err = binary.Read(buf, binary.BigEndian, &values)
	if err != nil {
		rpm.Logger.Println("extract_int32: binary.Read failed:", err)
		return
	}

	return
}

func (rpm *Rpm_file) extract_int64(store []byte, offset int32, count int32) (values int64, err error) {
	// for Ty = 5

	buf := bytes.NewReader(store[offset : offset+8*count])
	err = binary.Read(buf, binary.BigEndian, &values)
	if err != nil {
		rpm.Logger.Println("extract_int32: binary.Read failed:", err)
		return
	}

	return
}

func (rpm *Rpm_file) extract_string(store []byte, offset int32, count int32) (values string, err error) {
	// for Ty = 6 and 9
	if count > 1 {
		err = errors.New(fmt.Sprintf("extract_string failed: count = %s, count > 1", count))
		rpm.Logger.Println(err)
		return
	}
	//null_byte, err := hex.DecodeString("00")
	idx := bytes.Index(store[offset:], []byte{0})

	values = string(store[offset : offset+int32(idx)])

	return

}

func (rpm *Rpm_file) extract_bin(store []byte, offset int32, count int32) (values []byte, err error) {
	// for Ty = 7
	defer func() {
		if err := recover(); err != nil {
			rpm.Logger.Println("extract_bin failed:", err)
		}
	}()

	values = store[offset : offset+count]

	return
}

func (rpm *Rpm_file) extract_array(store []byte, offset int32, count int32) (values []string, err error) {
	// for Ty = 8
	//null_byte, err := hex.DecodeString("00")
	array_of_array_of_bytes := bytes.SplitN(store[offset:], []byte{0}, int(count)+1) // +1 to keep reminder in last item

	for i, byte_array := range array_of_array_of_bytes {
		if i >= int(count) {
			continue // skip last item
		}
		values = append(values, string(byte_array))
	}

	return

}

func (rpm *Rpm_file) read_entries_headers(ds rpm_data_structure) (err error) {

	var entries []common.Entry_header

	for n := 1; n <= int(ds.Get_number_of_entries()); n += 1 {
		//fmt.Printf("Number is: %d <= %d\n", n, int(number_of_entries))

		header_bytes := make([]byte, 16)
		rpm.rpm_file.Read(header_bytes)
		buf := bytes.NewReader(header_bytes)

		var header common.Entry_header
		err = binary.Read(buf, binary.BigEndian, &header)
		if err != nil && err != io.EOF {
			rpm.Logger.Println("read_entries_headers: binary.Read failed:", err)
			return
		}
		if err == io.EOF {
			break
		}

		entries = append(entries, header)
	}

	ds.Set_entries(entries)

	rpm.Logger.Printf("%#v\n", ds.Get_entries())

	return
}

func (rpm *Rpm_file) read_header_structure_size(ds rpm_data_structure) (err error) {

	var header_structure_size int32
	header_structure_size_bytes := make([]byte, 4)
	rpm.rpm_file.Read(header_structure_size_bytes)
	buf := bytes.NewReader(header_structure_size_bytes)

	err = binary.Read(buf, binary.BigEndian, &header_structure_size)
	if err != nil {
		rpm.Logger.Println("read_header_structure_size: binary.Read failed:", err)
		return
	}

	ds.Set_header_structure_size(header_structure_size)
	rpm.Logger.Printf("Header structure size: %d\n", ds.Get_header_structure_size())
	return
}

func (rpm *Rpm_file) read_number_of_entries(ds rpm_data_structure) (err error) {

	var number_of_entries int32
	num_entries_bytes := make([]byte, 4)
	rpm.rpm_file.Read(num_entries_bytes)
	buf := bytes.NewReader(num_entries_bytes)

	err = binary.Read(buf, binary.BigEndian, &number_of_entries)
	if err != nil {
		rpm.Logger.Println("read_number_of_entries: binary.Read failed:", err)
		return
	}

	ds.Set_number_of_entries(number_of_entries)
	rpm.Logger.Printf("Number of entries: %d\n", ds.Get_number_of_entries())
	return
}

func (rpm *Rpm_file) read_store(ds rpm_data_structure) (err error) {

	store := make([]byte, ds.Get_header_structure_size())
	_, err = rpm.rpm_file.Read(store)

	ds.Set_store(store)

	rpm.Logger.Printf("Store size is: %d\n", len(ds.Get_store()))
	return
}

func (rpm *Rpm_file) read_header(ds rpm_data_structure) (err error) {
	magic_byte, err := hex.DecodeString("8e")
	if err != nil {
		panic(err)
	}

	current_byte := make([]byte, 1)

	_, err_file_read := rpm.rpm_file.Read(current_byte)

	if err_file_read != nil && err_file_read != io.EOF {
		panic(err)
	}
	if err_file_read == io.EOF {
		return
	}

	current_byte_as_string := hex.EncodeToString(current_byte)
	for hex.EncodeToString(current_byte) != "8e" {
		rpm.Logger.Println(current_byte_as_string + " != 8e")

		_, err1 := rpm.rpm_file.Read(current_byte)

		if err1 != nil && err1 != io.EOF {
			panic(err)
		}
		if err1 == io.EOF {
			return
		}
	}

	two_bytes := make([]byte, 2)
	rpm.rpm_file.Read(two_bytes)

	var buffer2 bytes.Buffer
	buffer2.Write(magic_byte)
	buffer2.Write(two_bytes)
	test := buffer2.String()

	if hex.EncodeToString(buffer2.Bytes()) != "8eade8" {
		rpm.Logger.Println(test + " != 8eade8")
		return
	}

	version_byte := make([]byte, 1)
	rpm.rpm_file.Read(version_byte)
	//version := string(version_byte)
	rpm.Logger.Printf("Success version is: %d\n", version_byte)

	header_start, _ := rpm.rpm_file.Seek(0, os.SEEK_CUR)
	ds.Set_header_start(header_start - 4) // -4 for magic
	rpm.Logger.Printf("Header start: %d\n", ds.Get_header_start())

	magic_header := make([]byte, 4)
	rpm.rpm_file.Read(magic_header)

	err = rpm.read_number_of_entries(ds)
	if err != nil {
		panic(err)
	}

	err = rpm.read_header_structure_size(ds)
	if err != nil {
		panic(err)
	}

	err = rpm.read_entries_headers(ds)
	if err != nil {
		panic(err)
	}

	err = rpm.read_store(ds)
	if err != nil {
		panic(err)
	}

	rpm.extract_data(ds)

	return
}

func (rpm_file *Rpm_file) read_headers() {

	magic_offset := 96
	rpm_file.rpm_file.Seek(int64(magic_offset), os.SEEK_CUR)

	rpm_file.read_header(&rpm_file.Signature)
	rpm_file.read_header(&rpm_file.Header)

}

func (rpm *Rpm_file) Open(file_path string) {
	var err error
	// open input file
	rpm.rpm_file, err = os.Open(file_path)
	if err != nil {
		panic(err)
	}
	// close fi on exit and check for its returned error
	defer func() {
		if err := rpm.rpm_file.Close(); err != nil {
			panic(err)
		}
	}()

	if rpm.Logger == nil {
		rpm.Logger = log.New(ioutil.Discard, "RPM: ", log.Ldate|log.Ltime|log.Lshortfile)
	}

	rpm.read_headers()
}

func (rpm *Rpm_file) Set_debug() {
	rpm.Logger = log.New(os.Stdout, "RPM: ", log.Ldate|log.Ltime|log.Lshortfile)
}