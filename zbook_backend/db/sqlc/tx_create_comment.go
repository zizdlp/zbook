package db

import (
	"context"
	"fmt"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type CreateCommentTxParams struct {
	CreateCommentParams
}

type CreateCommentTxResult struct {
	Comment Comment
}

func (store *SQLStore) CreateCommentTx(ctx context.Context, arg CreateCommentTxParams) (CreateCommentTxResult, error) {
	var result CreateCommentTxResult

	err := store.execTx(ctx, func(q *Queries) error {
		var err error

		comment, err := q.GetMarkdownByID(ctx, arg.CreateCommentParams.MarkdownID)
		if err != nil {
			return err
		}
		result.Comment, err = q.CreateComment(ctx, arg.CreateCommentParams)

		if err != nil {
			if ErrorCode(err) == UniqueViolation || ErrorCode(err) == ForeignKeyViolation {
				return status.Errorf(codes.AlreadyExists, "Comment already exist: %s", err)
			}

			return status.Errorf(codes.Internal, "fail to create Comment: %s", err)
		}

		// 父评论user，root 评论user，repo user
		// notify post owner
		if comment.UserID != arg.CreateCommentParams.UserID {
			arg_noti_post := CreateCommentNotificationParams{
				UserID:    comment.UserID,
				CommentID: result.Comment.CommentID,
			}
			_, err = q.CreateCommentNotification(ctx, arg_noti_post)
			if err != nil {
				fmt.Println("mydebug:create post comment noti error:", err)
				return err
			} else {
				err = q.UpdateUnreadCount(ctx, comment.UserID)
				if err != nil {
					fmt.Println("mydebug:update unread count error:", err)
					return err
				}
			}
		}

		if arg.CreateCommentParams.ParentID.Valid {
			// parentid 应该校验
			pcomment, err := q.GetCommentBasicInfo(ctx, arg.CreateCommentParams.ParentID.Int64)
			if err != nil {
				return err
			}

			if pcomment.MarkdownID != result.Comment.MarkdownID {
				return status.Errorf(codes.Internal, "pcomment error: pcomment not belong to this post")
			}

			// notify pcomment user,
			if pcomment.UserID != comment.UserID && pcomment.UserID != arg.CreateCommentParams.UserID {
				arg_noti := CreateCommentNotificationParams{
					UserID:    pcomment.UserID,
					CommentID: result.Comment.CommentID,
				}
				_, err = q.CreateCommentNotification(ctx, arg_noti)
				if err != nil {
					return err
				}
				err = q.UpdateUnreadCount(ctx, pcomment.UserID)
				if err != nil {
					fmt.Println("mydebug:update unread count error:", err)
					return err
				}
			}

			// noti root comment user
			if pcomment.RootID.Valid && pcomment.RootID.Int64 != pcomment.CommentID {
				// rootid 应该校验
				rootComment, err := q.GetCommentBasicInfo(ctx, pcomment.RootID.Int64)
				if err != nil {
					return err
				}
				if rootComment.UserID != pcomment.UserID && rootComment.UserID != comment.UserID && rootComment.UserID != arg.UserID {
					arg_noti := CreateCommentNotificationParams{
						UserID:    rootComment.UserID,
						CommentID: result.Comment.CommentID,
					}
					_, err = q.CreateCommentNotification(ctx, arg_noti)
					if err != nil {
						return err
					}
					err = q.UpdateUnreadCount(ctx, rootComment.UserID)
					if err != nil {
						return err
					}
				}
			}
		}
		return nil
	})

	return result, err
}
