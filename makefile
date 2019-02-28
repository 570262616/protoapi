
out_Path = /Users/zhangpeng/Desktop/test
protoapi_Path = /Users/zhangpeng/Desktop/warehouse/blazer-api/apidoc/
alsApidoc_Path = /Users/zhangpeng/Desktop/warehouse/alsApidoc/
apidoc_Path = /Users/zhangpeng/Desktop/warehouse/blazer-api/apidoc/proto

srcfiles := $(wildcard $(alsApidoc_Path)/*/*/*.proto)

blazerSrcfiles := $(wildcard $(apidoc_Path)/*/*.proto)

gen:
	go generate
	go build
	cp /Users/zhangpeng/go/src/github.com/yoozoo/protoapi/protoapi /Users/zhangpeng/Desktop/warehouse/blazer-api/apidoc

ttt: 
	/Users/zhangpeng/Desktop/warehouse/blazer-api/apidoc/protoapi gen --lang=als /Users/zhangpeng/Desktop/test /Users/zhangpeng/Desktop/warehouse/blazer-api/apidoc/proto/coupon/coupon_public.proto --proto_path=/Users/zhangpeng/Desktop/warehouse/blazer-api/apidoc/proto/

# als: 
	#$(protoapi_Path)/protoapi gen --lang=als $(out_Path) /Users/zhangpeng/Desktop/warehouse/alsApidoc/basis/proto/link.proto --proto_path=$(alsApidoc_Path)

	#$(warning $(srcfiles))
	



# alsall:
# 	for file in $(srcfiles) ; \
# 	do \
#     $(protoapi_Path)/protoapi gen --lang=als $(out_Path) $$file --proto_path=$(alsApidoc_Path) ; \
# 	done

blazer:
	for file in $(blazerSrcfiles) ; \
	do \
		$(protoapi_Path)protoapi gen --lang=blazer $(out_Path) $$file --proto_path=$(apidoc_Path) ; \
	done