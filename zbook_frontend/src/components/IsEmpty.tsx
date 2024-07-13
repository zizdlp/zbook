import { ListDataType } from "@/fetchs/model";
import { ReactNode } from "react";
import { FaBox } from "react-icons/fa";
import { useTranslations } from "next-intl";
export default function IsEmpty({
  is_empty,
  listType,
  children,
}: {
  is_empty: boolean;
  listType: ListDataType;
  children: ReactNode;
}) {
  const t = useTranslations("DataList");
  let title = "";
  switch (listType) {
    case ListDataType.LIST_USER_REPO:
      title = t("NoListUserRepoTip");
      break;
    case ListDataType.LIST_USER_FAVORITE:
      title = t("NoListUserFavoriteTip");
      break;
    case ListDataType.LIST_PUBLIC_REPO:
      title = t("NoPublicRepoTip");
      break;
    case ListDataType.LIST_USER_FOLLOWER:
      title = t("NoListUserFollowerTip");
      break;
    case ListDataType.LIST_USER_FOLLOWING:
      title = t("NoListUserFollowingTip");
      break;

    case ListDataType.LIST_ADMIN_COMMENT:
      title = t("NoListAdminCommentTip");
      break;
    case ListDataType.LIST_ADMIN_COMMENT_REPORT:
      title = t("NoListAdminCommentReportTip");
      break;
    case ListDataType.LIST_ADMIN_SESSION:
      title = t("NoListAdminSessionTip");
      break;
    case ListDataType.LIST_ADMIN_REPO:
      title = t("NoListAdminRepoTip");
      break;
    case ListDataType.LIST_ADMIN_USER:
      title = t("NoListAdminUserTip");
      break;
    case ListDataType.LIST_REPO_VISI:
      title = t("NoListRepoUserTip");
      break;
    default:
      break;
      throw new Error("Unsupported oauth-party type");
  }
  if (is_empty) {
    return (
      <div className="min-h-[70vh] flex w-full items-center justify-center ">
        <div className="flex flex-col items-center space-y-2">
          <FaBox className="h-16 w-16 dark:text-slate-700/50 text-slate-400/50" />
          <p className="py-2 font-semibold text-xl dark:text-slate-600/50 text-slate-500/50">
            {title}
          </p>
        </div>
      </div>
    );
  } else {
    return children;
  }
}
