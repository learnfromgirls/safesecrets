# safesecrets
interfaces for deriving secrets from a master password and setting those secrets on components (write only)

This reduces the risk of an attacker finding your secret on the target machine. The secret only ever exists in
go application memory , never on disk and never in the environment.
It wont even exist in memory until derived via https POST by an external agent doing a bootstrap "vault" user login.

Use the secret to decrypt your API keys, verify user passwords etc.

## Usage
```go
import (
	"fmt"
	"github.com/learnfromgirls/safesecrets"
	"github.com/learnfromgirls/argon2-go-withsecret" //or use github.com/learnfromgirls/safesecrets/example
)

type MySecretSetter struct {
	secret [] byte
}

func (dss *MySecretSetter) SetSecret(secretIn [] byte) () {
	dss.secret = secretIn
}

func ExampleRun() () {
	dsd := argon2_go_withsecret.NewVaultContext() // or use &example.DefaultSecretDeriver{}
	dss := &MySecretSetter{}

	//the following password should have come from an https POST to avoid it being anywhere on the machine
	dsd.SetSecrets([] byte("password"), [] byte("salt"), dss) //massive slow hash computation happens

	fmt.Printf("secret now set to %v\n", dss.secret) //look no public reads! and a very unguessable secret

}
```

