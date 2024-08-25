import WikiInfo from "@/components/WikiInfo";
import NotFound from "@/components/loadings/NotFound";
import { fetchServerWithAuthWrapper } from "@/fetchs/server_with_auth";
import { FetchServerWithAuthWrapperEndPoint } from "@/fetchs/server_with_auth_util";
import { headers } from "next/headers";
import { Metadata } from "next";
import { SearchParams } from "@/types/interface";
import { FetchError } from "@/fetchs/util";
import { logger } from "@/utils/logger";

async function fetchMarkdownContent({
  username,
  reponame,
  href,
  locale,
}: {
  username: string;
  reponame: string;
  href: string;
  locale: string;
}) {
  const xforward = headers().get("x-forwarded-for") ?? "";
  const agent = headers().get("User-Agent") ?? "";
  const response = await fetchServerWithAuthWrapper({
    endpoint: FetchServerWithAuthWrapperEndPoint.GET_MARKDOWN_CONTENT,
    xforward,
    agent,
    tags: [],
    values: {
      username,
      repo_name: decodeURIComponent(reponame),
      relative_path: decodeURIComponent(href).replace(/,/g, "/"),
      lang: locale || "en",
    },
  });

  if (response.error) {
    throw new FetchError(response.message, response.status);
  }
  return response;
}

function parseMarkdownList(markdownList: string): string[] {
  const reg = /href="#(.*?)"/g;
  const matches = markdownList.match(reg);
  return matches ? matches.map((val) => val.slice(7, -1)) : [];
}

export async function generateMetadata({
  params,
}: {
  params: { href: string; username: string; reponame: string };
}): Promise<Metadata> {
  const url = decodeURIComponent(params.href);
  const parts = url.split(",");
  const [firstPart, lastPart] = [parts[0] ?? "", parts[parts.length - 1]];
  return {
    title: parts.length > 1 ? `${firstPart}: ${lastPart}` : lastPart,
  };
}

export default async function MarkdownPage({
  params,
  searchParams,
}: {
  params: { href: string; username: string; reponame: string; locale: string };
  searchParams?: SearchParams;
}) {
  try {
    const currentPage = Number(searchParams?.page) || 1;
    const { username, reponame, href, locale } = params;
    // const delay = Math.floor(Math.random() * 4000) + 1400;
    // await new Promise((resolve) => setTimeout(resolve, delay));
    const data = await fetchMarkdownContent({
      username,
      reponame,
      href,
      locale,
    });

    const { markdown, prev, next, footers, updated_at, theme_color } = data;
    const markdownText = markdown.main_content;
    const markdownID = markdown.markdown_id;
    const hrefSegments = markdown.relative_path.split("/");
    const prefixPath = hrefSegments.slice(0, -1).join("/");
    const markdownList = markdown.table_content || "";
    const sectionIds = markdown.table_content
      ? parseMarkdownList(markdown.table_content)
      : [];

    return (
      <WikiInfo
        markdowntext={markdownText}
        markdownlist={markdownList}
        prefixPath={prefixPath}
        NavBarOpen={true}
        sectionIds={sectionIds}
        markdown_id={markdownID}
        currentPage={currentPage}
        searchParams={searchParams}
        username={username}
        repo_name={reponame}
        prev={prev}
        next={next}
        footers={footers}
        updated_at={updated_at}
        theme_color={theme_color}
      />
    );
  } catch (error) {
    const fetchError = error as FetchError;
    logger.error(
      `Failed to fetch MarkdownPage: ${fetchError.message}`,
      fetchError.status
    );
    return <NotFound />;
  }
}
