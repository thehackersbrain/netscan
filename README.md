# NetScan
> Simple TCP Port Scanner<br/>
> Written in Go Lang by Gaurav Raj \[TheHackersBrain\]

![bannerImage](https://thehackersbrain.github.io/assets/hiddenwave/banner.png)

NetScan is a simple `TCP` Port Scanner written in Golang while learning it.

**NOTE:** This tool only supports TCP Protocol, will add more eventually ;)

## Version
**NetScan 1.0.0**

## Requirement
Nothing Specifically Required just Golang should be installed.

## Todo
- [ ] To Add `UDP` and other method supports.
- [ ] To Add Service Banner Grabbing for datails of the Port Service.
- [ ] To Improve it further :)

## Installation and Uses

### Automated Installation (Recommanded)
Change the directory to where you want to install this tool and run the follwing command
```bash
curl https://raw.githubusercontent.com/thehackersbrain/netscan/main/install.sh -s | bash
```

### Manual Installation

- Make Sure Golang in Installed
```bash
go --version
```
- Git clone this repo and change the directory
```bash
git clone https://github.com/thehackersbrain/netscan.git && cd netscan
```
- Now Build the package
```bash
go build main.go && mv main netscan
```

- Copy the binary in `/usr/bin/` for easy access \(optional\)
```bash
sudo cp netscan /usr/bin/
```

## How to Use

- For specifing port, default is 1024
    ```bash
    ./netscan --port 1024
    ```
- For specifing the target, default is `scanme.nmap.org`
    ```bash
    ./netscan --target 127.0.0.1
    ```


## Author

**Creator:** [Gaurav Raj](https://github.com/thehackersbrain/)<br/>
**Portfolio:** [Here](https://gauravraj.tech/about)<br/>
**Blog:** [TheHackersBrain Blog](https://gauravraj.tech/)<br/>
**Projects:** [Here](https://github.com/thehackersbrain?tab=repositories)<br/>
**Twitter:** [@thehackersbrain](https://twitter.com/thehackersbrain)<br/>
**TryHackMe:** [hackersbrain](https://tryhackme.com/p/hackersbrain)
