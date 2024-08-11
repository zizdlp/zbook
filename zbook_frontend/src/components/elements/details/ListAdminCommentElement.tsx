import { Link } from "@/navigation";
import ValueElement from "../ValueElement";
import TimeElement from "@/components/TimeElement";
import { useTranslations } from "next-intl";
import DeleteButton from "@/components/wrappers/DeleteButton";
import ListElementCard from "./ListElementCard";
import AvatarImageServer from "@/components/AvatarImageServer";
export default function ListAdminCommentElement({
  username,
  email,
  comment_id,
  comment_content,
  created_at,
}: {
  username: string;
  email: string;
  comment_id: number;
  comment_content: string;
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
          <DeleteButton
            comment_id={comment_id}
            username={""}
            repo_name=""
            dataType={"comment"}
          />
        </>
      }
      content={comment_content}
      footer={
        <>
          <ValueElement
            tip="comment id"
            content={comment_id as unknown as string}
          />

          <ValueElement
            tip={t("CreatedAt")}
            content={<TimeElement timeInfo={created_at} />}
          />
        </>
      }
    />
  );
}
