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
cp ../roll-smart-contracts/deploy/localhost/RollApprover.abi      ./social-money/abi/MultiSigWallet.abi
cp ../roll-smart-contracts/deploy/localhost/ERC20.abi      ./social-money/abi/ERC20.abi

```

```
abigen --abi=./social-money/abi/Manager.abi           --pkg=manager --out=./social-money/contracts/Manager.go
abigen --abi=./social-money/abi/Registry.abi          --pkg=registry --out=./social-money/contracts/Registry.go
abigen --abi=./social-money/abi/TokenFactory.abi      --pkg=tokenFactory --out=./social-money/contracts/TokenFactory.go
abigen --abi=./social-money/abi/TokenVesting.abi      --pkg=tokenVesting --out=./social-money/contracts/TokenVesting.go
abigen --abi=./social-money/abi/SocialMoney.abi       --pkg=socialMoney --out=./social-money/contracts/SocialMoney.go
abigen --abi=./social-money/abi/MultiSigWallet.abi       --pkg=multiSigWallet --out=./social-money/contracts/MultiSigWallet.go
abigen --abi=./social-money/abi/ERC20.abi       --pkg=ERC20 --out=./social-money/contracts/ERC20.go
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
