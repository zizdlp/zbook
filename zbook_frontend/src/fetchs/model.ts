interface ListFollowerNotificationInfo {
  username: string;
  email: string;
  readed: boolean;
  noti_id: number;
  created_at: string;
}
interface ListRepoNotificationInfo {
  username: string;
  email: string;
  readed: boolean;
  noti_id: number;
  created_at: string;
  repo_id: number;
  repo_name: string;
}
interface ListCommentNotificationInfo {
  username: string;
  email: string;
  readed: boolean;
  noti_id: number;
  created_at: string;
  comment_content: string;
  repo_id: number;
  relative_path: string;
  repo_name: string;
  repo_username: string;
}

interface ListSystemNotificationInfo {
  readed: boolean;
  noti_id: number;
  created_at: string;
  title: string;
  contents: string;
  redirect_url: string;
}
interface MarkdownBasicInfo {
  markdown_id: number;
  relative_path: string;
  user_id: number;
  repo_id: number;
  main_content: string;
  table_content: string;
  md5: string;
  version_key: string;
  created_at: string;
  username: string;
  repo_name: string;
}

enum ListDataType {
  LIST_DATA_Type_UNSPECIFIED = 0,
  LIST_USER_REPO = 1,
  LIST_USER_FAVORITE = 2,
  LIST_USER_FOLLOWING = 3,
  LIST_USER_FOLLOWER = 4,
  LIST_ADMIN_USER = 5,
  LIST_ADMIN_SESSION = 6,
  LIST_ADMIN_REPO = 7,
  LIST_ADMIN_COMMENT = 8,
  LIST_ADMIN_COMMENT_REPORT = 9,
  LIST_REPO_VISI = 10,
  LIST_PUBLIC_REPO = 11,
}
export { ListDataType };

enum AnimateDirection {
  ANI_DIRECTION_UNSPECIFIED = 0,
  ANI_DIRECTION_RIGHT = 1,
  ANI_DIRECTION_LEFT = 2,
  ANI_DIRECTION_TOP = 3,
  ANI_DIRECTION_BOTTOM = 4,
}
export { AnimateDirection };

export type {
  ListFollowerNotificationInfo,
  ListRepoNotificationInfo,
  ListSystemNotificationInfo,
  ListCommentNotificationInfo,
  MarkdownBasicInfo,
};
