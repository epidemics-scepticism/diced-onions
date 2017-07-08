# diced-onions
Generate vanity onions in Go

Generating a specific prefix

    $ go get github.com/epidemics-scepticism/diced-onions
    $ echo "lolol" > test
    $ ~/go/bin/diced-onions -wordlist=test -full=false
    2017/07/08 23:24:34 full search: false
    2017/07/08 23:24:34 wordlist: test
    2017/07/08 23:24:34 workers:2
    2017/07/08 23:24:34 added 1 words
    2017/07/08 23:24:34 thread 1 starting
    2017/07/08 23:24:34 thread 2 starting
    2017/07/08 23:24:34 press ctrl-c to exit
    2017/07/08 23:25:34 8216145 generated in 60 seconds. (136935 per sec)
    2017/07/08 23:26:34 16245248 generated in 120 seconds. (135377 per sec)
    2017/07/08 23:27:34 24178888 generated in 180 seconds. (134327 per sec)
    2017/07/08 23:28:34 32088631 generated in 240 seconds. (133702 per sec)
    2017/07/08 23:29:34 39956694 generated in 300 seconds. (133188 per sec)
    2017/07/08 23:30:09 match: lololfiooxrxhlip.onion
    ^C2017/07/08 23:30:14 signal received, exiting
    2017/07/08 23:30:14 thread 1 stopping
    2017/07/08 23:30:14 thread 2 stopping

Generating from only a sequence of pre-generated patterns or words

    $ ~/go/bin/diced-onions
    2017/07/08 23:32:44 full search: true
    2017/07/08 23:32:44 wordlist: words
    2017/07/08 23:32:44 workers:2
    2017/07/08 23:32:45 added 208088 words
    2017/07/08 23:32:45 thread 1 starting
    2017/07/08 23:32:45 thread 2 starting
    2017/07/08 23:32:45 press ctrl-c to exit
    2017/07/08 23:33:32 match: himtoowrytsrsibp.onion
    2017/07/08 23:33:43 match: keltiehinpussbaj.onion
    2017/07/08 23:33:45 6008233 generated in 60 seconds. (100137 per sec)
    2017/07/08 23:34:45 11890187 generated in 120 seconds. (99084 per sec)
    2017/07/08 23:35:40 match: adosdawsriwaneye.onion
    2017/07/08 23:35:45 17705726 generated in 180 seconds. (98365 per sec)
    2017/07/08 23:36:45 23486546 generated in 240 seconds. (97860 per sec)
    2017/07/08 23:37:45 29263325 generated in 300 seconds. (97544 per sec)
    2017/07/08 23:37:58 match: notshefauxhonoon.onion
    2017/07/08 23:38:45 35020446 generated in 360 seconds. (97279 per sec)
    2017/07/08 23:39:45 40784988 generated in 420 seconds. (97107 per sec)
    2017/07/08 23:40:45 46504392 generated in 480 seconds. (96884 per sec)
    2017/07/08 23:41:45 52242221 generated in 540 seconds. (96744 per sec)
    2017/07/08 23:42:38 match: ewerlamwowwhipow.onion
    2017/07/08 23:42:38 match: manwopprowzonkpr.onion
    2017/07/08 23:42:45 57964437 generated in 600 seconds. (96607 per sec)
    2017/07/08 23:43:01 match: bizportekealteug.onion
    2017/07/08 23:43:45 63667796 generated in 660 seconds. (96466 per sec)
    2017/07/08 23:44:45 69381087 generated in 720 seconds. (96362 per sec)
    2017/07/08 23:45:45 75074835 generated in 780 seconds. (96249 per sec)
    2017/07/08 23:46:10 match: honkfennimunauxw.onion
    2017/07/08 23:46:45 80764722 generated in 840 seconds. (96148 per sec)
    2017/07/08 23:46:59 match: jklbokinsisteucy.onion
    2017/07/08 23:47:21 match: tofubotakeylyeds.onion
    2017/07/08 23:47:45 86481232 generated in 900 seconds. (96090 per sec)
    2017/07/08 23:48:12 match: brapalsothiwisse.onion
    2017/07/08 23:48:45 92161745 generated in 960 seconds. (96001 per sec)
