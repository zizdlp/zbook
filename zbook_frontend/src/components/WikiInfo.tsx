import HtmlParser from "@/components/parsers/HtmlParser";
import { auth } from "@/auth";
interface WikiInfoProps {
  sectionIds: string[];
  markdownlist: string;
  markdowntext: string;
  prefixPath: string;
  NavBarOpen: boolean;
  markdown_id: number;
  currentPage: number;
  searchParams?: SearchParams;
  username: string;
  repo_name: string;
  prev: string;
  next: string;
  footers: FooterSocial[];
  updated_at: string;
  theme_color: ThemeColor;
}
import CreateComment from "./comments/CreateComment";
import MainContentWrapper from "./wrappers/MainContentWrapper";
import { FooterSocial, SearchParams } from "@/types/interface";
import ListLevelOneComment from "./comments/ListLevelOneComment";
import TableOfContent, { ThemeColor } from "./TableOfContent";
import { getTranslations } from "next-intl/server";
import MainContentFooter from "./MainContentFooter";
import { headers } from "next/headers";
export default async function WikiInfo(props: WikiInfoProps) {
  const session = await auth();
  const t = await getTranslations("Footer");
  const agent = headers().get("User-Agent") ?? "";
  return (
    <MainContentWrapper
      right_sidebar={
        <TableOfContent
          markdownlist={props.markdownlist}
          sectionIds={props.sectionIds}
          theme_color={props.theme_color}
        />
      }
    >
      <HtmlParser
        htmlString={props.markdowntext}
        prefixPath={props.prefixPath}
        username={props.username}
        repo_name={props.repo_name}
        theme_color={props.theme_color}
        agent={agent}
      />

      <MainContentFooter
        prev={props.prev}
        next={props.next}
        username={props.username}
        repo_name={props.repo_name}
        updated_at={props.updated_at}
        footers={props.footers}
        theme_color={props.theme_color}
      />

      {session?.access_token && (
        <div className="print:hidden">
          <CreateComment
            parentID={0}
            markdownID={props.markdown_id}
            username={session.username}
            theme_color={props.theme_color}
          />

          <div className="pb-16">
            <ListLevelOneComment
              markdown_id={props.markdown_id}
              authname={session.username}
            />
          </div>
        </div>
      )}
    </MainContentWrapper>
  );
}
