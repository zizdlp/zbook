import { Link } from "@/navigation";
import GitHost from "../../GitHost";
import ValueElement from "../ValueElement";

import UpdateLikeButton from "@/components/wrappers/UpdateLikeButton";
import TimeElement from "@/components/TimeElement";
import { useTranslations } from "next-intl";
import ListElementCard from "./ListElementCard";
import RepoButtons from "./RepoButtons";
import { ListDataType } from "@/fetchs/model";

export default function ListRepoElement({
  authname,
  repo_name,
  username,
  repo_description,
  visibility_level,
  git_host,
  like_count,
  is_liked,
  updated_at,
  created_at,
  listType,
  home,
}: {
  authname: string;
  repo_name: string;
  username: string;
  repo_description: string;
  visibility_level: string;
  git_host: string;
  like_count: number;
  is_liked: boolean;
  updated_at: string;
  created_at: string;
  listType: ListDataType;
  home: string;
}) {
  const t = useTranslations("DataList");
  return (
    <ListElementCard
      header={
        <>
          <div className="flex items-center justify-center space-x-2 py-1 overflow-hidden">
            <GitHost
              git_host={git_host}
              className="flex-none w-8 h-8 rounded-full object-cover"
            />

            <div
              className="flex flex-col justify-begin overflow-x-scroll
              scrollbar scrollbar-thumb-rounded-full scrollbar-track-rounded-full scrollbar-w-[6px]
            scrollbar-thumb-slate-200 scrollbar-track-slate-100
            dark:scrollbar-thumb-slate-500/50 dark:scrollbar-track-slate-500/[0.16]"
            >
              <Link
                href={`/workspace/${username}/o/${repo_name}/${home}`}
                className="cursor-pointer"
              >
                <div
                  className="text-sky-700 dark:text-sky-600 font-semibold text-base whitespace-nowrap overflow-x-scroll max-w-64
                          scrollbar scrollbar-thumb-rounded-full scrollbar-track-rounded-full scrollbar-w-[6px]
                        scrollbar-thumb-slate-200 scrollbar-track-slate-100
                        dark:scrollbar-thumb-slate-500/50 dark:scrollbar-track-slate-500/[0.16]
                          hover:underline underline-offset-4 hover:text-sky-500 hover:dark:text-sky-500"
                >
                  {repo_name}
                </div>
              </Link>
            </div>
          </div>
          <div className="flex space-x-1">
            <ValueElement
              tip={t("VisibilityLevel")}
              content={visibility_level}
            />
            <RepoButtons
              username={username}
              authname={authname}
              reponame={repo_name}
              listType={listType}
            />
          </div>
        </>
      }
      content={repo_description}
      footer={
        <>
          <UpdateLikeButton
            is_liked={is_liked}
            like_count={like_count}
            username={username}
            repo_name={repo_name}
          />
          <div className="flex items-center justify-center space-x-1">
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
