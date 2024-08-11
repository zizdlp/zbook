import { Link } from "@/navigation";
import UpdateBlockButton from "../../wrappers/UpdateBlockButton";
import ValueElement from "../ValueElement";
import TimeElement from "@/components/TimeElement";
import { useTranslations } from "next-intl";
import DeleteButton from "@/components/wrappers/DeleteButton";
import ListElementCard from "./ListElementCard";
import AvatarImageServer from "@/components/AvatarImageServer";
import UserButtons from "./UserButtons";
export default function ListAdminUserElement({
  username,
  email,
  blocked,
  verified,
  role,
  updated_at,
}: {
  username: string;
  email: string;
  blocked: boolean;
  verified: boolean;
  role: string;
  updated_at: string;
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
            <UserButtons username={username} is_blocked={blocked ?? false} />
          </div>
        </>
      }
      content={""}
      footer={
        <>
          <ValueElement
            tip={t("AccountVerification")}
            content={verified ? t("Verified") : t("UnVerified")}
          />
          <ValueElement
            tip={t("UpdatedAt")}
            content={<TimeElement timeInfo={updated_at} />}
          />
        </>
      }
    />
  );
}
