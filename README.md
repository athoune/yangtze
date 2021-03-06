Yangtze (长江 Yángzǐ Jiāng)
===========================

[![Build Status](https://drone.garambrogne.net/api/badges/athoune/yangtze/status.svg)](https://drone.garambrogne.net/athoune/yangtze)

Watch for text pattern using tokens and radix tree.

Yangtze target is filtering massive flow of logs.

Syntax
------

The pattern syntax is crude. The line is split on tokens. Tokens contain letter, digit, `_` and `-`, all other things disapear.

The pattern syntax use specific tokens.

 * `.` one token
 * `?` zero or one token
 * `...` one or more tokens

Example
-------

One boring example from my `/var/log/auth.log` :

    Mar  7 17:51:50 sd-127470 sshd[12455]: Failed password for invalid user cron from 51.15.72.126 port 59758 ssh2
    Mar  7 17:51:33 sd-127470 sshd[12453]: Failed password for root from 182.100.67.129 port 58472 ssh2

The pattern should be :

    Failed password for ... from ... port . ssh2

API
---

```golang
idx, err := index.NewSimpleIndex()
err = i.AddPatternBytes([]byte("Failed password for ... from ... port . ssh2"))
_, ok := i.ReadLine([]byte("Mar  7 17:51:50 sd-127470 sshd[12455]: Failed password for invalid user cron from 51.15.72.126 port 59758 ssh2"))
```

Licence
-------

3 terms BSD licence, ©2017 Mathieu Lecarme
