import { Link } from "@/navigation";
import UpdateCommentReportButton from "../../wrappers/UpdateCommentReportButton";
import ValueElement from "../ValueElement";
import TimeElement from "@/components/TimeElement";
import { useTranslations } from "next-intl";
import ListElementCard from "./ListElementCard";
import AvatarImageServer from "@/components/AvatarImageServer";
export default function ListCommentReportElement({
  report_id,
  repo_username,
  repo_name,
  relative_path,
  report_content,
  comment_content,
  created_at,
  processed,
  username,
}: {
  report_id: number;
  repo_name: string;
  repo_username: string;
  relative_path: string;
  report_content: string;
  comment_content: string;
  created_at: string;
  processed: boolean;
  username: string;
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
              <Link
                href={`/workspace/${repo_username}/o/${repo_name}/${relative_path}`}
                className="cursor-pointer"
              >
                <span
                  className="text-sky-700 dark:text-sky-600 font-semibold text-lg whitespace-nowrap overflow-x-auto scrollbar scrollbar-thumb-rounded-full scrollbar-track-rounded-full scrollbar-h-[6px]
                              scrollbar-thumb-slate-200 scrollbar-track-slate-100
                              dark:scrollbar-thumb-slate-500/50 dark:scrollbar-track-slate-500/[0.16]
                                hover:underline underline-offset-4 hover:text-sky-500 hover:dark:text-sky-500"
                >
                  {t("Comment")}
                </span>
              </Link>
            </div>
          </div>

          <UpdateCommentReportButton
            report_id={report_id}
            processed={processed}
          />
        </>
      }
      content={comment_content + "::" + report_content}
      footer={
        <>
          <div className="flex items-center justify-center space-x-1">
            <ValueElement
              tip="repo_user"
              content={repo_username as unknown as string}
            />
            <ValueElement
              tip="repo_name"
              content={repo_name as unknown as string}
            />
          </div>

          <ValueElement
            tip={t("CreatedAt")}
            content={<TimeElement timeInfo={created_at} />}
          />
        </>
      }
    />
  );
}
