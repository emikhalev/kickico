# KICKICO Test

The main goal of this cli application is to list and run public methods of the Ethereum contract. It uses infura.io.

To build and run the project you may use commands:
```
make build
bin/kickico
```

Configuration
---
Configuration file located in ```configs/config.yaml```

You can use following parameters:

```ContractABI``` - path to contract ABI file
```SmartContractAddress``` - address of the smart contract
```InfuraURL``` - infura URL of the smart contract
```CallMethodTimeout``` - timeout for method calling

Commands
---
Set path to config file (default: configs/config.yaml)
```
config=<path_to_config_file>
```

List methods of the contract:
```
kickico --list
```

Run contract method (use `-p` argument to set method parameters values, you can use comma to separate multiple values if needed):
```
kickico -method=<methodName> -p=<parameterValue1>,<parameterValue2>,...,<parameterValueN>
```


examples:
```
bin/kickico --list

bin/kickico -method=countAddresses

./bin/kickico -method=addAgingTime -p=1
```
