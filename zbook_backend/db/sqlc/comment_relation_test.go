package db

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/zizdlp/zbook/util"
)

func TestCreateCommentReport(t *testing.T) {
	user2 := createRandomUser(t)
	comment := testCreateRandomComment(t)

	arg_report := CreateCommentReportParams{
		UserID:        user2.UserID,
		CommentID:     comment.CommentID,
		ReportContent: util.RandomString(368),
	}
	err := testStore.CreateCommentReport(context.Background(), arg_report)
	require.NoError(t, err)
}
