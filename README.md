## Stail
*Tail over SSH*

#### What is stail
A command line wrapper to tail multiple files over ssh.

#### Why
Often, I need to monitor logs of multiple instances of a same service, but SSH'ing and then running tail command was tedious but I still wanted to be able to use my `~/.ssh/config`.

#### How to get it
`curl -sf https://gobinaries.com/sitaramshelke/stail | sh`

#### How to use it
`stail ssh-host-1,file-path-1 ssh-host-2,filepath`

#### Any prerequisites?
Yes!      
The host needs be accessible over ssh **without** a password prompt. *Well, the password prompt works when you mention a single host but in case of multiple hosts, your `stdin` can't be passed to the ssh prompt.*

Its expected that you have a direct SSH access to the host without needing to do anything else. Eg. If you're using a bastion/jumpbox, you have a `ProxyCommand` config in your `~/.ssh/config`.

#### Why not handle the ssh host configuration in a more cleaner way?
I did spend some time looking at Go's ssh library and while it provides excellent APIs but I felt two problems -   
 - I didn't want to write my own SSH client just for tailing logs.
 - I didn't want to pass all of these fancy feature flags and limit the use to just for one purpose.

#### Anything cool?
It does adds colors to hostnames if your terminal supports it.

#### Anything else?
Feel free to raise an Issue or a PR if you'd like to contribute.
