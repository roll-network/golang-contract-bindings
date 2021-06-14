# golang-contract-bindings

## Pre steps

Install solc

```
brew update
brew tap ethereum/ethereum
brew install solidity@7
```

Install abigen

```
go get -u github.com/ethereum/go-ethereum
cd $GOPATH/pkg/mod/github.com/ethereum/go-ethereum@v1.10.1
make
make devtools
```

## Social Money binding

Fetch latest roll-smart-contracts and deploy on localhost to generate the abi.
If you want any other network just update the following script.

```
cp ../roll-smart-contracts/deploy/localhost/Manager.abi          ./social-money/abi/Manager.abi
cp ../roll-smart-contracts/deploy/localhost/Registry.abi         ./social-money/abi/Registry.abi
cp ../roll-smart-contracts/deploy/localhost/TokenFactory.abi     ./social-money/abi/TokenFactory.abi
cp ../roll-smart-contracts/deploy/localhost/TokenVesting.abi     ./social-money/abi/TokenVesting.abi
cp ../roll-smart-contracts/deploy/localhost/SocialMoney.abi      ./social-money/abi/SocialMoney.abi

```

```
abigen --abi=./social-money/abi/Manager.abi           --pkg=contracts --out=./social-money/Manager.go
abigen --abi=./social-money/abi/Registry.abi          --pkg=contracts --out=./social-money/Registry.go
abigen --abi=./social-money/abi/TokenFactory.abi      --pkg=contracts --out=./social-money/TokenFactory.go
abigen --abi=./social-money/abi/TokenVesting.abi      --pkg=contracts --out=./social-money/TokenVesting.go
abigen --abi=./social-money/abi/SocialMoney.abi       --pkg=contracts --out=./social-money/SocialMoney.go
```

## Import using private repo

First configure ssh adding

`[url "git@github.com:"] insteadOf = https://github.com/`

in ~/.gitconfig

Then check is working ssh

```
> ssh -T git@github.com
> Hi andresaiello! You've successfully authenticated, but GitHub does not provide shell access.
```

And now configure go to read private repos

```
> go env -w GOPRIVATE=github.com/TuringAdvisoryGroup/golang-contract-bindings
> git ls-remote -q https://github.com/TuringAdvisoryGroup/golang-contract-bindings.git

> go get github.com/TuringAdvisoryGroup/golang-contract-bindings
```
