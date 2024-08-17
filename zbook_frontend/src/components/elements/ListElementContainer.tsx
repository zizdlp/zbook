import { ListDataType } from "@/fetchs/model";
import ListSessionElement from "./details/ListSessionElement";
import ListAdminCommentElement from "./details/ListAdminCommentElement";
import ListCommentReportElement from "./details/ListCommentReportElement";
import ListRepoElement from "./details/ListRepoElement";
import ListUserElement from "./details/ListUserElement";

export default function ListElementContainer({
  model,
  listType,
  authname,
  username,
  repo_name,
}: {
  model: any;
  listType: ListDataType;
  authname: string;
  username?: string;
  repo_name?: string;
}) {
  if (
    listType === ListDataType.LIST_USER_FOLLOWER ||
    listType == ListDataType.LIST_USER_FOLLOWING
  ) {
    return (
      <ListUserElement
        username={model.username}
        email={model.email}
        is_following={model.is_following ?? false}
        repo_count={model.repo_count ?? 0}
        updated_at={model.updated_at}
        created_at={model.created_at}
        listType={listType}
      />
    );
  } else if (listType === ListDataType.LIST_REPO_VISI) {
    return (
      <ListUserElement
        username={model.username}
        email={model.email}
        updated_at={model.updated_at}
        created_at={model.created_at}
        listType={listType}
        repo_username={username}
        repo_name={repo_name}
      />
    );
  } else if (
    listType === ListDataType.LIST_USER_FAVORITE ||
    listType === ListDataType.LIST_USER_REPO ||
    listType === ListDataType.LIST_PUBLIC_REPO ||
    listType === ListDataType.LIST_ADMIN_REPO
  ) {
    return (
      <ListRepoElement
        authname={authname}
        repo_name={model.repo_name}
        username={model.username}
        repo_description={model.repo_description}
        visibility_level={model.visibility_level}
        git_host={model.git_host}
        updated_at={model.updated_at}
        like_count={model.like_count ?? 0}
        is_liked={model.is_liked ?? false}
        created_at={model.created_at}
        listType={listType}
      />
    );
  } else if (listType === ListDataType.LIST_ADMIN_USER) {
    return (
      <ListUserElement
        username={model.username}
        email={model.email}
        blocked={model.blocked}
        verified={model.verified}
        role={model.role}
        created_at={model.updated_at}
        updated_at={model.updated_at}
        listType={ListDataType.LIST_ADMIN_USER}
      />
    );
  } else if (listType === ListDataType.LIST_ADMIN_SESSION) {
    return (
      <ListSessionElement
        username={model.username}
        email={model.email}
        user_agent={model.user_agent}
        client_ip={model.client_ip}
        created_at={model.created_at}
        expires_at={model.expires_at}
      />
    );
  } else if (listType === ListDataType.LIST_ADMIN_COMMENT) {
    return (
      <ListAdminCommentElement
        username={model.username}
        email={model.email}
        comment_content={model.comment_content}
        comment_id={model.comment_id}
        created_at={model.created_at}
      />
    );
  } else if (listType === ListDataType.LIST_ADMIN_COMMENT_REPORT) {
    return (
      <ListCommentReportElement
        report_id={model.report_id}
        repo_name={model.repo_name}
        repo_username={model.repo_username}
        relative_path={model.relative_path}
        report_content={model.report_content}
        comment_content={model.comment_content}
        created_at={model.created_at}
        processed={model.processed ?? false}
        username={model.username}
      />
    );
  }
}
