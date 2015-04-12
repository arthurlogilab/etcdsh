# etcd-console

etcd-console is a command line tool for [etcd](https://github.com/coreos/etcd).
etcd-console provides filesystem-like access to etcd structure. 

## examples
open etcd-console (by default connects to http://localhost:4001)
<pre>
<code>./etcd-console</code>
</pre>
<pre>
<code>./etcd-console --url http://localhost:4001</code>
</pre>
change dir to foo/bar,list directory, change to one directory up
<pre>
<code>>cd foo/bar</code>
<code>foo/bar>ls</code>
<code>...</code>
<code>foo/bar>cd ..</code>
<code>foo>...</code>
</pre>

more commands will be added soon
