package stun

import (
	"github.com/pkg/errors"
	"gitlab.com/pions/pion/pkg/go/stun"
)

type MessageBuilder struct {
	attributes []*Attribute
}

func Build(class MessageClass, method Method, transactionID []byte, attrs ...Attribute) (*Message, error) {

	m := &Message{
		Class:         class,
		Method:        method,
		TransactionID: transactionID,
	}

	m.Raw = make([]byte, messageHeaderLength)

	setMessageType(m.Raw[messageHeaderStart:2], class, method)

	copy(m.Raw[transactionIDStart:], m.TransactionID)

	attrs = append(attrs, &stun.Software{
		Software: "Pion",
	})
	for _, v := range attrs {
		err := v.Pack(m)
		if err != nil {
			return nil, errors.Wrap(err, "failed packing attribute")
		}
	}

	return m, nil
}
