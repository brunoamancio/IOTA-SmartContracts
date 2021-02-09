## IOTA smart contracts - Base Template used to create other templates

A simple template used to start developing your own smart contracts for ISCP (IOTA Smart Contract Protocol) in Rust and write unit tests in Go. In order to develop using the environment set up by this template, you need to have good understanding of the Rust and Go languages.

This is a repository for myself but I welcome anyone interested in playing around with the current state of development of the IOTA Smart Contract Protocol (ISCP). Feel free to contact me on IOTA Foundation's discord server under Th3B0Y#8380.

### Simple structure prepared to start with development right away
This is how the templated file structure looks like:

![View of the template on VSCode](VSCode-TemplateView.png)

---

### Requirements for development in Rust 
- [Rust](https://www.rust-lang.org/tools/install)
- [Wasm-pack](https://rustwasm.github.io/wasm-pack/installer/)
- [Go](https://golang.org/dl/) - [Why Go?](WhyGo.md)
- Gcc (or equivalent for Windows [(TDM-GCC)](https://jmeubank.github.io/tdm-gcc/)) - [Why Gcc?](WhyGo.md)
- [Visual Studio Code](https://code.visualstudio.com/Download) (VSCode)
  - [Rust extension](https://marketplace.visualstudio.com/items?itemName=rust-lang.rust)
  - [Better TOML](https://marketplace.visualstudio.com/items?itemName=bungcip.better-toml) *Optional nice to have 
  - [Go Extension](https://marketplace.visualstudio.com/items?itemName=golang.Go)

### Requirements for development in Go (only)
- [Go](https://golang.org/dl/)
- Gcc (or equivalent for Windows [(TDM-GCC)](https://jmeubank.github.io/tdm-gcc/)) - [Why Gcc?](WhyGo.md)
- [Visual Studio Code](https://code.visualstudio.com/Download) (VSCode)
  - [Go Extension](https://marketplace.visualstudio.com/items?itemName=golang.Go)

### Set code up!
- Use this template repository to create your own.
- Clone your git repository with:
```
git clone --recurse-submodules <your_git_repository>
```
- Open your git repository on VSCode
- For Windows *only*, open file ".vscode/settings.json" and uncomment the setting "go.testFlags" entry "-buildmode=exe":
![Go.testFlags for Windows](go-testflags.png)
- Open file `go.mod` and change the go module's name from `module github.com/brunoamancio/IOTA-SmartContracts-Template-Base` to `module github.com/<your_username>/<your_project_name>`

*Done!* Now you can write your smart contract in `smartcontract/src/`, [compile](rust-docs/Compile-SmartContract.md) it, [run and debug](UnitTest-and-debug-SmartContract.md) it in `tests/smartcontract/my_iota_smart_contract_test.go`!

---
[MIT License](LICENSE)
