"use client";
import { FaGithub } from "react-icons/fa";
import { FaGoogle } from "react-icons/fa";
import { FaLink, FaUnlink } from "react-icons/fa";
import { toast } from "react-toastify";
import SideBarLiContent from "./SideBarLiContent";
import { useRouter } from "@/navigation";
import { useTranslations } from "next-intl";
import {
  fetchServerWithAuthWrapper,
  refreshPage,
} from "@/fetchs/server_with_auth";
import { FetchServerWithAuthWrapperEndPoint } from "@/fetchs/server_with_auth_util";
import { FetchError } from "@/fetchs/util";
import { logger } from "@/utils/logger";

const getIconAndText = (oauthType: string) => {
  switch (oauthType) {
    case "github":
      return {
        icon: (
          <FaGithub className="text-slate-700 dark:text-slate-400 w-6 h-6" />
        ),
        text: "GitHub",
      };
    case "google":
      return {
        icon: <FaGoogle className="text-sky-700 dark:text-sky-400 w-6 h-6" />,
        text: "Google",
      };
    default:
      throw new Error("Unsupported oauth-party type");
  }
};
export default function OAuthWrapper({
  oauth_type,
  status,
  access_token,
}: {
  oauth_type: string;
  status: boolean;
  access_token: string;
}) {
  const t = useTranslations("SideBar");
  const router = useRouter();
  const { icon, text } = getIconAndText(oauth_type);
  const handleLink = async () => {
    if (status) {
      try {
        const data = await fetchServerWithAuthWrapper({
          endpoint: FetchServerWithAuthWrapperEndPoint.DELETE_OAUTH_LINK,
          xforward: "",
          agent: "",
          tags: [],
          values: {
            oauth_type: oauth_type,
          },
        });
        if (data.error) {
          throw new FetchError(data.message, data.status);
        }
        refreshPage("/workspace/[username]", true, false);
      } catch (error) {
        toast(t("FailedUnLinkAccount"), {
          type: "error",
          isLoading: false,
          autoClose: 1500,
        });
      }
    } else {
      switch (oauth_type) {
        case "github":
          router.push(`/oauth/github?access_token=${access_token}`);
          break;
        case "google":
          router.push(`/oauth/google?access_token=${access_token}`);
        default:
          logger.error(`Unsupported oauth-party type:${oauth_type}`);
      }
    }
  };

  return (
    <SideBarLiContent isSelected={false}>
      {icon}
      <span className="flex-1 ms-3 whitespace-nowrap">{text}</span>

      <button
        className="inline-flex items-center justify-center px-2 py-0.5 ms-3 text-xs font-medium text-gray-500 bg-gray-200 rounded dark:bg-gray-700 dark:text-gray-400"
        onClick={handleLink}
      >
        {status ? (
          <>
            <p className="p-1">{t("UnLink")}</p> <FaUnlink />
          </>
        ) : (
          <>
            <p className="p-1">{t("Link")}</p>
            <FaLink />
          </>
        )}
      </button>
    </SideBarLiContent>
  );
}
