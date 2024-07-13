interface CommentCountInfo {
  like_count: number;
  reply_count: number;
  is_liked: boolean;
  is_disliked: boolean;
  is_shared: boolean;
  is_reported: boolean;
}

interface ListUserInfo {
  username: string;
  email: string;
  avatar: string;
  repo_count: number;
  like_count: number;
  follower_count: number;
  following_count: number;
}

interface ListCommentInfo {
  comment_id: number;
  markdown_id: number;
  parent_id: number;
  username: string;
  pusername: string;
  comment_content: string;
  created_at: string;
  like_count: number;
  reply_count: number;
  is_liked: boolean;
  is_disliked: boolean;
  is_shared: boolean;
  is_reported: boolean;
}

export type { CommentCountInfo, ListUserInfo, ListCommentInfo };
