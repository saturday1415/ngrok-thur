    此包将任何文件转换为可管理的Go源代码。用于将二进制数据嵌入到go程序中。在转换为原始字节切片之前，文件数据可选地进行gzip压缩。
    https://github.com/jteeuwen/go-bindata
    
    #1 客户端: 
    	../../bin/go-bindata -nomemcopy -pkg=assets -tags=debug \
    		-debug=true \
    		-o=client/assets/assets_debug.go \
    		assets/client/...
    	or
    	../../bin/go-bindata -nomemcopy -pkg=assets -tags=release \
            		-debug=false \
            		-o=client/assets/assets_release.go \
            		assets/client/...
    --
    说明: -pkg    包名 string
         -debug  是否调试模式 true|false
         -o      生成文件路径 被生成文件 ...
    
    #build 编译客户端
        go build -tags 'release' ./main/ngrok
        
    #run 运行客户端
        ./ngrok -config=ngrok.cfg -log=ngrok.log start http
    
    
    
    server-assets: bin/go-bindata
    	../../bin/go-bindata -nomemcopy -pkg=assets -tags=release \
    		-debug=false \
    		-o=server/assets/assets_release.go \
    		assets/server/...


go build -tags 'release' ./main/ngrok

    ./ngrok -config ngrok.cfg start test

go build -tags 'release' ./main/ngrokd

    ./ngrokd -domain="thur.com"  -httpAddr=":8888" -httpsAddr=":4433" -tunnelAddr=":4443"

有问题:
go install ./main/ngrok    		
    		
没问题:
go install -tags 'release' ./main/ngrok

有问题:
go build -tags 'release' ./main/ngrok/ngrok.go

没问题:
go build -tags 'release' ./main/ngrok



生成证书

    cd assets
    mkdir cert
    cd cert
    openssl genrsa -out rootCA.key 2048
    openssl req -x509 -new -nodes -key rootCA.key -subj "/CN=thur.com" -days 5000 -out rootCA.pem
    openssl genrsa -out device.key 2048
    openssl req -new -key device.key -subj "/CN=thur.com" -out device.csr
    openssl x509 -req -in device.csr -CA rootCA.pem -CAkey rootCA.key -CAcreateserial -out device.crt -days 5000
    
    yes|cp rootCA.pem ../client/tls/ngrokroot.crt
    yes|cp device.crt ../server/tls/snakeoil.crt
    yes|cp device.key ../server/tls/snakeoil.key
