"use server";
import { revalidatePath } from "next/cache";
import { fetchServerWithAuthWrapper } from "@/fetchs/server_with_auth";
import { FetchServerWithAuthWrapperEndPoint } from "@/fetchs/server_with_auth_util";
import { headers } from "next/headers";
import { getTranslations } from "next-intl/server";
import { FetchError } from "@/fetchs/util";
import { logger } from "@/utils/logger";
export default async function UpdateBlockButton({
  username,
  is_blocked,
}: {
  username: string;
  is_blocked: boolean;
}) {
  const t = await getTranslations("Dialog");
  const xforward = headers().get("x-forwarded-for") ?? "";
  const agent = headers().get("User-Agent") ?? "";
  async function actionUpdateUserBlock(formData: FormData) {
    "use server";
    try {
      const data = await fetchServerWithAuthWrapper({
        endpoint: FetchServerWithAuthWrapperEndPoint.UPDATE_USER_BLOCK,
        xforward: xforward,
        agent: agent,
        tags: [],
        values: {
          username: username,
          blocked: !is_blocked,
        },
      });
      if (data.error) {
        throw new FetchError(data.message, data.status);
      }
    } catch (error) {
      let e = error as FetchError;
      logger.error(`UPDATE_USER_BLOCK failed:${e.message}`, e.status);
    }

    revalidatePath(`/workspace/${username}/admin_users`, "page");
  }

  return (
    <form action={actionUpdateUserBlock}>
      <button
        className="bg-sky-500 dark:bg-sky-700 text-white rounded-lg py-2 px-4 font-semibold text-sm cursor-pointer hover:bg-sky-600 dark:hover:bg-sky-500"
        type="submit"
      >
        <span className="flex-1 whitespace-nowrap text-xs">
          {is_blocked ? t("UnBlock") : t("Block")}
        </span>
      </button>
    </form>
  );
}
