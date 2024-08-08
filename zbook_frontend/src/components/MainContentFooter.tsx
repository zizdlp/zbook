import { FooterSocial } from "@/types/interface";
import { getTranslations } from "next-intl/server";
import { Link } from "@/navigation";
import IconFilter from "./IconFilter";
import TimeElement from "./TimeElement";
export default async function MainContentFooter({
  prev,
  next,
  username,
  repo_name,
  updated_at,
  footers,
}: {
  prev: string;
  next: string;
  username: string;
  repo_name: string;
  updated_at: string;
  footers: FooterSocial[];
}) {
  const t = await getTranslations("Footer");
  return (
    <div className="leading-6 mt-14">
      <div className="not-prose my-12 grid lg:grid-cols-2 gap-4">
        {prev && (
          <div className="group relative rounded-xl border border-slate-200 dark:border-slate-800">
            <div className="absolute -inset-px rounded-xl border-2 border-transparent opacity-0 [background:linear-gradient(var(--quick-links-hover-bg,theme(colors.purple.50)),var(--quick-links-hover-bg,theme(colors.purple.50)))_padding-box,linear-gradient(to_top,theme(colors.purple.400),theme(colors.purple.300),theme(colors.indigo.300))_border-box] group-hover:opacity-100 dark:[--quick-links-hover-bg:theme(colors.slate.800)]"></div>
            <div className="relative overflow-hidden rounded-xl py-3 px-6 text-left">
              <h2 className="mt-1 font-display text-base text-slate-900 dark:text-white">
                <a href={`/workspace/${username}/o/${repo_name}/${prev}`}>
                  <span className="absolute -inset-px rounded-xl"></span>
                  {t("PrevPage")}
                </a>
              </h2>
              <p className="mt-1 text-sm text-slate-700 dark:text-slate-400">
                {prev}
              </p>
            </div>
          </div>
        )}

        {next && (
          <div
            className={`group relative rounded-xl border border-slate-200 dark:border-slate-800 ${prev ? "" : "lg:col-start-2"}`}
          >
            <div className="absolute -inset-px rounded-xl border-2 border-transparent opacity-0 [background:linear-gradient(var(--quick-links-hover-bg,theme(colors.purple.50)),var(--quick-links-hover-bg,theme(colors.purple.50)))_padding-box,linear-gradient(to_top,theme(colors.purple.400),theme(colors.purple.300),theme(colors.indigo.300))_border-box] group-hover:opacity-100 dark:[--quick-links-hover-bg:theme(colors.slate.800)]"></div>
            <div className="relative overflow-hidden rounded-xl py-3 px-6 text-right">
              <h2 className="mt-1 font-display text-base text-slate-900 dark:text-white">
                <a href={`/workspace/${username}/o/${repo_name}/${next}`}>
                  <span className="absolute -inset-px rounded-xl"></span>
                  {t("NextPage")}
                </a>
              </h2>
              <p className="mt-1 text-sm text-slate-700 dark:text-slate-400">
                {next}
              </p>
            </div>
          </div>
        )}
      </div>

      <footer className="justify-between pt-10 border-t border-gray-100 sm:flex dark:border-gray-800/50 pb-10">
        <div className="flex mb-6 space-x-3 sm:mb-0">
          {footers?.map((footer: FooterSocial, index: any) => (
            <Link key={index} className="group" href={footer.url}>
              <IconFilter
                icon_name={footer.icon}
                class_name="h-5 w-5 fill-slate-500 group-hover:fill-slate-700"
              />
            </Link>
          ))}

          <div className="text-sm text-slate-500 hover:text-slate-700">
            {t("UpdatedAt")}
            <TimeElement timeInfo={updated_at ?? ""} />
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
  );
}
