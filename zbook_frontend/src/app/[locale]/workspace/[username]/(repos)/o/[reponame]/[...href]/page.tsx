import WikiInfo from "@/components/WikiInfo";
import TableOfContent from "@/components/TableOfContent";
import NotFound from "@/components/loadings/NotFound";
import { fetchServerWithAuthWrapper } from "@/fetchs/server_with_auth";
import { FetchServerWithAuthWrapperEndPoint } from "@/fetchs/server_with_auth_util";
import { headers } from "next/headers";
import { Metadata } from "next";
import { SearchParams } from "@/types/interface";
import { fetchServerWithoutAuthWrapper } from "@/fetchs/server_without_auth";
import { FetchServerWithoutAuthWrapperEndPoint } from "@/fetchs/server_without_auth_util";
import { FetchError } from "@/fetchs/util";
import { logger } from "@/utils/logger";
export async function generateMetadata({
  params,
}: {
  params: { href: string; username: string; reponame: string };
}): Promise<Metadata> {
  var url = decodeURIComponent(params.href); // 将 params.href 转换为字符串
  var parts = url.split(",");
  var lastPart = parts[parts.length - 1];
  var firstPart = parts[0] ?? "";
  return {
    title: parts.length > 1 ? firstPart + ": " + lastPart : lastPart,
  };
}

export default async function MarkdownPage({
  params,
  searchParams,
}: {
  params: { href: string; username: string; reponame: string };
  searchParams?: SearchParams;
}) {
  try {
    const currentPage = Number(searchParams?.page) || 1;
    // const delay = Math.floor(Math.random() * 4000) + 14000;
    // await new Promise((resolve) => setTimeout(resolve, delay));
    const xforward = headers().get("x-forwarded-for") ?? "";
    const agent = headers().get("User-Agent") ?? "";
    const data = await fetchServerWithoutAuthWrapper({
      endpoint: FetchServerWithoutAuthWrapperEndPoint.GET_REPO_ID,
      xforward: xforward,
      agent: agent,
      values: {
        repo_name: decodeURIComponent(params.reponame),
        username: params.username,
      },
    });

    if (data.error) {
      throw new FetchError(data.message, data.status);
    }
    const verify_result = await fetchServerWithAuthWrapper({
      endpoint: FetchServerWithAuthWrapperEndPoint.GET_MARKDOWN_CONTENT,
      xforward: xforward,
      agent: agent,
      tags: [],
      values: {
        repo_id: data.repo_id,
        relative_path: decodeURIComponent(params.href).split(",").join("/"),
      },
    });
    if (verify_result.error) {
      throw new FetchError(verify_result.message, verify_result.status);
    }

    const { markdown } = verify_result;
    const markdownText = markdown.main_content;
    const markdownID = markdown.markdown_id;
    const href_seg = markdown.relative_path.split("/");
    const href = href_seg.slice(0, -1).join("/");
    let markdownList = "";
    let sectionIds: string[] = [];

    if (markdown.table_content) {
      markdownList = markdown.table_content;
      const reg = /href="#(.*?)"/g;
      const res = markdownList.match(reg);

      if (res) {
        sectionIds = res.map((value: string) => value.slice(7, -1));
      }
    }

    return (
      <div>
        <WikiInfo
          markdowntext={markdownText}
          markdownlist={markdownList}
          prefixPath={href}
          NavBarOpen={true}
          sectionIds={sectionIds}
          markdown_id={markdownID}
          currentPage={currentPage}
          searchParams={searchParams}
          username={params.username}
          repo_name={params.reponame}
        />
        <TableOfContent markdownlist={markdownList} sectionIds={sectionIds} />
      </div>
    );
  } catch (error) {
    let e = error as FetchError;
    logger.error(`fetch MarkdownPage failed:${e.message}`, e.status);
    return <NotFound />;
  }
}
