package safesecrets

type SecretSetter interface {
	SetSecret([] byte) ()
}

type SecretDeriver interface {
	Hash(password []byte, salt []byte) (hash []byte, err error)
	SetSecrets(password []byte, initialsalt []byte, ssa ...SecretSetter) (err error)
}

