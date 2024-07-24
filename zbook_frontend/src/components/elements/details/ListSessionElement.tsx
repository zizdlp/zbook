import { Link } from "@/navigation";
import ValueElement from "../ValueElement";
import TimeElement from "@/components/TimeElement";
import { useTranslations } from "next-intl";
import ListElementCard from "./ListElementCard";
import { parseUserAgent } from "@/utils/util";
import AvatarImageServer from "@/components/AvatarImageServer";
export default function ListSessionElement({
  username,
  email,
  user_agent,
  client_ip,
  expires_at,
  created_at,
}: {
  username: string;
  email: string;
  user_agent: string;
  client_ip: string;
  expires_at: string;
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
          <div className="flex items-center justify-center space-x-1 ">
            <ValueElement
              tip="agent"
              content={parseUserAgent(user_agent).browser}
            />
          </div>
        </>
      }
      content={""}
      footer={
        <>
          <div className="flex items-center justify-center space-x-1">
            <ValueElement tip="client ip" content={client_ip.split(",")[0]} />
          </div>
          <div className="flex items-center justify-center space-x-1">
            <ValueElement
              tip={t("CreatedAt")}
              content={<TimeElement timeInfo={created_at} />}
            />
            <ValueElement
              tip={t("ExpiredAt")}
              content={<TimeElement timeInfo={expires_at} />}
            />
          </div>
        </>
      }
    />
  );
}
