import { Link } from "@/navigation";
import GitHost from "../../GitHost";

import { IoBookmarkOutline, IoBookmark } from "react-icons/io5";

import ToolTip from "@/components/ToolTip";

import ValueElement from "../ValueElement";
import TimeElement from "@/components/TimeElement";
import { useTranslations } from "next-intl";
import DeleteButton from "@/components/wrappers/DeleteButton";
import ListElementCard from "./ListElementCard";
import RepoButtons from "./RepoButtons";
export default function ListAdminRepoElement({
  repo_name,
  username,
  repo_description,
  visibility_level,
  git_host,
  like_count,
  is_liked,
  updated_at,
  created_at,
}: {
  repo_name: string;
  username: string;
  repo_description: string;
  visibility_level: string;
  git_host: string;
  like_count: number;
  is_liked: boolean;
  updated_at: string;
  created_at: string;
}) {
  const t = useTranslations("DataList");
  return (
    <ListElementCard
      header={
        <>
          <div className="flex items-center justify-center space-x-2">
            <GitHost
              git_host={git_host}
              className="flex-none w-12 h-12 rounded-full object-cover"
            />

            <div className="flex flex-col justify-begin">
              <Link
                href={`/workspace/${username}/o/${repo_name}`}
                className="cursor-pointer"
              >
                <div
                  className="text-sky-700 dark:text-sky-600 font-semibold text-lg whitespace-nowrap overflow-scroll max-w-64
                                  hover:underline underline-offset-4 hover:text-sky-500 hover:dark:text-sky-500"
                >
                  {repo_name}
                </div>
              </Link>
            </div>
          </div>
          <div className="flex space-x-1">
            <DeleteButton
              comment_id={0}
              username={username}
              repo_name={repo_name}
              dataType={"repo"}
            />
            {/* <RepoButtons
              username={username}
              authname={"authname"}
              reponame={repo_name}
            /> */}
          </div>
        </>
      }
      content={repo_description}
      footer={
        <>
          <ToolTip message={t("FavoriteCount")}>
            <div className="px-2 py-1 text-xs flex space-x-1 items-center justify-center">
              {is_liked ? (
                <IoBookmark className="h-5 w-5 text-slate-500 dark:text-slate-400" />
              ) : (
                <IoBookmarkOutline className="h-5 w-5 text-slate-500 dark:text-slate-400" />
              )}
              <p className="text-sm">{like_count}</p>
            </div>
          </ToolTip>

          <div className="flex items-center justify-center space-x-1">
            <ValueElement
              tip={t("CreatedAt")}
              content={<TimeElement timeInfo={created_at} />}
            />
            <ValueElement
              tip={t("UpdatedAt")}
              content={<TimeElement timeInfo={updated_at} />}
            />
            <ValueElement
              tip={t("VisibilityLevel")}
              content={visibility_level}
            />
          </div>
        </>
      }
    />
  );
}
