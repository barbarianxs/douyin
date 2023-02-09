update_user_kitex_gen:
	kitex -module "douyin" -I idl/ -type protobuf idl/user.proto

# update_note_kitex_gen:
	# kitex  idl/note.thrift