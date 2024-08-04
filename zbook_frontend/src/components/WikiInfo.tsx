/* eslint-disable react/jsx-no-literals */
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
}
import CreateComment from "./comments/CreateComment";
import MarkdownWrapper from "./wrappers/MarkdownWrapper";
import { SearchParams } from "@/types/interface";
import ListLevelOneComment from "./comments/ListLevelOneComment";
import TableOfContent from "./TableOfContent";
export default async function WikiInfo(props: WikiInfoProps) {
  const session = await auth();
  return (
    <MarkdownWrapper
      contentsidebar={
        <TableOfContent
          markdownlist={props.markdownlist}
          sectionIds={props.sectionIds}
        />
      }
    >
      <HtmlParser
        htmlString={props.markdowntext}
        prefixPath={props.prefixPath}
        username={props.username}
        repo_name={props.repo_name}
      />

      {session?.access_token && (
        <>
          <CreateComment
            parentID={0}
            markdownID={props.markdown_id}
            username={session.username}
          />

          <div className="pb-16">
            <ListLevelOneComment
              markdown_id={props.markdown_id}
              authname={session.username}
            />
          </div>
        </>
      )}
    </MarkdownWrapper>
  );
}
