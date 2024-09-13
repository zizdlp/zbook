import { Link } from "@/navigation";
import ValueElement from "../ValueElement";
import TimeElement from "@/components/TimeElement";
import { useTranslations } from "next-intl";
import ListElementCard from "./ListElementCard";
import AvatarImageServer from "@/components/AvatarImageServer";
import UserButtons from "./UserButtons";
import { ListDataType } from "@/fetchs/model";
import ToolTip from "@/components/ToolTip";
import { IoMdBookmarks } from "react-icons/io";
import UpdateVisibleButton from "@/components/wrappers/UpdateVisibilityButton";
export default function ListUserElement({
  username,
  email,
  blocked,
  verified,
  is_following,
  repo_count,
  role,
  repo_name,
  repo_username,
  created_at,
  updated_at,
  listType,
}: {
  username: string;
  email: string;
  blocked?: boolean;
  verified?: boolean;
  is_following?: boolean;
  repo_count?: number;
  role?: string;
  repo_name?: string;
  repo_username?: string;
  created_at: string;
  updated_at: string;
  listType: ListDataType;
}) {
  const t = useTranslations("DataList");
  return (
    <ListElementCard
      header={
        <>
          <div className="flex items-center justify-center space-x-2">
            <AvatarImageServer
              username={username}
              className="flex-none w-12 h-12 rounded-full  object-cover"
            />
            <div className="flex flex-col justify-begin">
              <Link href={`/workspace/${username}`} className="cursor-pointer">
                <span
                  className="text-sky-700 dark:text-sky-600 font-semibold text-lg whitespace-nowrap overflow-scroll
                                hover:underline underline-offset-4 hover:text-sky-500 hover:dark:text-sky-500"
                >
                  {username}
                </span>
              </Link>
              <span className=" text-slate-700 dark:text-slate-400 text-xs overflow-scroll">
                {email}
              </span>
            </div>
          </div>
          <div className="flex space-x-1">
            {(listType === ListDataType.LIST_USER_FOLLOWER ||
              listType === ListDataType.LIST_USER_FOLLOWING) && (
              <ValueElement
                tip={t("FollowStatus")}
                content={is_following ? t("Followed") : t("UnFollowed")}
              />
            )}
            {listType == ListDataType.LIST_ADMIN_USER && role != "admin" && (
              <UserButtons username={username} is_blocked={blocked ?? false} />
            )}
            {listType == ListDataType.LIST_REPO_VISI && (
              <UpdateVisibleButton
                username={username}
                repo_username={repo_username ?? ""}
                repo_name={repo_name ?? ""}
                is_visible={true}
              />
            )}
          </div>
        </>
      }
      content={""}
      footer={
        <>
          <div className="flex items-center justify-end space-x-1">
            {(listType === ListDataType.LIST_USER_FOLLOWER ||
              listType === ListDataType.LIST_USER_FOLLOWING) && (
              <ToolTip message={t("RepoCount")}>
                <div className="px-2 py-1 text-xs flex space-x-1 items-center justify-center">
                  <IoMdBookmarks className="h-5 w-5" />
                  <p className="text-sm">{repo_count}</p>
                </div>
              </ToolTip>
            )}

            {listType == ListDataType.LIST_ADMIN_USER && (
              <ValueElement
                tip={t("AccountVerification")}
                content={
                  <span className="whitespace-nowrap">
                    {verified ? t("Verified") : t("UnVerified")}
                  </span>
                }
              />
            )}
          </div>
          <div className="flex items-center justify-end space-x-1">
            <ValueElement
              tip={t("CreatedAt")}
              content={<TimeElement timeInfo={created_at} />}
            />
            <ValueElement
              tip={t("UpdatedAt")}
              content={<TimeElement timeInfo={updated_at} />}
            />
          </div>
        </>
      }
    />
  );
}
