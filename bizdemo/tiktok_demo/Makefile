init_api:
	sh hz_gen.sh init

update_api:
	sh hz_gen.sh

# 由于publish中接收文件使用*multipart.FileHeader，修改需手动修改model中的publish
# 由于relation中在FriendUser中嵌入匿名结构体common.User，修改需手动修改model中的relation