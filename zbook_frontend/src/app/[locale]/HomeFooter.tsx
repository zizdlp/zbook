import { FaDiscord, FaGithub } from "react-icons/fa";
import { FaBilibili } from "react-icons/fa6";
import { FaYoutube } from "react-icons/fa";
import { Link } from "@/navigation";
import { getTranslations } from "next-intl/server";
import { beian } from "@/utils/env_variable";

async function LinkElement({ url, title }: { url: string; title: string }) {
  return (
    <li>
      <Link
        className="hover:text-slate-900 dark:hover:text-slate-300"
        href={url}
      >
        {title}
      </Link>
    </li>
  );
}
export default async function HomeFooter() {
  const t = await getTranslations("HomePage");
  return (
    <footer className="bg-gradient-to-b from-[#e3d5b1] pt-16  to-[#ffffff] dark:from-transparent dark:via-transparent dark:to-transparent text-slate-700 dark:text-slate-500">
      <div className="mx-auto max-w-7xl px-4 sm:px-6 lg:px-8">
        <div className="grid grid-cols-8 gap-4 gap-x-8 px-2 py-8">
          <div className="lg:col-span-2 col-span-4">
            <h2 className="font-semibold text-slate-900 dark:text-slate-100">
              {t("GettingStarted")}
            </h2>
            <ul className="mt-3 space-y-2">
              <LinkElement url="/auth/login" title={t("Login")} />
              <LinkElement url="/auth/register" title={t("Register")} />
              <LinkElement url="/cases" title={t("Cases")} />
              <LinkElement
                url={`/workspace/${process.env.DOC_USERNAME}/o/${process.env.DOC_REPONAME}`}
                title={t("Documentation")}
              />
            </ul>
          </div>
          <div className="lg:col-span-2 col-span-4">
            <h2 className="font-semibold text-slate-900 dark:text-slate-100">
              {t("CoreConcepts")}
            </h2>
            <ul className="mt-3 space-y-2">
              <LinkElement
                url="#features_section"
                title={t("FeatureSection")}
              />
              <LinkElement url="#dashboard" title={t("DashboardSection")} />
              <LinkElement url="#create_repo" title={t("CreateRepoSection")} />
              <LinkElement
                url="#multi_user_section"
                title={t("MultiUserSection")}
              />
            </ul>
          </div>
          <div className="lg:col-span-2 col-span-4">
            <h2 className="font-semibold text-slate-900 dark:text-slate-100">
              {t("Community")}
            </h2>
            <ul className="mt-3 space-y-2">
              <LinkElement
                url="https://github.com/zizdlp/zbook"
                title="GitHub"
              />
              <LinkElement
                url="https://space.bilibili.com/1448262500"
                title="Bilibili"
              />
              <LinkElement
                url="https://www.youtube.com/channel/UC9D6VAJRoG7bD38dz8F9CSg"
                title="YouTube"
              />
              <LinkElement
                url="https://discord.com/channels/1250069935594536960/1250069935594536963"
                title="Discord"
              />
            </ul>
          </div>
          <div className="lg:col-span-2 col-span-4">
            <h2 className="font-semibold text-slate-900 dark:text-slate-100">
              {t("Links")}
            </h2>
            <ul className="mt-3 space-y-2">
              <li>
                <Link
                  className="hover:text-slate-900 dark:hover:text-slate-300"
                  href="/terms"
                >
                  {t("Terms")}
                </Link>
              </li>
              <li>
                <Link
                  className="hover:text-slate-900 dark:hover:text-slate-300"
                  href="/privacy"
                >
                  {t("Privacy")}
                </Link>
              </li>
              <li>
                <Link
                  className="hover:text-slate-900 dark:hover:text-slate-300"
                  href="https://linchat.zizdlp.com"
                >
                  {t("LinChatWeb")}
                </Link>
              </li>
              <li>
                <Link
                  className="hover:text-slate-900 dark:hover:text-slate-300"
                  href="https://apps.apple.com/cn/app/%E9%82%BB%E4%BF%A1-%E8%93%9D%E7%89%99%E7%95%85%E8%81%8A/id6472197439"
                >
                  {t("LinChatApp")}
                </Link>
              </li>
            </ul>
          </div>
        </div>

        <div className="flex flex-col items-center border-t-[0.01rem] border-slate-400 dark:border-slate-700 py-10 sm:flex-row-reverse sm:justify-between">
          <div className="flex gap-x-6">
            <Link
              className="group"
              aria-label="ZBook on Discord"
              href="https://discord.com/channels/1250069935594536960/1250069935594536963"
            >
              <FaDiscord className="h-5 w-5 fill-slate-500 group-hover:fill-slate-700" />
            </Link>
            <Link
              className="group"
              aria-label="ZBook on GitHub"
              href="https://github.com/zizdlp/zbook"
            >
              <FaGithub className="h-5 w-5 fill-slate-500 group-hover:fill-slate-700" />
            </Link>
            <Link
              className="group"
              aria-label="ZBook on Bilibili"
              href="https://space.bilibili.com/1448262500"
            >
              <FaBilibili className="h-5 w-5 fill-slate-500 group-hover:fill-slate-700" />
            </Link>
            <Link
              className="group"
              aria-label="ZBook on Youtube"
              href="https://www.youtube.com/channel/UC9D6VAJRoG7bD38dz8F9CSg"
            >
              <FaYoutube className="h-5 w-5 fill-slate-500 group-hover:fill-slate-700" />
            </Link>
          </div>
          {beian && (
            <Link
              href="https://beian.miit.gov.cn/"
              className="mt-6 text-sm sm:mt-0"
            >
              {beian}
            </Link>
          )}

          <p className="text-sm">{t("Copyright")}</p>
        </div>
      </div>
    </footer>
  );
}
