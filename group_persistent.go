package tox

/*
#include <stdint.h>
#include <stdlib.h>
#include <string.h>
#include <tox/tox.h>

*/
import "C"
import "encoding/hex"

const CONFERENCE_UID_SIZE = 32

func (this *Tox) ConferenceUidSize() uint32 {
	return uint32(C.tox_conference_uid_size(this.toxcore))
}

func (this *Tox) ConferenceEnter(conferenceNumber uint32) error {
	var TOX_ERR_CONFERENCE_ENTER err
	ok := C.tox_conference_enter(this.toxcore, C.uint32_t(conferenceNumber), &err)
	if ok == 0 {
		return toxerr(err)
	}
	return nil
}

func (this *Tox) ConferenceLeave(conferenceNumber uint32, keepLeave bool) error {
	keepLeavec := 0
	if keepLeave {
		keepLeavec = 1
	}
	var TOX_ERR_CONFERENCE_LEAVE err
	ok := C.tox_conference_leave(this.toxcore, C.uint32_t(conferenceNumber), C.int(keepLeavec), &err)
	if ok == 0 {
		return toxerr(err)
	}
	return nil
}

func (this *Tox) ConferenceGetUid(conferenceNumber uint32) (string, error) {
	uid := make([]byte, CONFERENCE_UID_SIZE)
	ok := C.tox_conference_get_uid(this.toxcore, C.uint32_t(conferenceNumber), (*C.uint8_t)(&uid[0]))
	if ok == 0 {
		return "", toxerr(-1)
	}
	return hex.EncodeToString(uid), nil
}

func (this *Tox) ConferenceGetByUid(uid string) (uint32, error) {
	uidbin, err := hex.DecodeString(uid)
	if err != nil {
		return 0, err
	}

	var TOX_ERR_CONFERENCE_BY_UID error
	cnum := C.tox_conference_by_uid(this.toxcore, (*C.uint8_t)(&uidbin[0]), &error)
	if error != 0 {
		return 0, toxerr(error)
	}
	return uint32(cnum), nil
}
