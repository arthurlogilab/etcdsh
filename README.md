# etcd-console

etcd-console is a command line tool for [etcd](https://github.com/coreos/etcd).
etcd-console provides filesystem-like access to etcd structure. 
Although there is official command line tool [etcdctl](https://github.com/coreos/etcd/tree/master/etcdctl), it is enoying you have to enter the same (ofter very long) keys again for every command. etcd-console tries to make it simpler and faster.

## Building
etcd-console is written using go language, it can be build using standard go build tool. 

## Downloads binaries
 * [mac](https://github.com/kamilhark/etcd-console/releases/download/0.0.1-ALPHA/etcd-console) 
 * linux (soon)
 * windows (soon)

## examples
open etcd-console (by default it connects to http://localhost:4001)
<pre>
<code>./etcd-console</code>
</pre>
<pre>
<code>./etcd-console --url http://localhost:4001</code>
</pre>
change dir to foo/bar and list content of it
<pre>
<code>>cd foo/bar</code>
<code>foo/bar>ls</code>
<code>...</code>
</pre>
go one directory up
<pre>
<code>foo/bar>cd ..</code>
<code>foo>...</code>
</pre>
set value
<pre>
<code>foo/bar>set key value</code>
</pre>
get value
<pre>
<code>foo/bar>get key</code>
<code>foo/bar>get /foo/bar/key</code>
<code>foo/bar>get /foo/bar/../bar/key</code>
</pre>
rm key or dir
<pre>
<code>foo/bar>rm key</code>
<code>foo/bar>rm key/abc/def</code>
</pre>
exit etcd-console
<pre>
<code>foo/bar>exit</code>
</pre>

