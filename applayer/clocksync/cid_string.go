// Code generated by "stringer -type=CID"; DO NOT EDIT.

package clocksync

import "strconv"

const _CID_name = "PackageVersionReqAppTimeReqDeviceAppTimePeriodicityReqForceDeviceResyncReq"

var _CID_index = [...]uint8{0, 17, 27, 54, 74}

func (i CID) String() string {
	if i >= CID(len(_CID_index)-1) {
		return "CID(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _CID_name[_CID_index[i]:_CID_index[i+1]]
}