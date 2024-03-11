package infra

import "math/rand"

type IdGenerator struct {
	charset string
}

func NewIdGenerator() *IdGenerator {
	return &IdGenerator{
		charset: "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZ",
	}
}

func (idg *IdGenerator) GenerateVerificationCode(length int) string {
	b := make([]byte, length)
	for i := range b {
		b[i] = idg.charset[rand.Intn(len(idg.charset))]
	}
	return string(b)
}
