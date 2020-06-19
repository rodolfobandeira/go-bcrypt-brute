go-bcrypt-brute
---

Pulls the list of active users (Admins and Managers) and test for weak passwords. It runs concurrently and defaults to the number of CPU cores available.


<p align="center">
  <img src="Screenshot%20from%202020-06-27%2017-58-32.png" />
</p>


---

#### Use:

```
./bin/go-bcrypt-brute 

Usage:
  go-bcrypt-brute [command]

Available Commands:
  config      Shows your database (MySQL) config
  help        Help about any command
  run         Runs go-bcrypt-brute using a passlist file
  show-users  Shows all users from your database (MySQL)

Flags:
  -h, --help   help for go-bcrypt-brute

Use "go-bcrypt-brute [command] --help" for more information about a command.
```

---

#### Example:

```bash
DB_USER=root DB_PASS=password DB_NAME=database ./bin/go-bcrypt-brute run 10k-most-common-passwords.txt | tee output.txt
```


```
Pulling all active users and checking for weak passwords...

Weak password user_id[439] -> password
Weak password user_id[365] -> password
Weak password user_id[2020] -> password
```

#### Variants:

- 10-most-common-passwords.txt
- 100-most-common-passwords.txt
- 10k-most-common-passwords.txt


#### Compiling:

It compiles to MacOS, Linux and Windows. You can find the binary files inside `/bin` folder. To compile:

```bash
make
```

