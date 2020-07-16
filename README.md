# natsStudy

https://github.com/nats-io/

https://github.com/nats-io/nats-server

https://github.com/nats-io/nats-server/releases/tag/v2.1.7

C:\> telnet xxx.xxx.xxx.xxx 4222

连上后Telnet立刻输出：

INFO{"server_id":"NADREYVHCIBQYI4LDIFKYM36JPZXJPNRG3FACPA3XL2S6H6FA3WSDDVU","version":"2.0.2","proto":1,"go":"go1.11.12","host":"0.0.0.0","port":4222,"max_payload":1048576,"client_id":1}

sub foo.* 90
+OK
