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
NimbusDB support few important datatypes.

## Strings
Nimbus supports the string data type, and you can use the `SET` and `GET` commands mentioned above to work with strings.

## Numbers
Nimbus supports the number data type and also offers a separate set of commands designed to work with it. Below is a detailed explanation of these commands.

### INCR
Increase the value by 1.
**Syntax**: `INCR key`
```zsh
nimbusdb > SET count 1
OK
nimbusdb > INCR count
2
```

### INCRBY
Increase the value by some `x` value.
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
Decrease the value by 1.
**Syntax**: `DECR key`
```zsh
nimbusdb > GET count
7
nimbusdb > DECR count
6
```

### DECRBY
Decrease the value by some `x` value.
**Syntax**: `DECRBY key value`
```zsh
nimbusdb > GET count
7
nimbusdb > DECRBY count 4
3
```

## Lists
List also offers it separate set of commands designed to work with it. Below are the detaled explanation of commands.

### LPUSH
Push the element to the left side of the list.
**Syntax**: `LPUSH key value`
```zsh
nimbusdb > LPUSH cities delhi
1
nimbusdb > LPUSH cities jaipur
2
```

### RPUSH
Push the element to the right side of the list.
**Syntax**: `RPUSH key value`
```zsh
nimbusdb > RPUSH cities lucknow
3
nimbusdb > RPUSH cities bhopal
4
```

### LLEN
Returns the length of the list.
**Syntax**: `LLEN key`
```zsh
nimbusdb > LLEN cities
4
```

### LPOP
Removes the leftmost element from the list.
**Syntax**: `LPOP key`
```zsh
nimbusdb > LPOP cities
jaipur
```

### RPOP
Removes the rightmost element from the list.
**Syntax**: `RPOP key`
```zsh
nimbusdb > RPOP cities
bhopal
```

### LINDEX
Returns the element from the list at that index.
**Syntax**: `LINDEX key value`
```zsh
nimbusdb > LINDEX cities 0
delhi
nimbusdb > LINDEX cities 1
lucknow
```

## Sets

Set also offers it separate set of commands designed to work with it. Below are the detaled explanation of commands.

ADD ALL COMMANDS WITH SYNTAX HERE...

WORK IN PROGRESS...

