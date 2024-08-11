import { IoMdBookmarks } from "react-icons/io";
import { Link } from "@/navigation";
import ToolTip from "@/components/ToolTip";
import ValueElement from "../ValueElement";
import TimeElement from "@/components/TimeElement";
import { useTranslations } from "next-intl";
import ListElementCard from "./ListElementCard";
import AvatarImageServer from "@/components/AvatarImageServer";
export default function ListRepoVisiElement({
  username,
  email,
  is_following,
  repo_count,
  updated_at,
  created_at,
}: {
  username: string;
  email: string;
  is_following: boolean;
  repo_count: number;
  updated_at: string;
  created_at: string;
}) {
  const t = useTranslations("DataList");
  return (
    <ListElementCard
      header={
        <>
          <div className="flex items-center justify-center space-x-2">
            <AvatarImageServer
              username={username}
              className="flex-none w-12 h-12 rounded-full object-cover"
            />
            <div className="flex flex-col justify-begin">
              <Link href={`/workspace/${username}`} className="cursor-pointer">
                <div
                  className="text-sky-700 dark:text-sky-600 font-semibold text-lg whitespace-nowrap overflow-scroll max-w-64
                                  hover:underline underline-offset-4 hover:text-sky-500 hover:dark:text-sky-500"
                >
                  {username}
                </div>
              </Link>
              <span className=" text-slate-700 dark:text-slate-400 text-xs overflow-scroll max-w-64">
                {email}
              </span>
            </div>
          </div>
          <ValueElement
            tip={t("FollowStatus")}
            content={is_following ? t("Followed") : t("UnFollowed")}
          />
        </>
      }
      content={""}
      footer={
        <>
          <ToolTip message={t("PublicRepoCount")}>
            <div className="px-2 py-1 text-xs flex space-x-1 items-center justify-center">
              <IoMdBookmarks className="h-5 w-5" />
              <p className="text-sm">{repo_count}</p>
            </div>
          </ToolTip>
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
