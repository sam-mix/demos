package msg

import (
	"bytes"
	"encoding/binary"
)

func (mh *ClientMsgHead) MsgPackBE() ([]byte, error) {
	var buf bytes.Buffer
	err := binary.Write(&buf, binary.BigEndian, *mh)
	return buf.Bytes(), err
}

func (mh *ClientMsgHead) UnMsgPackBE(DataBuf []byte) error {

	bufMsgLen := bytes.NewBuffer(DataBuf[:4])
	bufMsgID := bytes.NewBuffer(DataBuf[4:6])
	bufSvrID := bytes.NewBuffer(DataBuf[6:8])
	bufCCode := bytes.NewBuffer(DataBuf[8:9])

	var err error

	err = binary.Read(bufMsgLen, binary.BigEndian, &mh.MsgLen)
	if err != nil {
		return err
	}

	err = binary.Read(bufMsgID, binary.BigEndian, &mh.MsgID)
	if err != nil {
		return err
	}

	err = binary.Read(bufSvrID, binary.BigEndian, &mh.SvrID)
	if err != nil {
		return err
	}

	err = binary.Read(bufCCode, binary.BigEndian, &mh.CCode)
	if err != nil {
		return err
	}

	return err
}
