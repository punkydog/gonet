GS <---> HUB 

通信协议

|LENGTH|SEQNUM|PROTO|PAYLOAD|

|16|64|16|...|

包测试：

# get state
echo "000E 0000000000000001 0009 00000001" | xxd -r -ps |nc 127.0.0.1 8889 -q 2|hexdump -C

# forward content "ABCD" to user 1 
echo "0012 0000000000000001 0000 00000001 0002 ABCD" | xxd -r -ps |nc 127.0.0.1 8889 -q 2|hexdump -C