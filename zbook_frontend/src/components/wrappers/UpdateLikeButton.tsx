"use server";
import { revalidatePath } from "next/cache";
import { fetchServerWithAuthWrapper } from "@/fetchs/server_with_auth";
import { FetchServerWithAuthWrapperEndPoint } from "@/fetchs/server_with_auth_util";
import { headers } from "next/headers";
import { IoBookmarkOutline, IoBookmark } from "react-icons/io5";
import { FetchError } from "@/fetchs/util";
import { logger } from "@/utils/logger";

export default async function UpdateLikeButton({
  repo_id,
  is_liked,
  like_count,
}: {
  repo_id: number;
  is_liked: boolean;
  like_count: number;
}) {
  const xforward = headers().get("x-forwarded-for") ?? "";
  const agent = headers().get("User-Agent") ?? "";
  async function actionUpdateUserLike() {
    "use server";
    try {
      if (!is_liked) {
        const data = await fetchServerWithAuthWrapper({
          endpoint: FetchServerWithAuthWrapperEndPoint.CREATE_REPO_RELATION,
          xforward: xforward,
          agent: agent,
          tags: [],
          values: {
            repo_id: repo_id,
            relation_type: "like",
          },
        });
        if (data.error) {
          throw new FetchError(data.message, data.status);
        }
      } else {
        const data = await fetchServerWithAuthWrapper({
          endpoint: FetchServerWithAuthWrapperEndPoint.DELETE_REPO_RELATION,
          xforward: xforward,
          agent: agent,
          tags: [],
          values: {
            repo_id: repo_id,
            relation_type: "like",
          },
        });
        if (data.error) {
          throw new FetchError(data.message, data.status);
        }
      }
    } catch (error) {
      let e = error as FetchError;
      logger.error(`DELETE_REPO_RELATION failed:${e.message}`, e.status);
    }
    revalidatePath(`/workspace/`, "page");
  }

  return (
    <form action={actionUpdateUserLike}>
      <button
        className=" py-1 text-xs flex space-x-1 items-center justify-center"
        type="submit"
      >
        {is_liked ? (
          <IoBookmark className="h-5 w-5 text-slate-500 dark:text-slate-400" />
        ) : (
          <IoBookmarkOutline className="h-5 w-5 text-slate-500 dark:text-slate-400" />
        )}
        <p className="text-sm">{like_count}</p>
      </button>
    </form>
  );
}
