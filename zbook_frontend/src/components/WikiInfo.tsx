import HtmlParser from "@/components/parsers/HtmlParser";
import { auth } from "@/auth";
import { Link } from "@/navigation";
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
}
import CreateComment from "./comments/CreateComment";
import MainContentWrapper from "./wrappers/MainContentWrapper";
import { FooterSocial, SearchParams } from "@/types/interface";
import ListLevelOneComment from "./comments/ListLevelOneComment";
import TableOfContent from "./TableOfContent";
import IconFilter from "./IconFilter";
import TimeElement from "./TimeElement";
import { getTranslations } from "next-intl/server";
export default async function WikiInfo(props: WikiInfoProps) {
  const session = await auth();
  const t = await getTranslations("Footer");
  return (
    <MainContentWrapper
      right_sidebar={
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

      <div className="leading-6 mt-14">
        <div className="not-prose my-12 grid lg:grid-cols-2 gap-4">
          {props.prev && (
            <div className="group relative rounded-xl border border-slate-200 dark:border-slate-800">
              <div className="absolute -inset-px rounded-xl border-2 border-transparent opacity-0 [background:linear-gradient(var(--quick-links-hover-bg,theme(colors.purple.50)),var(--quick-links-hover-bg,theme(colors.purple.50)))_padding-box,linear-gradient(to_top,theme(colors.purple.400),theme(colors.purple.300),theme(colors.indigo.300))_border-box] group-hover:opacity-100 dark:[--quick-links-hover-bg:theme(colors.slate.800)]"></div>
              <div className="relative overflow-hidden rounded-xl py-3 px-6 text-left">
                <h2 className="mt-1 font-display text-base text-slate-900 dark:text-white">
                  <a
                    href={`/workspace/${props.username}/o/${props.repo_name}/${props.prev}`}
                  >
                    <span className="absolute -inset-px rounded-xl"></span>
                    {t("PrevPage")}
                  </a>
                </h2>
                <p className="mt-1 text-sm text-slate-700 dark:text-slate-400">
                  {props.prev}
                </p>
              </div>
            </div>
          )}

          {props.next && (
            <div
              className={`group relative rounded-xl border border-slate-200 dark:border-slate-800 ${props.prev ? "" : "col-start-2"}`}
            >
              <div className="absolute -inset-px rounded-xl border-2 border-transparent opacity-0 [background:linear-gradient(var(--quick-links-hover-bg,theme(colors.purple.50)),var(--quick-links-hover-bg,theme(colors.purple.50)))_padding-box,linear-gradient(to_top,theme(colors.purple.400),theme(colors.purple.300),theme(colors.indigo.300))_border-box] group-hover:opacity-100 dark:[--quick-links-hover-bg:theme(colors.slate.800)]"></div>
              <div className="relative overflow-hidden rounded-xl py-3 px-6 text-right">
                <h2 className="mt-1 font-display text-base text-slate-900 dark:text-white">
                  <a
                    href={`/workspace/${props.username}/o/${props.repo_name}/${props.next}`}
                  >
                    <span className="absolute -inset-px rounded-xl"></span>
                    {t("NextPage")}
                  </a>
                </h2>
                <p className="mt-1 text-sm text-slate-700 dark:text-slate-400">
                  {props.next}
                </p>
              </div>
            </div>
          )}
        </div>

        <footer className="justify-between pt-10 border-t border-gray-100 sm:flex dark:border-gray-800/50 pb-10">
          <div className="flex mb-6 space-x-3 sm:mb-0">
            {props.footers?.map((footer: FooterSocial, index: any) => (
              <Link key={index} className="group" href={footer.url}>
                <IconFilter
                  icon_name={footer.icon}
                  class_name="h-5 w-5 fill-slate-500 group-hover:fill-slate-700"
                />
              </Link>
            ))}

            <div className="text-sm text-slate-500 hover:text-slate-700">
              {t("UpdatedAt")}
              <TimeElement timeInfo={props.updated_at ?? ""} />
            </div>
          </div>
          <div className="sm:flex">
            <Link
              href="https://github.com/zizdlp/zbook"
              target="_blank"
              rel="noreferrer"
              className="text-sm text-slate-500 hover:text-slate-700"
            >
              {t("PowerBy")}
            </Link>
          </div>
        </footer>
      </div>

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
    </MainContentWrapper>
  );
}
