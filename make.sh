#!/bin/bash
cd ../XMagicHooker-Cmd && git pull && cd ../wechat-hacker && protoc --proto_path=../XMagicHooker-Cmd --go_out=proto ../XMagicHooker-Cmd/Cmd.proto ../XMagicHooker-Cmd/WwCmd.proto ../XMagicHooker-Cmd/WxCmd.proto
