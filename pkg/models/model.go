package models

const NoneModelDiscriminator ModelDiscriminator = iota
const NoneMessageDiscriminator MessageDiscriminator = iota

type ModelDiscriminator int

type Model interface {
	ModelDiscriminator() ModelDiscriminator
}

type NoneModel struct{}

func (m *NoneModel) ModelDiscriminator() ModelDiscriminator {
	return NoneModelDiscriminator
}

type MessageDiscriminator int

type Message interface {
	MessageDiscriminator() MessageDiscriminator
}

type NoneMessage struct{}

func (m *NoneMessage) MessageDiscriminator() MessageDiscriminator {
	return NoneMessageDiscriminator
}
