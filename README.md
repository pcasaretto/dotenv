# dotenv

A small command tool line that loads enviroment variables from a file and executes a command.

# Why?

There were others, but they all had dependencies (ruby, node).

## Usage

```
dotenv command [arg ...]:
  -file string
    	file to load variables from (default ".env")
```

## Example

```
$ cat .env
SECRET=VALUE
$ env | grep SECRET 
$ dotenv env | grep SECRET
SECRET=VALUE
```

## Instalation

Grab the [latest binary](https://github.com/pcasaretto/dotenv/releases/latest) and put it somewhere in your $PATH
