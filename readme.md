### High-Pass Filter
This tool takes a list of subdomains, counts the frequency of FQDNs within the list, then prints out a sorted frequency analysis of the FQDNs in the subdomain list. The goal of hpf is to help filter out FQDNs with loads of garbage subdomains. 

By default, hpf will print out the sorted list, however, if you pass the `-f` parameter, it will filter out subdomains which belong to FQDNs which appear more than `-f` times in the list. 

##### Installation 
```
go get github.com/Cgboal/hpf
```

##### Usage
```
λ cgboal [github.com/Cgboal/hpf] → ./hpf -h
Usage of ./hpf:
  -f int
    	Frequency cut off, FQDNs which appear more than this many times will be excluded
```

###### Performance
Proccessing a file with `7661809` domains takes around 20 seconds. Not the best, but it's a computationally expensive process.
```
λ cgboal [github.com/Cgboal/hpf] → time cat ~/BugBounty/MassBB/chaos_subdomains |wc -l                      
7661809

λ cgboal [github.com/Cgboal/hpf] → time cat ~/BugBounty/MassBB/chaos_subdomains | ./hpf -f 100000 >/dev/null  
cat ~/BugBounty/MassBB/chaos_subdomains  0.02s user 0.89s system 4% cpu 19.658 total
```
