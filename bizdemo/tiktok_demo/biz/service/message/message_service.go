/*
 * Copyright 2023 CloudWeGo Authors
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package service

import (
	"context"
	"offer_tiktok/biz/dal/db"
	"offer_tiktok/biz/model/social/message"
	"offer_tiktok/pkg/errno"
	"offer_tiktok/pkg/utils"

	"github.com/cloudwego/hertz/pkg/app"
)

type MessageService struct {
	ctx context.Context
	c   *app.RequestContext
}

func NewMessageService(ctx context.Context, c *app.RequestContext) *MessageService {
	return &MessageService{ctx: ctx, c: c}
}

func (m *MessageService) GetMessageChat(req *message.DouyinMessageChatRequest) ([]*message.Message, error) {
	messages := make([]*message.Message, 0)
	from_user_id, _ := m.c.Get("current_user_id")
	to_user_id := req.ToUserId
	pre_msg_time := req.PreMsgTime
	db_messages, err := db.GetMessageByIdPair(from_user_id.(int64), to_user_id, utils.MillTimeStampToTime(pre_msg_time))
	if err != nil {
		return messages, err
	}
	for _, db_message := range db_messages {
		messages = append(messages, &message.Message{
			Id:         db_message.ID,
			ToUserId:   db_message.ToUserId,
			FromUserId: db_message.FromUserId,
			Content:    db_message.Content,
			CreateTime: db_message.CreatedAt.UnixNano() / 1000000,
		})
	}
	return messages, nil
}

func (m *MessageService) MessageAction(req *message.DouyinMessageActionRequest) error {
	from_user_id, _ := m.c.Get("current_user_id")
	to_user_id := req.ToUserId
	// action_type := req.ActionType
	content := req.Content

	ok, err := db.AddNewMessage(&db.Messages{
		FromUserId: from_user_id.(int64),
		ToUserId:   to_user_id,
		Content:    content,
	})
	if err != nil {
		return err
	}
	if !ok {
		return errno.MessageAddFailedErr
	}
	return nil
}
