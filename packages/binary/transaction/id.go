package transaction

type Id [IdLength]byte

var EmptyId = Id{}

const IdLength = 64