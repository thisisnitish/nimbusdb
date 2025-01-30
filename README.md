# NimbusDB

NimbusDB is a key value store. The aim of this project is to understand the core behaviour 
of Redis and its design. It provides various set of commands similar to Redis along with 
various data structures and more.

# Core Features
- **Key-Value Store** Supports commands like  `SET`, `GET`, `DEL`, `INCR`, `INCRBY` etc.
- **Data type support** String, Numbers, Set, List.
- **Pub/Sub support** For Publishers and Subscribers channels.
- **Disk Persistence** Supports Append Only Logs (AOF) for data recovery mechanism.


# Commands
## SET
Set the key in the store.
**Syntax**: `SET key value`
```zsh
nimbusdb > SET name john
OK
```

## GET
Get the key from the store.
**Syntax**: `GET key`
```zsh
nimbusdb > SET name john
OK
nimbusdb > GET name
john
```

## DELETE
Delete key from the store.
**Syntax**: `DEL key`
```zsh
nimbusdb > SET name john
OK
nimbusdb > GET name
john
nimbusdb > DEL name
john Deleted
```
> [!NOTE]
>
> It works with smallcase `set`, `get`, `del` commands also.

## Data Types
Nimbus db support few important datatypes.

## Strings
Nimbus supports the string data type, and you can use the `SET` and `GET` commands mentioned above to work with strings.

## Numbers
Nimbus supports the number data type and also offers a separate set of commands designed to work with it. Below is a detailed explanation of these commands.

### INCR
Increase the value by 1
**Syntax**: `INCR key`
```zsh
nimbusdb > SET count 1
OK
nimbusdb > INCR count
2
```

### INCRBY
Increase the value by some `x` value
**Syntax**: `INCRBY key x`
```zsh
nimbusdb > SET count 1
OK
nimbusdb > INCR count
2
nimbusdb > INCRBY count 5
7
```

### DECR
Decrease the value by 1
**Syntax**: `DECR key`
```zsh
nimbusdb > GET count
7
nimbusdb > DECR count
6
```

### DECRBY
Decrease the value by some `x` value
**Syntax**: `DECRBY key x`
```zsh
nimbusdb > GET count
7
nimbusdb > DECRBY count 4
3
```

## Lists
List also offers it separate set of commands designed to work with it. Below are the detaled explanation of commands.

ADD ALL COMMANDS WITH SYNTAX HERE...

## Sets

Set also offers it separate set of commands designed to work with it. Below are the detaled explanation of commands.

ADD ALL COMMANDS WITH SYNTAX HERE...

WORK IN PROGRESS...

