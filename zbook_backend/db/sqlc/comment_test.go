package db

import (
	"context"
	"fmt"
	"testing"

	"github.com/jackc/pgx/v5/pgtype"
	"github.com/stretchr/testify/require"
)

func testCreateRandomComment(t *testing.T) Comment {
	user := createRandomUser(t)
	markdown := testCreateRandomMarkdown(t)
	arg := CreateCommentParams{
		UserID:         user.UserID,
		RepoID:         markdown.RepoID,
		MarkdownID:     markdown.MarkdownID,
		ParentID:       pgtype.Int8{Int64: int64(0), Valid: false},
		RootID:         pgtype.Int8{Int64: int64(0), Valid: false},
		CommentContent: "hhhh",
	}
	comment, err := testStore.CreateComment(context.Background(), arg)
	require.NoError(t, err)
	require.Equal(t, comment.UserID, user.UserID)

	arg_second := CreateCommentParams{
		UserID:         user.UserID,
		RepoID:         markdown.RepoID,
		MarkdownID:     markdown.MarkdownID,
		ParentID:       pgtype.Int8{Int64: int64(0), Valid: false},
		RootID:         pgtype.Int8{Int64: int64(0), Valid: false},
		CommentContent: "bbb",
	}
	comment, err = testStore.CreateComment(context.Background(), arg_second)
	require.NoError(t, err)
	require.Equal(t, comment.UserID, user.UserID)
	return comment
}
func testCreateRandomCommentOnMd(t *testing.T, markdown Markdown) Comment {
	user := createRandomUser(t)

	arg := CreateCommentParams{
		UserID:         user.UserID,
		RepoID:         markdown.RepoID,
		MarkdownID:     markdown.MarkdownID,
		ParentID:       pgtype.Int8{Int64: int64(0), Valid: false},
		RootID:         pgtype.Int8{Int64: int64(0), Valid: false},
		CommentContent: "hhhh",
	}
	comment, err := testStore.CreateComment(context.Background(), arg)
	require.NoError(t, err)
	require.Equal(t, comment.UserID, user.UserID)

	arg_second := CreateCommentParams{
		UserID:         user.UserID,
		RepoID:         markdown.RepoID,
		MarkdownID:     markdown.MarkdownID,
		ParentID:       pgtype.Int8{Int64: int64(0), Valid: false},
		RootID:         pgtype.Int8{Int64: int64(0), Valid: false},
		CommentContent: "bbb",
	}
	comment, err = testStore.CreateComment(context.Background(), arg_second)
	require.NoError(t, err)
	require.Equal(t, comment.UserID, user.UserID)
	return comment
}
func TestCreateComment(t *testing.T) {
	testCreateRandomComment(t)
}
func TestListCommentLevelOne(t *testing.T) {
	md := testCreateRandomMarkdown(t)
	testCreateRandomCommentOnMd(t, md)
	testCreateRandomCommentOnMd(t, md)
	testCreateRandomCommentOnMd(t, md)
	arg := ListCommentLevelOneParams{
		MarkdownID: md.MarkdownID,
		Limit:      10,
		Offset:     0,
		UserID:     int64(md.UserID),
	}
	comments, error := testStore.ListCommentLevelOne(context.Background(), arg)
	require.NoError(t, error)
	require.Equal(t, len(comments), 6)
	for i := 0; i < len(comments); i++ {
		comment := comments[i]
		fmt.Println("comment i:", comment.IsLiked, comment.IsDisliked)
	}
}
func TestGetCommentDetail(t *testing.T) {
	commenta := testCreateRandomComment(t)
	arg := GetCommentDetailParams{
		CommentID: commenta.CommentID,
		UserID:    commenta.UserID,
	}
	comment, err := testStore.GetCommentDetail(context.Background(), arg)
	require.NoError(t, err)
	fmt.Println("comment:", comment.IsLiked)
	fmt.Print("comment:", comment.IsDisliked)
}

func TestDeleteComment(t *testing.T) {
	commenta := testCreateRandomComment(t)
	arg := GetCommentDetailParams{
		CommentID: commenta.CommentID,
		UserID:    commenta.UserID,
	}
	err := testStore.DeleteComment(context.Background(), commenta.CommentID)
	require.NoError(t, err)
	_, err = testStore.GetCommentDetail(context.Background(), arg)
	require.Error(t, err, ErrRecordNotFound)

}
